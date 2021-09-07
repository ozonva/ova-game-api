package server

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-game-api/internal/flusher"
	"github.com/ozonva/ova-game-api/internal/logs"
	"github.com/ozonva/ova-game-api/internal/mocks"
	"github.com/ozonva/ova-game-api/internal/saver"
	"github.com/ozonva/ova-game-api/pkg/game"
	ovagameapi "github.com/ozonva/ova-game-api/pkg/hero-api"
	"github.com/rs/zerolog/log"
	"time"
)

func newHeroTest(userId uint64, typeHero game.TypeHero) game.Hero {
	hero := game.NewHero(userId, typeHero, "")
	hero.GenerateName()

	return hero
}

func createHeroes(userId uint64, count int) []game.Hero {
	list := make([]game.Hero, count)

	for index := range list {
		var typeHero game.TypeHero = game.GetTypeHeroesEnums()[index%3]
		list[index] = newHeroTest(userId, typeHero)
	}

	return list
}

var _ = Describe("HeroServer", func() {
	const (
		capacitySize                 = 2
		flushTimeoutShortMillisecond = 2 * time.Millisecond
		flushTimeoutLongMillisecond  = 5 * time.Millisecond
	)
	var (
		ctrl        *gomock.Controller
		mockRepo    *mocks.MockHeroRepo
		mockMetrics *mocks.MockMetrics
		mockKafka   *mocks.MockKafkaProducer
		api         ovagameapi.HeroApiServer
		ctx         context.Context
	)

	BeforeEach(func() {
		logs.InitLogger()
		ctrl = gomock.NewController(GinkgoT())
		mockMetrics = mocks.NewMockMetrics(ctrl)
		mockKafka = mocks.NewMockKafkaProducer(ctrl)
		mockRepo = mocks.NewMockHeroRepo(ctrl)
		ctx = context.Background()
		flusher_ := flusher.NewFlusher(10, mockRepo)
		saver := saver.NewSaver(ctx, capacitySize, flusher_, flushTimeoutShortMillisecond)

		api = NewHeroApiServer(&log.Logger, mockRepo, saver, mockMetrics, mockKafka)
	})

	AfterEach(func() {
		ctrl.Finish()
		logs.FileLogger.Close()
	})

	Context("CreateHero", func() {
		It("save hero", func() {
			typeHero := game.GetTypeHeroesEnums()[0]
			hero := newHeroTest(123, typeHero)
			mockRepo.EXPECT().AddHeroes(ctx, []game.Hero{hero}).Times(1)

			res, err := api.CreateHero(ctx, &ovagameapi.CreateHeroRequest{
				UserId:      123,
				Name:        "awesome",
				TypeHero:    "Fighter",
				Description: "Anything...",
			})
			time.Sleep(flushTimeoutLongMillisecond)

			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.Hero.Id).To(Equal(hero.ID))
			Expect(res.Hero.Name).To(Equal(hero.Name))
			Expect(res.Hero.UserId).To(Equal(hero.UserID))
		})
		It("update hero", func() {
			typeHero := game.GetTypeHeroesEnums()[0]
			hero := newHeroTest(123, typeHero)
			mockRepo.EXPECT().AddHeroes(ctx, []game.Hero{hero}).Times(1)

			res, err := api.CreateHero(ctx, &ovagameapi.CreateHeroRequest{
				UserId:      123,
				Name:        "awesome",
				TypeHero:    "Fighter",
				Description: "Anything...",
			})

			Expect(err).ShouldNot(HaveOccurred())

			res2, err := api.UpdateHero(ctx, &ovagameapi.UpdateHeroRequest{
				Id:          hero.ID.String(),
				UserId:      125,
				Name:        "awesome2",
				TypeHero:    "Magician",
				Description: "Anything update...",
			})
			time.Sleep(flushTimeoutLongMillisecond)

			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.Hero.Id).To(Equal(res2.Hero.Id))
			Expect(res.Hero.Name).ToNot(Equal(res2.Hero.Name))
			Expect(res.Hero.UserId).ToNot(Equal(res2.Hero.UserId))
			Expect(res.Hero.TypeHero).ToNot(Equal(res2.Hero.TypeHero))
			Expect(res.Hero.Description).ToNot(Equal(res2.Hero.Description))
		})

		It("DescribeHero", func() {
			typeHero := game.GetTypeHeroesEnums()[0]
			hero := newHeroTest(123, typeHero)
			mockRepo.EXPECT().DescribeHero(ctx, hero.ID).Return(&hero, nil).Times(1)

			_, err := api.DescribeHero(ctx, &ovagameapi.DescribeHeroRequest{
				Id: hero.ID.String(),
			})
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("ListHeroes", func() {
			typeHero := game.GetTypeHeroesEnums()[0]
			hero := newHeroTest(123, typeHero)
			mockRepo.EXPECT().ListHeroes(ctx, 0, 10).Return([]game.Hero{hero}, nil).Times(1)

			res, err := api.ListHeroes(ctx, &ovagameapi.ListHeroRequest{
				Limit:  uint64(10),
				Offset: uint64(0),
			})
			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.Heroes[0].Id).To(Equal(hero.ID))
			Expect(res.Heroes[0].Name).To(Equal(hero.Name))
			Expect(res.Heroes[0].UserId).To(Equal(hero.UserID))
		})

		It("RemoveHero", func() {
			typeHero := game.GetTypeHeroesEnums()[0]
			hero := newHeroTest(123, typeHero)
			mockRepo.EXPECT().RemoveHero(ctx, hero.ID).Return(nil).Times(1)

			_, err := api.RemoveHero(ctx, &ovagameapi.RemoveHeroRequest{
				Id: hero.ID.String(),
			})
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
	Context("MULTI create heroes", func() {
		It("heroes count > batch size", func() {
			heroes := createHeroes(1, 3)
			heroRequest := make([]*ovagameapi.HeroRequest, len(heroes))
			for _, item := range heroes {
				heroRequest = append(heroRequest, &ovagameapi.HeroRequest{
					UserId:      item.UserID,
					TypeHero:    item.TypeHero.String(),
					Name:        item.Name,
					Description: item.Description,
				})
			}

			mockRepo.EXPECT().AddHeroes(gomock.Any(), heroes[:2]).Return(nil).Times(1)
			mockRepo.EXPECT().AddHeroes(gomock.Any(), heroes[2:]).Return(nil).Times(1)

			req := ovagameapi.MultiCreateHeroRequest{
				Heroes: heroRequest,
			}
			_, err := api.MultiCreateHero(ctx, &req)
			Expect(err).To(BeNil())
		})
	})
})

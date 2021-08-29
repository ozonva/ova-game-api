package server

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-game-api/internal/flusher"
	"github.com/ozonva/ova-game-api/internal/mocks"
	"github.com/ozonva/ova-game-api/internal/saver"
	"github.com/ozonva/ova-game-api/pkg/game"
	ovagameapi "github.com/ozonva/ova-game-api/pkg/hero-api"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func newHeroTest(userId uint64, typeHero game.TypeHero) game.Hero {
	hero := game.NewHero(userId, typeHero, "")
	hero.GenerateName()

	return hero
}

var _ = Describe("HeroServer", func() {
	const (
		capacitySize                 = 2
		flushTimeoutShortMillisecond = 2 * time.Millisecond
		flushTimeoutLongMillisecond  = 5 * time.Millisecond
	)
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockHeroRepo
		api      ovagameapi.HeroApiServer
		ctx      context.Context
	)

	BeforeEach(func() {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		log := zerolog.New(output).With().Timestamp().Logger()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockHeroRepo(ctrl)
		ctx = context.Background()
		flusher_ := flusher.NewFlusher(10, mockRepo)
		saver := saver.NewSaver(capacitySize, flusher_, flushTimeoutShortMillisecond)

		api = NewHeroApiServer(&log, mockRepo, saver)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("CreateHero", func() {
		It("save hero", func() {
			typeHero := game.GetTypeHeroesEnums()[0]
			hero := newHeroTest(123, typeHero)
			mockRepo.EXPECT().AddHeroes([]game.Hero{hero}).Times(1)

			_, err := api.CreateHero(ctx, &ovagameapi.CreateHeroRequest{
				UserId:      123,
				Name:        "awesome",
				TypeHero:    "Fighter",
				Description: "Anything...",
			})
			time.Sleep(flushTimeoutLongMillisecond)

			Expect(err).ShouldNot(HaveOccurred())
		})

		It("DescribeHero", func() {
			typeHero := game.GetTypeHeroesEnums()[0]
			hero := newHeroTest(123, typeHero)
			mockRepo.EXPECT().DescribeHero(hero.ID).Return(&hero, nil).Times(1)

			_, err := api.DescribeHero(ctx, &ovagameapi.DescribeHeroRequest{
				Id: hero.ID.String(),
			})
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("ListRules", func() {
			typeHero := game.GetTypeHeroesEnums()[0]
			hero := newHeroTest(123, typeHero)
			mockRepo.EXPECT().ListHeroes(0, 10).Return([]game.Hero{hero}, nil).Times(1)

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
			mockRepo.EXPECT().RemoveHero(hero.ID).Return(nil).Times(1)

			_, err := api.RemoveHero(ctx, &ovagameapi.RemoveHeroRequest{
				Id: hero.ID.String(),
			})
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})

package flusher

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-game-api/internal/mocks"
	"github.com/ozonva/ova-game-api/pkg/game"
	"testing"
)

func TestFlusher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flusher")
}

func newHeroTest(userId uint64, typeHero game.TypeHero) game.Hero {
	hero := game.NewHero(userId, typeHero, "")
	hero.GenerateName()

	return hero
}

var _ = Describe("Flusher", func() {
	const chunkSize = 2
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *mocks.MockHeroRepo
		testFlusher Flusher
		ctx         context.Context
	)
	typeHero := game.GetTypeHeroesEnums()[0]
	heroes := []game.Hero{
		newHeroTest(1, typeHero),
		newHeroTest(2, typeHero),
		newHeroTest(3, typeHero),
		newHeroTest(4, typeHero),
		newHeroTest(5, typeHero),
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockHeroRepo(mockCtrl)
		testFlusher = NewFlusher(chunkSize, mockRepo)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Writing data to storage", func() {
		When("Write success", func() {
			AssertReturnNil := func(heroes []game.Hero) {
				Expect(testFlusher.Flush(ctx, heroes)).To(BeNil())
			}
			Context("Write count less than chunkSize", func() {
				oneHero := heroes[:1]
				BeforeEach(func() {
					mockRepo.EXPECT().AddHeroes(ctx, oneHero).Return(nil).Times(1)
				})
				It("Should return nil", func() {
					AssertReturnNil(oneHero)
				})
			})
			Context("Write count equal chunkSize", func() {
				heroes := heroes[:chunkSize]
				BeforeEach(func() {
					mockRepo.EXPECT().AddHeroes(ctx, heroes).Return(nil).Times(1)
				})
				It("Should return nil", func() {
					AssertReturnNil(heroes)
				})
			})
			Context("Write count more than chunkSize", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddHeroes(ctx, heroes[:chunkSize]).Return(nil).Times(1),
						mockRepo.EXPECT().AddHeroes(ctx, heroes[chunkSize:chunkSize*2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddHeroes(ctx, heroes[chunkSize*2:]).Return(nil).Times(1),
					)
				})
				It("Should return nil", func() {
					AssertReturnNil(heroes)
				})
			})
		})
		When("Write error", func() {
			err := fmt.Errorf("error writing data")
			AssertReturnHeroes := func(returnHeroes []game.Hero) {
				Expect(testFlusher.Flush(ctx, heroes)).To(Equal(returnHeroes))
			}
			Context("All data", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddHeroes(ctx, heroes[:chunkSize]).Return(err).Times(1),
						mockRepo.EXPECT().AddHeroes(ctx, heroes[chunkSize:chunkSize*2]).Return(err).Times(1),
						mockRepo.EXPECT().AddHeroes(ctx, heroes[chunkSize*2:]).Return(err).Times(1),
					)
				})
				It("Should return all data", func() {
					AssertReturnHeroes(heroes)
				})
			})
			Context("Error write first chunk", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddHeroes(ctx, heroes[:chunkSize]).Return(err).Times(1),
						mockRepo.EXPECT().AddHeroes(ctx, heroes[chunkSize:chunkSize*2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddHeroes(ctx, heroes[chunkSize*2:]).Return(nil).Times(1),
					)
				})
				It("Should return first chunk", func() {
					AssertReturnHeroes(heroes[:chunkSize])
				})
			})
		})
	})
})

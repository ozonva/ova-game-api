package flusher

import (
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

var _ = Describe("Flusher", func() {
	const chunkSize = 2
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *mocks.MockHeroRepo
		testFlusher Flusher
	)
	typeHero := game.GetTypeHeroesEnums()[0]
	heroes := []game.Hero{
		game.NewHero(1, typeHero),
		game.NewHero(2, typeHero),
		game.NewHero(3, typeHero),
		game.NewHero(4, typeHero),
		game.NewHero(5, typeHero),
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
				Expect(testFlusher.Flush(heroes)).To(BeNil())
			}
			Context("Write count less than chunkSize", func() {
				oneHero := heroes[:1]
				BeforeEach(func() {
					mockRepo.EXPECT().AddHeroes(oneHero).Return(nil).Times(1)
				})
				It("Should return nil", func() {
					AssertReturnNil(oneHero)
				})
			})
			Context("Write count equal chunkSize", func() {
				heroes := heroes[:chunkSize]
				BeforeEach(func() {
					mockRepo.EXPECT().AddHeroes(heroes).Return(nil).Times(1)
				})
				It("Should return nil", func() {
					AssertReturnNil(heroes)
				})
			})
			Context("Write count more than chunkSize", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddHeroes(heroes[:chunkSize]).Return(nil).Times(1),
						mockRepo.EXPECT().AddHeroes(heroes[chunkSize:chunkSize*2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddHeroes(heroes[chunkSize*2:]).Return(nil).Times(1),
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
				Expect(testFlusher.Flush(heroes)).To(Equal(returnHeroes))
			}
			Context("All data", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddHeroes(heroes[:chunkSize]).Return(err).Times(1),
						mockRepo.EXPECT().AddHeroes(heroes[chunkSize:chunkSize*2]).Return(err).Times(1),
						mockRepo.EXPECT().AddHeroes(heroes[chunkSize*2:]).Return(err).Times(1),
					)
				})
				It("Should return all data", func() {
					AssertReturnHeroes(heroes)
				})
			})
			Context("Error write first chunk", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddHeroes(heroes[:chunkSize]).Return(err).Times(1),
						mockRepo.EXPECT().AddHeroes(heroes[chunkSize:chunkSize*2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddHeroes(heroes[chunkSize*2:]).Return(nil).Times(1),
					)
				})
				It("Should return first chunk", func() {
					AssertReturnHeroes(heroes[:chunkSize])
				})
			})
		})
	})
})

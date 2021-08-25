package saver

import (
	_ "fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-game-api/internal/mocks"
	"github.com/ozonva/ova-game-api/pkg/game"
	"testing"
	"time"
)

func TestSaver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Saver")
}

var _ = Describe("Saver", func() {
	const (
		capacitySize                 = 2
		flushTimeoutSecond           = 2 * time.Second
		flushTimeoutShortMillisecond = 2 * time.Millisecond
		flushTimeoutLongMillisecond  = 5 * time.Millisecond
	)
	var (
		mockCtrl    *gomock.Controller
		mockFlusher *mocks.MockFlusher
		testSaver   Saver
	)
	typeHero := game.GetTypeHeroesEnums()[0]
	heroes := []game.Hero{
		game.NewHero(1, typeHero),
		game.NewHero(2, typeHero),
		game.NewHero(3, typeHero),
		game.NewHero(4, typeHero),
		game.NewHero(5, typeHero),
		game.NewHero(6, typeHero),
		game.NewHero(7, typeHero),
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Save all data", func() {
		When("Save success", func() {
			Context("Save count zero", func() {
				It("Not Saved", func() {
					mockFlusher = mocks.NewMockFlusher(mockCtrl)
					testSaver = NewSaver(capacitySize, mockFlusher, flushTimeoutSecond)
					mockFlusher.EXPECT().Flush(gomock.Any()).Times(0)
					testSaver.Close()
				})
			})
			Context("Save count less than capacitySize", func() {
				It("Save hero", func() {
					mockFlusher = mocks.NewMockFlusher(mockCtrl)
					testSaver = NewSaver(capacitySize, mockFlusher, flushTimeoutSecond)

					mockFlusher.EXPECT().Flush(gomock.Any()).Times(1)

					for _, hero := range heroes {
						testSaver.Save(hero)
						break
					}

					testSaver.Close()
				})
			})
			Context("Save count more equal than capacitySize", func() {
				It("Save hero", func() {
					mockFlusher = mocks.NewMockFlusher(mockCtrl)
					testSaver = NewSaver(capacitySize, mockFlusher, flushTimeoutSecond)
					mockFlusher.EXPECT().Flush(gomock.Any()).Times(4)

					for _, hero := range heroes {
						testSaver.Save(hero)
					}

					testSaver.Close()
				})
			})
			Context("Save count equal capacitySize", func() {
				It("Save hero", func() {
					mockFlusher = mocks.NewMockFlusher(mockCtrl)
					testSaver = NewSaver(capacitySize, mockFlusher, flushTimeoutSecond)
					mockFlusher.EXPECT().Flush(gomock.Any()).Times(1)
					heroesList := []game.Hero{
						game.NewHero(1, typeHero),
						game.NewHero(2, typeHero),
					}

					for _, hero := range heroesList {
						testSaver.Save(hero)
					}

					testSaver.Close()
				})
			})
			Context("Save for time", func() {
				It("Save hero", func() {
					mockFlusher = mocks.NewMockFlusher(mockCtrl)
					testSaver = NewSaver(capacitySize, mockFlusher, flushTimeoutShortMillisecond)
					mockFlusher.EXPECT().Flush(gomock.Any()).Times(1)
					heroesList := []game.Hero{
						game.NewHero(1, typeHero),
					}

					for _, hero := range heroesList {
						testSaver.Save(hero)
					}

					time.Sleep(flushTimeoutLongMillisecond)

					testSaver.Close()
				})
			})
			Context("Save for time no data", func() {
				It("Save no hero", func() {
					mockFlusher = mocks.NewMockFlusher(mockCtrl)
					testSaver = NewSaver(capacitySize, mockFlusher, flushTimeoutShortMillisecond)
					mockFlusher.EXPECT().Flush(gomock.Any()).Times(0)
					time.Sleep(flushTimeoutLongMillisecond)
					testSaver.Close()
				})
			})
		})
	})
})

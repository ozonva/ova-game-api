package saver

import (
	"github.com/ozonva/ova-game-api/internal/flusher"
	"github.com/ozonva/ova-game-api/pkg/game"
	"log"
	"sync"
	"time"
)

type Saver interface {
	Save(hero game.Hero)
	Close()
}

func NewSaver(capacity uint, flusher flusher.Flusher, flushTimeout time.Duration) Saver {
	saver := &saver{
		localStorage: make([]game.Hero, 0, capacity),
		flusher:      flusher,
		flushTimeout: flushTimeout,
	}

	saver.init()

	return saver
}

type saver struct {
	sync.Mutex
	signalChannel chan struct{}
	localStorage  []game.Hero
	flusher       flusher.Flusher
	flushTimeout  time.Duration
}

func (s *saver) Save(hero game.Hero) {
	if len(s.localStorage) == cap(s.localStorage) {
		s.flush()
	}
	s.localStorage = append(s.localStorage, hero)
	s.signalChannel <- struct{}{}
}

func (s *saver) Close() {
	s.flush()
	close(s.signalChannel)
}

func (s *saver) init() {
	s.signalChannel = make(chan struct{})

	go s.handlerChannel(s.signalChannel)
}

func (s *saver) flush() {
	s.Lock()
	defer s.Unlock()

	if len(s.localStorage) == 0 {
		return
	}

	unsaved := s.flusher.Flush(s.localStorage)
	if len(unsaved) > 0 {
		log.Printf("warning: some entities can't be saved to database and will be discraded: \n%v\n", unsaved)
	}
	s.localStorage = make([]game.Hero, 0, cap(s.localStorage))
}

func (s *saver) handlerChannel(ch <-chan struct{}) {
	ticker := time.NewTicker(s.flushTimeout)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.flush()
		case _, ok := <-ch:
			if !ok {
				s.flush()
				return
			}
		}
	}
}

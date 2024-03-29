package saver

import (
	"context"
	"github.com/ozonva/ova-game-api/internal/flusher"
	"github.com/ozonva/ova-game-api/pkg/game"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

const countUnsafe = 3

type Saver interface {
	Save(ctx context.Context, hero game.Hero)
	Close(ctx context.Context)
}

func NewSaver(ctx context.Context, capacity uint, flusher flusher.Flusher, flushTimeout time.Duration) Saver {
	saver := &saver{
		localStorage:  make([]game.Hero, 0, capacity),
		flusher:       flusher,
		flushTimeout:  flushTimeout,
		signalChannel: make(chan struct{}, capacity),
		countUnsafe:   0,
	}

	go saver.handlerChannel(ctx, saver.signalChannel)

	return saver
}

type saver struct {
	sync.Mutex
	signalChannel chan struct{}
	localStorage  []game.Hero
	flusher       flusher.Flusher
	flushTimeout  time.Duration
	countUnsafe   uint8
}

func (s *saver) Save(ctx context.Context, hero game.Hero) {
	s.countUnsafe = 0
	if len(s.localStorage) == cap(s.localStorage) {
		s.flush(ctx)
	}
	s.localStorage = append(s.localStorage, hero)
	s.signalChannel <- struct{}{}
}

func (s *saver) Close(ctx context.Context) {
	s.flush(ctx)
	close(s.signalChannel)
}

func (s *saver) flush(ctx context.Context) {
	s.Lock()
	defer s.Unlock()

	if len(s.localStorage) == 0 {
		return
	}

	unsaved := s.flusher.Flush(ctx, s.localStorage)
	s.localStorage = make([]game.Hero, 0, cap(s.localStorage))

	if len(unsaved) > 0 {
		s.countUnsafe++
		log.Info().Msgf("warning: some entities can't be saved to database and will be discraded: \n%v\n", unsaved)
		if s.countUnsafe <= countUnsafe {
			s.localStorage = append(s.localStorage, unsaved...)
		}
	}
}

func (s *saver) handlerChannel(ctx context.Context, ch <-chan struct{}) {
	ticker := time.NewTicker(s.flushTimeout)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.flush(ctx)
		case _, ok := <-ch:
			if !ok {
				s.flush(ctx)
				return
			}
		}
	}
}

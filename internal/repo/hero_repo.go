package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ozonva/ova-game-api/pkg/game"
	"github.com/rs/zerolog/log"
)

// HeroRepo - интерфейс хранилища для сущности Hero
type HeroRepo interface {
	AddHeroes(ctx context.Context, heroes []game.Hero) error
	ListHeroes(ctx context.Context, limit, offset uint64) ([]game.Hero, error)
	DescribeHero(ctx context.Context, heroId uuid.UUID) (*game.Hero, error)
	RemoveHero(ctx context.Context, heroId uuid.UUID) error
}

func NewHeroRepo(pool *pgxpool.Pool) HeroRepo {
	return &repo{
		pool: pool,
	}
}

type repo struct {
	pool *pgxpool.Pool
}

func (r *repo) AddHeroes(ctx context.Context, heroes []game.Hero) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query := psql.Insert("heroes").Columns("id", "name", "user_id", "type_hero")
	for _, hero := range heroes {
		query = query.Values(hero.ID, hero.Name, hero.UserID, hero.TypeHero)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		log.Info().Msg(err.Error())
		return err
	}

	log.Info().Msgf("query: %s; args: %s", sql, args)

	_, err = conn.Exec(ctx, sql, args...)
	if err != nil {
		log.Info().Msg(err.Error())
		return err
	}

	return nil
}

func (r *repo) ListHeroes(ctx context.Context, limit, offset uint64) ([]game.Hero, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	sql, _, err := sq.Select("id, user_id, type_hero, name").
		From("heroes").
		Limit(limit).
		Offset(offset).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("query: %s", sql)

	var heroesList []*game.Hero
	if err = pgxscan.Select(ctx, conn, &heroesList, sql); err != nil {
		return nil, err
	}

	heroes := make([]game.Hero, len(heroesList))
	for i, ptr := range heroesList {
		heroes[i] = *ptr
	}

	return heroes, nil
}

func (r *repo) DescribeHero(ctx context.Context, heroId uuid.UUID) (*game.Hero, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.Select("id, name, user_id").
		From("heroes").
		Where(sq.Eq{"id": heroId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("query: %s; args: %s", sql, args)

	hero := game.Hero{}
	if err = pgxscan.Get(ctx, conn, &hero, sql, args...); err != nil {
		return nil, err
	}

	return &hero, nil
}

func (r *repo) RemoveHero(ctx context.Context, heroId uuid.UUID) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.Delete("heroes").
		Where(sq.Eq{"id": heroId}).
		ToSql()
	if err != nil {
		return err
	}

	log.Info().Msgf("query: %s; args: %s", sql, args)

	_, err = conn.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

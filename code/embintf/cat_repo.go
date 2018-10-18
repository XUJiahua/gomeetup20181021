package embintf

import (
	"github.com/sirupsen/logrus"
	"time"
)

type Cat struct {
	ID int64
}

type CatRepo interface {
	Get(id int64) (*Cat, error)
	Set(cat *Cat) error
}

// db impl
type dbCatRepo struct{}

// inject dependencies
func NewDBCatRepo(conn string) *dbCatRepo {
	return &dbCatRepo{}
}

func (dbCatRepo) Get(id int64) (*Cat, error) {
	logrus.Debug("reading from DB")
	time.Sleep(time.Microsecond * 200)
	return &Cat{id}, nil
}

func (dbCatRepo) Set(cat *Cat) error {
	logrus.Debug("writing to DB")
	time.Sleep(time.Second)
	return nil
}

// proxy by embedding + interface
// cache proxy impl
type cachedCatRepo struct {
	*dbCatRepo
	cache map[int64]*Cat
}

// inject dependencies
func NewCachedCatRepo(mCatRepo *dbCatRepo) *cachedCatRepo {
	return &cachedCatRepo{
		dbCatRepo: mCatRepo,
		cache:     make(map[int64]*Cat),
	}
}

// overwrite
func (r cachedCatRepo) Get(id int64) (*Cat, error) {
	logrus.Debug("reading from cache")
	if cat, ok := r.cache[id]; ok {
		return cat, nil
	}

	cat, err := r.dbCatRepo.Get(id)
	if err != nil {
		return nil, err
	}

	logrus.Debug("writing to cache")
	r.cache[id] = cat

	return cat, nil
}

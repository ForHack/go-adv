package link

import "go/adv-demo/pkg/db"

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(databse *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: databse,
	}
}

func (repo *LinkRepository) Create(link Link) {

}

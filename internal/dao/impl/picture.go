package impl

import (
	"context"
	"github.com/KawashiroNitori/MoeManager/internal/dao"
	"github.com/KawashiroNitori/MoeManager/internal/ent"
	"github.com/KawashiroNitori/MoeManager/internal/ent/picture"
	"github.com/KawashiroNitori/MoeManager/internal/storage"
	_ "github.com/mattn/go-sqlite3"
)

type PictureDAOImpl struct {
	db *ent.Client
}

var DefaultPictureDAO = NewPictureDAO()

func NewPictureDAO() dao.PictureDAO {
	return &PictureDAOImpl{
		db: storage.SqliteDB,
	}
}

func (p *PictureDAOImpl) GetPictureByFilename(ctx context.Context, filename string) *ent.Picture {
	pic, _ := p.db.Picture.Query().Where(picture.Filename(filename)).Only(ctx)
	return pic
}

func (p *PictureDAOImpl) GetPicturesByDigest(ctx context.Context, digest string) []*ent.Picture {
	pics, _ := p.db.Picture.Query().Where(picture.Digest(digest)).All(ctx)
	return pics
}

func (p *PictureDAOImpl) Create(ctx context.Context, pic *ent.Picture) (*ent.Picture, error) {
	pic, err := p.db.Picture.
		Create().
		SetFilename(pic.Filename).
		SetOriginalFilename(pic.OriginalFilename).
		SetDir(pic.Dir).
		SetDigest(pic.Digest).
		SetOriginalWidth(pic.OriginalWidth).
		SetOriginalHeight(pic.OriginalHeight).
		SetStatus(pic.Status).
		SetCreatedAt(pic.CreatedAt).
		Save(ctx)
	return pic, err
}

func (p *PictureDAOImpl) Remove(ctx context.Context, filename string) error {
	_, err := p.db.Picture.Delete().Where(picture.Filename(filename)).Exec(ctx)
	return err
}

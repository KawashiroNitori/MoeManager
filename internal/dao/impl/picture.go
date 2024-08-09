package impl

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/KawashiroNitori/MoeManager/internal/dao"
	"github.com/KawashiroNitori/MoeManager/internal/ent"
	"github.com/KawashiroNitori/MoeManager/internal/ent/picture"
	"github.com/KawashiroNitori/MoeManager/internal/macro"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
	"path/filepath"
)

type PictureDAOImpl struct {
	db *ent.Client
}

func getSqliteConnectionURL(path string) string {
	return path + "?_pragma=foreign_keys(1)"
}

func NewPictureDAO() dao.PictureDAO {
	db := lo.Must(sql.Open("sqlite", getSqliteConnectionURL(viper.GetString(macro.ConfigKeyDatabasePath))))
	drv := entsql.OpenDB(dialect.SQLite, db)
	sqliteDB := ent.NewClient(ent.Driver(drv))
	lo.Must0(sqliteDB.Schema.Create(context.Background()))
	return &PictureDAOImpl{
		db: sqliteDB,
	}
}

func (p *PictureDAOImpl) GetPictureByFilename(ctx context.Context, filename string) *ent.Picture {
	filename = filepath.Base(filename)
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

func (p *PictureDAOImpl) Close() error {
	return p.db.Close()
}

package dao

import (
	"context"
	"github.com/KawashiroNitori/MoeManager/internal/ent"
)

type PictureDAO interface {
	GetPictureByFilename(ctx context.Context, filename string) *ent.Picture
	GetPicturesByDigest(ctx context.Context, digest string) []*ent.Picture

	Create(ctx context.Context, pic *ent.Picture) (*ent.Picture, error)
	Remove(ctx context.Context, filename string) error
}

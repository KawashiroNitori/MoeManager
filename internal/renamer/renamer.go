package renamer

import (
	"context"
	"fmt"
	"github.com/KawashiroNitori/MoeManager/internal/ent"
	"github.com/KawashiroNitori/MoeManager/internal/macro"
	"github.com/KawashiroNitori/MoeManager/internal/util"
	"github.com/alexflint/go-filemutex"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Renamer interface {
	Rename(ctx context.Context, path string, fileMu *filemutex.FileMutex, pic *ent.Picture) (*ent.Picture, error)
}

type renamer struct {
	validFilenameRe *regexp.Regexp
}

func NewRenamer() Renamer {
	return &renamer{
		validFilenameRe: regexp.MustCompile(`^\d{8}-\d{6}-[\dA-F]{6}(?:-.+)?$`),
	}
}

func (r *renamer) getValidFilename(path string, pic *ent.Picture) string {
	lastModify := pic.CreatedAt
	digest := pic.Digest
	filename := fmt.Sprintf("%s-%s%s",
		lastModify.Format("20060102-150405"),
		strings.ToUpper(digest[:6]), filepath.Ext(path))
	duplicateCount := 0
	for util.IsExists(filepath.Join(filepath.Dir(path), filename)) {
		duplicateCount++
		filename = fmt.Sprintf("%s-%s-%d%s",
			lastModify.Format("20060102-150405"),
			strings.ToUpper(digest[:6]),
			duplicateCount, filepath.Ext(path))
	}
	return filename
}

func (r *renamer) Rename(ctx context.Context, path string, fileMu *filemutex.FileMutex, pic *ent.Picture) (*ent.Picture, error) {
	if !viper.GetBool(macro.ConfigKeyRenameEnabled) {
		return pic, nil
	}
	filename := filepath.Base(path)
	if !util.IsSupportedExtension(viper.GetStringSlice(macro.ConfigKeyRenameExtensions), filename) {
		return pic, nil
	}
	filenameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
	if r.validFilenameRe.MatchString(filenameWithoutExt) {
		return pic, nil
	}
	newFilename := r.getValidFilename(path, pic)
	pic, err := pic.Update().
		SetFilename(newFilename).
		SetOriginalFilename(filename).
		Save(ctx)
	if err != nil {
		return pic, err
	}
	return pic, os.Rename(path, filepath.Join(filepath.Dir(path), newFilename))
}

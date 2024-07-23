package upscaler

import (
	"context"
	"fmt"
	"github.com/KawashiroNitori/MoeManager/internal/ent"
	"github.com/KawashiroNitori/MoeManager/internal/macro"
	"github.com/KawashiroNitori/MoeManager/internal/util"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Upscaler interface {
	Upscale(ctx context.Context, path string, pic *ent.Picture) error
}

type upscaler struct {
}

func NewUpscaler() Upscaler {
	return &upscaler{}
}

func (u *upscaler) findRatio(longSide int) int {
	ratioRange := viper.GetIntSlice(macro.ConfigKeyUpscaleRatioRange)
	if len(ratioRange) != 2 {
		return 1
	}
	for ratio := ratioRange[0]; ratio <= ratioRange[1]; ratio++ {
		if longSide*ratio >= viper.GetInt(macro.ConfigKeyUpscaleTargetLongSide) {
			return ratio
		}
	}
	return ratioRange[1]
}

func (u *upscaler) Upscale(ctx context.Context, path string, pic *ent.Picture) error {
	if !viper.GetBool(macro.ConfigKeyUpscaleEnabled) {
		return nil
	}
	if !util.IsSupportedExtension(viper.GetStringSlice(macro.ConfigKeyUpscaleExtensions), path) {
		return nil
	}
	if pic.IsUpscaled {
		return nil
	}
	width, height := pic.OriginalWidth, pic.OriginalHeight
	if width <= 0 || height <= 0 {
		return fmt.Errorf("invalid picture size: %dx%d", width, height)
	}
	longSide := max(width, height)
	ratio := u.findRatio(longSide)
	if longSide < viper.GetInt(macro.ConfigKeyUpscaleMinLongSide) || longSide >= viper.GetInt(macro.ConfigKeyUpscaleTargetLongSide) || ratio == 1 {
		_, err := pic.Update().
			SetIsUpscaled(true).
			SetUpscaledWidth(width).
			SetUpscaledHeight(height).
			SetUpscaleRatio(1).
			Save(ctx)
		return err
	}
	outputPath := filepath.Join(
		os.TempDir(),
		fmt.Sprintf("%s.%s", lo.RandomString(10, lo.AlphanumericCharset), viper.GetString(macro.ConfigKeyUpscaleFormat)),
	)
	defer os.Remove(outputPath)

	args := viper.GetStringSlice(macro.ConfigKeyUpscalerArgs)
	args = append(args, viper.GetString(macro.ConfigKeyUpscaleFormatArg), viper.GetString(macro.ConfigKeyUpscaleFormat))
	args = append(args, viper.GetString(macro.ConfigKeyUpscaleRatioArg), cast.ToString(ratio))
	args = append(args, "-i", path, "-o", outputPath)
	upscaleCmd := exec.Command(viper.GetString(macro.ConfigKeyUpscalerPath), args...)
	if err := upscaleCmd.Run(); err != nil {
		return err
	}
	newFilename := strings.TrimSuffix(pic.Filename, filepath.Ext(pic.Filename)) + "." + viper.GetString(macro.ConfigKeyUpscaleFormat)
	newPath := filepath.Join(filepath.Dir(path), newFilename)
	newDigest, err := util.GetDigest(outputPath)
	if err != nil {
		return err
	}
	pic, err = pic.Update().
		SetIsUpscaled(true).
		SetUpscaledWidth(width * ratio).
		SetUpscaledHeight(height * ratio).
		SetUpscaleRatio(ratio).
		SetDigest(newDigest).
		Save(ctx)
	if err != nil {
		return err
	}
	_ = os.Remove(path)
	return os.Rename(outputPath, newPath)
}

package macro

import "github.com/spf13/viper"

const (
	ConfigKeyDatabasePath = "database"
	ConfigKeyIncludeDirs  = "includes"

	ConfigKeyRenameEnabled    = "rename.enabled"
	ConfigKeyRenameExtensions = "rename.extensions"

	ConfigKeyUpscaleEnabled        = "upscale.enabled"
	ConfigKeyUpscaleExtensions     = "upscale.extensions"
	ConfigKeyUpscalerPath          = "upscale.upscaler"
	ConfigKeyUpscalerArgs          = "upscale.arguments"
	ConfigKeyUpscaleFormatArg      = "upscale.format.argument"
	ConfigKeyUpscaleFormat         = "upscale.format.format"
	ConfigKeyUpscaleRatioArg       = "upscale.ratio.argument"
	ConfigKeyUpscaleRatioRange     = "upscale.ratio.range"
	ConfigKeyUpscaleMinLongSide    = "upscale.min_long_side"
	ConfigKeyUpscaleTargetLongSide = "upscale.target_long_side"
)

var configDefaultMap = map[string]any{
	ConfigKeyDatabasePath:      "moe_manager1.db",
	ConfigKeyRenameEnabled:     true,
	ConfigKeyUpscaleEnabled:    true,
	ConfigKeyUpscaleFormatArg:  "-f",
	ConfigKeyUpscaleFormat:     "png",
	ConfigKeyUpscaleRatioArg:   "-s",
	ConfigKeyUpscaleRatioRange: []int{2, 4},
}

func init() {
	for k, v := range configDefaultMap {
		viper.SetDefault(k, v)
	}
}

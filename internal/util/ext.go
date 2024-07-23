package util

import (
	"github.com/KawashiroNitori/MoeManager/internal/macro"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func GetAllSupportedExtensions() []string {
	extensions := make([]string, 0)
	extensions = append(extensions, viper.GetStringSlice(macro.ConfigKeyRenameExtensions)...)
	extensions = append(extensions, viper.GetStringSlice(macro.ConfigKeyUpscaleExtensions)...)
	return extensions
}

func IsSupportedExtension(extensions []string, path string) bool {
	extension := strings.TrimLeft(filepath.Ext(path), ".")
	if len(extension) == 0 {
		return false
	}
	for _, ext := range extensions {
		if strings.TrimLeft(ext, ".") == extension {
			return true
		}
	}
	return false
}

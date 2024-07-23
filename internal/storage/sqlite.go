package storage

import (
	"github.com/KawashiroNitori/MoeManager/internal/ent"
	"github.com/KawashiroNitori/MoeManager/internal/macro"
	_ "github.com/mattn/go-sqlite3"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"net/url"
)

var (
	SqliteDB *ent.Client
)

func getSqliteConnectionURL(path string) string {
	u := url.URL{
		Scheme:   "file",
		Path:     path,
		RawQuery: "cache=shared&mode=rwc",
	}
	return u.String()
}

func init() {
	SqliteDB = lo.Must(ent.Open("sqlite3", getSqliteConnectionURL(viper.GetString(macro.ConfigKeyDatabasePath))))
}

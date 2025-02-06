package src

import (
	"fmt"

	"github.com/clidey/whodb/core/src/engine"
	"github.com/clidey/whodb/core/src/env"
	"github.com/clidey/whodb/core/src/plugins/clickhouse"
	"github.com/clidey/whodb/core/src/plugins/elasticsearch"
	"github.com/clidey/whodb/core/src/plugins/mongodb"
	"github.com/clidey/whodb/core/src/plugins/mysql"
	"github.com/clidey/whodb/core/src/plugins/postgres"
	"github.com/clidey/whodb/core/src/plugins/redis"
	"github.com/clidey/whodb/core/src/plugins/sqlite3"
)

var MainEngine *engine.Engine

func InitializeEngine() *engine.Engine {
	MainEngine = &engine.Engine{}
	MainEngine.RegistryPlugin(postgres.NewPostgresPlugin())
	MainEngine.RegistryPlugin(mysql.NewMySQLPlugin())
	MainEngine.RegistryPlugin(mysql.NewMyMariaDBPlugin())
	MainEngine.RegistryPlugin(sqlite3.NewSqlite3Plugin())
	MainEngine.RegistryPlugin(mongodb.NewMongoDBPlugin())
	MainEngine.RegistryPlugin(redis.NewRedisPlugin())
	MainEngine.RegistryPlugin(elasticsearch.NewElasticSearchPlugin())
	MainEngine.RegistryPlugin(clickhouse.NewClickHousePlugin())
	return MainEngine
}

func GetLoginProfiles() []env.DatabaseCredentials {
	profiles := []env.DatabaseCredentials{}
	for _, plugin := range MainEngine.Plugins {
		databaseProfiles := env.GetDefaultDatabaseCredentials(string(plugin.Type))
		for _, databaseProfile := range databaseProfiles {
			databaseProfile.Type = string(plugin.Type)
			databaseProfile.IsProfile = true
			profiles = append(profiles, databaseProfile)
		}
	}
	return profiles
}

func GetLoginProfileId(index int, profile env.DatabaseCredentials) string {
	if len(profile.Alias) > 0 {
		return profile.Alias
	}
	return fmt.Sprintf("#%v - %v@%v [%v]", index+1, profile.Username, profile.Hostname, profile.Database)
}

func GetLoginCredentials(profile env.DatabaseCredentials) *engine.Credentials {
	advanced := []engine.Record{
		{
			Key:   "Port",
			Value: profile.Port,
		},
	}

	for key, value := range profile.Config {
		advanced = append(advanced, engine.Record{
			Key:   key,
			Value: value,
		})
	}

	return &engine.Credentials{
		Type:      profile.Type,
		Hostname:  profile.Hostname,
		Username:  profile.Username,
		Password:  profile.Password,
		Database:  profile.Database,
		Advanced:  advanced,
		IsProfile: profile.IsProfile,
	}
}

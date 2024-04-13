package config

type Configs struct {
	DBDriverName     string
	DBDataSourceName string
}

func NewConfigs() *Configs {
	return &Configs{
        DBDriverName: "sqlite",
        DBDataSourceName: "scripts/simplify",
    }
}

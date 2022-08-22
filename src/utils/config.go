package utils

type databaseEnv struct {
	User     string `json:"db_user,omitempty"`
	Password string `json:"db_password,omitempty"`
	Name     string `json:"db_name,omitempty"`
	Host     string `json:"db_host,omitempty"`
}

var (
	DATABASE databaseEnv
)

// LoadEnv loads all environment variables
func LoadEnv() {
	ParseFile("./.credentials/database.json", &DATABASE)
}

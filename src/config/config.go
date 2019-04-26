package config

type Tables struct {
	User    string
	Tenant  string
	Product string
	Order   string
}

type Roles struct {
	Sns string
}

type Configuration struct {
	Environment   string
	AwsRegion     string
	CognitoRegion string
	AwsAccount    string
	Domain        string
	ServiceUrl    string
	Name          string
	Table         *Tables
	UserRole      string
	Role          *Roles
	Tier          string
	Port          string
	LogLevel      string
}

func Configure(env string) *Configuration {
	c := &Configuration{
		Table: &Tables{
			User: "User",
		},
	}
	switch env {
	case "PRODUNCTION":

		break
	case "DEVELOPMENT":
		break
	}
	return c
}

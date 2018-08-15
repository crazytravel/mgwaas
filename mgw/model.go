package mgw

// User model
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// EdgeProxy model
type EdgeProxy struct {
	ID        string `json:"id"`
	ProxyName string `json:"proxyName"`
	TargetURL string `json:"targetURL"`
}

// EdgeConfig struct is part of template config properties
type EdgeConfig struct {
	Bootstrap        string `json:"bootstrap"`
	JwtPublicKey     string `json:"jwtPublicKey"`
	ManagementURI    string `json:"managementURI"`
	VaultName        string `json:"vaultName"`
	AuthURI          string `json:"authURI"`
	BaseURI          string `json:"baseURI"`
	BootstrapMessage string `json:"bootstrapMessage"`
	KeySecretMessage string `json:"keySecretMessage"`
	Products         string `json:"products"`
}

type EdgeMicro struct {
	Port                      uint     `json:"port"`
	MaxConnections            uint     `json:"maxConnections"`
	ConfigChangePollInterval  uint     `json:"configChangePollInterval"`
	DisableConfigPollInterval bool     `json:"disableConfigPollInterval"`
	Logging                   *Logging `json:"logging"`
	Plugins                   *Plugins `json:"Plugins"`
}

type Headers struct {
	XForwardedFor  bool `json:"xForwardedFor"`
	XForwardedHost bool `json:"xForwardedHost"`
	XRequestID     bool `json:"xRequestID"`
	XResponseTime  bool `json:"xResponseTime"`
	Via            bool `json:"via"`
}

type OAuth struct {
	AllowNoAuthorization      bool   `json:"allowNoAuthorization"`
	AllowInvalidAuthorization bool   `json:"allowInvalidAuthorization"`
	AuthorizationHeader       string `json:"authorizationHeader"`
	APIKeyHeader              string `json:"APIKeyHeader"`
	KeepAuthHeader            bool   `json:"k eepAuthHeader"`
	VerifyAPIKeyURL           string `json:"verifyAPIKeyURL"`
}

type Analytics struct {
	URI           string `json:"uri"`
	BufferSize    uint   `json:"bufferSize"`
	BatchSize     uint   `json:"batchSize"`
	FlushInterval uint   `json:"flushInterval"`
}

type Logging struct {
	Level            string `json:"level"`
	Dir              string `json:"dir"`
	StatsLogInterval uint   `json:"statsLogInterval"`
	RotateInterval   uint   `json:"rotateInterval"`
}

type Plugins struct {
	Dir      string   `json:"dir"`
	Sequence []string `json:"sequence"`
}

type Debug struct {
	Port uint   `json:"port"`
	Args string `json:"args"`
}

// MicroGatewayTemplate struct
type MicroGatewayTemplate struct {
	*EdgeConfig `json:"edgeConfig"`
	*EdgeMicro  `json:"edgeMicro"`
	*Headers    `json:"headers"`
	*OAuth      `json:"oAuth"`
	*Analytics  `json:"analytics"`
}

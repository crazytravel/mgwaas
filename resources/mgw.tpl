edge_config:
  bootstrap: {{.EdgeConfig.Bootstrap}}
  jwt_public_key: {{.EdgeConfig.JwtPublicKey}}
  managementUri: {{.EdgeConfig.ManagementURI}}
  vaultName: {{.EdgeConfig.VaultName}}
  authUri: {{.EdgeConfig.AuthURI}}
  baseUri: {{.EdgeConfig.BaseURI}}
  bootstrapMessage: {{.EdgeConfig.BootstrapMessage}}
  keySecretMessage: {{.EdgeConfig.KeySecretMessage}}
  products: {{.EdgeConfig.Products}}
edgemicro:
  port: {{.EdgeMicro.Port}}
  max_connections: {{.EdgeMicro.MaxConnections}}
  config_change_poll_interval: {{.EdgeMicro.ConfigChangePollInterval}}
  disable_config_poll_interval: {{.EdgeMicro.DisableConfigPollInterval}}
  logging:
    level: {{.EdgeMicro.Logging.Level}}
    dir: {{.EdgeMicro.Logging.Dir}}
    stats_log_interval: {{.EdgeMicro.Logging.StatsLogInterval}}
    rotate_interval: {{.EdgeMicro.Logging.RotateInterval}}
  plugins:
    sequence: 
      {{range $k, $v := .EdgeMicro.Plugins.Sequence}} - {{$v}}{{end}}
headers:
  x-forwarded-for: {{.Headers.XForwardedFor}}
  x-forwarded-host: {{.Headers.XForwardedHost}}
  x-request-id: {{.Headers.XRequestID}}
  x-response-time: {{.Headers.XResponseTime}}
  via: {{.Headers.Via}}
oauth:
  allowNoAuthorization: {{.OAuth.AllowNoAuthorization}}
  allowInvalidAuthorization: {{.OAuth.AllowInvalidAuthorization}}
  authorization-header: {{.OAuth.AuthorizationHeader}}
  api-key-header: {{.OAuth.APIKeyHeader}}
  keepAuthHeader: {{.OAuth.KeepAuthHeader}}
  verify_api_key_url: {{.OAuth.VerifyAPIKeyURL}}
analytics:
  uri: {{.Analytics.URI}}
  bufferSize: {{.Analytics.BufferSize}}
  batchSize: {{.Analytics.BatchSize}}
  flushInterval: {{.Analytics.FlushInterval}}

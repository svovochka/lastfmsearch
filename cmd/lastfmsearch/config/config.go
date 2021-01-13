package config

type Config struct {
	HttpPort          string `long:"http-port" description:"Http port" env:"CL_HTTP_PORT" default:"80"`
	EnablePlayground  bool   `long:"enable-playground" description:"Enable playground component" env:"CL_ENABLE_PLAYGROUND"`
	LastfmApiEndpoint string `long:"lastfm-api-endpoint" description:"LastFm API endpoint" env:"CL_LASTFM_API_ENDPOINT" default:"http://ws.audioscrobbler.com/2.0"`
	LastfmApiKey      string `long:"lastfm-api-key" description:"LastFm API key" env:"CL_LASTFM_API_KEY" default:""`
}

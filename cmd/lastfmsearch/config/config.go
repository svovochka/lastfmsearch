package config

type Config struct {
	HttpPort          string `long:"http-port" description:"Http port" env:"CL_HTTP_PORT" default:"80"`
	LastFmApiEndpoint string `long:"last-fm-api-endpoint" description:"LastFm API endpoint" env:"CL_LASTFM_API_ENDPOINT" default:"https://ws.audioscrobbler.com/2.0"`
	LastFmApiPort     string `long:"last-fm-api-port" description:"LastFm API port" env:"CL_LASTFM_API_PORT" default:"433"`
	LastFmApiKey      string `long:"last-fm-api-key" description:"LastFm API key" env:"CL_LASTFM_API_KEY" default:""`
}

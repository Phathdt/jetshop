package common

import "github.com/namsral/flag"

var (
	defaultRedisMaxActive = 0 // 0 is unlimited max active connection
	defaultRedisMaxIdle   = 10
)
var (
	ConsulHost = ""
	RedisUri   = ""
	MaxActive  = 0
	MaxIde     = 0
)

func init() {
	flag.StringVar(&ConsulHost, "consul_host", "localhost:8500", "consult host, should be localhost:8500")
	flag.StringVar(&RedisUri, "redis-uri", "redis://localhost:6379", "(For go-redis) Redis connection-string. Ex: redis://localhost/0")
	flag.IntVar(&MaxActive, "redis-pool-max-active", defaultRedisMaxActive, "(For go-redis) Override redis pool MaxActive")
	flag.IntVar(&MaxIde, "redis-pool-max-idle", defaultRedisMaxIdle, "(For go-redis) Override redis pool MaxIdle")
	flag.Parse()
}

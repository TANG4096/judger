module judger

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4 // indirect
	github.com/antihax/optional v0.0.0-20180407024304-ca021399b1a6 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-kit/kit v0.9.0 // indirect
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.3.3
	github.com/google/btree v1.0.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/websocket v0.0.0-20170926233335-4201258b820c // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.7.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v0.9.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.4.1
	github.com/prometheus/procfs v0.2.0 // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/sta-golang/go-lib-utils v0.2.8
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/mod v0.1.1-0.20191107180719-034126e5016b // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a1e // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	golang.org/x/tools v0.0.0-20200207183749-b753a1ba74fa
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.26.0-rc.1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.8
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace google.golang.org/grpc v1.34.0 => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4

replace (
	google.golang.org/api => google.golang.org/api v0.14.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

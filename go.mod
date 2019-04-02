module github.com/sapk/rtme-browser

go 1.12

require (
	github.com/gin-contrib/gzip v0.0.1
	github.com/gin-gonic/gin v1.3.1-0.20190328061400-ce20f107f5dc
	github.com/mattn/go-colorable v0.1.1
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/rs/zerolog v1.13.0
	github.com/sapk/go-genesys v0.0.0-20190401001325-ea895b59ef21
	github.com/sapk-fork/webview v0.0.0-20190125144901-4293066e9dda
)

replace github.com/golang/lint => golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3

//Enfore some version to de-dup
replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.2

replace cloud.google.com/go => cloud.google.com/go v0.37.2

replace github.com/denisenkom/go-mssqldb => github.com/denisenkom/go-mssqldb v0.0.0-20190328043727-2183450503ad

replace honnef.co/go/tools => honnef.co/go/tools v0.0.0-20190106161140-3f1c8253044a

replace github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.3.1-0.20190328061400-ce20f107f5dc

replace github.com/gin-contrib/sse => github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3

replace github.com/go-xorm/builder => github.com/go-xorm/builder v0.3.4

replace github.com/go-xorm/core => github.com/go-xorm/core v0.6.2

replace github.com/go-xorm/xorm => github.com/go-xorm/xorm v0.7.2-0.20190331113943-617e0ae295d7

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.1

replace github.com/jackc/pgx => github.com/jackc/pgx v3.3.0+incompatible

replace github.com/mattn/go-isatty => github.com/mattn/go-isatty v0.0.6

replace github.com/openzipkin/zipkin-go => github.com/openzipkin/zipkin-go v0.1.6

replace github.com/pkg/errors => github.com/pkg/errors v0.8.1

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c

replace golang.org/x/lint => golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3

replace github.com/mattn/go-sqlite3 => github.com/mattn/go-sqlite3 v1.10.0

replace golang.org/x/net => golang.org/x/net v0.0.0-20190311183353-d8887717615a

replace golang.org/x/oauth2 => golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421

replace golang.org/x/sync => golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65

replace golang.org/x/tools => golang.org/x/tools v0.0.0-20190312170243-e65039ee4138

replace google.golang.org/api => google.golang.org/api v0.3.0

replace google.golang.org/appengine => google.golang.org/appengine v1.5.0

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19

replace google.golang.org/grpc => google.golang.org/grpc v1.19.0

replace gopkg.in/check.v1 => gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127

replace github.com/go-sql-driver/mysql => github.com/go-sql-driver/mysql v1.4.1

# Fabulous Go

This application parses github repos from [Fabulous Go](https://github.com/avelino/awesome-go) and populates the lists with additional data, like stars count and commit activity.

How it works and how to rebuild the list is described at `BUILDING.md`.

## Contents
* [Fabulous Go](#fabulous-go)
  * [Audio/Music](#audiomusic)
  * [Authentication & OAuth](#authentication--oauth)
  * [Command Line](#command-line)
    * [Standard CLI](#standard-cli)
    * [Advanced Console UIs](#advanced-console-uis)
  * [Configuration](#configuration)
  * [Continuous Integration](#continuous-integration)
  * [CSS Preprocessors](#css-preprocessors)
  * [Data Structures](#data-structures)
  * [Database](#database)
  * [Database Drivers](#database-drivers)
  * [Date & Time](#date--time)
  * [Distributed Systems](#distributed-systems)
  * [Email](#email)
  * [Embeddable Scripting Languages](#embeddable-scripting-languages)
  * [Financial](#financial)
  * [Forms](#forms)
  * [Game Development](#game-development)
  * [Generation & Generics](#generation--generics)
  * [Go Compilers](#go-compilers)
  * [Goroutines](#goroutines)
  * [GUI](#gui)
  * [Images](#images)
  * [Logging](#logging)
  * [Machine Learning](#machine-learning)
  * [Messaging](#messaging)
  * [Miscellaneous](#miscellaneous)
  * [Natural Language Processing](#natural-language-processing)
  * [Networking](#networking)
  * [OpenGL](#opengl)
  * [ORM](#orm)
  * [Package Management](#package-management)
  * [Query Language](#query-language)
  * [Resource Embedding](#resource-embedding)
  * [Science and Data Analysis](#science-and-data-analysis)
  * [Security](#security)
  * [Serialization](#serialization)
  * [Server Applications](#server-applications)
  * [Template Engines](#template-engines)
  * [Testing](#testing)
  * [Text Processing](#text-processing)
  * [Third-party APIs](#third-party-apis)
  * [Utilities](#utilities)
  * [Validation](#validation)
  * [Version Control](#version-control)
  * [Video](#video)
  * [Web Frameworks](#web-frameworks)
    * [Middlewares](#middlewares)
      * [Actual middlewares](#actual-middlewares)
      * [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
* [Tools](#tools)
  * [Code Analysis](#code-analysis)
  * [Editor Plugins](#editor-plugins)
  * [Go Tools](#go-tools)
  * [Software Packages](#software-packages)
    * [DevOps Tools](#devops-tools)
    * [Other Software](#other-software)
* [Resources](#resources)
  * [Benchmarks](#benchmarks)
  * [E-Books](#e-books)
  * [Websites](#websites)
    * [Tutorials](#tutorials)
  * [Windows](#windows)


# Fabulous Go
## Audio/Music
*Libraries for manipulating audio.*

  * [eaburns/flac](https://github.com/eaburns/flac) :star:43 - A Free Lossless Audio Codec decoder in Go *(**0** commits last month)*
  * [outrightmental/go-atomix](https://github.com/outrightmental/go-atomix) :star:2 - Sequence-based Go-native audio mixer for Music apps *(**0** commits last month)*
  * [krig/go-sox](https://github.com/krig/go-sox) :star:33 - libsox bindings for go *(**1** commits last month)*
  * [zhulik/go_mediainfo](https://github.com/zhulik/go_mediainfo) :star:1 - Golang bindings for libmediainfo *(**0** commits last month)*
  * [tcolgate/mp3](https://github.com/tcolgate/mp3) :star:6 - golang mp3 frame parser *(**5** commits last month)*
  * [gordonklaus/portaudio](https://github.com/gordonklaus/portaudio) :star:21 - Go bindings for the PortAudio audio I/O library *(**0** commits last month)*
  * [rakyll/portmidi](https://github.com/rakyll/portmidi) :star:73 - Go bindings for libportmidi *(**0** commits last month)*
  * [wtolson/go-taglib](https://github.com/wtolson/go-taglib) :star:38 - Go wrapper for taglib *(**0** commits last month)*
  * [mccoyst/vorbis](https://github.com/mccoyst/vorbis) :star:9 - A "native" ogg vorbis decoder for Go (uses inline stb_vorbis) *(**0** commits last month)*
  * [mdlayher/waveform](https://github.com/mdlayher/waveform) :star:107 - Go package capable of generating waveform images from audio streams. MIT Licensed. *(**0** commits last month)*

## Authentication & OAuth
*Libraries for implementing authentications schemes.*

  * [smartystreets/go-aws-auth](https://github.com/smartystreets/go-aws-auth) :star:93 - Signs requests to Amazon Web Services (AWS) using IAM roles or signed signature versions 2, 3, and 4. Supports S3 and STS. *(**0** commits last month)*
  * [square/go-jose](https://github.com/square/go-jose) :star:451 - An implementation of JOSE standards in Golang. *(**18** commits last month)*
  * [bradrydzewski/go.auth](https://github.com/bradrydzewski/go.auth) :star:288 - authentication API for Go web applications *(**0** commits last month)*
  * [dghubble/gologin](https://github.com/dghubble/gologin) :star:587 - Chainable Go login handlers for authentication providers (oauth1, oauth2) *(**0** commits last month)*
  * [mikespook/gorbac](https://github.com/mikespook/gorbac) :star:272 - goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. *(**0** commits last month)*
  * [markbates/goth](https://github.com/markbates/goth) :star:566 - Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. *(**13** commits last month)*
  * [goji/httpauth](https://github.com/goji/httpauth) :star:85 - HTTP Authentication middlewares *(**0** commits last month)*
  * [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) :star:906 - Golang implementation of JSON Web Tokens (JWT) *(**0** commits last month)*
  * [golang/oauth2](https://github.com/golang/oauth2) :star:628 - Go OAuth2 *(**18** commits last month)*
  * [RangelReale/osin](https://github.com/RangelReale/osin) :star:786 - Golang OAuth2 server library *(**8** commits last month)*
  * [xyproto/permissions2](https://github.com/xyproto/permissions2) :star:140 -   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions *(**3** commits last month)*
  * [GeertJohan/yubigo](https://github.com/GeertJohan/yubigo) :star:51 -   Yubigo is a Yubikey client API library that provides an easy way to integrate the Yubico Yubikey into your existing Go-based user authentication infrastructure. *(**0** commits last month)*

## Command Line
### Standard CLI
*Libraries for building standard or basic Command Line applications*

  * [tcnksm/gcli](https://github.com/tcnksm/gcli) :star:435 - The easy way to build Golang command-line application. *(**1** commits last month)*
  * [tucnak/climax](https://github.com/tucnak/climax) :star:90 - Climax is an alternative CLI with "human face" *(**0** commits last month)*
  * [spf13/cobra](https://github.com/spf13/cobra) :star:1690 - A Commander for modern Go CLI interactions *(**9** commits last month)*
  * [codegangsta/cli](https://github.com/codegangsta/cli) :star:3728 - A small package for building command line apps in Go *(**6** commits last month)*
  * [docopt/docopt.go](https://github.com/docopt/docopt.go) :star:634 - A command-line arguments parser that will make you smile. *(**0** commits last month)*
  * [jessevdk/go-flags](https://github.com/jessevdk/go-flags) :star:526 - go command line option parser *(**2** commits last month)*
  * [alecthomas/kingpin](https://github.com/alecthomas/kingpin) :star:483 - A Go (golang) command line and flag parser *(**11** commits last month)*
  * [peterh/liner](https://github.com/peterh/liner) :star:272 - Pure Go line editor with history, inspired by linenoise *(**16** commits last month)*
  * [mitchellh/cli](https://github.com/mitchellh/cli) :star:372 - A Go library for implementing command-line interfaces. *(**0** commits last month)*
  * [jawher/mow.cli](https://github.com/jawher/mow.cli) :star:320 - A versatile library for building CLI applications in Go *(**4** commits last month)*
  * [chzyer/readline](https://github.com/chzyer/readline) :star:435 - Readline is a pure go(golang) implementation for GNU-Readline kind library *(**0** commits last month)*
  * [ukautz/clif](https://github.com/ukautz/clif) :star:60 - Another CLI framework for Go. It works on my machine. *(**0** commits last month)*

### Advanced Console UIs
*Libraries for building Console Applications and Console User Interfaces*

  * [ttacon/chalk](https://github.com/ttacon/chalk) :star:145 - Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk *(**0** commits last month)*
  * [fatih/color](https://github.com/fatih/color) :star:692 - Color package for Go (golang) *(**0** commits last month)*
  * [TreyBastian/colourize](https://github.com/TreyBastian/colourize) :star:2 - An ANSI colour terminal package for Go *(**0** commits last month)*
  * [daviddengcn/go-colortext](https://github.com/daviddengcn/go-colortext) :star:138 - Change the color of console text. *(**0** commits last month)*
  * [jroimartin/gocui](https://github.com/jroimartin/gocui) :star:358 - Minimalist Go library aimed at creating Console User Interfaces. *(**32** commits last month)*
  * [nsf/termbox-go](https://github.com/nsf/termbox-go) :star:1292 - Pure Go termbox implementation *(**6** commits last month)*
  * [apcera/termtables](https://github.com/apcera/termtables) :star:42 - Go ASCII Table Generator, ported from the Ruby terminal-tables library *(**0** commits last month)*
  * [gizak/termui](https://github.com/gizak/termui) :star:4309 - Golang terminal dashboard *(**12** commits last month)*
  * [gosuri/uilive](https://github.com/gosuri/uilive) :star:407 - uilive is a go library for updating terminal output in realtime *(**0** commits last month)*
  * [gosuri/uiprogress](https://github.com/gosuri/uiprogress) :star:761 - A go library to render progress bars in terminal applications *(**0** commits last month)*
  * [gosuri/uitable](https://github.com/gosuri/uitable) :star:269 - A go library to improve readability in terminal apps using tabular data *(**0** commits last month)*

## Configuration
*Libraries for configuration parsing*

  * [olebedev/config](https://github.com/olebedev/config) :star:57 - JSON or YAML configuration wrapper with convenient access methods. *(**0** commits last month)*
  * [paked/configure](https://github.com/paked/configure) :star:6 - Configure is a Go package that gives you easy configuration of your project through redundancy *(**0** commits last month)*
  * [caarlos0/env](https://github.com/caarlos0/env) :star:77 - A KISS way to deal with environment variables in Go. *(**0** commits last month)*
  * [tomazk/envcfg](https://github.com/tomazk/envcfg) :star:75 - Un-marshaling environment variables to Go structs *(**0** commits last month)*
  * [ian-kent/envconf](https://github.com/ian-kent/envconf) :star:2 - Configure Go applications from the environment *(**0** commits last month)*
  * [vrischmann/envconfig](https://github.com/vrischmann/envconfig) :star:86 - Small library to read your configuration from environment variables *(**0** commits last month)*
  * [go-gcfg/gcfg](https://github.com/go-gcfg/gcfg) :star:23 - read INI-style configuration files into Go structs; supports user-defined types and subsections *(**0** commits last month)*
  * [ian-kent/gofigure](https://github.com/ian-kent/gofigure) :star:28 - Go configuration made easy! *(**0** commits last month)*
  * [go-ini/ini](https://github.com/go-ini/ini) :star:203 - Package ini provides INI file read and write functionality in Go. *(**3** commits last month)*
  * [FogCreek/mini](https://github.com/FogCreek/mini) :star:82 - A golang package for parsing ini-style configuration files *(**6** commits last month)*
  * [tucnak/store](https://github.com/tucnak/store) :star:36 - A dead simple configuration manager for Go applications *(**0** commits last month)*
  * [spf13/viper](https://github.com/spf13/viper) :star:1242 - Go configuration with fangs *(**5** commits last month)*

## Continuous Integration
*Tools for help with continuous integration*

  * [drone/drone](https://github.com/drone/drone) :star:6133 - Drone is a Continuous Delivery platform built on Docker, written in Go *(**73** commits last month)*
  * [mattn/goveralls](https://github.com/mattn/goveralls) :star:218 -  *(**2** commits last month)*
  * [go-playground/overalls](https://github.com/go-playground/overalls) :star:15 - :jeans:Multi-Package go project coverprofile for tools like goveralls *(**0** commits last month)*

## CSS Preprocessors
*Libraries for preprocessing CSS files*

  * [c9s/c6](https://github.com/c9s/c6) :star:324 - Compile SASS Faster ! C6 is a SASS-compatible compiler written in Go. *(**0** commits last month)*
  * [yosssi/gcss](https://github.com/yosssi/gcss) :star:333 - Pure Go CSS Preprocessor *(**0** commits last month)*
  * [wellington/go-libsass](https://github.com/wellington/go-libsass) :star:16 - Go wrapper for libsass, the only Sass 3.3 compiler for Go *(**0** commits last month)*

## Data Structures
*Generic datastructures and algorithms in Go.*

  * [willf/bitset](https://github.com/willf/bitset) :star:186 - Go package implementing bitsets *(**2** commits last month)*
  * [zhenjl/bloom](https://github.com/zhenjl/bloom) :star:75 - Bloom filters implemented in Go. *(**0** commits last month)*
  * [tylertreat/BoomFilters](https://github.com/tylertreat/BoomFilters) :star:660 - Probabilistic data structures for processing continuous, unbounded streams. *(**80** commits last month)*
  * [seiflotfy/cuckoofilter](https://github.com/seiflotfy/cuckoofilter) :star:135 - Cuckoo Filter: Practically Better Than Bloom *(**0** commits last month)*
  * [zhenjl/encoding](https://github.com/zhenjl/encoding) :star:40 - Integer Compression Libraries for Go *(**0** commits last month)*
  * [Workiva/go-datastructures](https://github.com/Workiva/go-datastructures) :star:2177 -  *(**174** commits last month)*
  * [hailocab/go-geoindex](https://github.com/hailocab/go-geoindex) :star:127 - Go native library for fast point tracking and K-Nearest queries *(**11** commits last month)*
  * [deckarep/golang-set](https://github.com/deckarep/golang-set) :star:333 - A simple set type for the Go language. *(**0** commits last month)*
  * [ryszard/goskiplist](https://github.com/ryszard/goskiplist) :star:97 - A skip list implementation in Go *(**0** commits last month)*
  * [smartystreets/mafsa](https://github.com/smartystreets/mafsa) :star:202 - Package mafsa implements Minimal Acyclic Finite State Automata in Go, essentially a high-speed, memory-efficient, Unicode-friendly set of strings. *(**0** commits last month)*
  * [gansidui/skiplist](https://github.com/gansidui/skiplist) :star:25 - skiplist for golang *(**0** commits last month)*
  * [derekparker/trie](https://github.com/derekparker/trie) :star:104 - Data structure and relevant algorithms for extremely fast prefix/fuzzy string searching. *(**5** commits last month)*
  * [diegobernardes/ttlcache](https://github.com/diegobernardes/ttlcache) :star:17 - An in-memory LRU string-interface{} map with expiration for golang *(**0** commits last month)*
  * [willf/bloom](https://github.com/willf/bloom) :star:218 - Go package implementing Bloom filters *(**6** commits last month)*

## Database
*Databases implemented in Go.*

  * [boltdb/bolt](https://github.com/boltdb/bolt) :star:3616 - An embedded key/value database for Go. *(**12** commits last month)*
  * [muesli/cache2go](https://github.com/muesli/cache2go) :star:55 - Simple go caching library with expiration capabilities and access counters *(**1** commits last month)*
  * [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach) :star:6109 - A Scalable, Survivable, Strongly-Consistent SQL Database *(**125** commits last month)*
  * [codingsince1985/couchcache](https://github.com/codingsince1985/couchcache) :star:14 - A RESTful caching micro-service in Go backed by Couchbase server *(**0** commits last month)*
  * [peterbourgon/diskv](https://github.com/peterbourgon/diskv) :star:327 - A disk-backed key-value store. *(**0** commits last month)*
  * [couchbase/goforestdb](https://github.com/couchbase/goforestdb) :star:18 - Go bindings for ForestDB *(**4** commits last month)*
  * [bluele/gcache](https://github.com/bluele/gcache) :star:89 - Cache library for golang. It supports expirable Cache, LFU, LRU and ARC. *(**2** commits last month)*
  * [pmylund/go-cache](https://github.com/pmylund/go-cache) :star:409 - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. *(**0** commits last month)*
  * [syndtr/goleveldb](https://github.com/syndtr/goleveldb) :star:929 - LevelDB key/value database in Go. *(**0** commits last month)*
  * [golang/groupcache](https://github.com/golang/groupcache) :star:4166 - groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. *(**0** commits last month)*
  * [influxdb/influxdb](https://github.com/influxdb/influxdb) :star:7275 - Scalable datastore for metrics, events, and real-time analytics *(**638** commits last month)*
  * [siddontang/ledisdb](https://github.com/siddontang/ledisdb) :star:1470 - a high performance NoSQL powered by Go *(**25** commits last month)*
  * [jmhodges/levigo](https://github.com/jmhodges/levigo) :star:260 - levigo is a Go wrapper for LevelDB *(**1** commits last month)*
  * [prometheus/prometheus](https://github.com/prometheus/prometheus) :star:3806 - The Prometheus monitoring system and time series database. *(**101** commits last month)*
  * [pingcap/tidb](https://github.com/pingcap/tidb) :star:3359 - TiDB is a distributed NewSQL database compatible with MySQL protocol  *(**0** commits last month)*
  * [HouzuoGuo/tiedot](https://github.com/HouzuoGuo/tiedot) :star:1446 - Your NoSQL database powered by Golang *(**5** commits last month)*

*Database tools.*

  * [siddontang/go-mysql](https://github.com/siddontang/go-mysql) :star:161 - a powerful mysql toolset with Go *(**2** commits last month)*
  * [siddontang/go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) :star:200 - Sync MySQL data into elasticsearch  *(**0** commits last month)*
  * [flike/kingshard](https://github.com/flike/kingshard) :star:1467 - A high-performance MySQL proxy *(**0** commits last month)*
  * [mattes/migrate](https://github.com/mattes/migrate) :star:471 - Database migration handling in Golang *(**7** commits last month)*
  * [2tvenom/myreplication](https://github.com/2tvenom/myreplication) :star:45 - Golang MySql binary log replication listener *(**7** commits last month)*
  * [outbrain/orchestrator](https://github.com/outbrain/orchestrator) :star:345 - MySQL replication topology manager/visualizer *(**44** commits last month)*
  * [sosedoff/pgweb](https://github.com/sosedoff/pgweb) :star:3195 - Cross-platform client for PostgreSQL databases *(**10** commits last month)*
  * [pravasan/pravasan](https://github.com/pravasan/pravasan) :star:10 - Simple Migration Tool - written in Go *(**13** commits last month)*
  * [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate) :star:384 - SQL schema migration tool for Go. *(**2** commits last month)*
  * [youtube/vitess](https://github.com/youtube/vitess) :star:3201 - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. *(**190** commits last month)*

*SQL query builder, libraries for building and using SQL.*

  * [mgutz/dat](https://github.com/mgutz/dat) :star:252 - Go Postgres Data Access Toolkit *(**55** commits last month)*
  * [gchaincl/dotsql](https://github.com/gchaincl/dotsql) :star:212 - A Golang library for using SQL. *(**2** commits last month)*
  * [doug-martin/goqu](https://github.com/doug-martin/goqu) :star:237 - SQL builder and query library for golang *(**6** commits last month)*
  * [go-ozzo/ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) :star:53 - A Go package that enhances the standard database/sql package by providing powerful data retrieval methods as well as DB-agnostic query building capabilities. *(**0** commits last month)*
  * [variadico/scaneo](https://github.com/variadico/scaneo) :star:53 - Generate Go code to convert database rows into arbitrary structs. *(**0** commits last month)*
  * [elgris/sqrl](https://github.com/elgris/sqrl) :star:24 - Fluent SQL generation for golang *(**10** commits last month)*
  * [Masterminds/squirrel](https://github.com/Masterminds/squirrel) :star:732 - Fluent SQL generation for golang *(**10** commits last month)*

## Database Drivers
*Libraries for connecting and operating databases.*
* Relational Databases

  * [nakagami/firebirdsql](https://github.com/nakagami/firebirdsql) :star:40 - Firebird RDBMS sql driver for Go (golang) *(**12** commits last month)*
  * [mattn/go-adodb](https://github.com/mattn/go-adodb) :star:30 - Microsoft ActiveX Object DataBase driver for go that using exp/sql *(**0** commits last month)*
  * [rounds/go-bqstreamer](https://github.com/rounds/go-bqstreamer) :star:72 - Stream data into Google BigQuery concurrently using InsertAll() *(**0** commits last month)*
  * [denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb) :star:284 - Microsoft SQL server driver written in go language *(**1** commits last month)*
  * [mattn/go-oci8](https://github.com/mattn/go-oci8) :star:115 - oracle driver for go that using database/sql *(**16** commits last month)*
  * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) :star:2035 - Go-MySQL-Driver is a lightweight and fast MySQL-Driver for Go's (golang) database/sql package *(**51** commits last month)*
  * [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) :star:1068 - sqlite3 driver for go that using database/sql *(**1** commits last month)*
  * [minus5/gofreetds](https://github.com/minus5/gofreetds) :star:33 - Go Sql Server database driver. *(**0** commits last month)*
  * [jackc/pgx](https://github.com/jackc/pgx) :star:481 - PostgreSQL client library for Go *(**4** commits last month)*
  * [lib/pq](https://github.com/lib/pq) :star:1880 - Pure Go Postgres driver for database/sql *(**8** commits last month)*

* NoSQL Databases

  * [aerospike/aerospike-client-go](https://github.com/aerospike/aerospike-client-go) :star:151 - Aerospike Client Go  *(**8** commits last month)*
  * [solher/arangolite](https://github.com/solher/arangolite) :star:15 - Lightweight golang driver for ArangoDB *(**0** commits last month)*
  * [google/cayley](https://github.com/google/cayley) :star:6849 - An open-source graph database *(**38** commits last month)*
  * [underarmour/dynago](https://github.com/underarmour/dynago) :star:11 - A DynamoDB client for Go *(**0** commits last month)*
  * [couchbase/go-couchbase](https://github.com/couchbase/go-couchbase) :star:198 - Couchbase client in Go *(**8** commits last month)*
  * [fjl/go-couchdb](https://github.com/fjl/go-couchdb) :star:29 - Yet another CouchDB HTTP API wrapper for Go *(**0** commits last month)*
  * [couchbase/gocb](https://github.com/couchbase/gocb) :star:152 - The Couchbase Go SDK *(**16** commits last month)*
  * [dancannon/gorethink](https://github.com/dancannon/gorethink) :star:748 - Go language driver for RethinkDB *(**46** commits last month)*
  * [cihangir/neo4j](https://github.com/cihangir/neo4j) :star:12 - Neo4j Rest API Client for Go lang *(**0** commits last month)*
  * [davemeehan/Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) :star:60 - Neo4j REST Client in golang *(**0** commits last month)*
  * [jmcvetta/neoism](https://github.com/jmcvetta/neoism) :star:223 - Neo4j client for Golang *(**2** commits last month)*
  * [garyburd/redigo](https://github.com/garyburd/redigo) :star:1922 - Go client for Redis *(**0** commits last month)*
  * [go-redis/redis](https://github.com/go-redis/redis) :star:642 - Redis client for Golang. *(**12** commits last month)*
  * [hoisie/redis](https://github.com/hoisie/redis) :star:487 - A simple, powerful Redis client for Go *(**0** commits last month)*
  * [bsm/redeo](https://github.com/bsm/redeo) :star:32 - High-performance framework for building redis-protocol compatible TCP servers/services *(**7** commits last month)*

* Search and Analytic Databases

  * [blevesearch/bleve](https://github.com/blevesearch/bleve) :star:2295 - A modern text indexing library for go *(**19** commits last month)*
  * [olivere/elastic](https://github.com/olivere/elastic) :star:497 - Elasticsearch client for Go. *(**10** commits last month)*
  * [mattbaird/elastigo](https://github.com/mattbaird/elastigo) :star:681 - A Go (golang) based Elasticsearch client library. *(**7** commits last month)*
  * [belogik/goes](https://github.com/belogik/goes) :star:72 - A library to interact with Elasticsearch in Go! *(**5** commits last month)*
  * [seiflotfy/skizze](https://github.com/seiflotfy/skizze) :star:7 - A probabilistic data structure service and storage *(**0** commits last month)*

## Date & Time
*Libraries for working with dates and times.*

  * [grsmv/goweek](https://github.com/grsmv/goweek) :star:3 - ISO 8601 compatible library for working with week entities for Go *(**0** commits last month)*
  * [jinzhu/now](https://github.com/jinzhu/now) :star:544 - Now is a time toolkit for golang *(**2** commits last month)*
  * [leekchan/timeutil](https://github.com/leekchan/timeutil) :star:96 - timeutil - useful extensions (Timedelta, Strftime, ...) to the golang's time package *(**0** commits last month)*

## Distributed Systems
*Packages that help with building Distributed Systems.*

  * [svcavallar/celeriac.v1](https://github.com/svcavallar/celeriac.v1) :star:7 - Golang client library for adding support for interacting and monitoring Celery workers, tasks and events. *(**0** commits last month)*
  * [vectaport/flowgraph](https://github.com/vectaport/flowgraph) :star:6 - Ready-send coordination layer on top of goroutines. *(**0** commits last month)*
  * [chrislusf/glow](https://github.com/chrislusf/glow) :star:596 - Glow is an easy-to-use distributed computation system written in Go, similar to Hadoop Map Reduce, Spark, Flink, Samza, etc. Currently just started and not feature rich yet, but should be reliable to run most common cases. *(**0** commits last month)*
  * [dgryski/go-jump](https://github.com/dgryski/go-jump) :star:53 - go-jump: Jump consistent hashing *(**0** commits last month)*
  * [valyala/gorpc](https://github.com/valyala/gorpc) :star:185 - Simple, fast and scalable golang rpc library for high load *(**10** commits last month)*
  * [grpc/grpc-go](https://github.com/grpc/grpc-go) :star:1429 - The Go language implementation of gRPC. HTTP/2 based RPC *(**149** commits last month)*
  * [micro/micro](https://github.com/micro/micro) :star:636 - A microservice toolkit *(**2** commits last month)*
  * [hashicorp/raft](https://github.com/hashicorp/raft) :star:451 - Golang implementation of the Raft consensus protocol *(**4** commits last month)*
  * [anacrolix/torrent](https://github.com/anacrolix/torrent) :star:1029 - Full-featured BitTorrent client package and utilities *(**46** commits last month)*
  * [Sioro-Neoku/go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) :star:117 - Go Peerflix *(**0** commits last month)*

## Email
*Libraries that implement email creation and sending*

  * [aymerick/douceur](https://github.com/aymerick/douceur) :star:34 - A simple CSS parser and inliner in Go *(**0** commits last month)*
  * [jordan-wright/email](https://github.com/jordan-wright/email) :star:643 - Robust and flexible email library for Go *(**0** commits last month)*
  * [toorop/go-dkim](https://github.com/toorop/go-dkim) :star:12 - DKIM package for golang *(**0** commits last month)*
  * [hectane/hectane](https://github.com/hectane/hectane) :star:34 - Lightweight SMTP client written in Go *(**0** commits last month)*
  * [mailhog/MailHog](https://github.com/mailhog/MailHog) :star:591 - Web and API based SMTP testing *(**6** commits last month)*
  * [sendgrid/sendgrid-go](https://github.com/sendgrid/sendgrid-go) :star:153 - SendGrid Library to Interface through Go *(**2** commits last month)*
  * [mailhog/smtp](https://github.com/mailhog/smtp) :star:18 - MailHog SMTP Protocol *(**0** commits last month)*

## Embeddable Scripting Languages
*Embedding other languages inside your go code*

  * [PuerkitoBio/agora](https://github.com/PuerkitoBio/agora) :star:203 - a dynamically typed, garbage collected, embeddable programming language built with Go *(**0** commits last month)*
  * [mattn/anko](https://github.com/mattn/anko) :star:232 - Scriptable interpreter written in golang *(**3** commits last month)*
  * [jcla1/gisp](https://github.com/jcla1/gisp) :star:313 - Simple LISP in Go *(**0** commits last month)*
  * [olebedev/go-duktape](https://github.com/olebedev/go-duktape) :star:264 - Duktape JavaScript engine bindings for Go *(**10** commits last month)*
  * [Shopify/go-lua](https://github.com/Shopify/go-lua) :star:570 - A Lua VM in Go *(**12** commits last month)*
  * [deuill/go-php](https://github.com/deuill/go-php) :star:89 - PHP bindings for the Go programming language (Golang) *(**0** commits last month)*
  * [sbinet/go-python](https://github.com/sbinet/go-python) :star:336 - naive go bindings to the CPython C-API *(**6** commits last month)*
  * [aarzilli/golua](https://github.com/aarzilli/golua) :star:254 - Go bindings for Lua C API - in progress *(**0** commits last month)*
  * [yuin/gopher-lua](https://github.com/yuin/gopher-lua) :star:1073 - GopherLua: VM and compiler for Lua in Go *(**26** commits last month)*
  * [robertkrimen/otto](https://github.com/robertkrimen/otto) :star:1933 - A JavaScript interpreter in Go (golang) *(**15** commits last month)*
  * [ian-kent/purl](https://github.com/ian-kent/purl) :star:8 - Perl, but fluffy like a cat! *(**0** commits last month)*

## Financial
*Packages for accounting and finance*

  * [leekchan/accounting](https://github.com/leekchan/accounting) :star:183 - money and currency formatting for golang *(**0** commits last month)*
  * [shopspring/decimal](https://github.com/shopspring/decimal) :star:222 - Arbitrary-precision fixed-point decimal numbers in go *(**6** commits last month)*

## Forms
*Libraries for working with forms.*

  * [robfig/bind](https://github.com/robfig/bind) :star:9 -  *(**0** commits last month)*
  * [mholt/binding](https://github.com/mholt/binding) :star:424 - Reflectionless data binding for Go's net/http (not yet a stable 1.0, but not likely to change much either) *(**0** commits last month)*
  * [monoculum/formam](https://github.com/monoculum/formam) :star:42 - a package for decode form's values into struct in Go *(**1** commits last month)*
  * [albrow/forms](https://github.com/albrow/forms) :star:51 - A lightweight go library for parsing form data or json from an http.Request. *(**5** commits last month)*
  * [gorilla/csrf](https://github.com/gorilla/csrf) :star:90 - gorilla/csrf provides Cross Site Request Forgery (CSRF) prevention middleware for Go web applications & services. *(**0** commits last month)*
  * [justinas/nosurf](https://github.com/justinas/nosurf) :star:612 - CSRF protection middleware for Go. *(**0** commits last month)*

## Game Development
*Awesome game development libraries.*

  * [paked/engi](https://github.com/paked/engi) :star:74 - A cross-platform game engine written in Go following my interpretation of the Entity Component System paradigm, A fork of ajhager/engi *(**1** commits last month)*
  * [vova616/GarageEngine](https://github.com/vova616/GarageEngine) :star:206 - Game engine written in Go (golang). *(**0** commits last month)*
  * [runningwild/glop](https://github.com/runningwild/glop) :star:71 - Bare-bones osx alternative to sdl *(**0** commits last month)*
  * [beefsack/go-astar](https://github.com/beefsack/go-astar) :star:135 - Go implementation of the A* search algorithm *(**0** commits last month)*
  * [GlenKelley/go-collada](https://github.com/GlenKelley/go-collada) :star:8 - Go package for working with the Collada file format. *(**0** commits last month)*
  * [veandco/go-sdl2](https://github.com/veandco/go-sdl2) :star:299 - SDL2 binding for Go *(**56** commits last month)*
  * [ungerik/go3d](https://github.com/ungerik/go3d) :star:92 - A performance oriented 2D/3D math package for Go *(**5** commits last month)*
  * [xtaci/gonet](https://github.com/xtaci/gonet) :star:606 - a game server skeleton in golang *(**0** commits last month)*
  * [name5566/leaf](https://github.com/name5566/leaf) :star:529 - A game server framework in Go (golang) *(**23** commits last month)*
  * [JoelOtter/termloop](https://github.com/JoelOtter/termloop) :star:573 - Terminal-based game engine for Go, built on top of Termbox *(**0** commits last month)*

## Generation & Generics
*Tools to enhance the language with features like generics via code generation*

  * [clipperhouse/gen](https://github.com/clipperhouse/gen) :star:635 - Type-driven code generation for Go *(**2** commits last month)*
  * [ahmetalpbalkan/go-linq](https://github.com/ahmetalpbalkan/go-linq) :star:700 - .NET LINQ-like query methods for Go *(**0** commits last month)*
  * [rjeczalik/interfaces](https://github.com/rjeczalik/interfaces) :star:60 - Code generation tools for Go. *(**0** commits last month)*
  * [ungerik/pkgreflect](https://github.com/ungerik/pkgreflect) :star:29 - A Go preprocessor for package scoped reflection *(**0** commits last month)*

## Go Compilers
*Tools for compiling Go to other languages*

  * [gopherjs/gopherjs](https://github.com/gopherjs/gopherjs) :star:3442 - A compiler from Go to JavaScript for running Go code in a browser *(**78** commits last month)*
  * [go-llvm/llgo](https://github.com/go-llvm/llgo) :star:717 - LLVM-based compiler for Go *(**0** commits last month)*
  * [tardisgo/tardisgo](https://github.com/tardisgo/tardisgo) :star:301 - Golang->Haxe->CPP/CSharp/Java/JavaScript transpiler   *(**28** commits last month)*

## Goroutines
*Tools for managing and working with Goroutines*

  * [ivpusic/grpool](https://github.com/ivpusic/grpool) :star:85 - Lightweight Goroutine pool *(**0** commits last month)*
  * [go-playground/pool](https://github.com/go-playground/pool) :star:34 - :speedboat: Go consumer goroutine pool for easy goroutine handling + time saving *(**0** commits last month)*
  * [Jeffail/tunny](https://github.com/Jeffail/tunny) :star:265 - A goroutine pool for golang *(**0** commits last month)*

## GUI
*Libraries for building GUI Applications*

  * [go-qml/qml](https://github.com/go-qml/qml) :star:1595 - QML support for the Go language *(**2** commits last month)*
  * [visualfc/goqt](https://github.com/visualfc/goqt) :star:847 - Golang bindings to the Qt cross-platform application framework. *(**0** commits last month)*
  * [deckarep/gosx-notifier](https://github.com/deckarep/gosx-notifier) :star:265 - gosx-notifier is a Go framework for sending desktop notifications to OSX 10.8 or higher *(**0** commits last month)*
  * [gotk3/gotk3](https://github.com/gotk3/gotk3) :star:46 - Go bindings for GTK3 *(**0** commits last month)*
  * [oskca/sciter](https://github.com/oskca/sciter) :star:116 - Golang bindings of Sciter: the Embeddable HTML/CSS/script engine for modern UI development *(**0** commits last month)*
  * [getlantern/systray](https://github.com/getlantern/systray) :star:40 - a cross platfrom Go library to place an icon and menu in the notification area *(**69** commits last month)*
  * [shurcooL/trayhost](https://github.com/shurcooL/trayhost) :star:28 - Cross-platform Go library to place an icon in the host operating system's taskbar. *(**0** commits last month)*
  * [andlabs/ui](https://github.com/andlabs/ui) :star:2631 - Platform-native GUI library for Go. *(**84** commits last month)*
  * [lxn/walk](https://github.com/lxn/walk) :star:1191 - A Windows GUI toolkit for the Go Programming Language *(**0** commits last month)*

## Images
*Libraries for manipulating images.*

  * [h2non/bimg](https://github.com/h2non/bimg) :star:137 - Small Go package for fast high-level image processing using libvips via C bindings *(**0** commits last month)*
  * [pravj/geopattern](https://github.com/pravj/geopattern) :star:832 - Create beautiful generative image patterns from a string in golang. *(**0** commits last month)*
  * [disintegration/gift](https://github.com/disintegration/gift) :star:684 - Go Image Filtering Toolkit *(**5** commits last month)*
  * [ungerik/go-cairo](https://github.com/ungerik/go-cairo) :star:48 - Go binding for the cairo graphics library *(**2** commits last month)*
  * [bolknote/go-gd](https://github.com/bolknote/go-gd) :star:38 - Go bingings for GD (http://www.boutell.com/gd/) *(**0** commits last month)*
  * [koyachi/go-nude](https://github.com/koyachi/go-nude) :star:164 - Nudity detection with Go. *(**0** commits last month)*
  * [lazywei/go-opencv](https://github.com/lazywei/go-opencv) :star:337 - Go bindings for OpenCV / 2.x API in gocv / 1.x API in opencv *(**39** commits last month)*
  * [jyotiska/go-webcolors](https://github.com/jyotiska/go-webcolors) :star:17 - Port of webcolors library from Python to Go *(**0** commits last month)*
  * [gographics/imagick](https://github.com/gographics/imagick) :star:377 - naive Go binding to ImageMagick's MagickWand C API *(**0** commits last month)*
  * [h2non/imaginary](https://github.com/h2non/imaginary) :star:655 - Fast HTTP microservice for high-level image processing. Perfectly fitted for Docker and Heroku. *(**0** commits last month)*
  * [disintegration/imaging](https://github.com/disintegration/imaging) :star:741 - Simple Go image processing package *(**2** commits last month)*
  * [hawx/img](https://github.com/hawx/img) :star:72 - A selection of image manipulation tools *(**1** commits last month)*
  * [donatj/mpo](https://github.com/donatj/mpo) :star:4 - Simple Go-lang JPEG MPO (Multi Picture Object Decoder) *(**0** commits last month)*
  * [thoas/picfit](https://github.com/thoas/picfit) :star:423 - An image resizing server written in Go *(**11** commits last month)*
  * [nfnt/resize](https://github.com/nfnt/resize) :star:969 - Pure golang image resizing  *(**0** commits last month)*
  * [bamiaux/rez](https://github.com/bamiaux/rez) :star:115 - Image resizing in pure Go and SIMD *(**1** commits last month)*
  * [muesli/smartcrop](https://github.com/muesli/smartcrop) :star:174 - smartcrop implementation in Go *(**1** commits last month)*
  * [ajstarks/svgo](https://github.com/ajstarks/svgo) :star:642 - Go Language Library for SVG generation *(**0** commits last month)*
  * [ftrvxmtrx/tga](https://github.com/ftrvxmtrx/tga) :star:12 - Go package for decoding and encoding TARGA image format *(**0** commits last month)*

## Logging
*Libraries for generating and working with log files.*

  * [golang/glog](https://github.com/golang/glog) :star:907 - Leveled execution logs for Go *(**0** commits last month)*
  * [siddontang/go-log](https://github.com/siddontang/go-log) :star:13 - a golang log lib supports level and multi handlers *(**0** commits last month)*
  * [ian-kent/go-log](https://github.com/ian-kent/go-log) :star:16 - A logger, for Go *(**0** commits last month)*
  * [ventu-io/go-log-interface](https://github.com/ventu-io/go-log-interface) :star:2 - Leveled log interface and the matching facade for the Go stdlib log.Logger *(**0** commits last month)*
  * [apsdehal/go-logger](https://github.com/apsdehal/go-logger) :star:149 - Simple logger for Go programs *(**0** commits last month)*
  * [sadlil/gologger](https://github.com/sadlil/gologger) :star:16 - Simple Logger for golang. Logs Into console, file or ElasticSearch. Simple, easy to use.  *(**0** commits last month)*
  * [apex/log](https://github.com/apex/log) :star:162 - Structured logging package for Go. *(**0** commits last month)*
  * [firstrow/logvoyage](https://github.com/firstrow/logvoyage) :star:47 - LogVoyage - logging SaaS written in GoLang *(**1** commits last month)*
  * [inconshreveable/log15](https://github.com/inconshreveable/log15) :star:383 - Powerful, composable logging for Go *(**1** commits last month)*
  * [chzyer/logex](https://github.com/chzyer/logex) :star:22 - An golang log lib, supports tracking and level, wrap by standard log lib *(**0** commits last month)*
  * [azer/logger](https://github.com/azer/logger) :star:55 - Minimalistic logging library for Go. *(**0** commits last month)*
  * [Sirupsen/logrus](https://github.com/Sirupsen/logrus) :star:2253 - Structured, pluggable logging for Go. *(**20** commits last month)*
  * [sebest/logrusly](https://github.com/sebest/logrusly) :star:6 - Loggly Hooks for GO Logrus logger *(**0** commits last month)*
  * [hashicorp/logutils](https://github.com/hashicorp/logutils) :star:123 - Utilities for slightly better logging in Go (Golang). *(**0** commits last month)*
  * [mgutz/logxi](https://github.com/mgutz/logxi) :star:224 - A 12-factor app logger built for performance and happy development *(**0** commits last month)*
  * [natefinch/lumberjack](https://github.com/natefinch/lumberjack) :star:257 - lumberjack is a rolling logger for Go *(**0** commits last month)*
  * [jbrodriguez/mlog](https://github.com/jbrodriguez/mlog) :star:2 -  *(**0** commits last month)*
  * [cihub/seelog](https://github.com/cihub/seelog) :star:591 - Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting. *(**3** commits last month)*
  * [alexcesaro/log](https://github.com/alexcesaro/log) :star:27 - Logging packages for Go *(**0** commits last month)*
  * [hpcloud/tail](https://github.com/hpcloud/tail) :star:416 - Go package for reading from continously updated files (tail -f) *(**0** commits last month)*
  * [rs/xlog](https://github.com/rs/xlog) :star:55 - xlog is a logger for net/context aware HTTP applications *(**0** commits last month)*

## Machine Learning
*Libraries for Machine Learning.*

  * [jbrukh/bayesian](https://github.com/jbrukh/bayesian) :star:282 - Naive Bayesian Classification for Golang. *(**0** commits last month)*
  * [ryanbressler/CloudForest](https://github.com/ryanbressler/CloudForest) :star:321 - Ensembles of decision trees in go/golang. *(**0** commits last month)*
  * [white-pony/go-fann](https://github.com/white-pony/go-fann) :star:68 - Go bindings for FANN, library for artificial neural networks *(**1** commits last month)*
  * [thoj/go-galib](https://github.com/thoj/go-galib) :star:106 - Genetic Algorithms library written in Go / golang *(**0** commits last month)*
  * [daviddengcn/go-pr](https://github.com/daviddengcn/go-pr) :star:33 - Pattern recognition package in Go lang. *(**0** commits last month)*
  * [goml/gobrain](https://github.com/goml/gobrain) :star:68 - Neural Networks written in go *(**0** commits last month)*
  * [e-dard/godist](https://github.com/e-dard/godist) :star:6 - Probability distributions and associated methods in Go *(**0** commits last month)*
  * [tomcraven/goga](https://github.com/tomcraven/goga) :star:21 - Golang Genetic Algorithm *(**0** commits last month)*
  * [sjwhitworth/golearn](https://github.com/sjwhitworth/golearn) :star:2515 - Machine Learning for Go *(**0** commits last month)*
  * [cdipaolo/goml](https://github.com/cdipaolo/goml) :star:432 - On-line Machine Learning in Go (and so much more) *(**0** commits last month)*
  * [timkaye11/goRecommend](https://github.com/timkaye11/goRecommend) :star:31 - Collaborative Filtering (CF) Algorithms in Go!  *(**0** commits last month)*
  * [datastream/libsvm](https://github.com/datastream/libsvm) :star:33 - libsvm go version *(**0** commits last month)*
  * [schuyler/neural-go](https://github.com/schuyler/neural-go) :star:40 - A multilayer perceptron network implemented in Go, with training via backpropagation. *(**0** commits last month)*
  * [muesli/regommend](https://github.com/muesli/regommend) :star:86 - Recommendation engine for Go *(**0** commits last month)*
  * [eaigner/shield](https://github.com/eaigner/shield) :star:74 - Bayesian text classifier with flexible tokenizers and storage backends for Go *(**0** commits last month)*

## Messaging
*Libraries that implement messaging systems*

  * [godbus/dbus](https://github.com/godbus/dbus) :star:67 - Native Go bindings for D-Bus *(**0** commits last month)*
  * [olebedev/emitter](https://github.com/olebedev/emitter) :star:26 - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins *(**0** commits last month)*
  * [asaskevich/EventBus](https://github.com/asaskevich/EventBus) :star:121 - [Go] Lightweight eventbus with async compatibility for Go *(**0** commits last month)*
  * [ventu-io/go-longpoll](https://github.com/ventu-io/go-longpoll) :star:3 - PubSub queuing with long-polling subscribers for web and other distributed applications *(**0** commits last month)*
  * [TheCreeper/go-notify](https://github.com/TheCreeper/go-notify) :star:12 - Package notify provides an implementation of the Gnome DBus Notification Specification. *(**0** commits last month)*
  * [nsqio/go-nsq](https://github.com/nsqio/go-nsq) :star:525 - The official Go package for NSQ *(**27** commits last month)*
  * [Terry-Mao/gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) :star:1102 - Golang push server cluster *(**0** commits last month)*
  * [RichardKnop/machinery](https://github.com/RichardKnop/machinery) :star:691 - Machinery is an asynchronous task queue/job queue based on distributed message passing. *(**0** commits last month)*
  * [nats-io/nats](https://github.com/nats-io/nats) :star:572 - Golang client for NATS, the cloud native messaging system. *(**5** commits last month)*
  * [dailymotion/oplog](https://github.com/dailymotion/oplog) :star:58 - A generic oplog/replication system for microservices *(**24** commits last month)*
  * [tuxychandru/pubsub](https://github.com/tuxychandru/pubsub) :star:75 - A simple pubsub package for go. *(**0** commits last month)*
  * [Shopify/sarama](https://github.com/Shopify/sarama) :star:817 - Sarama is a Go library for Apache Kafka 0.8 and 0.9  *(**112** commits last month)*
  * [uniqush/uniqush-push](https://github.com/uniqush/uniqush-push) :star:767 - Uniqush is a free and open source software which provides a unified push service for server-side notification to apps on mobile devices. *(**0** commits last month)*
  * [pebbe/zmq4](https://github.com/pebbe/zmq4) :star:350 - A Go interface to ZeroMQ version 4 *(**0** commits last month)*

## Miscellaneous
*These libraries were placed here because none of the other categories seemed to fit*

  * [artyom/autoflags](https://github.com/artyom/autoflags) :star:12 - Populate go command line app flags from config struct *(**0** commits last month)*
  * [digitalcrab/browscap_go](https://github.com/digitalcrab/browscap_go) :star:12 - GoLang Library for Browser Capabilities Project *(**0** commits last month)*
  * [miolini/datacounter](https://github.com/miolini/datacounter) :star:4 - Golang counters for readers/writers *(**0** commits last month)*
  * [go-chat-bot/bot](https://github.com/go-chat-bot/bot) :star:23 - IRC, Slack and Telegram bot written in go *(**0** commits last month)*
  * [hashicorp/go-multierror](https://github.com/hashicorp/go-multierror) :star:178 - A Go (golang) package for representing a list of errors as a single error. *(**0** commits last month)*
  * [ventu-io/go-shortid](https://github.com/ventu-io/go-shortid) :star:32 - Super short, unique, non sequential, URL friendly Ids for Go *(**0** commits last month)*
  * [shirou/gopsutil](https://github.com/shirou/gopsutil) :star:902 - psutil for golang *(**26** commits last month)*
  * [haxpax/gosms](https://github.com/haxpax/gosms) :star:897 - :mailbox_closed: Your own local SMS gateway in Go *(**3** commits last month)*
  * [albrow/jobs](https://github.com/albrow/jobs) :star:306 - A persistent and flexible background jobs library for go. *(**107** commits last month)*
  * [zhulik/margelet](https://github.com/zhulik/margelet) :star:25 - Telegram Bot Framework for Go *(**0** commits last month)*
  * [rjeczalik/notify](https://github.com/rjeczalik/notify) :star:110 - File system event notification library on steroids. *(**82** commits last month)*
  * [go-playground/stats](https://github.com/go-playground/stats) :star:10 - :chart_with_upwards_trend: Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... *(**0** commits last month)*
  * [go-xkg/xkg](https://github.com/go-xkg/xkg) :star:9 - X Keyboard Grabber *(**0** commits last month)*
  * [huandu/xstrings](https://github.com/huandu/xstrings) :star:299 - xstrings: A collection of useful string functions for Go. *(**1** commits last month)*

## Natural Language Processing
*Libraries for working with human languages.*

  * [nuance/go-nlp](https://github.com/nuance/go-nlp) :star:41 - Utilities for working with discrete probability distributions and other tools useful for doing NLP work *(**0** commits last month)*
  * [agonopol/go-stem](https://github.com/agonopol/go-stem) :star:33 - Word Stemming in Go *(**0** commits last month)*
  * [danieldk/go2vec](https://github.com/danieldk/go2vec) :star:5 - Read and use word2vec vectors in Go *(**1** commits last month)*
  * [rjohnsondev/golibstemmer](https://github.com/rjohnsondev/golibstemmer) :star:10 - Go bindings for the snowball libstemmer library including porter 2 *(**0** commits last month)*
  * [fiam/gounidecode](https://github.com/fiam/gounidecode) :star:35 - Unicode transliterator for #golang *(**0** commits last month)*
  * [goodsign/icu](https://github.com/goodsign/icu) :star:15 - Cgo binding for icu4c library *(**0** commits last month)*
  * [goodsign/libtextcat](https://github.com/goodsign/libtextcat) :star:8 - Cgo binding for libtextcat C library *(**0** commits last month)*
  * [awsong/MMSEGO](https://github.com/awsong/MMSEGO) :star:44 - Chinese word splitting algorithm MMSEG in GO *(**0** commits last month)*
  * [rookii/paicehusk](https://github.com/rookii/paicehusk) :star:10 - Golang implementation of the Paice/Husk Stemming Algorithm *(**0** commits last month)*
  * [a2800276/porter](https://github.com/a2800276/porter) :star:2 - porter stemmer *(**0** commits last month)*
  * [zhenjl/porter2](https://github.com/zhenjl/porter2) :star:19 - High Performance Porter2 Stemmer *(**0** commits last month)*
  * [blevesearch/segment](https://github.com/blevesearch/segment) :star:14 - A Go library for performing Unicode Text Segmentation as described in Unicode Standard Annex #29 *(**0** commits last month)*
  * [neurosnap/sentences](https://github.com/neurosnap/sentences) :star:103 - A multilingual command line sentence tokenizer in Golang *(**0** commits last month)*
  * [goodsign/snowball](https://github.com/goodsign/snowball) :star:13 - Cgo binding for Snowball C library *(**0** commits last month)*
  * [dchest/stemmer](https://github.com/dchest/stemmer) :star:25 - Stemmer packages for Go programming language. Includes English and German stemmers. *(**0** commits last month)*
  * [pebbe/textcat](https://github.com/pebbe/textcat) :star:39 - A Go package for n-gram based text categorization, with support for utf-8 and raw text *(**0** commits last month)*

## Networking
*Libraries for working with various layers of the network*

  * [mdlayher/arp](https://github.com/mdlayher/arp) :star:26 - Package arp implements the ARP protocol, as described in RFC 826. MIT Licensed. *(**0** commits last month)*
  * [stabbycutyou/buffstreams](https://github.com/stabbycutyou/buffstreams) :star:122 - A library to simplify writing applications using TCP sockets to stream protobuff messages *(**0** commits last month)*
  * [zubairhamed/canopus](https://github.com/zubairhamed/canopus) :star:19 - CoAP Client/Server implementing RFC 7252 (For Go/Golang) *(**12** commits last month)*
  * [mdlayher/dhcp6](https://github.com/mdlayher/dhcp6) :star:7 - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. MIT Licensed. *(**0** commits last month)*
  * [miekg/dns](https://github.com/miekg/dns) :star:1320 - DNS library in Go *(**42** commits last month)*
  * [mdlayher/ethernet](https://github.com/mdlayher/ethernet) :star:9 - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. *(**0** commits last month)*
  * [valyala/fasthttp](https://github.com/valyala/fasthttp) :star:1635 - Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http *(**0** commits last month)*
  * [jlaffaye/ftp](https://github.com/jlaffaye/ftp) :star:110 - FTP client package for Go *(**0** commits last month)*
  * [hashicorp/go-getter](https://github.com/hashicorp/go-getter) :star:202 - Package for downloading things from a string URL using a variety of protocols. *(**0** commits last month)*
  * [ccding/go-stun](https://github.com/ccding/go-stun) :star:62 - a go implementation of the STUN client (RFC 3489 and RFC 5389) *(**0** commits last month)*
  * [sunwxg/golibwireshark](https://github.com/sunwxg/golibwireshark) :star:3 -  *(**0** commits last month)*
  * [google/gopacket](https://github.com/google/gopacket) :star:472 - Provides packet processing capabilities for Go *(**0** commits last month)*
  * [akrennmair/gopcap](https://github.com/akrennmair/gopcap) :star:197 - A simple wrapper around libpcap for the Go programming language *(**2** commits last month)*
  * [sunwxg/goshark](https://github.com/sunwxg/goshark) :star:2 -  *(**0** commits last month)*
  * [soniah/gosnmp](https://github.com/soniah/gosnmp) :star:97 - An SNMP library written in GoLang. *(**3** commits last month)*
  * [gansidui/gotcp](https://github.com/gansidui/gotcp) :star:128 - A Go package for quickly building tcp servers *(**0** commits last month)*
  * [koofr/graval](https://github.com/koofr/graval) :star:8 - An experimental go FTP server framework *(**4** commits last month)*
  * [ian-kent/linkio](https://github.com/ian-kent/linkio) :star:15 - Simulate network link speed *(**0** commits last month)*
  * [hashicorp/mdns](https://github.com/hashicorp/mdns) :star:206 - Simple mDNS client/server library in Golang *(**1** commits last month)*
  * [aybabtme/portproxy](https://github.com/aybabtme/portproxy) :star:17 - TCP proxy, highjacks HTTP to allow CORS *(**0** commits last month)*
  * [mdlayher/raw](https://github.com/mdlayher/raw) :star:14 - Package raw enables reading and writing data at the device driver level for a network interface.  MIT Licensed. *(**0** commits last month)*
  * [pkg/sftp](https://github.com/pkg/sftp) :star:220 - SFTP support for the go.crypto/ssh package *(**1** commits last month)*
  * [eduardonunesp/sslb](https://github.com/eduardonunesp/sslb) :star:56 - Golang Super Simple Load Balance *(**0** commits last month)*
  * [firstrow/tcp_server](https://github.com/firstrow/tcp_server) :star:30 - GoLang simple TCP server *(**2** commits last month)*
  * [anacrolix/utp](https://github.com/anacrolix/utp) :star:42 - Implements uTP, the micro transport protocol as used with Bittorrent *(**9** commits last month)*

## OpenGL
*Libraries for using OpenGL in Go.*

  * [go-gl/gl](https://github.com/go-gl/gl) :star:154 - Go bindings for OpenGL (generated via glow) *(**3** commits last month)*
  * [go-gl/glfw](https://github.com/go-gl/glfw) :star:212 - Go bindings for GLFW 3 *(**6** commits last month)*
  * [goxjs/gl](https://github.com/goxjs/gl) :star:57 - Go cross-platform OpenGL bindings. *(**0** commits last month)*
  * [goxjs/glfw](https://github.com/goxjs/glfw) :star:21 - Go cross-platform glfw library for creating an OpenGL context and receiving events. *(**7** commits last month)*
  * [go-gl/mathgl](https://github.com/go-gl/mathgl) :star:104 - A pure Go 3D math library. *(**0** commits last month)*

## ORM
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

  * [astaxie/beedb](https://github.com/astaxie/beedb) :star:586 - beedb is a go ORM,support database/sql interface，pq/mysql/sqlite *(**0** commits last month)*
  * [gosuri/go-store](https://github.com/gosuri/go-store) :star:69 - A simple and fast Redis backed key-value store library for Go *(**0** commits last month)*
  * [cosiner/gomodel](https://github.com/cosiner/gomodel) :star:46 - A lightweight, fast, orm-like library helps interactive with database *(**0** commits last month)*
  * [jinzhu/gorm](https://github.com/jinzhu/gorm) :star:3202 - The fantastic ORM library for Golang, aims to be developer friendly *(**69** commits last month)*
  * [go-gorp/gorp](https://github.com/go-gorp/gorp) :star:1922 - Go Relational Persistence - an ORM-ish library for Go *(**46** commits last month)*
  * [coocood/qbs](https://github.com/coocood/qbs) :star:373 - QBS stands for Query By Struct. A Go ORM. *(**0** commits last month)*
  * [upper/db](https://github.com/upper/db) :star:304 - Magic-free ORM-like package for Go. *(**12** commits last month)*
  * [go-xorm/xorm](https://github.com/go-xorm/xorm) :star:1024 - Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle *(**16** commits last month)*
  * [albrow/zoom](https://github.com/albrow/zoom) :star:119 - A blazing-fast datastore and querying engine for Go built on Redis. *(**0** commits last month)*

## Package Management
*Libraries for package and dependency management.*

  * [LyricalSecurity/gigo](https://github.com/LyricalSecurity/gigo) :star:188 - GIGO: PIP for GO *(**0** commits last month)*
  * [Masterminds/glide](https://github.com/Masterminds/glide) :star:930 - Vendor Package Management for Golang *(**0** commits last month)*
  * [tools/godep](https://github.com/tools/godep) :star:3063 - dependency tool for go *(**0** commits last month)*
  * [mattn/gom](https://github.com/mattn/gom) :star:1009 - Go Manager - bundle for go *(**0** commits last month)*
  * [nitrous-io/goop](https://github.com/nitrous-io/goop) :star:741 - A simple dependency manager for Go (golang), inspired by Bundler. *(**0** commits last month)*
  * [gpmgo/gopm](https://github.com/gpmgo/gopm) :star:862 - Go Package Manager (gopm) is a package manager and build tool for Go. *(**0** commits last month)*
  * [pote/gpm](https://github.com/pote/gpm) :star:891 - Barebones dependency manager for Go. *(**1** commits last month)*
  * [VividCortex/johnny-deps](https://github.com/VividCortex/johnny-deps) :star:215 - Barebones dependency manager for Go. *(**0** commits last month)*
  * [jingweno/nut](https://github.com/jingweno/nut) :star:243 - Vendor Go dependencies *(**14** commits last month)*
  * [DamnWidget/VenGO](https://github.com/DamnWidget/VenGO) :star:84 - Create and manage Isolated Virtual Environments for Go *(**3** commits last month)*

## Query Language

  * [tmc/graphql](https://github.com/tmc/graphql) :star:30 - graphql parser + utilities *(**0** commits last month)*
  * [sevki/graphql](https://github.com/sevki/graphql) :star:27 - GraphQL implementation in go *(**0** commits last month)*
  * [chris-ramon/graphql-go](https://github.com/chris-ramon/graphql-go) :star:276 - An implementation of GraphQL for Go / Golang *(**0** commits last month)*
  * [elgs/jsonql](https://github.com/elgs/jsonql) :star:24 - JSON query expression library in Golang. *(**0** commits last month)*

## Resource Embedding

  * [UnnoTed/fileb0x](https://github.com/UnnoTed/fileb0x) :star:19 - simple customizable tool to embed files in go *(**0** commits last month)*
  * [jteeuwen/go-bindata](https://github.com/jteeuwen/go-bindata) :star:1561 - A small utility which generates Go code from any file. Useful for embedding binary data in a Go program. *(**2** commits last month)*
  * [pyros2097/go-embed](https://github.com/pyros2097/go-embed) :star:1 - Generates go code to embed resource files into your library or executable *(**0** commits last month)*
  * [omeid/go-resources](https://github.com/omeid/go-resources) :star:95 - Unfancy resources embedding for Go with out of box http.FileSystem support. *(**3** commits last month)*
  * [GeertJohan/go.rice](https://github.com/GeertJohan/go.rice) :star:663 - go.rice is a Go package that makes working with resources such as html,js,css,images,templates, etc very easy. *(**11** commits last month)*
  * [go-playground/statics](https://github.com/go-playground/statics) :star:30 - :file_folder: Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks *(**0** commits last month)*
  * [shurcooL/vfsgen](https://github.com/shurcooL/vfsgen) :star:65 - Takes an input http.FileSystem (likely at go generate time) and generates Go code that statically implements it. *(**2** commits last month)*

## Science and Data Analysis
*Libraries for scientific computing and data analyzing.*

  * [ziutek/blas](https://github.com/ziutek/blas) :star:88 - Go implementation of BLAS (Basic Linear Algebra Subprograms) *(**0** commits last month)*
  * [vdobler/chart](https://github.com/vdobler/chart) :star:276 - Provide basic charts in go *(**0** commits last month)*
  * [soniah/evaler](https://github.com/soniah/evaler) :star:20 - Implements a simple floating point arithmetic expression evaluator in Go (golang). *(**0** commits last month)*
  * [VividCortex/ewma](https://github.com/VividCortex/ewma) :star:144 - Exponentially Weighted Moving Average algorithms for Go. *(**0** commits last month)*
  * [skelterjohn/geom](https://github.com/skelterjohn/geom) :star:32 - 2d geometry for golang *(**0** commits last month)*
  * [mjibson/go-dsp](https://github.com/mjibson/go-dsp) :star:367 - Digital Signal Processing for Go *(**0** commits last month)*
  * [skelterjohn/go.matrix](https://github.com/skelterjohn/go.matrix) :star:233 - linear algebra for go *(**0** commits last month)*
  * [anschelsc/gofrac](https://github.com/anschelsc/gofrac) :star:2 - A fractions library for go (http://golang.org) *(**0** commits last month)*
  * [VividCortex/gohistogram](https://github.com/VividCortex/gohistogram) :star:61 - Streaming approximate histograms in Go *(**0** commits last month)*
  * [gonum/matrix](https://github.com/gonum/matrix) :star:185 - Matrix packages for the Go language. *(**20** commits last month)*
  * [gonum/plot](https://github.com/gonum/plot) :star:178 - A repository for plotting and visualizing data *(**22** commits last month)*
  * [gyuho/goraph](https://github.com/gyuho/goraph) :star:181 - Package goraph implements graph, tree data structures and algorithms. *(**8** commits last month)*
  * [alixaxel/pagerank](https://github.com/alixaxel/pagerank) :star:10 - Weighted PageRank implementation in Go *(**0** commits last month)*
  * [montanaflynn/stats](https://github.com/montanaflynn/stats) :star:557 - A statistics package with common functions that are missing from the Golang standard library.  *(**0** commits last month)*
  * [nytlabs/streamtools](https://github.com/nytlabs/streamtools) :star:1266 - tools for working with streams of data *(**0** commits last month)*
  * [spate/vectormath](https://github.com/spate/vectormath) :star:53 - Vectormath for Go *(**0** commits last month)*

## Security
*Libraries that are used to help make your application more secure.*

  * [hlandau/acme](https://github.com/hlandau/acme) :star:525 - :lock: Automatic certificate acquisition tool for ACME (Let's Encrypt) *(**0** commits last month)*
  * [jaredfolkins/badactor](https://github.com/jaredfolkins/badactor) :star:158 - BadActor.org An in-memory application driven jailer written in Go *(**1** commits last month)*
  * [hillu/go-yara](https://github.com/hillu/go-yara) :star:20 - Go bindings for YARA *(**30** commits last month)*
  * [xenolf/lego](https://github.com/xenolf/lego) :star:890 - Let's Encrypt client and ACME library written in Go *(**0** commits last month)*
  * [hlandau/passlib](https://github.com/hlandau/passlib) :star:74 - :key: Idiotproof golang password validation library inspired by Python's passlib *(**0** commits last month)*
  * [elithrar/simple-scrypt](https://github.com/elithrar/simple-scrypt) :star:58 - A convenience library for generating, comparing and inspecting password hashes using the scrypt KDF in Go. *(**0** commits last month)*

## Serialization
*Libraries and tools for binary serialization*

  * [glycerine/go-capnproto](https://github.com/glycerine/go-capnproto) :star:239 - Cap'n Proto library and parser for go. This is go-capnproto-1.0, and does not have rpc. See https://github.com/zombiezen/go-capnproto2 for 2.0 which has rpc and capabilities. *(**12** commits last month)*
  * [glycerine/bambam](https://github.com/glycerine/bambam) :star:51 - auto-generate capnproto schema from your golang source files. Depends on go-capnproto-1.0 at https://github.com/glycerine/go-capnproto *(**9** commits last month)*
  * [ugorji/go](https://github.com/ugorji/go) :star:503 - idiomatic codec and rpc lib for msgpack, cbor, json, etc. msgpack.org[Go] *(**2** commits last month)*
  * [gogo/protobuf](https://github.com/gogo/protobuf) :star:327 - Protocol Buffers for Go with Gadgets *(**3** commits last month)*
  * [golang/protobuf](https://github.com/golang/protobuf) :star:717 - Go support for Google's protocol buffers *(**8** commits last month)*
  * [mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) :star:515 - Go library for decoding generic map values into native Go structures. *(**0** commits last month)*
  * [yvasiyarov/php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) :star:44 - PHP session encoder/decoder written in Go *(**2** commits last month)*
  * [tuvistavie/structomap](https://github.com/tuvistavie/structomap) :star:26 - Easily and dynamically generate maps from Go static structures *(**0** commits last month)*

## Server Applications

  * [xyproto/algernon](https://github.com/xyproto/algernon) :star:201 - :diamond_shape_with_a_dot_inside: HTTP/2 web/application server with Lua support *(**0** commits last month)*
  * [mholt/caddy](https://github.com/mholt/caddy) :star:3826 - Fast, cross-platform HTTP/2 web server with automatic HTTPS *(**23** commits last month)*
  * [cortesi/devd](https://github.com/cortesi/devd) :star:1883 -  A local webserver for developers *(**0** commits last month)*
  * [coreos/etcd](https://github.com/coreos/etcd) :star:8600 - Distributed reliable key-value store for the most critical data of a distributed system *(**193** commits last month)*
  * [minio/minio](https://github.com/minio/minio) :star:500 - Minio is a distributed object storage server written in Golang. *(**213** commits last month)*
  * [minio/minio-xl](https://github.com/minio/minio-xl) :star:14 - Cloud Storage Server for Petascale. *(**0** commits last month)*
  * [sci4me/yakvs](https://github.com/sci4me/yakvs) :star:3 - YAKVS (Yet Another Key Value Store) *(**0** commits last month)*

## Template Engines
*Libraries and tools for templating and lexing.*

  * [yosssi/ace](https://github.com/yosssi/ace) :star:409 - HTML template engine for Go *(**0** commits last month)*
  * [eknkc/amber](https://github.com/eknkc/amber) :star:556 - Amber is an elegant templating engine for Go Programming Language, inspired from HAML and Jade *(**0** commits last month)*
  * [dskinner/damsel](https://github.com/dskinner/damsel) :star:16 - Package damsel provides html outlining via css-selectors and common template functionality. *(**0** commits last month)*
  * [benbjohnson/ego](https://github.com/benbjohnson/ego) :star:235 - An ERB-style templating language for Go. *(**1** commits last month)*
  * [ziutek/kasia.go](https://github.com/ziutek/kasia.go) :star:66 - Templating system for HTML and other text documents - go implementation *(**0** commits last month)*
  * [hoisie/mustache](https://github.com/hoisie/mustache) :star:778 - The mustache template language in Go *(**0** commits last month)*
  * [flosch/pongo2](https://github.com/flosch/pongo2) :star:644 - Django-syntax like template-engine for Go *(**5** commits last month)*
  * [aymerick/raymond](https://github.com/aymerick/raymond) :star:69 - Handlebars for golang *(**0** commits last month)*
  * [sipin/gorazor](https://github.com/sipin/gorazor) :star:473 - Razor view engine for Golang *(**0** commits last month)*
  * [robfig/soy](https://github.com/robfig/soy) :star:106 - Go implementation for Soy templates (Google Closure templates) *(**0** commits last month)*

## Testing
*Libraries for testing codebases and generating test data.*
* Testing Frameworks

  * [go-playground/assert](https://github.com/go-playground/assert) :star:5 - :exclamation:Basic Assertion Library used along side native go testing, with building blocks for custom assertions *(**0** commits last month)*
  * [bmizerany/assert](https://github.com/bmizerany/assert) :star:119 - Asserts to Go testing *(**0** commits last month)*
  * [marioidival/bro](https://github.com/marioidival/bro) :star:10 - bro watch files in directory and run tests for them *(**0** commits last month)*
  * [verdverm/frisby](https://github.com/verdverm/frisby) :star:120 - API testing framework inspired by frisby-js *(**0** commits last month)*
  * [zimmski/go-mutesting](https://github.com/zimmski/go-mutesting) :star:33 - Mutation testing for Go source code *(**4** commits last month)*
  * [franela/goblin](https://github.com/franela/goblin) :star:246 - Minimal and Beautiful Go testing framework *(**0** commits last month)*
  * [DATA-DOG/godog](https://github.com/DATA-DOG/godog) :star:35 - BDD Behat, Cucumber like gherkin feature runner for golang *(**0** commits last month)*
  * [orfjackal/gospec](https://github.com/orfjackal/gospec) :star:110 - Testing framework for Go. Allows writing self-documenting tests/specifications, and executes them concurrently and safely isolated. [UNMAINTAINED] *(**0** commits last month)*
  * [stesla/gospecify](https://github.com/stesla/gospecify) :star:50 - A BDD library for Go *(**0** commits last month)*
  * [rdrdr/hamcrest](https://github.com/rdrdr/hamcrest) :star:21 - Hamcrest matchers for the Go programming language *(**0** commits last month)*
  * [yookoala/restit](https://github.com/yookoala/restit) :star:22 - A Go micro framework to help writing RESTful API integration test *(**0** commits last month)*
  * [stretchr/testify](https://github.com/stretchr/testify) :star:1589 - A sacred extension to the standard go testing package *(**10** commits last month)*

* Mock

  * [maxbrunsfeld/counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) :star:49 - A tool for generating self-contained, type-safe test doubles in go *(**0** commits last month)*
  * [DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) :star:160 - Sql mock driver for golang to test database interactions *(**0** commits last month)*
  * [DATA-DOG/go-txdb](https://github.com/DATA-DOG/go-txdb) :star:1 - Single transaction sql.Driver for GO *(**0** commits last month)*
  * [golang/mock](https://github.com/golang/mock) :star:172 - GoMock is a mocking framework for the Go programming language. *(**0** commits last month)*
  * [tv42/mockhttp](https://github.com/tv42/mockhttp) :star:17 - Mock object for Go http.ResponseWriter *(**0** commits last month)*

* Fuzzing and delta-debugging/reducing/shrinking

  * [dvyukov/go-fuzz](https://github.com/dvyukov/go-fuzz) :star:1257 - Randomized testing for Go *(**0** commits last month)*
  * [google/gofuzz](https://github.com/google/gofuzz) :star:205 - Fuzz testing for go. *(**0** commits last month)*
  * [zimmski/tavor](https://github.com/zimmski/tavor) :star:137 - A generic fuzzing and delta-debugging framework *(**51** commits last month)*

## Text Processing
*Libraries for parsing and manipulating texts.*
* Specific Formats

  * [russross/blackfriday](https://github.com/russross/blackfriday) :star:1680 - Blackfriday: a markdown processor for Go *(**6** commits last month)*
  * [microcosm-cc/bluemonday](https://github.com/microcosm-cc/bluemonday) :star:305 - bluemonday: a fast golang HTML sanitizer (inspired by the OWASP Java HTML Sanitizer) to scrub user generated content of XSS *(**0** commits last month)*
  * [endeveit/enca](https://github.com/endeveit/enca) :star:2 - Minimal cgo bindings for libenca *(**0** commits last month)*
  * [alixaxel/genex](https://github.com/alixaxel/genex) :star:35 - Genex package for Go *(**0** commits last month)*
  * [dustin/go-humanize](https://github.com/dustin/go-humanize) :star:606 - Go Humans! (formatters for units to human friendly sizes) *(**0** commits last month)*
  * [adrianmo/go-nmea](https://github.com/adrianmo/go-nmea) :star:12 - NMEA parser library for Golang *(**0** commits last month)*
  * [jteeuwen/go-pkg-rss](https://github.com/jteeuwen/go-pkg-rss) :star:247 - This package reads RSS and Atom feeds and provides a caching mechanism that adheres to the feed specs. *(**2** commits last month)*
  * [jteeuwen/go-pkg-xmlx](https://github.com/jteeuwen/go-pkg-xmlx) :star:100 - Extension to the standard Go XML package. Maintains a node tree that allows forward/backwards browsing and exposes some simple single/multi-node search functions. *(**2** commits last month)*
  * [mattn/go-runewidth](https://github.com/mattn/go-runewidth) :star:51 -  *(**2** commits last month)*
  * [awalterschulze/gographviz](https://github.com/awalterschulze/gographviz) :star:32 - Parses the Graphviz DOT language in golang *(**0** commits last month)*
  * [polera/gonameparts](https://github.com/polera/gonameparts) :star:15 - Takes a full name and splits it into individual name parts *(**0** commits last month)*
  * [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) :star:2332 - A little like that j-thing, only in Go. *(**3** commits last month)*
  * [zach-klippenstein/goregen](https://github.com/zach-klippenstein/goregen) :star:14 - randexp for Go. *(**0** commits last month)*
  * [endeveit/guesslanguage](https://github.com/endeveit/guesslanguage) :star:16 - Guess the natural language of a text in Go *(**0** commits last month)*
  * [clbanning/mxj](https://github.com/clbanning/mxj) :star:100 - Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages. *(**0** commits last month)*
  * [gosimple/slug](https://github.com/gosimple/slug) :star:65 - URL-friendly slugify with multiple languages support. *(**0** commits last month)*
  * [avelino/slugify](https://github.com/avelino/slugify) :star:9 - A Go slugify application that handles string *(**0** commits last month)*
  * [BurntSushi/toml](https://github.com/BurntSushi/toml) :star:755 - TOML parser for Golang with reflection. *(**2** commits last month)*

* Utility

  * [bndr/gotabulate](https://github.com/bndr/gotabulate) :star:111 - Gotabulate - Easily pretty-print your tabular data with Go *(**0** commits last month)*
  * [codemodus/kace](https://github.com/codemodus/kace) :star:1 - Common case conversions covering common initialisms. *(**0** commits last month)*
  * [codemodus/parth](https://github.com/codemodus/parth) :star:5 - Path parsing / URL segmentation handling. *(**0** commits last month)*
  * [mvdan/xurls](https://github.com/mvdan/xurls) :star:172 - Extract urls from text *(**40** commits last month)*

## Third-party APIs
*Libraries for accessing third party APIs.*

  * [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda) :star:473 - A Go client library for the Twitter 1.1 API *(**14** commits last month)*
  * [aws/aws-sdk-go](https://github.com/aws/aws-sdk-go) :star:2290 - AWS SDK for the Go programming language. *(**98** commits last month)*
  * [naegelejd/brewerydb](https://github.com/naegelejd/brewerydb) :star:10 - Go library for http://www.brewerydb.com/ API *(**0** commits last month)*
  * [samuelcouch/clarifai](https://github.com/samuelcouch/clarifai) :star:27 - Clarifai library for Go *(**0** commits last month)*
  * [bwmarrin/discordgo](https://github.com/bwmarrin/discordgo) :star:31 -  (Golang) Go bindings for Discord *(**0** commits last month)*
  * [huandu/facebook](https://github.com/huandu/facebook) :star:271 - A Facebook Graph API SDK Library For Golang *(**0** commits last month)*
  * [emiddleton/gads](https://github.com/emiddleton/gads) :star:11 - Google Adwords API for Go *(**1** commits last month)*
  * [bit4bit/gami](https://github.com/bit4bit/gami) :star:11 - GO - Asterisk AMI Interface *(**12** commits last month)*
  * [Aorioli/gcm](https://github.com/Aorioli/gcm) :star:26 - Google Cloud Messaging for application servers implemented using the Go programming language. *(**0** commits last month)*
  * [codingsince1985/geo-golang](https://github.com/codingsince1985/geo-golang) :star:108 - (reverse) geocoding service in Go *(**17** commits last month)*
  * [neuegram/ghost](https://github.com/neuegram/ghost) :star:13 - A Go library for Snapchat's API *(**3** commits last month)*
  * [google/go-github](https://github.com/google/go-github) :star:1463 - Go library for accessing the GitHub API *(**3** commits last month)*
  * [gambol99/go-marathon](https://github.com/gambol99/go-marathon) :star:62 - A GO API library for working with Marathon *(**64** commits last month)*
  * [mitchellh/goamz](https://github.com/mitchellh/goamz) :star:610 - Golang Amazon Library *(**11** commits last month)*
  * [michiwend/gomusicbrainz](https://github.com/michiwend/gomusicbrainz) :star:15 - a Go (Golang) MusicBrainz WS2 client library - work in progress *(**0** commits last month)*
  * [google/google-api-go-client](https://github.com/google/google-api-go-client) :star:526 - Auto-generated Google APIs for Go *(**5** commits last month)*
  * [chonthu/go-google-analytics](https://github.com/chonthu/go-google-analytics) :star:4 - Simple Reporting for Google Analytics *(**0** commits last month)*
  * [GoogleCloudPlatform/gcloud-golang](https://github.com/GoogleCloudPlatform/gcloud-golang) :star:250 - Google Cloud APIs Go Client Library *(**8** commits last month)*
  * [jsgilmore/gostorm](https://github.com/jsgilmore/gostorm) :star:61 - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. *(**1** commits last month)*
  * [andybons/hipchat](https://github.com/andybons/hipchat) :star:94 - This project implements a golang client library for the Hipchat API. *(**0** commits last month)*
  * [daneharrigan/hipchat](https://github.com/daneharrigan/hipchat) :star:87 - A golang package to communicate with HipChat over XMPP *(**0** commits last month)*
  * [Medium/medium-sdk-go](https://github.com/Medium/medium-sdk-go) :star:34 - A Golang SDK for Medium's OAuth2 API *(**0** commits last month)*
  * [minio/minio-go](https://github.com/minio/minio-go) :star:71 - Minio Go Library for Amazon S3 compatible cloud storage *(**0** commits last month)*
  * [dukex/mixpanel](https://github.com/dukex/mixpanel) :star:12 - Golang Mixpanel Client *(**0** commits last month)*
  * [logpacker/paypalsdk](https://github.com/logpacker/paypalsdk) :star:59 - Golang client for PayPal REST API *(**0** commits last month)*
  * [playlyfe/playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) :star:0 - This is the official Playlyfe Golang Sdk *(**0** commits last month)*
  * [gregdel/pushover](https://github.com/gregdel/pushover) :star:10 - Go wrapper arround the Pushover API *(**7** commits last month)*
  * [Omie/rrdaclient](https://github.com/Omie/rrdaclient) :star:3 - Go bindings for RRDA https://github.com/fcambus/rrda *(**0** commits last month)*
  * [rapito/go-shopify](https://github.com/rapito/go-shopify) :star:9 - Simple Shopify API for the Go Programming Language *(**0** commits last month)*
  * [nlopes/slack](https://github.com/nlopes/slack) :star:299 - Slack API in Go *(**17** commits last month)*
  * [sergiotapia/smitego](https://github.com/sergiotapia/smitego) :star:6 - SmiteGo is an API wrapper for the Smite game from HiRez. It is written in Go! *(**0** commits last month)*
  * [rapito/go-spotify](https://github.com/rapito/go-spotify) :star:5 - Go library for the Spotify Web API *(**0** commits last month)*
  * [sostronk/go-steam](https://github.com/sostronk/go-steam) :star:7 - Go library for querying Source servers *(**6** commits last month)*
  * [stripe/stripe-go](https://github.com/stripe/stripe-go) :star:330 - Go client for the Stripe API *(**14** commits last month)*
  * [tucnak/telebot](https://github.com/tucnak/telebot) :star:103 - Telegram bot framework written in Go *(**0** commits last month)*
  * [Syfaro/telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) :star:117 - Golang bindings for the Telegram Bot API *(**0** commits last month)*
  * [jbrodriguez/go-tmdb](https://github.com/jbrodriguez/go-tmdb) :star:3 -  *(**0** commits last month)*
  * [poorny/translate](https://github.com/poorny/translate) :star:9 - Go online translation package *(**0** commits last month)*
  * [mattcunningham/gumblr](https://github.com/mattcunningham/gumblr) :star:0 - A Go Wrapper for the Tumblr v2 API *(**0** commits last month)*
  * [go-playground/webhooks](https://github.com/go-playground/webhooks) :star:8 - :fishing_pole_and_fish: Webhook receiver for GitHub and Bitbucket *(**0** commits last month)*

## Utilities
*General utilities and tools to make your life easier.*

  * [topfreegames/apm](https://github.com/topfreegames/apm) :star:25 - APM is a process manager for Golang applications. *(**0** commits last month)*
  * [tmrts/boilr](https://github.com/tmrts/boilr) :star:19 - :zap: Blazingly fast CLI tool for creating projects from boilerplate templates *(**0** commits last month)*
  * [rakyll/coop](https://github.com/rakyll/coop) :star:1127 - Cheat sheet for some of the common concurrent flows in Go *(**0** commits last month)*
  * [ulule/deepcopier](https://github.com/ulule/deepcopier) :star:79 - simple struct copying for golang *(**0** commits last month)*
  * [derekparker/delve](https://github.com/derekparker/delve) :star:3012 - Delve is a debugger for the Go programming language. *(**21** commits last month)*
  * [digitalcrab/fastlz](https://github.com/digitalcrab/fastlz) :star:4 - Wrap over FastLz for GoLang *(**1** commits last month)*
  * [h2non/filetype](https://github.com/h2non/filetype) :star:45 - Small Go package to infer the file type checking the magic numbers signature *(**0** commits last month)*
  * [junegunn/fzf](https://github.com/junegunn/fzf) :star:3838 - :cherry_blossom: A command-line fuzzy finder written in Go *(**11** commits last month)*
  * [go-playground/generate](https://github.com/go-playground/generate) :star:1 - :runner:runs go generate recursively on a specified path or environment variable and can filter by regex *(**0** commits last month)*
  * [rk/go-cron](https://github.com/rk/go-cron) :star:104 - A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. *(**0** commits last month)*
  * [tj/go-debug](https://github.com/tj/go-debug) :star:296 - Conditional debug logging for Golang libraries & applications *(**0** commits last month)*
  * [ungerik/go-dry](https://github.com/ungerik/go-dry) :star:139 - DRY (don't repeat yourself) package for Go *(**0** commits last month)*
  * [beefsack/go-rate](https://github.com/beefsack/go-rate) :star:184 - A timed rate limiter for Go *(**0** commits last month)*
  * [ikeikeikeike/go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) :star:3 - go-sitemap-generator is the easiest way to generate Sitemaps in Go. *(**0** commits last month)*
  * [sadlil/go-trigger](https://github.com/sadlil/go-trigger) :star:13 - A Global event triggerer for golang. Defines functions as event with id string. Trigger the event anywhere from your project. *(**0** commits last month)*
  * [tobyhede/go-underscore](https://github.com/tobyhede/go-underscore) :star:824 -  Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness.  *(**0** commits last month)*
  * [carlescere/goback](https://github.com/carlescere/goback) :star:21 - Golang simple exponential backoff package. *(**0** commits last month)*
  * [VividCortex/godaemon](https://github.com/VividCortex/godaemon) :star:232 - Daemonize Go applications deviously. *(**0** commits last month)*
  * [joho/godotenv](https://github.com/joho/godotenv) :star:247 - A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) *(**0** commits last month)*
  * [dropbox/godropbox](https://github.com/dropbox/godropbox) :star:2953 - Common libraries for writing Go services/applications. *(**18** commits last month)*
  * [cosiner/gohper](https://github.com/cosiner/gohper) :star:201 - common libs here. *(**22** commits last month)*
  * [elgs/gojq](https://github.com/elgs/gojq) :star:18 - JSON query in Golang *(**0** commits last month)*
  * [msempere/golarm](https://github.com/msempere/golarm) :star:13 - Fire alarms with system events *(**0** commits last month)*
  * [mlimaloureiro/golog](https://github.com/mlimaloureiro/golog) :star:11 - Easy and simple CLI time tracker for your tasks *(**0** commits last month)*
  * [bndr/gopencils](https://github.com/bndr/gopencils) :star:329 - Easily consume REST APIs with Go (golang) *(**0** commits last month)*
  * [michiwend/goplaceholder](https://github.com/michiwend/goplaceholder) :star:10 - a small golang lib to generate placeholder images *(**0** commits last month)*
  * [franela/goreq](https://github.com/franela/goreq) :star:417 - Minimal and simple request library for Go language *(**2** commits last month)*
  * [smallnest/goreq](https://github.com/smallnest/goreq) :star:11 - A Simplified Golang Http Client *(**0** commits last month)*
  * [parnurzeal/gorequest](https://github.com/parnurzeal/gorequest) :star:575 - GoRequest -- Simplified HTTP client ( inspired by nodejs SuperAgent ) *(**11** commits last month)*
  * [subosito/gotenv](https://github.com/subosito/gotenv) :star:48 - Load environment variables dynamically in Go. *(**0** commits last month)*
  * [levigross/grequests](https://github.com/levigross/grequests) :star:687 - A Go "clone" of the great and famous Requests library *(**0** commits last month)*
  * [htcat/htcat](https://github.com/htcat/htcat) :star:415 - Parallel and Pipelined HTTP GET Utility *(**0** commits last month)*
  * [facebookgo/httpcontrol](https://github.com/facebookgo/httpcontrol) :star:380 - Package httpcontrol allows for HTTP transport level control around timeouts and retries. *(**0** commits last month)*
  * [afex/hystrix-go](https://github.com/afex/hystrix-go) :star:350 - Netflix's Hystrix latency and fault tolerance library, for Go  *(**3** commits last month)*
  * [bamzi/jobrunner](https://github.com/bamzi/jobrunner) :star:255 - Framework for performing work asynchronously, outside of the request flow *(**0** commits last month)*
  * [miolini/jsonf](https://github.com/miolini/jsonf) :star:20 - Console JSON formatter with query feature *(**0** commits last month)*
  * [ricardolonga/jsongo](https://github.com/ricardolonga/jsongo) :star:44 - Fluent API to make it easier to create Json objects. *(**0** commits last month)*
  * [jaschaephraim/lrserver](https://github.com/jaschaephraim/lrserver) :star:58 - LiveReload server for Go [golang] *(**0** commits last month)*
  * [minio/mc](https://github.com/minio/mc) :star:170 - Minio Client is a replacement for ls, cp, mkdir, diff and rsync commands for filesystems and object storage. *(**66** commits last month)*
  * [sanbornm/mp](https://github.com/sanbornm/mp) :star:9 - Simple Email Parser *(**0** commits last month)*
  * [VividCortex/multitick](https://github.com/VividCortex/multitick) :star:43 - A multiplexor for aligned time.Time tickers in Go *(**0** commits last month)*
  * [e-dard/netbug](https://github.com/e-dard/netbug) :star:25 - Package netbug provides a handler for registering profilers on your own ServeMux. *(**0** commits last month)*
  * [inconshreveable/ngrok](https://github.com/inconshreveable/ngrok) :star:6445 - Introspected tunnels to localhost *(**0** commits last month)*
  * [xta/okrun](https://github.com/xta/okrun) :star:10 - ok, run your gofile *(**0** commits last month)*
  * [maruel/panicparse](https://github.com/maruel/panicparse) :star:1209 - Crash your app in style (Golang) *(**2** commits last month)*
  * [peco/peco](https://github.com/peco/peco) :star:2811 - Simplistic interactive filtering tool *(**9** commits last month)*
  * [sethgrid/pester](https://github.com/sethgrid/pester) :star:50 - Go (golang) http calls with retries and backoff  *(**0** commits last month)*
  * [VividCortex/pm](https://github.com/VividCortex/pm) :star:35 - Processlist manager with TCP listener *(**0** commits last month)*
  * [davecheney/profile](https://github.com/davecheney/profile) :star:414 - A simple profiling support package for Go *(**0** commits last month)*
  * [mozillazg/request](https://github.com/mozillazg/request) :star:106 - Go HTTP Requests for Humans™. *(**7** commits last month)*
  * [ivpusic/rerun](https://github.com/ivpusic/rerun) :star:96 - Configurable recompiling and rerunning go apps when source changes *(**0** commits last month)*
  * [go-resty/resty](https://github.com/go-resty/resty) :star:258 - Simple HTTP and REST client for Go inspired by Ruby rest-client *(**0** commits last month)*
  * [VividCortex/robustly](https://github.com/VividCortex/robustly) :star:101 - Run functions resiliently in Go, catching and restarting panics *(**0** commits last month)*
  * [carlescere/scheduler](https://github.com/carlescere/scheduler) :star:75 - Job scheduling made easy. *(**20** commits last month)*
  * [dghubble/sling](https://github.com/dghubble/sling) :star:258 - A Go HTTP client library for creating and sending API requests *(**0** commits last month)*
  * [briandowns/spinner](https://github.com/briandowns/spinner) :star:212 - Go (golang) package for providing a terminal spinner/progress indicator with options. *(**5** commits last month)*
  * [jmoiron/sqlx](https://github.com/jmoiron/sqlx) :star:1515 - general purpose extensions to golang's database/sql *(**1** commits last month)*
  * [tealeg/xlsx](https://github.com/tealeg/xlsx) :star:819 - Google Go (golang) library for reading and writing XLSX files. *(**13** commits last month)*

## Validation
*Libraries for validation.*

  * [asaskevich/govalidator](https://github.com/asaskevich/govalidator) :star:834 - [Go] Package of validators and sanitizers for strings, numerics, slices and structs *(**5** commits last month)*
  * [go-playground/validator](https://github.com/go-playground/validator) :star:384 - :100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving *(**63** commits last month)*

## Version Control
*Libraries for version control.*

  * [rjeczalik/gh](https://github.com/rjeczalik/gh) :star:31 - Scriptable server and net/http middleware for GitHub Webhooks. *(**0** commits last month)*
  * [libgit2/git2go](https://github.com/libgit2/git2go) :star:698 - Git to Go. Like McDonald's but tastier. *(**10** commits last month)*
  * [sourcegraph/go-vcs](https://github.com/sourcegraph/go-vcs) :star:38 - manipulate and inspect VCS repositories in Go *(**6** commits last month)*
  * [beyang/hgo](https://github.com/beyang/hgo) :star:7 - Hgo is a collection of Go packages providing read-access to local Mercurial repositories. *(**0** commits last month)*

## Video
*Libraries for manipulating video.*

  * [nareix/codec](https://github.com/nareix/codec) :star:139 - Golang libav codec bindings (h264,aac) *(**0** commits last month)*
  * [3d0c/gmf](https://github.com/3d0c/gmf) :star:171 - Go Media Framework *(**0** commits last month)*
  * [giorgisio/goav](https://github.com/giorgisio/goav) :star:90 - Golang bindings for FFmpeg *(**0** commits last month)*
  * [ziutek/gst](https://github.com/ziutek/gst) :star:109 - Go bindings for GStreamer *(**0** commits last month)*

## Web Frameworks
*Full stack web frameworks.*

  * [astaxie/beego](https://github.com/astaxie/beego) :star:6184 - beego is an open-source, high-performance web framework for the Go programming language. *(**29** commits last month)*
  * [go-zoo/bone](https://github.com/go-zoo/bone) :star:854 - Lightning Fast HTTP Multiplexer *(**0** commits last month)*
  * [pressly/chi](https://github.com/pressly/chi) :star:247 - small, fast and expressive router / mux for Go HTTP services built with net/context *(**0** commits last month)*
  * [labstack/echo](https://github.com/labstack/echo) :star:2941 - Echo is  a fast :rocket: and unfancy micro web framework for Go. *(**0** commits last month)*
  * [gin-gonic/gin](https://github.com/gin-gonic/gin) :star:5366 - Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. *(**33** commits last month)*
  * [NYTimes/gizmo](https://github.com/NYTimes/gizmo) :star:1068 - A Microservice Toolkit from The New York Times *(**0** commits last month)*
  * [desertbit/glue](https://github.com/desertbit/glue) :star:124 - Glue - Robust Go and Javascript Socket Library (Alternative to Socket.io) *(**0** commits last month)*
  * [ant0ine/go-json-rest](https://github.com/ant0ine/go-json-rest) :star:2111 - A quick and easy way to setup a RESTful JSON API *(**25** commits last month)*
  * [go-kit/kit](https://github.com/go-kit/kit) :star:3019 - A standard library for microservices. *(**34** commits last month)*
  * [codehack/go-relax](https://github.com/codehack/go-relax) :star:115 - Framework for building RESTful API's in Go *(**0** commits last month)*
  * [ungerik/go-rest](https://github.com/ungerik/go-rest) :star:84 - A small and evil REST framework for Go *(**0** commits last month)*
  * [googollee/go-socket.io](https://github.com/googollee/go-socket.io) :star:994 - socket.io library for golang, a realtime application framework. *(**2** commits last month)*
  * [raphael/goa](https://github.com/raphael/goa) :star:552 - Design-based HTTP microservices in Go *(**0** commits last month)*
  * [bahlo/goat](https://github.com/bahlo/goat) :star:76 - :goat: A minimalistic REST API server in Go *(**0** commits last month)*
  * [gocraft/web](https://github.com/gocraft/web) :star:960 - Go Router + Middleware. Your Contexts. *(**0** commits last month)*
  * [goji/goji](https://github.com/goji/goji) :star:71 - Goji is a minimalistic and flexible HTTP request multiplexer for Go (golang) *(**0** commits last month)*
  * [jcuga/golongpoll](https://github.com/jcuga/golongpoll) :star:244 - golang HTTP longpolling library.  Makes web pub-sub easy :smiley: :coffee: :computer: *(**0** commits last month)*
  * [rainycape/gondola](https://github.com/rainycape/gondola) :star:292 - The web framework for writing faster sites, faster *(**0** commits last month)*
  * [ian-kent/goose](https://github.com/ian-kent/goose) :star:13 - Server-Sent Events in Go *(**0** commits last month)*
  * [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) :star:2700 - A high performance HTTP request router that scales well *(**16** commits last month)*
  * [dimfeld/httptreemux](https://github.com/dimfeld/httptreemux) :star:119 - High-speed, flexible tree-based HTTP router for Go. *(**0** commits last month)*
  * [Unknwon/macaron](https://github.com/Unknwon/macaron) :star:7 - :exclamation::exclamation::exclamation: [deprecated] Moved to https://github.com/go-macaron/macaron *(**5** commits last month)*
  * [paulbellamy/mango](https://github.com/paulbellamy/mango) :star:296 - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. *(**0** commits last month)*
  * [imdario/medeina](https://github.com/imdario/medeina) :star:15 - Go HTTP routing tree based on HttpRouter. Inspired by Roda and Cuba. *(**0** commits last month)*
  * [gorilla/mux](https://github.com/gorilla/mux) :star:1951 - A powerful URL router and dispatcher for golang. *(**1** commits last month)*
  * [ivpusic/neo](https://github.com/ivpusic/neo) :star:202 - Go Web Framework *(**19** commits last month)*
  * [bmizerany/pat](https://github.com/bmizerany/pat) :star:900 -  *(**0** commits last month)*
  * [resoursea/api](https://github.com/resoursea/api) :star:24 - A REST framework for quickly writing resource based services in Golang. *(**8** commits last month)*
  * [revel/revel](https://github.com/revel/revel) :star:6315 - A high productivity, full-stack web framework for the Go language. *(**12** commits last month)*
  * [goanywhere/rex](https://github.com/goanywhere/rex) :star:10 - Pleasures for Web in Golang *(**0** commits last month)*
  * [VividCortex/siesta](https://github.com/VividCortex/siesta) :star:325 - Composable framework for writing HTTP handlers in Go. *(**13** commits last month)*
  * [lunny/tango](https://github.com/lunny/tango) :star:345 - Micro & pluggable web framework for Go *(**6** commits last month)*
  * [rcrowley/go-tigertonic](https://github.com/rcrowley/go-tigertonic) :star:914 - A Go framework for building JSON web services inspired by Dropwizard *(**1** commits last month)*
  * [pilu/traffic](https://github.com/pilu/traffic) :star:458 - Sinatra inspired regexp/pattern mux and web framework for Go *(**0** commits last month)*
  * [volatile/core](https://github.com/volatile/core) :star:78 - Pure handlers stack *(**0** commits last month)*
  * [hoisie/web](https://github.com/hoisie/web) :star:2582 - The easiest way to create web applications with Go *(**0** commits last month)*
  * [rs/xmux](https://github.com/rs/xmux) :star:40 - xmux is a httprouter fork on top of xhandler (net/context aware) *(**0** commits last month)*
  * [cosiner/zerver](https://github.com/cosiner/zerver) :star:119 - a RESTful API framework *(**14** commits last month)*
  * [daryl/zeus](https://github.com/daryl/zeus) :star:91 - Go HTTP router. *(**0** commits last month)*

### Middlewares
#### Actual middlewares

  * [rs/cors](https://github.com/rs/cors) :star:282 - Go net/http configurable handler to handle CORS requests *(**7** commits last month)*
  * [rs/formjson](https://github.com/rs/formjson) :star:17 - Go net/http handler to transparently manage posted JSON *(**0** commits last month)*
  * [ulule/limiter](https://github.com/ulule/limiter) :star:207 - Dead simple rate limit middleware for Go. *(**0** commits last month)*
  * [didip/tollbooth](https://github.com/didip/tollbooth) :star:351 - Simple middleware to rate-limit HTTP requests. *(**0** commits last month)*
  * [sebest/xff](https://github.com/sebest/xff) :star:41 - A Golang Middleware to handle X-Forwarded-For Header *(**10** commits last month)*

#### Libraries for creating HTTP middlewares

  * [justinas/alice](https://github.com/justinas/alice) :star:800 - Painless middleware chaining for Go *(**0** commits last month)*
  * [codemodus/catena](https://github.com/codemodus/catena) :star:2 - http.Handler wrapper catenation. *(**0** commits last month)*
  * [codemodus/chain](https://github.com/codemodus/chain) :star:37 - Handler wrapper chaining with scoped data. *(**0** commits last month)*
  * [go-on/wrap](https://github.com/go-on/wrap) :star:49 - Go http.Hander based middleware stack with context sharing *(**0** commits last month)*
  * [carbocation/interpose](https://github.com/carbocation/interpose) :star:219 - Minimalist net/http middleware for golang *(**4** commits last month)*
  * [stephens2424/muxchain](https://github.com/stephens2424/muxchain) :star:192 - Lightweight Middleware for net/http *(**0** commits last month)*
  * [codegangsta/negroni](https://github.com/codegangsta/negroni) :star:3328 - Idiomatic HTTP Middleware for Golang *(**0** commits last month)*
  * [unrolled/render](https://github.com/unrolled/render) :star:612 - Go package for easily rendering JSON, XML, binary data, and HTML templates responses. *(**1** commits last month)*
  * [thoas/stats](https://github.com/thoas/stats) :star:305 - A Go middleware that stores various information about your web application (response time, status code count, etc.) *(**0** commits last month)*

# Tools
## Code Analysis

  * [mibk/dupl](https://github.com/mibk/dupl) :star:40 - a tool for code clone detection *(**1** commits last month)*
  * [kisielk/errcheck](https://github.com/kisielk/errcheck) :star:501 - errcheck checks that you checked errors. *(**4** commits last month)*
  * [davecheney/gcvis](https://github.com/davecheney/gcvis) :star:472 - Visualise Go program GC trace data in real time *(**0** commits last month)*
  * [alecthomas/gometalinter](https://github.com/alecthomas/gometalinter) :star:684 - Concurrently run Go lint tools and normalise their output *(**16** commits last month)*
  * [qiniu/checkstyle](https://github.com/qiniu/checkstyle) :star:22 - checkstyle for go *(**0** commits last month)*
  * [yuroyoro/goast-viewer](https://github.com/yuroyoro/goast-viewer) :star:130 - Golang AST visualizer *(**0** commits last month)*
  * [golang/lint](https://github.com/golang/lint) :star:1209 - This is a linter for Go source code. *(**2** commits last month)*
  * [shurcooL/gostatus](https://github.com/shurcooL/gostatus) :star:119 - A command line tool that shows the status of Go repositories. *(**1** commits last month)*
  * [mvdan/interfacer](https://github.com/mvdan/interfacer) :star:203 - A linter that suggests interface types *(**0** commits last month)*
  * [mccoyst/validate](https://github.com/mccoyst/validate) :star:55 - A Go package to automatically validate fields with tags *(**0** commits last month)*

## Editor Plugins

  * [go-lang-plugin-org/go-lang-idea-plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) :star:2904 - Google Go language IDE built using the IntelliJ Platform *(**256** commits last month)*
  * [joefitzgerald/go-plus](https://github.com/joefitzgerald/go-plus) :star:722 - An Enhanced Go Experience For The Atom Editor *(**28** commits last month)*
  * [GoClipse/goclipse](https://github.com/GoClipse/goclipse) :star:473 - Eclipse IDE for the Go programming language *(**76** commits last month)*
  * [nsf/gocode](https://github.com/nsf/gocode) :star:2912 - An autocompletion daemon for the Go programming language *(**2** commits last month)*
  * [DisposaBoy/GoSublime](https://github.com/DisposaBoy/GoSublime) :star:2086 - A  Golang plugin collection for the text editor SublimeText 2 providing code completion and other IDE-like features. *(**0** commits last month)*
  * [velour/velour](https://github.com/velour/velour) :star:12 - An IRC client for acme — the project that started it all. *(**0** commits last month)*
  * [rjohnsondev/vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) :star:60 - Vim compiler plugin for Go (golang) *(**0** commits last month)*
  * [fatih/vim-go](https://github.com/fatih/vim-go) :star:4200 - Go development plugin for Vim *(**38** commits last month)*
  * [eaburns/Watch](https://github.com/eaburns/Watch) :star:102 - Watches for changes in a directory tree and reruns a command in an acme win or just on the terminal. *(**2** commits last month)*

## Go Tools

  * [songgao/colorgo](https://github.com/songgao/colorgo) :star:57 - Colorize (highlight) `go build` command output *(**0** commits last month)*
  * [skelterjohn/go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) :star:25 - bash completion for go and wgo *(**0** commits last month)*

## Software Packages
### DevOps Tools

  * [smira/aptly](https://github.com/smira/aptly) :star:971 - aptly - Debian repository management tool *(**74** commits last month)*
  * [soniah/awsenv](https://github.com/soniah/awsenv) :star:5 - AWS environment config loader *(**0** commits last month)*
  * [rakyll/boom](https://github.com/rakyll/boom) :star:3871 - HTTP(S) load generator, ApacheBench (ab) replacement, written in Go *(**0** commits last month)*
  * [bosun-monitor/bosun](https://github.com/bosun-monitor/bosun) :star:1574 - Time Series Alerting Framework *(**76** commits last month)*
  * [liudng/dogo](https://github.com/liudng/dogo) :star:83 - Monitoring changes in the source file and automatically compile and run (restart). *(**6** commits last month)*
  * [chrismckenzie/dropship](https://github.com/chrismckenzie/dropship) :star:20 - instead of jumping ship... Dropship *(**0** commits last month)*
  * [hypersleep/easyssh](https://github.com/hypersleep/easyssh) :star:73 - Golang package for easy remote execution through SSH *(**11** commits last month)*
  * [rcrowley/go-metrics](https://github.com/rcrowley/go-metrics) :star:1069 - Go port of Coda Hale's Metrics library *(**1** commits last month)*
  * [sanbornm/go-selfupdate](https://github.com/sanbornm/go-selfupdate) :star:323 - Enable your Golang applications to self update *(**3** commits last month)*
  * [cryptojuice/gobrew](https://github.com/cryptojuice/gobrew) :star:137 - Shell script to download and set GO environmental paths to allow multiple versions. *(**6** commits last month)*
  * [sirnewton01/godbg](https://github.com/sirnewton01/godbg) :star:180 - Web-based gdb front-end application *(**0** commits last month)*
  * [inconshreveable/gonative](https://github.com/inconshreveable/gonative) :star:219 - Build Go Toolchains /w native libs for cross-compilation *(**1** commits last month)*
  * [mitchellh/gox](https://github.com/mitchellh/gox) :star:1430 - A dead simple, no frills Go cross compile tool *(**0** commits last month)*
  * [laher/goxc](https://github.com/laher/goxc) :star:1208 - a build tool for Go, with a focus on cross-compiling, packaging and deployment *(**14** commits last month)*
  * [moovweb/gvm](https://github.com/moovweb/gvm) :star:1736 - Go Version Manager *(**4** commits last month)*
  * [heroku/hk](https://github.com/heroku/hk) :star:766 - Fast Heroku command-line interface *(**2** commits last month)*
  * [ajvb/kala](https://github.com/ajvb/kala) :star:757 - Modern Job Scheduler *(**0** commits last month)*
  * [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes) :star:12425 - Container Cluster Manager from Google *(**1217** commits last month)*
  * [emicklei/mora](https://github.com/emicklei/mora) :star:133 - MongoDB generic REST server in Go *(**0** commits last month)*
  * [ostrost/ostent](https://github.com/ostrost/ostent) :star:58 - ostent collects system metrics to display and relay *(**0** commits last month)*
  * [mitchellh/packer](https://github.com/mitchellh/packer) :star:4952 - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. *(**62** commits last month)*
  * [alouche/rodent](https://github.com/alouche/rodent) :star:31 - Manage Go Versions/Projects/Dependencies *(**0** commits last month)*
  * [rlmcpherson/s3gof3r](https://github.com/rlmcpherson/s3gof3r) :star:608 - Fast, concurrent, streaming access to Amazon S3, including gof3r, a CLI. http://godoc.org/github.com/rlmcpherson/s3gof3r *(**5** commits last month)*
  * [tsenart/vegeta](https://github.com/tsenart/vegeta) :star:2909 - HTTP load testing tool and library. It's over 9000! *(**5** commits last month)*
  * [adnanh/webhook](https://github.com/adnanh/webhook) :star:311 - webhook is a lightweight configurable tool written in Go, that allows you to easily create HTTP endpoints (hooks) on your server, which you can use to execute configured commands. *(**11** commits last month)*

### Other Software

  * [tejo/boxed](https://github.com/tejo/boxed) :star:41 - dropbox based blog engine, written in go. *(**36** commits last month)*
  * [gocircuit/circuit](https://github.com/gocircuit/circuit) :star:1209 - Circuit: Dynamic cloud orchestration http://gocircuit.org *(**52** commits last month)*
  * [tylertreat/Comcast](https://github.com/tylertreat/Comcast) :star:4099 - Simulating shitty network connections so you can build better systems. *(**9** commits last month)*
  * [kelseyhightower/confd](https://github.com/kelseyhightower/confd) :star:2314 - Manage local application configuration files using templates and data from etcd or consul *(**10** commits last month)*
  * [coreos/fleet](https://github.com/coreos/fleet) :star:2056 - A Distributed init System *(**7** commits last month)*
  * [goccmack/gocc](https://github.com/goccmack/gocc) :star:15 - Parser / Scanner Generator *(**0** commits last month)*
  * [diankong/GoDocTooltip](https://github.com/diankong/GoDocTooltip) :star:7 - A Chrome extension for golang users.When you're at golang's official doc site, it will show function's description as tooltip on function list *(**0** commits last month)*
  * [buger/gor](https://github.com/buger/gor) :star:4588 - Gor is an open-source tool for capturing and replaying live HTTP traffic into a test environment in order to continuously test your system with real data. It can be used to increase confidence in code deployments, configuration changes and infrastructure changes. *(**4** commits last month)*
  * [mozilla-services/heka](https://github.com/mozilla-services/heka) :star:2831 - Data collection and processing made easy. *(**161** commits last month)*
  * [dimiro1/ipe](https://github.com/dimiro1/ipe) :star:129 - An open source Pusher server implementation compatible with Pusher client libraries written in GO *(**11** commits last month)*
  * [visualfc/liteide](https://github.com/visualfc/liteide) :star:2535 - LiteIDE is a simple, open source, cross-platform Go IDE.  *(**96** commits last month)*
  * [unix4fun/naclpipe](https://github.com/unix4fun/naclpipe) :star:4 - NaCL pipe *(**0** commits last month)*
  * [fogleman/nes](https://github.com/fogleman/nes) :star:2136 - NES emulator written in Go. *(**0** commits last month)*
  * [noraesae/orange-cat](https://github.com/noraesae/orange-cat) :star:134 - A Markdown previewer written in Go *(**0** commits last month)*
  * [pointlander/peg](https://github.com/pointlander/peg) :star:267 - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. *(**1** commits last month)*
  * [zachlatta/postman](https://github.com/zachlatta/postman) :star:622 - Command-line utility for batch-sending email. *(**0** commits last month)*
  * [restic/restic](https://github.com/restic/restic) :star:441 - restic backup program *(**90** commits last month)*
  * [coreos/rkt](https://github.com/coreos/rkt) :star:5048 - rkt is an App Container runtime for Linux *(**146** commits last month)*
  * [chrislusf/seaweedfs](https://github.com/chrislusf/seaweedfs) :star:2111 - SeaweedFS is a simple and highly scalable distributed file system. There are two objectives:  to store billions of files! to serve the files fast! Instead of supporting full POSIX file system semantics, Seaweed-FS choose to implement only a key~file mapping. Similar to the word "NoSQL", you can call it as "NoFS". *(**39** commits last month)*
  * [msoap/shell2http](https://github.com/msoap/shell2http) :star:35 - Executing shell commands via http server (GoLang) *(**0** commits last month)*
  * [intelsdi-x/snap](https://github.com/intelsdi-x/snap) :star:518 - A powerful telemetry framework *(**131** commits last month)*
  * [pressly/sup](https://github.com/pressly/sup) :star:804 - Super simple deployment tool - just Unix - think of it like 'make' for a network of servers *(**8** commits last month)*
  * [kyleterry/tenyks](https://github.com/kyleterry/tenyks) :star:152 - The Tenyks IRC bot. *(**0** commits last month)*
  * [shopify/toxiproxy](https://github.com/shopify/toxiproxy) :star:1085 - :alarm_clock: :fire: A proxy to simulate network and system conditions *(**16** commits last month)*
  * [ian-kent/websysd](https://github.com/ian-kent/websysd) :star:17 - Like Marathon or Upstart - for your desktop! *(**0** commits last month)*
  * [wellington/wellington](https://github.com/wellington/wellington) :star:160 - Spriting that sass has been missing *(**43** commits last month)*

# Resources
## Benchmarks

  * [davecheney/autobench](https://github.com/davecheney/autobench) :star:83 - Go benchmark harness.  *(**0** commits last month)*
  * [julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) :star:621 - Go HTTP request router and web framework benchmark *(**5** commits last month)*
  * [hgfischer/go-type-assertion-benchmark](https://github.com/hgfischer/go-type-assertion-benchmark) :star:2 - Naive performance test of two ways to do type assertion in Go. *(**0** commits last month)*
  * [alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) :star:260 - Benchmarks of Go serialization methods *(**2** commits last month)*
  * [PuerkitoBio/gocostmodel](https://github.com/PuerkitoBio/gocostmodel) :star:46 - Benchmarks of common basic operations for the Go language. *(**0** commits last month)*
  * [amscanne/golang-micro-benchmarks](https://github.com/amscanne/golang-micro-benchmarks) :star:5 - Tiny collection of micro benchmarks. *(**0** commits last month)*
  * [tyler-smith/golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) :star:18 - A benchmarking shootout of various db/SQL utilities for Go *(**0** commits last month)*
  * [feyeleanor/GoSpeed](https://github.com/feyeleanor/GoSpeed) :star:44 - Go micro-benchmarks for calculating the speed of language constructs *(**0** commits last month)*
  * [jimrobinson/kvbench](https://github.com/jimrobinson/kvbench) :star:9 - Key/Value database benchmark *(**0** commits last month)*
  * [fawick/speedtest-resize](https://github.com/fawick/speedtest-resize) :star:60 - Compare various Image resize algorithms for the Go language *(**0** commits last month)*

## E-Books

  * [dariubs/GoBooks](https://github.com/dariubs/GoBooks) :star:1851 - List of Golang books *(**0** commits last month)*

## Websites

  * [lukasz-madon/awesome-remote-job](https://github.com/lukasz-madon/awesome-remote-job) :star:4853 - A curated list of awesome remote jobs and resources. Inspired by https://github.com/vinta/awesome-python *(**43** commits last month)*
  * [bayandin/awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) :star:15222 - A curated list of awesome awesomeness *(**0** commits last month)*
  * [mholt/golang-graphics](https://github.com/mholt/golang-graphics) :star:70 - Community-contributed Go graphics files *(**1** commits last month)*

### Tutorials

  * [mkaz/working-with-go](https://github.com/mkaz/working-with-go) :star:471 - A set of example golang code to start learning Go *(**0** commits last month)*

## Windows

  * [go-ole/go-ole](https://github.com/go-ole/go-ole) :star:175 - win32 ole implementation for golang *(**22** commits last month)*



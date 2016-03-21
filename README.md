# Fabulous Go

This application parses github repos from [Awesome Go](https://github.com/avelino/awesome-go) and populates the lists with additional data, like stars count and commit activity.

To learn how it works and how to refresh the list, see [`BUILDING.md`](https://github.com/yamnikov-oleg/fabulous-go/blob/master/BUILDING.md).

**clm** stands for 'commits last month'.

## Contents
* [Packages](#packages)
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


# Packages
## Audio/Music
*Libraries for manipulating audio.*

  * [mdlayher/**waveform**](https://github.com/mdlayher/waveform) - Go package capable of generating waveform images from audio streams. MIT Licensed. *(:star:114 | 0clm)*
  * [rakyll/**portmidi**](https://github.com/rakyll/portmidi) - Go bindings for libportmidi *(:star:80 | 0clm)*
  * [eaburns/**flac**](https://github.com/eaburns/flac) - A Free Lossless Audio Codec decoder in Go *(:star:47 | 0clm)*
  * [wtolson/**go-taglib**](https://github.com/wtolson/go-taglib) - Go wrapper for taglib *(:star:38 | 1clm)*
  * [krig/**go-sox**](https://github.com/krig/go-sox) - libsox bindings for go *(:star:35 | 0clm)*
  * [gordonklaus/**portaudio**](https://github.com/gordonklaus/portaudio) - Go bindings for the PortAudio audio I/O library *(:star:31 | 0clm)*
  * [mewkiz/**flac**](https://github.com/mewkiz/flac) - Package flac provides access to FLAC (Free Lossless Audio Codec) streams. *(:star:21 | 0clm)*
  * [go-ontomix/**ontomix**](https://github.com/go-ontomix/ontomix) - Go-native audio mixer for Music apps *(:star:14 | 0clm)*
  * [tcolgate/**mp3**](https://github.com/tcolgate/mp3) - golang mp3 frame parser *(:star:11 | 1clm)*
  * [mccoyst/**vorbis**](https://github.com/mccoyst/vorbis) - A "native" ogg vorbis decoder for Go (uses inline stb_vorbis) *(:star:9 | 0clm)*
  * [zhulik/**go_mediainfo**](https://github.com/zhulik/go_mediainfo) - Golang bindings for libmediainfo *(:star:2 | 0clm)*

## Authentication & OAuth
*Libraries for implementing authentications schemes.*

  * [dgrijalva/**jwt-go**](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens (JWT) *(:star:1006 | 7clm)*
  * [RangelReale/**osin**](https://github.com/RangelReale/osin) - Golang OAuth2 server library *(:star:836 | 0clm)*
  * [golang/**oauth2**](https://github.com/golang/oauth2) - Go OAuth2 *(:star:665 | 6clm)*
  * [markbates/**goth**](https://github.com/markbates/goth) - Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. *(:star:636 | 3clm)*
  * [dghubble/**gologin**](https://github.com/dghubble/gologin) - Chainable Go login handlers for authentication providers (oauth1, oauth2) *(:star:618 | 0clm)*
  * [square/**go-jose**](https://github.com/square/go-jose) - An implementation of JOSE standards in Golang. *(:star:485 | 5clm)*
  * [mikespook/**gorbac**](https://github.com/mikespook/gorbac) - goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. *(:star:307 | 0clm)*
  * [bradrydzewski/**go.auth**](https://github.com/bradrydzewski/go.auth) - authentication API for Go web applications *(:star:301 | 0clm)*
  * [xyproto/**permissions2**](https://github.com/xyproto/permissions2) -   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions *(:star:148 | 4clm)*
  * [smartystreets/**go-aws-auth**](https://github.com/smartystreets/go-aws-auth) - Signs requests to Amazon Web Services (AWS) using IAM roles or signed signature versions 2, 3, and 4. Supports S3 and STS. *(:star:117 | 0clm)*
  * [goji/**httpauth**](https://github.com/goji/httpauth) - HTTP Authentication middlewares *(:star:88 | 2clm)*
  * [GeertJohan/**yubigo**](https://github.com/GeertJohan/yubigo) -   Yubigo is a Yubikey client API library that provides an easy way to integrate the Yubico Yubikey into your existing Go-based user authentication infrastructure. *(:star:53 | 0clm)*

## Command Line
### Standard CLI
*Libraries for building standard or basic Command Line applications*

  * [codegangsta/**cli**](https://github.com/codegangsta/cli) - A small package for building command line apps in Go *(:star:3934 | 4clm)*
  * [spf13/**cobra**](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions *(:star:1862 | 28clm)*
  * [docopt/**docopt.go**](https://github.com/docopt/docopt.go) - A command-line arguments parser that will make you smile. *(:star:653 | 0clm)*
  * [jessevdk/**go-flags**](https://github.com/jessevdk/go-flags) - go command line option parser *(:star:555 | 2clm)*
  * [alecthomas/**kingpin**](https://github.com/alecthomas/kingpin) - A Go (golang) command line and flag parser *(:star:536 | 5clm)*
  * [chzyer/**readline**](https://github.com/chzyer/readline) - Readline is a pure go(golang) implementation for GNU-Readline kind library *(:star:522 | 0clm)*
  * [tcnksm/**gcli**](https://github.com/tcnksm/gcli) - The easy way to build Golang command-line application. *(:star:469 | 0clm)*
  * [mitchellh/**cli**](https://github.com/mitchellh/cli) - A Go library for implementing command-line interfaces. *(:star:386 | 1clm)*
  * [jawher/**mow.cli**](https://github.com/jawher/mow.cli) - A versatile library for building CLI applications in Go *(:star:335 | 0clm)*
  * [peterh/**liner**](https://github.com/peterh/liner) - Pure Go line editor with history, inspired by linenoise *(:star:289 | 2clm)*
  * [tucnak/**climax**](https://github.com/tucnak/climax) - Climax is an alternative CLI with "human face" *(:star:92 | 0clm)*
  * [ukautz/**clif**](https://github.com/ukautz/clif) - Another CLI framework for Go. It works on my machine. *(:star:63 | 0clm)*
  * [mkideal/**cli**](https://github.com/mkideal/cli) - Golang cli,best command line interface for go (go cli) *(:star:22 | 0clm)*

### Advanced Console UIs
*Libraries for building Console Applications and Console User Interfaces*

  * [gizak/**termui**](https://github.com/gizak/termui) - Golang terminal dashboard *(:star:4523 | 40clm)*
  * [nsf/**termbox-go**](https://github.com/nsf/termbox-go) - Pure Go termbox implementation *(:star:1360 | 2clm)*
  * [fatih/**color**](https://github.com/fatih/color) - Color package for Go (golang) *(:star:795 | 2clm)*
  * [gosuri/**uiprogress**](https://github.com/gosuri/uiprogress) - A go library to render progress bars in terminal applications *(:star:776 | 0clm)*
  * [jroimartin/**gocui**](https://github.com/jroimartin/gocui) - Minimalist Go package aimed at creating Console User Interfaces. *(:star:507 | 0clm)*
  * [gosuri/**uilive**](https://github.com/gosuri/uilive) - uilive is a go library for updating terminal output in realtime *(:star:423 | 0clm)*
  * [gosuri/**uitable**](https://github.com/gosuri/uitable) - A go library to improve readability in terminal apps using tabular data *(:star:280 | 0clm)*
  * [ttacon/**chalk**](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk *(:star:150 | 0clm)*
  * [daviddengcn/**go-colortext**](https://github.com/daviddengcn/go-colortext) - Change the color of console text. *(:star:141 | 0clm)*
  * [apcera/**termtables**](https://github.com/apcera/termtables) - Go ASCII Table Generator, ported from the Ruby terminal-tables library *(:star:48 | 1clm)*
  * [TreyBastian/**colourize**](https://github.com/TreyBastian/colourize) - An ANSI colour terminal package for Go *(:star:3 | 0clm)*

## Configuration
*Libraries for configuration parsing*

  * [spf13/**viper**](https://github.com/spf13/viper) - Go configuration with fangs *(:star:1402 | 8clm)*
  * [go-ini/**ini**](https://github.com/go-ini/ini) - Package ini provides INI file read and write functionality in Go. *(:star:242 | 0clm)*
  * [vrischmann/**envconfig**](https://github.com/vrischmann/envconfig) - Small library to read your configuration from environment variables *(:star:89 | 0clm)*
  * [FogCreek/**mini**](https://github.com/FogCreek/mini) - A golang package for parsing ini-style configuration files *(:star:82 | 0clm)*
  * [caarlos0/**env**](https://github.com/caarlos0/env) - A KISS way to deal with environment variables in Go. *(:star:82 | 0clm)*
  * [tomazk/**envcfg**](https://github.com/tomazk/envcfg) - Un-marshaling environment variables to Go structs *(:star:75 | 0clm)*
  * [olebedev/**config**](https://github.com/olebedev/config) - JSON or YAML configuration wrapper with convenient access methods. *(:star:63 | 0clm)*
  * [tucnak/**store**](https://github.com/tucnak/store) - A dead simple configuration manager for Go applications *(:star:37 | 0clm)*
  * [go-gcfg/**gcfg**](https://github.com/go-gcfg/gcfg) - read INI-style configuration files into Go structs; supports user-defined types and subsections *(:star:33 | 0clm)*
  * [ian-kent/**gofigure**](https://github.com/ian-kent/gofigure) - Go configuration made easy! *(:star:32 | 0clm)*
  * [paked/**configure**](https://github.com/paked/configure) - Configure is a Go package that gives you easy configuration of your project through redundancy *(:star:6 | 0clm)*
  * [ian-kent/**envconf**](https://github.com/ian-kent/envconf) - Configure Go applications from the environment *(:star:2 | 0clm)*

## Continuous Integration
*Tools for help with continuous integration*

  * [drone/**drone**](https://github.com/drone/drone) - Drone is a Continuous Delivery platform built on Docker, written in Go *(:star:6442 | 62clm)*
  * [mattn/**goveralls**](https://github.com/mattn/goveralls) -  *(:star:236 | 3clm)*
  * [go-playground/**overalls**](https://github.com/go-playground/overalls) - :jeans:Multi-Package go project coverprofile for tools like goveralls *(:star:16 | 0clm)*

## CSS Preprocessors
*Libraries for preprocessing CSS files*

  * [yosssi/**gcss**](https://github.com/yosssi/gcss) - Pure Go CSS Preprocessor *(:star:335 | 0clm)*
  * [c9s/**c6**](https://github.com/c9s/c6) - Compile SASS Faster ! C6 is a SASS-compatible compiler written in Go. *(:star:329 | 18clm)*
  * [wellington/**go-libsass**](https://github.com/wellington/go-libsass) - Go wrapper for libsass, the only Sass 3.3 compiler for Go *(:star:22 | 3clm)*

## Data Structures
*Generic datastructures and algorithms in Go.*

  * [Workiva/**go-datastructures**](https://github.com/Workiva/go-datastructures) -  *(:star:2369 | 7clm)*
  * [tylertreat/**BoomFilters**](https://github.com/tylertreat/BoomFilters) - Probabilistic data structures for processing continuous, unbounded streams. *(:star:676 | 4clm)*
  * [deckarep/**golang-set**](https://github.com/deckarep/golang-set) - A simple set type for the Go language. *(:star:356 | 0clm)*
  * [willf/**bloom**](https://github.com/willf/bloom) - Go package implementing Bloom filters *(:star:226 | 9clm)*
  * [smartystreets/**mafsa**](https://github.com/smartystreets/mafsa) - Package mafsa implements Minimal Acyclic Finite State Automata in Go, essentially a high-speed, memory-efficient, Unicode-friendly set of strings. *(:star:203 | 0clm)*
  * [willf/**bitset**](https://github.com/willf/bitset) - Go package implementing bitsets *(:star:192 | 4clm)*
  * [hailocab/**go-geoindex**](https://github.com/hailocab/go-geoindex) - Go native library for fast point tracking and K-Nearest queries *(:star:157 | 0clm)*
  * [seiflotfy/**cuckoofilter**](https://github.com/seiflotfy/cuckoofilter) - Cuckoo Filter: Practically Better Than Bloom *(:star:139 | 0clm)*
  * [derekparker/**trie**](https://github.com/derekparker/trie) - Data structure and relevant algorithms for extremely fast prefix/fuzzy string searching. *(:star:113 | 3clm)*
  * [ryszard/**goskiplist**](https://github.com/ryszard/goskiplist) - A skip list implementation in Go *(:star:102 | 0clm)*
  * [zhenjl/**bloom**](https://github.com/zhenjl/bloom) - Bloom filters implemented in Go. *(:star:79 | 0clm)*
  * [zhenjl/**encoding**](https://github.com/zhenjl/encoding) - Integer Compression Libraries for Go *(:star:40 | 0clm)*
  * [gansidui/**skiplist**](https://github.com/gansidui/skiplist) - skiplist for golang *(:star:26 | 0clm)*
  * [seiflotfy/**count-min-log**](https://github.com/seiflotfy/count-min-log) - Go implementation of Count-Min-Log *(:star:18 | 0clm)*
  * [diegobernardes/**ttlcache**](https://github.com/diegobernardes/ttlcache) - An in-memory LRU string-interface{} map with expiration for golang *(:star:17 | 0clm)*
  * [zhuangsirui/**binpacker**](https://github.com/zhuangsirui/binpacker) - A binary packer and unpacker *(:star:7 | 0clm)*

## Database
*Databases implemented in Go.*

  * [influxdb/**influxdb**](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics *(:star:7644 | 593clm)*
  * [cockroachdb/**cockroach**](https://github.com/cockroachdb/cockroach) - A Scalable, Survivable, Strongly-Consistent SQL Database *(:star:6405 | 393clm)*
  * [golang/**groupcache**](https://github.com/golang/groupcache) - groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. *(:star:4323 | 0clm)*
  * [prometheus/**prometheus**](https://github.com/prometheus/prometheus) - The Prometheus monitoring system and time series database. *(:star:4146 | 65clm)*
  * [boltdb/**bolt**](https://github.com/boltdb/bolt) - An embedded key/value database for Go. *(:star:3846 | 22clm)*
  * [pingcap/**tidb**](https://github.com/pingcap/tidb) - TiDB is a distributed NewSQL database compatible with MySQL protocol  *(:star:3559 | 0clm)*
  * [siddontang/**ledisdb**](https://github.com/siddontang/ledisdb) - a high performance NoSQL powered by Go *(:star:1536 | 5clm)*
  * [HouzuoGuo/**tiedot**](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang *(:star:1483 | 1clm)*
  * [otoolep/**rqlite**](https://github.com/otoolep/rqlite) - Replicating SQLite using the Raft consensus protocol *(:star:1193 | 36clm)*
  * [syndtr/**goleveldb**](https://github.com/syndtr/goleveldb) - LevelDB key/value database in Go. *(:star:977 | 5clm)*
  * [dgraph-io/**dgraph**](https://github.com/dgraph-io/dgraph) - Scalable, Distributed, Low Latency Graph Database *(:star:829 | 0clm)*
  * [pmylund/**go-cache**](https://github.com/pmylund/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. *(:star:450 | 0clm)*
  * [peterbourgon/**diskv**](https://github.com/peterbourgon/diskv) - A disk-backed key-value store. *(:star:338 | 0clm)*
  * [jmhodges/**levigo**](https://github.com/jmhodges/levigo) - levigo is a Go wrapper for LevelDB *(:star:266 | 0clm)*
  * [bluele/**gcache**](https://github.com/bluele/gcache) - Cache library for golang. It supports expirable Cache, LFU, LRU and ARC. *(:star:91 | 0clm)*
  * [muesli/**cache2go**](https://github.com/muesli/cache2go) - Simple go caching library with expiration capabilities and access counters *(:star:66 | 0clm)*
  * [couchbase/**goforestdb**](https://github.com/couchbase/goforestdb) - Go bindings for ForestDB *(:star:18 | 0clm)*
  * [codingsince1985/**couchcache**](https://github.com/codingsince1985/couchcache) - A RESTful caching micro-service backed by Couchbase server *(:star:15 | 14clm)*

*Database tools.*

  * [sosedoff/**pgweb**](https://github.com/sosedoff/pgweb) - Cross-platform client for PostgreSQL databases *(:star:3329 | 12clm)*
  * [youtube/**vitess**](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. *(:star:3329 | 221clm)*
  * [flike/**kingshard**](https://github.com/flike/kingshard) - A high-performance MySQL proxy *(:star:1578 | 0clm)*
  * [mattes/**migrate**](https://github.com/mattes/migrate) - Database migration handling in Golang *(:star:543 | 0clm)*
  * [rubenv/**sql-migrate**](https://github.com/rubenv/sql-migrate) - SQL schema migration tool for Go. *(:star:425 | 5clm)*
  * [outbrain/**orchestrator**](https://github.com/outbrain/orchestrator) - MySQL replication topology manager/visualizer *(:star:389 | 27clm)*
  * [siddontang/**go-mysql-elasticsearch**](https://github.com/siddontang/go-mysql-elasticsearch) - Sync MySQL data into elasticsearch  *(:star:233 | 9clm)*
  * [siddontang/**go-mysql**](https://github.com/siddontang/go-mysql) - a powerful mysql toolset with Go *(:star:189 | 42clm)*
  * [2tvenom/**myreplication**](https://github.com/2tvenom/myreplication) - Golang MySql binary log replication listener *(:star:47 | 0clm)*
  * [pravasan/**pravasan**](https://github.com/pravasan/pravasan) - Simple Migration Tool - written in Go *(:star:10 | 0clm)*
  * [steinbacher/**goose**](https://github.com/steinbacher/goose) -  *(:star:4 | 0clm)*

*SQL query builder, libraries for building and using SQL.*

  * [Masterminds/**squirrel**](https://github.com/Masterminds/squirrel) - Fluent SQL generation for golang *(:star:768 | 0clm)*
  * [knq/**xo**](https://github.com/knq/xo) - Command line tool to generate idiomatic Go code for SQL databases supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server *(:star:331 | 0clm)*
  * [mgutz/**dat**](https://github.com/mgutz/dat) - Go Postgres Data Access Toolkit *(:star:274 | 11clm)*
  * [doug-martin/**goqu**](https://github.com/doug-martin/goqu) - SQL builder and query library for golang *(:star:246 | 0clm)*
  * [gchaincl/**dotsql**](https://github.com/gchaincl/dotsql) - A Golang library for using SQL. *(:star:221 | 1clm)*
  * [go-ozzo/**ozzo-dbx**](https://github.com/go-ozzo/ozzo-dbx) - A Go (golang) package that enhances the standard database/sql package by providing powerful data retrieval methods as well as DB-agnostic query building capabilities. *(:star:78 | 0clm)*
  * [variadico/**scaneo**](https://github.com/variadico/scaneo) - Generate Go code to convert database rows into arbitrary structs. *(:star:61 | 0clm)*
  * [elgris/**sqrl**](https://github.com/elgris/sqrl) - Fluent SQL generation for golang *(:star:34 | 0clm)*

## Database Drivers
*Libraries for connecting and operating databases.*
* Relational Databases

  * [go-sql-driver/**mysql**](https://github.com/go-sql-driver/mysql) - Go MySQL Driver is a lightweight and fast MySQL driver for Go's (golang) database/sql package *(:star:2185 | 2clm)*
  * [lib/**pq**](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql *(:star:1974 | 7clm)*
  * [mattn/**go-sqlite3**](https://github.com/mattn/go-sqlite3) - sqlite3 driver for go that using database/sql *(:star:1140 | 10clm)*
  * [jackc/**pgx**](https://github.com/jackc/pgx) - PostgreSQL client library for Go *(:star:519 | 16clm)*
  * [denisenkom/**go-mssqldb**](https://github.com/denisenkom/go-mssqldb) - Microsoft SQL server driver written in go language *(:star:307 | 8clm)*
  * [mattn/**go-oci8**](https://github.com/mattn/go-oci8) - oracle driver for go that using database/sql *(:star:123 | 5clm)*
  * [rounds/**go-bqstreamer**](https://github.com/rounds/go-bqstreamer) - Stream data into Google BigQuery concurrently using InsertAll() *(:star:73 | 11clm)*
  * [nakagami/**firebirdsql**](https://github.com/nakagami/firebirdsql) - Firebird RDBMS sql driver for Go (golang) *(:star:40 | 2clm)*
  * [minus5/**gofreetds**](https://github.com/minus5/gofreetds) - Go Sql Server database driver. *(:star:33 | 1clm)*
  * [mattn/**go-adodb**](https://github.com/mattn/go-adodb) - Microsoft ActiveX Object DataBase driver for go that using exp/sql *(:star:33 | 0clm)*

* NoSQL Databases

  * [google/**cayley**](https://github.com/google/cayley) - An open-source graph database *(:star:7230 | 54clm)*
  * [garyburd/**redigo**](https://github.com/garyburd/redigo) - Go client for Redis *(:star:2036 | 0clm)*
  * [dancannon/**gorethink**](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB *(:star:820 | 44clm)*
  * [go-redis/**redis**](https://github.com/go-redis/redis) - Redis client for Golang. *(:star:725 | 30clm)*
  * [hoisie/**redis**](https://github.com/hoisie/redis) - A simple, powerful Redis client for Go *(:star:498 | 0clm)*
  * [jmcvetta/**neoism**](https://github.com/jmcvetta/neoism) - Neo4j client for Golang *(:star:233 | 2clm)*
  * [couchbase/**go-couchbase**](https://github.com/couchbase/go-couchbase) - Couchbase client in Go *(:star:203 | 8clm)*
  * [couchbase/**gocb**](https://github.com/couchbase/gocb) - The Couchbase Go SDK *(:star:165 | 0clm)*
  * [aerospike/**aerospike-client-go**](https://github.com/aerospike/aerospike-client-go) - Aerospike Client Go  *(:star:159 | 15clm)*
  * [davemeehan/**Neo4j-GO**](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang *(:star:60 | 0clm)*
  * [bsm/**redeo**](https://github.com/bsm/redeo) - High-performance framework for building redis-protocol compatible TCP servers/services *(:star:35 | 0clm)*
  * [fjl/**go-couchdb**](https://github.com/fjl/go-couchdb) - Yet another CouchDB HTTP API wrapper for Go *(:star:30 | 0clm)*
  * [solher/**arangolite**](https://github.com/solher/arangolite) - Lightweight golang driver for ArangoDB *(:star:22 | 0clm)*
  * [underarmour/**dynago**](https://github.com/underarmour/dynago) - A DynamoDB client for Go *(:star:13 | 0clm)*
  * [cihangir/**neo4j**](https://github.com/cihangir/neo4j) - Neo4j Rest API Client for Go lang *(:star:12 | 15clm)*

* Search and Analytic Databases

  * [blevesearch/**bleve**](https://github.com/blevesearch/bleve) - A modern text indexing library for go *(:star:2382 | 38clm)*
  * [mattbaird/**elastigo**](https://github.com/mattbaird/elastigo) - A Go (golang) based Elasticsearch client library. *(:star:705 | 41clm)*
  * [olivere/**elastic**](https://github.com/olivere/elastic) - Elasticsearch client for Go. *(:star:568 | 18clm)*
  * [belogik/**goes**](https://github.com/belogik/goes) - A library to interact with Elasticsearch in Go! *(:star:70 | 0clm)*
  * [seiflotfy/**skizze**](https://github.com/seiflotfy/skizze) - A probabilistic data structure service and storage *(:star:17 | 0clm)*

## Date & Time
*Libraries for working with dates and times.*

  * [jinzhu/**now**](https://github.com/jinzhu/now) - Now is a time toolkit for golang *(:star:572 | 0clm)*
  * [leekchan/**timeutil**](https://github.com/leekchan/timeutil) - timeutil - useful extensions (Timedelta, Strftime, ...) to the golang's time package *(:star:98 | 0clm)*
  * [grsmv/**goweek**](https://github.com/grsmv/goweek) - ISO 8601 compatible library for working with week entities for Go *(:star:4 | 0clm)*
  * [yaa110/**go-persian-calendar**](https://github.com/yaa110/go-persian-calendar) - The implementation of the Persian (Solar Hijri) Calendar in Go (golang) *(:star:3 | 0clm)*
  * [kirillDanshin/**nulltime**](https://github.com/kirillDanshin/nulltime) -  *(:star:0 | 0clm)*

## Distributed Systems
*Packages that help with building Distributed Systems.*

  * [grpc/**grpc-go**](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC *(:star:1576 | 85clm)*
  * [micro/**micro**](https://github.com/micro/micro) - A microservice toolkit *(:star:1113 | 13clm)*
  * [anacrolix/**torrent**](https://github.com/anacrolix/torrent) - Full-featured BitTorrent client package and utilities *(:star:1090 | 57clm)*
  * [chrislusf/**glow**](https://github.com/chrislusf/glow) - Glow is an easy-to-use distributed computation system written in Go, similar to Hadoop Map Reduce, Spark, Flink, Storm, etc. *(:star:807 | 0clm)*
  * [hashicorp/**raft**](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol *(:star:538 | 2clm)*
  * [valyala/**gorpc**](https://github.com/valyala/gorpc) - Simple, fast and scalable golang rpc library for high load *(:star:209 | 7clm)*
  * [Sioro-Neoku/**go-peerflix**](https://github.com/Sioro-Neoku/go-peerflix) - Go Peerflix *(:star:135 | 0clm)*
  * [dgryski/**go-jump**](https://github.com/dgryski/go-jump) - go-jump: Jump consistent hashing *(:star:61 | 0clm)*
  * [vectaport/**flowgraph**](https://github.com/vectaport/flowgraph) - Ready-send coordination layer on top of goroutines. *(:star:10 | 95clm)*
  * [svcavallar/**celeriac.v1**](https://github.com/svcavallar/celeriac.v1) - Golang client library for adding support for interacting and monitoring Celery workers, tasks and events. *(:star:10 | 0clm)*

## Email
*Libraries that implement email creation and sending*

  * [mailhog/**MailHog**](https://github.com/mailhog/MailHog) - Web and API based SMTP testing *(:star:726 | 6clm)*
  * [jordan-wright/**email**](https://github.com/jordan-wright/email) - Robust and flexible email library for Go *(:star:660 | 0clm)*
  * [sendgrid/**sendgrid-go**](https://github.com/sendgrid/sendgrid-go) - SendGrid Library to Interface through Go *(:star:164 | 0clm)*
  * [hectane/**hectane**](https://github.com/hectane/hectane) - Lightweight SMTP client written in Go *(:star:43 | 0clm)*
  * [aymerick/**douceur**](https://github.com/aymerick/douceur) - A simple CSS parser and inliner in Go *(:star:37 | 33clm)*
  * [mailhog/**smtp**](https://github.com/mailhog/smtp) - MailHog SMTP Protocol *(:star:19 | 0clm)*
  * [toorop/**go-dkim**](https://github.com/toorop/go-dkim) - DKIM package for golang *(:star:13 | 0clm)*

## Embeddable Scripting Languages
*Embedding other languages inside your go code*

  * [robertkrimen/**otto**](https://github.com/robertkrimen/otto) - A JavaScript interpreter in Go (golang) *(:star:2032 | 3clm)*
  * [yuin/**gopher-lua**](https://github.com/yuin/gopher-lua) - GopherLua: VM and compiler for Lua in Go *(:star:1117 | 9clm)*
  * [Shopify/**go-lua**](https://github.com/Shopify/go-lua) - A Lua VM in Go *(:star:620 | 0clm)*
  * [sbinet/**go-python**](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API *(:star:353 | 0clm)*
  * [jcla1/**gisp**](https://github.com/jcla1/gisp) - Simple LISP in Go *(:star:318 | 0clm)*
  * [olebedev/**go-duktape**](https://github.com/olebedev/go-duktape) - Duktape JavaScript engine bindings for Go *(:star:299 | 2clm)*
  * [aarzilli/**golua**](https://github.com/aarzilli/golua) - Go bindings for Lua C API - in progress *(:star:260 | 0clm)*
  * [mattn/**anko**](https://github.com/mattn/anko) - Scriptable interpreter written in golang *(:star:247 | 1clm)*
  * [PuerkitoBio/**agora**](https://github.com/PuerkitoBio/agora) - a dynamically typed, garbage collected, embeddable programming language built with Go *(:star:214 | 0clm)*
  * [deuill/**go-php**](https://github.com/deuill/go-php) - PHP bindings for the Go programming language (Golang) *(:star:107 | 0clm)*
  * [ian-kent/**purl**](https://github.com/ian-kent/purl) - Perl, but fluffy like a cat! *(:star:11 | 0clm)*

## Financial
*Packages for accounting and finance*

  * [shopspring/**decimal**](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers in go *(:star:248 | 13clm)*
  * [leekchan/**accounting**](https://github.com/leekchan/accounting) - money and currency formatting for golang *(:star:191 | 0clm)*

## Forms
*Libraries for working with forms.*

  * [justinas/**nosurf**](https://github.com/justinas/nosurf) - CSRF protection middleware for Go. *(:star:624 | 3clm)*
  * [mholt/**binding**](https://github.com/mholt/binding) - Reflectionless data binding for Go's net/http (not yet a stable 1.0, but not likely to change much either) *(:star:446 | 0clm)*
  * [gorilla/**csrf**](https://github.com/gorilla/csrf) - gorilla/csrf provides Cross Site Request Forgery (CSRF) prevention middleware for Go web applications & services. *(:star:97 | 0clm)*
  * [albrow/**forms**](https://github.com/albrow/forms) - A lightweight go library for parsing form data or json from an http.Request. *(:star:56 | 0clm)*
  * [monoculum/**formam**](https://github.com/monoculum/formam) - a package for decode form's values into struct in Go *(:star:45 | 0clm)*
  * [robfig/**bind**](https://github.com/robfig/bind) -  *(:star:10 | 0clm)*
  * [leebenson/**conform**](https://github.com/leebenson/conform) - Trims, sanitizes & scrubs data based on struct tags (go, golang) *(:star:9 | 0clm)*

## Game Development
*Awesome game development libraries.*

  * [xtaci/**gonet**](https://github.com/xtaci/gonet) - a game server skeleton in golang *(:star:623 | 0clm)*
  * [JoelOtter/**termloop**](https://github.com/JoelOtter/termloop) - Terminal-based game engine for Go, built on top of Termbox *(:star:601 | 0clm)*
  * [name5566/**leaf**](https://github.com/name5566/leaf) - A game server framework in Go (golang) *(:star:584 | 3clm)*
  * [veandco/**go-sdl2**](https://github.com/veandco/go-sdl2) - SDL2 binding for Go *(:star:318 | 28clm)*
  * [vova616/**GarageEngine**](https://github.com/vova616/GarageEngine) - Game engine written in Go (golang). *(:star:213 | 0clm)*
  * [beefsack/**go-astar**](https://github.com/beefsack/go-astar) - Go implementation of the A* search algorithm *(:star:141 | 0clm)*
  * [ungerik/**go3d**](https://github.com/ungerik/go3d) - A performance oriented 2D/3D math package for Go *(:star:99 | 4clm)*
  * [paked/**engi**](https://github.com/paked/engi) - A cross-platform game engine written in Go following my interpretation of the Entity Component System paradigm, A fork of ajhager/engi *(:star:93 | 0clm)*
  * [runningwild/**glop**](https://github.com/runningwild/glop) - Bare-bones osx alternative to sdl *(:star:71 | 0clm)*
  * [GlenKelley/**go-collada**](https://github.com/GlenKelley/go-collada) - Go package for working with the Collada file format. *(:star:8 | 0clm)*

## Generation & Generics
*Tools to enhance the language with features like generics via code generation*

  * [ahmetalpbalkan/**go-linq**](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go *(:star:735 | 2clm)*
  * [clipperhouse/**gen**](https://github.com/clipperhouse/gen) - Type-driven code generation for Go *(:star:662 | 0clm)*
  * [rjeczalik/**interfaces**](https://github.com/rjeczalik/interfaces) - Code generation tools for Go. *(:star:66 | 0clm)*
  * [ungerik/**pkgreflect**](https://github.com/ungerik/pkgreflect) - A Go preprocessor for package scoped reflection *(:star:29 | 0clm)*

## Go Compilers
*Tools for compiling Go to other languages*

  * [gopherjs/**gopherjs**](https://github.com/gopherjs/gopherjs) - A compiler from Go to JavaScript for running Go code in a browser *(:star:3665 | 22clm)*
  * [go-llvm/**llgo**](https://github.com/go-llvm/llgo) - LLVM-based compiler for Go *(:star:730 | 0clm)*
  * [tardisgo/**tardisgo**](https://github.com/tardisgo/tardisgo) - Golang->Haxe->CPP/CSharp/Java/JavaScript transpiler   *(:star:309 | 12clm)*

## Goroutines
*Tools for managing and working with Goroutines*

  * [Jeffail/**tunny**](https://github.com/Jeffail/tunny) - A goroutine pool for golang *(:star:287 | 0clm)*
  * [ivpusic/**grpool**](https://github.com/ivpusic/grpool) - Lightweight Goroutine pool *(:star:94 | 0clm)*
  * [go-playground/**pool**](https://github.com/go-playground/pool) - :speedboat: Go consumer goroutine pool for easy goroutine handling + time saving *(:star:37 | 0clm)*

## GUI
*Libraries for building GUI Applications*

  * [andlabs/**ui**](https://github.com/andlabs/ui) - Platform-native GUI library for Go. *(:star:2735 | 275clm)*
  * [go-qml/**qml**](https://github.com/go-qml/qml) - QML support for the Go language *(:star:1622 | 0clm)*
  * [lxn/**walk**](https://github.com/lxn/walk) - A Windows GUI toolkit for the Go Programming Language *(:star:1232 | 0clm)*
  * [visualfc/**goqt**](https://github.com/visualfc/goqt) - Golang bindings to the Qt cross-platform application framework. *(:star:943 | 0clm)*
  * [deckarep/**gosx-notifier**](https://github.com/deckarep/gosx-notifier) - gosx-notifier is a Go framework for sending desktop notifications to OSX 10.8 or higher *(:star:275 | 0clm)*
  * [oskca/**sciter**](https://github.com/oskca/sciter) - Golang bindings of Sciter: the Embeddable HTML/CSS/script engine for modern UI development *(:star:203 | 0clm)*
  * [gotk3/**gotk3**](https://github.com/gotk3/gotk3) - Go bindings for GTK3 *(:star:60 | 6clm)*
  * [getlantern/**systray**](https://github.com/getlantern/systray) - a cross platfrom Go library to place an icon and menu in the notification area *(:star:51 | 0clm)*
  * [shurcooL/**trayhost**](https://github.com/shurcooL/trayhost) - Cross-platform Go library to place an icon in the host operating system's taskbar. *(:star:35 | 0clm)*

## Images
*Libraries for manipulating images.*

  * [nfnt/**resize**](https://github.com/nfnt/resize) - Pure golang image resizing  *(:star:1021 | 8clm)*
  * [pravj/**geopattern**](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string in golang. *(:star:842 | 0clm)*
  * [disintegration/**imaging**](https://github.com/disintegration/imaging) - Simple Go image processing package *(:star:774 | 0clm)*
  * [h2non/**imaginary**](https://github.com/h2non/imaginary) - Fast HTTP microservice for high-level image processing. Perfectly fitted for Docker and Heroku. *(:star:708 | 70clm)*
  * [disintegration/**gift**](https://github.com/disintegration/gift) - Go Image Filtering Toolkit *(:star:695 | 10clm)*
  * [ajstarks/**svgo**](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation *(:star:661 | 5clm)*
  * [thoas/**picfit**](https://github.com/thoas/picfit) - An image resizing server written in Go *(:star:462 | 42clm)*
  * [gographics/**imagick**](https://github.com/gographics/imagick) - naive Go binding to ImageMagick's MagickWand C API *(:star:394 | 0clm)*
  * [lazywei/**go-opencv**](https://github.com/lazywei/go-opencv) - Go bindings for OpenCV / 2.x API in gocv / 1.x API in opencv *(:star:365 | 0clm)*
  * [muesli/**smartcrop**](https://github.com/muesli/smartcrop) - smartcrop implementation in Go *(:star:189 | 0clm)*
  * [koyachi/**go-nude**](https://github.com/koyachi/go-nude) - Nudity detection with Go. *(:star:170 | 1clm)*
  * [h2non/**bimg**](https://github.com/h2non/bimg) - Small Go package for fast high-level image processing using libvips via C bindings *(:star:147 | 119clm)*
  * [bamiaux/**rez**](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD *(:star:115 | 0clm)*
  * [hawx/**img**](https://github.com/hawx/img) - A selection of image manipulation tools *(:star:73 | 6clm)*
  * [ungerik/**go-cairo**](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library *(:star:50 | 0clm)*
  * [bolknote/**go-gd**](https://github.com/bolknote/go-gd) - Go bingings for GD (http://www.boutell.com/gd/) *(:star:39 | 3clm)*
  * [jyotiska/**go-webcolors**](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go *(:star:17 | 0clm)*
  * [ftrvxmtrx/**tga**](https://github.com/ftrvxmtrx/tga) - Go package for decoding and encoding TARGA image format *(:star:11 | 0clm)*
  * [donatj/**mpo**](https://github.com/donatj/mpo) - Simple Golang JPEG MPO (Multi Picture Object Decoder) *(:star:4 | 24clm)*

## Logging
*Libraries for generating and working with log files.*

  * [Sirupsen/**logrus**](https://github.com/Sirupsen/logrus) - Structured, pluggable logging for Go. *(:star:2473 | 11clm)*
  * [golang/**glog**](https://github.com/golang/glog) - Leveled execution logs for Go *(:star:947 | 0clm)*
  * [cihub/**seelog**](https://github.com/cihub/seelog) - Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting. *(:star:621 | 21clm)*
  * [hpcloud/**tail**](https://github.com/hpcloud/tail) - Go package for reading from continously updated files (tail -f) *(:star:454 | 0clm)*
  * [inconshreveable/**log15**](https://github.com/inconshreveable/log15) - Powerful, composable logging for Go *(:star:407 | 0clm)*
  * [natefinch/**lumberjack**](https://github.com/natefinch/lumberjack) - lumberjack is a rolling logger for Go *(:star:277 | 0clm)*
  * [mgutz/**logxi**](https://github.com/mgutz/logxi) - A 12-factor app logger built for performance and happy development *(:star:237 | 14clm)*
  * [apex/**log**](https://github.com/apex/log) - Structured logging package for Go. *(:star:195 | 0clm)*
  * [apsdehal/**go-logger**](https://github.com/apsdehal/go-logger) - Simple logger for Go programs *(:star:152 | 0clm)*
  * [hashicorp/**logutils**](https://github.com/hashicorp/logutils) - Utilities for slightly better logging in Go (Golang). *(:star:129 | 3clm)*
  * [azer/**logger**](https://github.com/azer/logger) - Minimalistic logging library for Go. *(:star:58 | 0clm)*
  * [rs/**xlog**](https://github.com/rs/xlog) - xlog is a logger for net/context aware HTTP applications *(:star:56 | 0clm)*
  * [firstrow/**logvoyage**](https://github.com/firstrow/logvoyage) - LogVoyage - logging SaaS written in GoLang *(:star:50 | 82clm)*
  * [go-ozzo/**ozzo-log**](https://github.com/go-ozzo/ozzo-log) - A Go (golang) package providing high-performance asynchronous logging, message filtering by severity and category, and multiple message targets. *(:star:36 | 0clm)*
  * [alexcesaro/**log**](https://github.com/alexcesaro/log) - Logging packages for Go *(:star:28 | 0clm)*
  * [chzyer/**logex**](https://github.com/chzyer/logex) - An golang log lib, supports tracking and level, wrap by standard log lib *(:star:22 | 10clm)*
  * [ian-kent/**go-log**](https://github.com/ian-kent/go-log) - A logger, for Go *(:star:17 | 0clm)*
  * [sadlil/**gologger**](https://github.com/sadlil/gologger) - Simple Logger for golang. Logs Into console, file or ElasticSearch. Simple, easy to use.  *(:star:16 | 0clm)*
  * [siddontang/**go-log**](https://github.com/siddontang/go-log) - a golang log lib supports level and multi handlers *(:star:13 | 0clm)*
  * [sebest/**logrusly**](https://github.com/sebest/logrusly) - Loggly Hooks for GO Logrus logger *(:star:6 | 0clm)*
  * [jbrodriguez/**mlog**](https://github.com/jbrodriguez/mlog) -  *(:star:3 | 0clm)*

## Machine Learning
*Libraries for Machine Learning.*

  * [sjwhitworth/**golearn**](https://github.com/sjwhitworth/golearn) - Machine Learning for Go *(:star:2629 | 0clm)*
  * [cdipaolo/**goml**](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go (and so much more) *(:star:453 | 0clm)*
  * [ryanbressler/**CloudForest**](https://github.com/ryanbressler/CloudForest) - Ensembles of decision trees in go/golang. *(:star:338 | 0clm)*
  * [jbrukh/**bayesian**](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang. *(:star:291 | 0clm)*
  * [thoj/**go-galib**](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang *(:star:114 | 0clm)*
  * [MaxHalford/**gago**](https://github.com/MaxHalford/gago) - :turtle: Multi-population, flexible, parallel genetic algorithm. *(:star:99 | 0clm)*
  * [muesli/**regommend**](https://github.com/muesli/regommend) - Recommendation engine for Go *(:star:92 | 0clm)*
  * [eaigner/**shield**](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go *(:star:76 | 0clm)*
  * [goml/**gobrain**](https://github.com/goml/gobrain) - Neural Networks written in go *(:star:74 | 0clm)*
  * [white-pony/**go-fann**](https://github.com/white-pony/go-fann) - Go bindings for FANN, library for artificial neural networks *(:star:71 | 0clm)*
  * [schuyler/**neural-go**](https://github.com/schuyler/neural-go) - A multilayer perceptron network implemented in Go, with training via backpropagation. *(:star:41 | 0clm)*
  * [timkaye11/**goRecommend**](https://github.com/timkaye11/goRecommend) - Collaborative Filtering (CF) Algorithms in Go!  *(:star:34 | 0clm)*
  * [datastream/**libsvm**](https://github.com/datastream/libsvm) - libsvm go version *(:star:34 | 0clm)*
  * [daviddengcn/**go-pr**](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang. *(:star:33 | 0clm)*
  * [tomcraven/**goga**](https://github.com/tomcraven/goga) - Golang Genetic Algorithm *(:star:32 | 0clm)*
  * [danieldk/**golinear**](https://github.com/danieldk/golinear) - liblinear bindings for Go *(:star:23 | 1clm)*
  * [e-dard/**godist**](https://github.com/e-dard/godist) - Probability distributions and associated methods in Go *(:star:6 | 0clm)*
  * [ThePaw/**probab**](https://github.com/ThePaw/probab) - Automatically exported from code.google.com/p/probab *(:star:1 | 0clm)*
  * [NullHypothesis/**mlgo**](https://github.com/NullHypothesis/mlgo) - Automatically exported from code.google.com/p/mlgo *(:star:0 | 0clm)*

## Messaging
*Libraries that implement messaging systems*

  * [Terry-Mao/**gopush-cluster**](https://github.com/Terry-Mao/gopush-cluster) - Golang push server cluster *(:star:1142 | 0clm)*
  * [centrifugal/**centrifugo**](https://github.com/centrifugal/centrifugo) - Real-time messaging (Websockets or SockJS) server in Go *(:star:958 | 46clm)*
  * [Shopify/**sarama**](https://github.com/Shopify/sarama) - Sarama is a Go library for Apache Kafka 0.8 and 0.9  *(:star:884 | 93clm)*
  * [uniqush/**uniqush-push**](https://github.com/uniqush/uniqush-push) - Uniqush is a free and open source software which provides a unified push service for server-side notification to apps on mobile devices. *(:star:797 | 0clm)*
  * [go-mangos/**mangos**](https://github.com/go-mangos/mangos) - package mangos is an implementation in pure Go of the SP ("Scalable Protocols") protocols. *(:star:794 | 5clm)*
  * [RichardKnop/**machinery**](https://github.com/RichardKnop/machinery) - Machinery is an asynchronous task queue/job queue based on distributed message passing. *(:star:762 | 34clm)*
  * [nats-io/**nats**](https://github.com/nats-io/nats) - Golang client for NATS, the cloud native messaging system. *(:star:626 | 1clm)*
  * [nsqio/**go-nsq**](https://github.com/nsqio/go-nsq) - The official Go package for NSQ *(:star:549 | 20clm)*
  * [pebbe/**zmq4**](https://github.com/pebbe/zmq4) - A Go interface to ZeroMQ version 4 *(:star:362 | 0clm)*
  * [asaskevich/**EventBus**](https://github.com/asaskevich/EventBus) - [Go] Lightweight eventbus with async compatibility for Go *(:star:131 | 2clm)*
  * [tuxychandru/**pubsub**](https://github.com/tuxychandru/pubsub) - A simple pubsub package for go. *(:star:81 | 0clm)*
  * [godbus/**dbus**](https://github.com/godbus/dbus) - Native Go bindings for D-Bus *(:star:78 | 2clm)*
  * [dailymotion/**oplog**](https://github.com/dailymotion/oplog) - A generic oplog/replication system for microservices *(:star:64 | 0clm)*
  * [olebedev/**emitter**](https://github.com/olebedev/emitter) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins *(:star:43 | 0clm)*
  * [TheCreeper/**go-notify**](https://github.com/TheCreeper/go-notify) - Package notify provides an implementation of the Gnome DBus Notifications Specification. *(:star:12 | 8clm)*
  * [ventu-io/**go-longpoll**](https://github.com/ventu-io/go-longpoll) - PubSub queuing with long-polling subscribers for web and other distributed applications *(:star:7 | 0clm)*

## Miscellaneous
*These libraries were placed here because none of the other categories seemed to fit*

  * [shirou/**gopsutil**](https://github.com/shirou/gopsutil) - psutil for golang *(:star:959 | 2clm)*
  * [haxpax/**gosms**](https://github.com/haxpax/gosms) - :mailbox_closed: Your own local SMS gateway in Go *(:star:906 | 0clm)*
  * [albrow/**jobs**](https://github.com/albrow/jobs) - A persistent and flexible background jobs library for go. *(:star:314 | 7clm)*
  * [huandu/**xstrings**](https://github.com/huandu/xstrings) - xstrings: A collection of useful string functions for Go. *(:star:308 | 0clm)*
  * [jolestar/**go-commons-pool**](https://github.com/jolestar/go-commons-pool) - a generic object pool for golang *(:star:223 | 0clm)*
  * [hashicorp/**go-multierror**](https://github.com/hashicorp/go-multierror) - A Go (golang) package for representing a list of errors as a single error. *(:star:185 | 0clm)*
  * [dimiro1/**health**](https://github.com/dimiro1/health) - An easy to use, extensible health check library for Go applications. *(:star:147 | 0clm)*
  * [rjeczalik/**notify**](https://github.com/rjeczalik/notify) - File system event notification library on steroids. *(:star:141 | 0clm)*
  * [pariz/**gountries**](https://github.com/pariz/gountries) - Gountries provides: Countries (ISO-3166-1), Country Subdivisions(ISO-3166-2), Currencies (ISO 4217), Geo Coordinates(ISO-6709) as well as translations, country borders and other stuff exposed as struct data. *(:star:83 | 0clm)*
  * [ventu-io/**go-shortid**](https://github.com/ventu-io/go-shortid) - Super short, unique, non sequential, URL friendly Ids for Go *(:star:46 | 0clm)*
  * [go-chat-bot/**bot**](https://github.com/go-chat-bot/bot) - IRC, Slack and Telegram bot written in go *(:star:34 | 0clm)*
  * [zhulik/**margelet**](https://github.com/zhulik/margelet) - Telegram Bot Framework for Go *(:star:26 | 0clm)*
  * [digitalcrab/**browscap_go**](https://github.com/digitalcrab/browscap_go) - GoLang Library for Browser Capabilities Project *(:star:15 | 0clm)*
  * [go-playground/**stats**](https://github.com/go-playground/stats) - :chart_with_upwards_trend: Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... *(:star:15 | 0clm)*
  * [artyom/**autoflags**](https://github.com/artyom/autoflags) - Populate go command line app flags from config struct *(:star:12 | 0clm)*
  * [go-xkg/**xkg**](https://github.com/go-xkg/xkg) - X Keyboard Grabber *(:star:10 | 0clm)*
  * [txgruppi/**werr**](https://github.com/txgruppi/werr) - Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. *(:star:4 | 0clm)*
  * [miolini/**datacounter**](https://github.com/miolini/datacounter) - Golang counters for readers/writers *(:star:4 | 0clm)*

## Natural Language Processing
*Libraries for working with human languages.*

  * [neurosnap/**sentences**](https://github.com/neurosnap/sentences) - A multilingual command line sentence tokenizer in Golang *(:star:119 | 0clm)*
  * [awsong/**MMSEGO**](https://github.com/awsong/MMSEGO) - Chinese word splitting algorithm MMSEG in GO *(:star:46 | 0clm)*
  * [nuance/**go-nlp**](https://github.com/nuance/go-nlp) - Utilities for working with discrete probability distributions and other tools useful for doing NLP work *(:star:44 | 0clm)*
  * [pebbe/**textcat**](https://github.com/pebbe/textcat) - A Go package for n-gram based text categorization, with support for utf-8 and raw text *(:star:40 | 0clm)*
  * [fiam/**gounidecode**](https://github.com/fiam/gounidecode) - Unicode transliterator for #golang *(:star:36 | 0clm)*
  * [agonopol/**go-stem**](https://github.com/agonopol/go-stem) - Word Stemming in Go *(:star:34 | 0clm)*
  * [dchest/**stemmer**](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers. *(:star:28 | 0clm)*
  * [zhenjl/**porter2**](https://github.com/zhenjl/porter2) - High Performance Porter2 Stemmer *(:star:21 | 0clm)*
  * [blevesearch/**segment**](https://github.com/blevesearch/segment) - A Go library for performing Unicode Text Segmentation as described in Unicode Standard Annex #29 *(:star:16 | 0clm)*
  * [goodsign/**icu**](https://github.com/goodsign/icu) - Cgo binding for icu4c library *(:star:15 | 0clm)*
  * [goodsign/**snowball**](https://github.com/goodsign/snowball) - Cgo binding for Snowball C library *(:star:13 | 0clm)*
  * [rookii/**paicehusk**](https://github.com/rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm *(:star:11 | 0clm)*
  * [rjohnsondev/**golibstemmer**](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2 *(:star:10 | 0clm)*
  * [goodsign/**libtextcat**](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library *(:star:8 | 0clm)*
  * [danieldk/**go2vec**](https://github.com/danieldk/go2vec) - Read and use word2vec vectors in Go *(:star:6 | 1clm)*
  * [a2800276/**porter**](https://github.com/a2800276/porter) - porter stemmer *(:star:3 | 0clm)*
  * [ThePaw/**go-eco**](https://github.com/ThePaw/go-eco) - Automatically exported from code.google.com/p/go-eco *(:star:1 | 0clm)*

## Networking
*Libraries for working with various layers of the network*

  * [valyala/**fasthttp**](https://github.com/valyala/fasthttp) - Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http *(:star:1904 | 0clm)*
  * [miekg/**dns**](https://github.com/miekg/dns) - DNS library in Go *(:star:1418 | 12clm)*
  * [google/**gopacket**](https://github.com/google/gopacket) - Provides packet processing capabilities for Go *(:star:526 | 21clm)*
  * [pkg/**sftp**](https://github.com/pkg/sftp) - SFTP support for the go.crypto/ssh package *(:star:231 | 0clm)*
  * [hashicorp/**mdns**](https://github.com/hashicorp/mdns) - Simple mDNS client/server library in Golang *(:star:218 | 0clm)*
  * [hashicorp/**go-getter**](https://github.com/hashicorp/go-getter) - Package for downloading things from a string URL using a variety of protocols. *(:star:218 | 0clm)*
  * [akrennmair/**gopcap**](https://github.com/akrennmair/gopcap) - A simple wrapper around libpcap for the Go programming language *(:star:205 | 0clm)*
  * [gansidui/**gotcp**](https://github.com/gansidui/gotcp) - A Go package for quickly building tcp servers *(:star:139 | 0clm)*
  * [stabbycutyou/**buffstreams**](https://github.com/stabbycutyou/buffstreams) - A library to simplify writing applications using TCP sockets to stream protobuff messages *(:star:125 | 0clm)*
  * [jlaffaye/**ftp**](https://github.com/jlaffaye/ftp) - FTP client package for Go *(:star:118 | 0clm)*
  * [soniah/**gosnmp**](https://github.com/soniah/gosnmp) - An SNMP library written in GoLang. *(:star:101 | 0clm)*
  * [ccding/**go-stun**](https://github.com/ccding/go-stun) - a go implementation of the STUN client (RFC 3489 and RFC 5389) *(:star:71 | 0clm)*
  * [eduardonunesp/**sslb**](https://github.com/eduardonunesp/sslb) - Golang Super Simple Load Balance *(:star:60 | 0clm)*
  * [anacrolix/**utp**](https://github.com/anacrolix/utp) - Implements uTP, the micro transport protocol as used with Bittorrent *(:star:45 | 9clm)*
  * [firstrow/**tcp_server**](https://github.com/firstrow/tcp_server) - GoLang simple TCP server *(:star:36 | 1clm)*
  * [mdlayher/**arp**](https://github.com/mdlayher/arp) - Package arp implements the ARP protocol, as described in RFC 826. MIT Licensed. *(:star:31 | 0clm)*
  * [zubairhamed/**canopus**](https://github.com/zubairhamed/canopus) - CoAP Client/Server implementing RFC 7252 (For Go/Golang) *(:star:24 | 23clm)*
  * [aybabtme/**portproxy**](https://github.com/aybabtme/portproxy) - TCP proxy, highjacks HTTP to allow CORS *(:star:18 | 0clm)*
  * [cavaliercoder/**grab**](https://github.com/cavaliercoder/grab) - Go package for downloading large files from the internet *(:star:17 | 0clm)*
  * [ian-kent/**linkio**](https://github.com/ian-kent/linkio) - Simulate network link speed *(:star:16 | 0clm)*
  * [mdlayher/**raw**](https://github.com/mdlayher/raw) - Package raw enables reading and writing data at the device driver level for a network interface.  MIT Licensed. *(:star:16 | 0clm)*
  * [mdlayher/**ethernet**](https://github.com/mdlayher/ethernet) - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. *(:star:9 | 0clm)*
  * [koofr/**graval**](https://github.com/koofr/graval) - An experimental go FTP server framework *(:star:8 | 1clm)*
  * [mdlayher/**dhcp6**](https://github.com/mdlayher/dhcp6) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. MIT Licensed. *(:star:7 | 0clm)*
  * [sunwxg/**golibwireshark**](https://github.com/sunwxg/golibwireshark) -  *(:star:3 | 0clm)*
  * [sunwxg/**goshark**](https://github.com/sunwxg/goshark) -  *(:star:2 | 0clm)*
  * [kirillDanshin/**llb**](https://github.com/kirillDanshin/llb) -  *(:star:1 | 0clm)*

## OpenGL
*Libraries for using OpenGL in Go.*

  * [go-gl/**glfw**](https://github.com/go-gl/glfw) - Go bindings for GLFW 3 *(:star:220 | 3clm)*
  * [go-gl/**gl**](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow) *(:star:182 | 0clm)*
  * [go-gl/**mathgl**](https://github.com/go-gl/mathgl) - A pure Go 3D math library. *(:star:111 | 4clm)*
  * [goxjs/**gl**](https://github.com/goxjs/gl) - Go cross-platform OpenGL bindings. *(:star:64 | 0clm)*
  * [goxjs/**glfw**](https://github.com/goxjs/glfw) - Go cross-platform glfw library for creating an OpenGL context and receiving events. *(:star:23 | 5clm)*

## ORM
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

  * [jinzhu/**gorm**](https://github.com/jinzhu/gorm) - The fantastic ORM library for Golang, aims to be developer friendly *(:star:3530 | 24clm)*
  * [go-gorp/**gorp**](https://github.com/go-gorp/gorp) - Go Relational Persistence - an ORM-ish library for Go *(:star:2000 | 10clm)*
  * [go-xorm/**xorm**](https://github.com/go-xorm/xorm) - Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle *(:star:1106 | 8clm)*
  * [coocood/**qbs**](https://github.com/coocood/qbs) - QBS stands for Query By Struct. A Go ORM. *(:star:380 | 0clm)*
  * [upper/**db**](https://github.com/upper/db) - Magic-free ORM-like package for Go. *(:star:319 | 0clm)*
  * [albrow/**zoom**](https://github.com/albrow/zoom) - A blazing-fast datastore and querying engine for Go built on Redis. *(:star:127 | 34clm)*
  * [gosuri/**go-store**](https://github.com/gosuri/go-store) - A simple and fast Redis backed key-value store library for Go *(:star:70 | 26clm)*
  * [cosiner/**gomodel**](https://github.com/cosiner/gomodel) - A lightweight, fast, orm-like library helps interactive with database *(:star:46 | 8clm)*

## Package Management
*Libraries for package and dependency management.*

  * [tools/**godep**](https://github.com/tools/godep) - dependency tool for go *(:star:3262 | 0clm)*
  * [Masterminds/**glide**](https://github.com/Masterminds/glide) - Package Management for Golang *(:star:1501 | 0clm)*
  * [mattn/**gom**](https://github.com/mattn/gom) - Go Manager - bundle for go *(:star:1046 | 0clm)*
  * [pote/**gpm**](https://github.com/pote/gpm) - Barebones dependency manager for Go. *(:star:925 | 1clm)*
  * [gpmgo/**gopm**](https://github.com/gpmgo/gopm) - Go Package Manager (gopm) is a package manager and build tool for Go. *(:star:918 | 1clm)*
  * [nitrous-io/**goop**](https://github.com/nitrous-io/goop) - A simple dependency manager for Go (golang), inspired by Bundler. *(:star:753 | 0clm)*
  * [jingweno/**nut**](https://github.com/jingweno/nut) - Vendor Go dependencies *(:star:245 | 4clm)*
  * [VividCortex/**johnny-deps**](https://github.com/VividCortex/johnny-deps) - Barebones dependency manager for Go. *(:star:215 | 0clm)*
  * [LyricalSecurity/**gigo**](https://github.com/LyricalSecurity/gigo) - GIGO: PIP for GO *(:star:190 | 0clm)*
  * [DamnWidget/**VenGO**](https://github.com/DamnWidget/VenGO) - Create and manage Isolated Virtual Environments for Go *(:star:84 | 0clm)*

## Query Language

  * [chris-ramon/**graphql-go**](https://github.com/chris-ramon/graphql-go) - An implementation of GraphQL for Go / Golang *(:star:353 | 0clm)*
  * [elgs/**jsonql**](https://github.com/elgs/jsonql) - JSON query expression library in Golang. *(:star:46 | 0clm)*
  * [tmc/**graphql**](https://github.com/tmc/graphql) - graphql parser + utilities *(:star:35 | 11clm)*
  * [sevki/**graphql**](https://github.com/sevki/graphql) - GraphQL implementation in go *(:star:28 | 0clm)*

## Resource Embedding

  * [jteeuwen/**go-bindata**](https://github.com/jteeuwen/go-bindata) - A small utility which generates Go code from any file. Useful for embedding binary data in a Go program. *(:star:1686 | 4clm)*
  * [GeertJohan/**go.rice**](https://github.com/GeertJohan/go.rice) - go.rice is a Go package that makes working with resources such as html,js,css,images,templates, etc very easy. *(:star:683 | 4clm)*
  * [omeid/**go-resources**](https://github.com/omeid/go-resources) - Unfancy resources embedding for Go with out of box http.FileSystem support. *(:star:97 | 11clm)*
  * [shurcooL/**vfsgen**](https://github.com/shurcooL/vfsgen) - Takes an input http.FileSystem (likely at go generate time) and generates Go code that statically implements it. *(:star:73 | 0clm)*
  * [UnnoTed/**fileb0x**](https://github.com/UnnoTed/fileb0x) - simple customizable tool to embed files in go *(:star:39 | 0clm)*
  * [go-playground/**statics**](https://github.com/go-playground/statics) - :file_folder: Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks *(:star:31 | 0clm)*
  * [pyros2097/**go-embed**](https://github.com/pyros2097/go-embed) - Generates go code to embed resource files into your library or executable *(:star:2 | 0clm)*

## Science and Data Analysis
*Libraries for scientific computing and data analyzing.*

  * [nytlabs/**streamtools**](https://github.com/nytlabs/streamtools) - tools for working with streams of data *(:star:1274 | 4clm)*
  * [montanaflynn/**stats**](https://github.com/montanaflynn/stats) - A statistics package with common functions that are missing from the Golang standard library.  *(:star:580 | 0clm)*
  * [mjibson/**go-dsp**](https://github.com/mjibson/go-dsp) - Digital Signal Processing for Go *(:star:378 | 0clm)*
  * [vdobler/**chart**](https://github.com/vdobler/chart) - Provide basic charts in go *(:star:292 | 1clm)*
  * [skelterjohn/**go.matrix**](https://github.com/skelterjohn/go.matrix) - linear algebra for go *(:star:240 | 0clm)*
  * [gonum/**matrix**](https://github.com/gonum/matrix) - Matrix packages for the Go language. *(:star:230 | 10clm)*
  * [gonum/**plot**](https://github.com/gonum/plot) - A repository for plotting and visualizing data *(:star:226 | 15clm)*
  * [gyuho/**goraph**](https://github.com/gyuho/goraph) - Package goraph implements graph, tree data structures and algorithms. *(:star:197 | 30clm)*
  * [VividCortex/**ewma**](https://github.com/VividCortex/ewma) - Exponentially Weighted Moving Average algorithms for Go. *(:star:152 | 0clm)*
  * [ziutek/**blas**](https://github.com/ziutek/blas) - Go implementation of BLAS (Basic Linear Algebra Subprograms) *(:star:90 | 0clm)*
  * [VividCortex/**gohistogram**](https://github.com/VividCortex/gohistogram) - Streaming approximate histograms in Go *(:star:67 | 0clm)*
  * [spate/**vectormath**](https://github.com/spate/vectormath) - Vectormath for Go *(:star:53 | 0clm)*
  * [skelterjohn/**geom**](https://github.com/skelterjohn/geom) - 2d geometry for golang *(:star:32 | 0clm)*
  * [soniah/**evaler**](https://github.com/soniah/evaler) - Implements a simple floating point arithmetic expression evaluator in Go (golang). *(:star:20 | 2clm)*
  * [alixaxel/**pagerank**](https://github.com/alixaxel/pagerank) - Weighted PageRank implementation in Go *(:star:10 | 0clm)*
  * [anschelsc/**gofrac**](https://github.com/anschelsc/gofrac) - A fractions library for go (http://golang.org) *(:star:6 | 0clm)*
  * [ematvey/**gostat**](https://github.com/ematvey/gostat) - Collection of statistical routines in golang *(:star:4 | 0clm)*
  * [varver/**gocomplex**](https://github.com/varver/gocomplex) - Automatically exported from code.google.com/p/gocomplex *(:star:1 | 0clm)*
  * [ThePaw/**go-gt**](https://github.com/ThePaw/go-gt) - Automatically exported from code.google.com/p/go-gt *(:star:1 | 0clm)*
  * [ematvey/**go-fn**](https://github.com/ematvey/go-fn) - Automatically exported from code.google.com/p/go-fn *(:star:1 | 0clm)*
  * [pwil3058/**mudlark-go-pkgs**](https://github.com/pwil3058/mudlark-go-pkgs) - Automatically exported from code.google.com/p/mudlark-go-pkgs *(:star:0 | 0clm)*

## Security
*Libraries that are used to help make your application more secure.*

  * [xenolf/**lego**](https://github.com/xenolf/lego) - Let's Encrypt client and ACME library written in Go *(:star:1033 | 0clm)*
  * [hlandau/**acme**](https://github.com/hlandau/acme) - :lock: acmetool, an automatic certificate acquisition tool for ACME (Let's Encrypt) *(:star:676 | 0clm)*
  * [jaredfolkins/**badactor**](https://github.com/jaredfolkins/badactor) - BadActor.org An in-memory application driven jailer written in Go *(:star:164 | 0clm)*
  * [hlandau/**passlib**](https://github.com/hlandau/passlib) - :key: Idiotproof golang password validation library inspired by Python's passlib *(:star:85 | 0clm)*
  * [elithrar/**simple-scrypt**](https://github.com/elithrar/simple-scrypt) - A convenience library for generating, comparing and inspecting password hashes using the scrypt KDF in Go. *(:star:59 | 6clm)*
  * [hillu/**go-yara**](https://github.com/hillu/go-yara) - Go bindings for YARA *(:star:27 | 1clm)*

## Serialization
*Libraries and tools for binary serialization*

  * [golang/**protobuf**](https://github.com/golang/protobuf) - Go support for Google's protocol buffers *(:star:820 | 2clm)*
  * [mitchellh/**mapstructure**](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures. *(:star:561 | 0clm)*
  * [ugorji/**go**](https://github.com/ugorji/go) - idiomatic codec and rpc lib for msgpack, cbor, json, etc. msgpack.org[Go] *(:star:530 | 0clm)*
  * [gogo/**protobuf**](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets *(:star:354 | 29clm)*
  * [glycerine/**go-capnproto**](https://github.com/glycerine/go-capnproto) - Cap'n Proto library and parser for go. This is go-capnproto-1.0, and does not have rpc. See https://github.com/zombiezen/go-capnproto2 for 2.0 which has rpc and capabilities. *(:star:241 | 1clm)*
  * [glycerine/**bambam**](https://github.com/glycerine/bambam) - auto-generate capnproto schema from your golang source files. Depends on go-capnproto-1.0 at https://github.com/glycerine/go-capnproto *(:star:53 | 7clm)*
  * [yvasiyarov/**php_session_decoder**](https://github.com/yvasiyarov/php_session_decoder) - PHP session encoder/decoder written in Go *(:star:47 | 0clm)*
  * [tuvistavie/**structomap**](https://github.com/tuvistavie/structomap) - Easily and dynamically generate maps from Go static structures *(:star:29 | 0clm)*
  * [PromonLogicalis/**asn1**](https://github.com/PromonLogicalis/asn1) - Asn.1 BER and DER encoding library for golang. *(:star:0 | 0clm)*

## Server Applications

  * [coreos/**etcd**](https://github.com/coreos/etcd) - Distributed reliable key-value store for the most critical data of a distributed system *(:star:9023 | 152clm)*
  * [mholt/**caddy**](https://github.com/mholt/caddy) - Fast, cross-platform HTTP/2 web server with automatic HTTPS *(:star:5072 | 55clm)*
  * [cortesi/**devd**](https://github.com/cortesi/devd) -  A local webserver for developers *(:star:1988 | 0clm)*
  * [minio/**minio**](https://github.com/minio/minio) - Minio is an object storage server compatible with Amazon S3 and licensed under Apache 2.0 License *(:star:782 | 261clm)*
  * [xyproto/**algernon**](https://github.com/xyproto/algernon) - :diamond_shape_with_a_dot_inside: HTTP/2 web/application server with Lua support *(:star:220 | 108clm)*
  * [sci4me/**yakvs**](https://github.com/sci4me/yakvs) - YAKVS (Yet Another Key Value Store) *(:star:5 | 0clm)*

## Template Engines
*Libraries and tools for templating and lexing.*

  * [hoisie/**mustache**](https://github.com/hoisie/mustache) - The mustache template language in Go *(:star:786 | 0clm)*
  * [flosch/**pongo2**](https://github.com/flosch/pongo2) - Django-syntax like template-engine for Go *(:star:678 | 0clm)*
  * [eknkc/**amber**](https://github.com/eknkc/amber) - Amber is an elegant templating engine for Go Programming Language, inspired from HAML and Jade *(:star:570 | 0clm)*
  * [sipin/**gorazor**](https://github.com/sipin/gorazor) - Razor view engine for Golang *(:star:493 | 2clm)*
  * [yosssi/**ace**](https://github.com/yosssi/ace) - HTML template engine for Go *(:star:434 | 2clm)*
  * [benbjohnson/**ego**](https://github.com/benbjohnson/ego) - An ERB-style templating language for Go. *(:star:248 | 2clm)*
  * [valyala/**quicktemplate**](https://github.com/valyala/quicktemplate) - Fast, powerful, yet easy to use template engine for Go. Optimized for speed, zero memory allocations in hot paths. Up to 20x faster than html/template *(:star:224 | 0clm)*
  * [robfig/**soy**](https://github.com/robfig/soy) - Go implementation for Soy templates (Google Closure templates) *(:star:110 | 0clm)*
  * [aymerick/**raymond**](https://github.com/aymerick/raymond) - Handlebars for golang *(:star:77 | 0clm)*
  * [ziutek/**kasia.go**](https://github.com/ziutek/kasia.go) - Templating system for HTML and other text documents - go implementation *(:star:67 | 0clm)*
  * [valyala/**fasttemplate**](https://github.com/valyala/fasttemplate) - Simple and fast template engine for Go *(:star:51 | 0clm)*
  * [dskinner/**damsel**](https://github.com/dskinner/damsel) - Package damsel provides html outlining via css-selectors and common template functionality. *(:star:16 | 6clm)*

## Testing
*Libraries for testing codebases and generating test data.*
* Testing Frameworks

  * [stretchr/**testify**](https://github.com/stretchr/testify) - A sacred extension to the standard go testing package *(:star:1717 | 11clm)*
  * [franela/**goblin**](https://github.com/franela/goblin) - Minimal and Beautiful Go testing framework *(:star:259 | 0clm)*
  * [msoap/**go-carpet**](https://github.com/msoap/go-carpet) - go-carpet - show test coverage for Go source files *(:star:134 | 0clm)*
  * [verdverm/**frisby**](https://github.com/verdverm/frisby) - API testing framework inspired by frisby-js *(:star:127 | 0clm)*
  * [bmizerany/**assert**](https://github.com/bmizerany/assert) - Asserts to Go testing *(:star:120 | 0clm)*
  * [orfjackal/**gospec**](https://github.com/orfjackal/gospec) - Testing framework for Go. Allows writing self-documenting tests/specifications, and executes them concurrently and safely isolated. [UNMAINTAINED] *(:star:109 | 0clm)*
  * [dnaeon/**go-vcr**](https://github.com/dnaeon/go-vcr) - Record and replay your HTTP interactions for fast, deterministic and accurate tests *(:star:83 | 0clm)*
  * [stesla/**gospecify**](https://github.com/stesla/gospecify) - A BDD library for Go *(:star:50 | 0clm)*
  * [DATA-DOG/**godog**](https://github.com/DATA-DOG/godog) - BDD Behat, Cucumber like gherkin feature runner for golang *(:star:46 | 0clm)*
  * [zimmski/**go-mutesting**](https://github.com/zimmski/go-mutesting) - Mutation testing for Go source code *(:star:35 | 0clm)*
  * [yookoala/**restit**](https://github.com/yookoala/restit) - A Go micro framework to help writing RESTful API integration test *(:star:24 | 0clm)*
  * [rdrdr/**hamcrest**](https://github.com/rdrdr/hamcrest) - Hamcrest matchers for the Go programming language *(:star:21 | 0clm)*
  * [marioidival/**bro**](https://github.com/marioidival/bro) - bro watch files in directory and run tests for them *(:star:11 | 0clm)*
  * [go-playground/**assert**](https://github.com/go-playground/assert) - :exclamation:Basic Assertion Library used along side native go testing, with building blocks for custom assertions *(:star:5 | 0clm)*
  * [cavaliercoder/**badio**](https://github.com/cavaliercoder/badio) - Extensions to Go's testing/iotest package *(:star:1 | 0clm)*
  * [go-testfixtures/**testfixtures**](https://github.com/go-testfixtures/testfixtures) - Rails' like test fixtures for Go Lang *(:star:0 | 0clm)*

* Mock

  * [golang/**mock**](https://github.com/golang/mock) - GoMock is a mocking framework for the Go programming language. *(:star:228 | 0clm)*
  * [DATA-DOG/**go-sqlmock**](https://github.com/DATA-DOG/go-sqlmock) - Sql mock driver for golang to test database interactions *(:star:198 | 2clm)*
  * [h2non/**gock**](https://github.com/h2non/gock) - Versatile HTTP mocking made easy in Go ༼ʘ̚ل͜ʘ̚༽ *(:star:192 | 0clm)*
  * [maxbrunsfeld/**counterfeiter**](https://github.com/maxbrunsfeld/counterfeiter) - A tool for generating self-contained, type-safe test doubles in go *(:star:54 | 0clm)*
  * [tv42/**mockhttp**](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter *(:star:18 | 0clm)*
  * [DATA-DOG/**go-txdb**](https://github.com/DATA-DOG/go-txdb) - Single transaction sql.Driver for GO *(:star:3 | 0clm)*

* Fuzzing and delta-debugging/reducing/shrinking

  * [dvyukov/**go-fuzz**](https://github.com/dvyukov/go-fuzz) - Randomized testing for Go *(:star:1326 | 10clm)*
  * [google/**gofuzz**](https://github.com/google/gofuzz) - Fuzz testing for go. *(:star:224 | 0clm)*
  * [zimmski/**tavor**](https://github.com/zimmski/tavor) - A generic fuzzing and delta-debugging framework *(:star:140 | 3clm)*

## Text Processing
*Libraries for parsing and manipulating texts.*
* Specific Formats

  * [PuerkitoBio/**goquery**](https://github.com/PuerkitoBio/goquery) - A little like that j-thing, only in Go. *(:star:2469 | 7clm)*
  * [russross/**blackfriday**](https://github.com/russross/blackfriday) - Blackfriday: a markdown processor for Go *(:star:1756 | 10clm)*
  * [BurntSushi/**toml**](https://github.com/BurntSushi/toml) - TOML parser for Golang with reflection. *(:star:812 | 0clm)*
  * [dustin/**go-humanize**](https://github.com/dustin/go-humanize) - Go Humans! (formatters for units to human friendly sizes) *(:star:652 | 6clm)*
  * [microcosm-cc/**bluemonday**](https://github.com/microcosm-cc/bluemonday) - bluemonday: a fast golang HTML sanitizer (inspired by the OWASP Java HTML Sanitizer) to scrub user generated content of XSS *(:star:335 | 0clm)*
  * [jteeuwen/**go-pkg-rss**](https://github.com/jteeuwen/go-pkg-rss) - This package reads RSS and Atom feeds and provides a caching mechanism that adheres to the feed specs. *(:star:258 | 3clm)*
  * [clbanning/**mxj**](https://github.com/clbanning/mxj) - Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages. *(:star:109 | 12clm)*
  * [jteeuwen/**go-pkg-xmlx**](https://github.com/jteeuwen/go-pkg-xmlx) - Extension to the standard Go XML package. Maintains a node tree that allows forward/backwards browsing and exposes some simple single/multi-node search functions. *(:star:106 | 0clm)*
  * [gosimple/**slug**](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support. *(:star:72 | 0clm)*
  * [mattn/**go-runewidth**](https://github.com/mattn/go-runewidth) -  *(:star:52 | 0clm)*
  * [awalterschulze/**gographviz**](https://github.com/awalterschulze/gographviz) - Parses the Graphviz DOT language in golang *(:star:37 | 0clm)*
  * [alixaxel/**genex**](https://github.com/alixaxel/genex) - Genex package for Go *(:star:36 | 1clm)*
  * [endeveit/**guesslanguage**](https://github.com/endeveit/guesslanguage) - Guess the natural language of a text in Go *(:star:17 | 0clm)*
  * [polera/**gonameparts**](https://github.com/polera/gonameparts) - Takes a full name and splits it into individual name parts *(:star:16 | 0clm)*
  * [zach-klippenstein/**goregen**](https://github.com/zach-klippenstein/goregen) - randexp for Go. *(:star:15 | 0clm)*
  * [adrianmo/**go-nmea**](https://github.com/adrianmo/go-nmea) - NMEA parser library for Golang *(:star:14 | 2clm)*
  * [avelino/**slugify**](https://github.com/avelino/slugify) - A Go slugify application that handles string *(:star:11 | 10clm)*
  * [endeveit/**enca**](https://github.com/endeveit/enca) - Minimal cgo bindings for libenca *(:star:2 | 0clm)*

* Utility

  * [mvdan/**xurls**](https://github.com/mvdan/xurls) - Extract urls from text *(:star:177 | 64clm)*
  * [bndr/**gotabulate**](https://github.com/bndr/gotabulate) - Gotabulate - Easily pretty-print your tabular data with Go *(:star:113 | 0clm)*
  * [codemodus/**parth**](https://github.com/codemodus/parth) - Path parsing / URL segmentation handling. *(:star:5 | 2clm)*
  * [nproc/**parseargs-go**](https://github.com/nproc/parseargs-go) - A string argument parser that understands quotes and backslashes *(:star:3 | 0clm)*
  * [codemodus/**kace**](https://github.com/codemodus/kace) - Common case conversions covering common initialisms. *(:star:2 | 0clm)*

## Third-party APIs
*Libraries for accessing third party APIs.*

  * [aws/**aws-sdk-go**](https://github.com/aws/aws-sdk-go) - AWS SDK for the Go programming language. *(:star:2400 | 98clm)*
  * [google/**go-github**](https://github.com/google/go-github) - Go library for accessing the GitHub API *(:star:1556 | 4clm)*
  * [mitchellh/**goamz**](https://github.com/mitchellh/goamz) - Golang Amazon Library *(:star:615 | 0clm)*
  * [google/**google-api-go-client**](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go *(:star:588 | 8clm)*
  * [ChimeraCoder/**anaconda**](https://github.com/ChimeraCoder/anaconda) - A Go client library for the Twitter 1.1 API *(:star:508 | 5clm)*
  * [nlopes/**slack**](https://github.com/nlopes/slack) - Slack API in Go *(:star:364 | 8clm)*
  * [stripe/**stripe-go**](https://github.com/stripe/stripe-go) - Go client for the Stripe API *(:star:350 | 11clm)*
  * [huandu/**facebook**](https://github.com/huandu/facebook) - A Facebook Graph API SDK Library For Golang *(:star:294 | 1clm)*
  * [GoogleCloudPlatform/**gcloud-golang**](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library *(:star:268 | 8clm)*
  * [Syfaro/**telegram-bot-api**](https://github.com/Syfaro/telegram-bot-api) - Golang bindings for the Telegram Bot API *(:star:139 | 0clm)*
  * [tucnak/**telebot**](https://github.com/tucnak/telebot) - Telegram bot framework written in Go *(:star:127 | 0clm)*
  * [codingsince1985/**geo-golang**](https://github.com/codingsince1985/geo-golang) - (reverse) geocoding service in Go *(:star:114 | 0clm)*
  * [andybons/**hipchat**](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API. *(:star:96 | 0clm)*
  * [daneharrigan/**hipchat**](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP *(:star:89 | 3clm)*
  * [minio/**minio-go**](https://github.com/minio/minio-go) - Minio Go Library for Amazon S3 compatible cloud storage *(:star:85 | 0clm)*
  * [logpacker/**paypalsdk**](https://github.com/logpacker/paypalsdk) - Golang client for PayPal REST API *(:star:79 | 0clm)*
  * [andygrunwald/**go-trending**](https://github.com/andygrunwald/go-trending) - Go library for accessing trending repositories and developers at Github. *(:star:70 | 0clm)*
  * [gambol99/**go-marathon**](https://github.com/gambol99/go-marathon) - A GO API library for working with Marathon *(:star:67 | 3clm)*
  * [jsgilmore/**gostorm**](https://github.com/jsgilmore/gostorm) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. *(:star:65 | 0clm)*
  * [bwmarrin/**discordgo**](https://github.com/bwmarrin/discordgo) -  (Golang) Go bindings for Discord *(:star:51 | 0clm)*
  * [dghubble/**go-twitter**](https://github.com/dghubble/go-twitter) - Go Twitter REST and Streaming API v1.1 *(:star:43 | 3clm)*
  * [Medium/**medium-sdk-go**](https://github.com/Medium/medium-sdk-go) - A Golang SDK for Medium's OAuth2 API *(:star:40 | 0clm)*
  * [samuelcouch/**clarifai**](https://github.com/samuelcouch/clarifai) - Clarifai library for Go *(:star:30 | 0clm)*
  * [Aorioli/**gcm**](https://github.com/Aorioli/gcm) - Google Cloud Messaging for application servers implemented using the Go programming language. *(:star:26 | 0clm)*
  * [michiwend/**gomusicbrainz**](https://github.com/michiwend/gomusicbrainz) - a Go (Golang) MusicBrainz WS2 client library - work in progress *(:star:17 | 2clm)*
  * [neuegram/**ghost**](https://github.com/neuegram/ghost) - A Go library for Snapchat's API *(:star:14 | 0clm)*
  * [emiddleton/**gads**](https://github.com/emiddleton/gads) - Google Adwords API for Go *(:star:13 | 1clm)*
  * [bit4bit/**gami**](https://github.com/bit4bit/gami) - GO - Asterisk AMI Interface *(:star:13 | 1clm)*
  * [dukex/**mixpanel**](https://github.com/dukex/mixpanel) - Golang Mixpanel Client *(:star:12 | 0clm)*
  * [go-playground/**webhooks**](https://github.com/go-playground/webhooks) - :fishing_pole_and_fish: Webhook receiver for GitHub and Bitbucket *(:star:12 | 0clm)*
  * [gregdel/**pushover**](https://github.com/gregdel/pushover) - Go wrapper arround the Pushover API *(:star:10 | 0clm)*
  * [poorny/**translate**](https://github.com/poorny/translate) - Go online translation package *(:star:10 | 0clm)*
  * [naegelejd/**brewerydb**](https://github.com/naegelejd/brewerydb) - Go library for http://www.brewerydb.com/ API *(:star:10 | 1clm)*
  * [nishanths/**go-xkcd**](https://github.com/nishanths/go-xkcd) - xkcd API client in Golang *(:star:9 | 0clm)*
  * [rapito/**go-shopify**](https://github.com/rapito/go-shopify) - Simple Shopify API for the Go Programming Language *(:star:9 | 0clm)*
  * [sostronk/**go-steam**](https://github.com/sostronk/go-steam) - Go library for querying Source servers *(:star:7 | 1clm)*
  * [sergiotapia/**smitego**](https://github.com/sergiotapia/smitego) - SmiteGo is an API wrapper for the Smite game from HiRez. It is written in Go! *(:star:6 | 0clm)*
  * [rapito/**go-spotify**](https://github.com/rapito/go-spotify) - Go library for the Spotify Web API *(:star:5 | 4clm)*
  * [dietsche/**textbelt**](https://github.com/dietsche/textbelt) - golang library for textbelt.com *(:star:4 | 0clm)*
  * [jbrodriguez/**go-tmdb**](https://github.com/jbrodriguez/go-tmdb) -  *(:star:4 | 0clm)*
  * [chonthu/**go-google-analytics**](https://github.com/chonthu/go-google-analytics) - Simple Reporting for Google Analytics *(:star:4 | 0clm)*
  * [Omie/**rrdaclient**](https://github.com/Omie/rrdaclient) - Go bindings for RRDA https://github.com/fcambus/rrda *(:star:3 | 0clm)*
  * [playlyfe/**playlyfe-go-sdk**](https://github.com/playlyfe/playlyfe-go-sdk) - This is the official Playlyfe Golang Sdk *(:star:0 | 0clm)*
  * [mattcunningham/**gumblr**](https://github.com/mattcunningham/gumblr) - A Go Wrapper for the Tumblr v2 API *(:star:0 | 0clm)*

## Utilities
*General utilities and tools to make your life easier.*

  * [inconshreveable/**ngrok**](https://github.com/inconshreveable/ngrok) - Introspected tunnels to localhost *(:star:6718 | 1clm)*
  * [junegunn/**fzf**](https://github.com/junegunn/fzf) - :cherry_blossom: A command-line fuzzy finder written in Go *(:star:4404 | 75clm)*
  * [derekparker/**delve**](https://github.com/derekparker/delve) - Delve is a debugger for the Go programming language. *(:star:3335 | 22clm)*
  * [dropbox/**godropbox**](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications. *(:star:2998 | 13clm)*
  * [peco/**peco**](https://github.com/peco/peco) - Simplistic interactive filtering tool *(:star:2883 | 7clm)*
  * [jmoiron/**sqlx**](https://github.com/jmoiron/sqlx) - general purpose extensions to golang's database/sql *(:star:1638 | 11clm)*
  * [maruel/**panicparse**](https://github.com/maruel/panicparse) - Crash your app in style (Golang) *(:star:1252 | 2clm)*
  * [rakyll/**coop**](https://github.com/rakyll/coop) - Cheat sheet for some of the common concurrent flows in Go *(:star:1138 | 0clm)*
  * [tealeg/**xlsx**](https://github.com/tealeg/xlsx) - Google Go (golang) library for reading and writing XLSX files. *(:star:895 | 10clm)*
  * [tobyhede/**go-underscore**](https://github.com/tobyhede/go-underscore) -  Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness.  *(:star:838 | 0clm)*
  * [levigross/**grequests**](https://github.com/levigross/grequests) - A Go "clone" of the great and famous Requests library *(:star:702 | 0clm)*
  * [parnurzeal/**gorequest**](https://github.com/parnurzeal/gorequest) - GoRequest -- Simplified HTTP client ( inspired by nodejs SuperAgent ) *(:star:636 | 7clm)*
  * [franela/**goreq**](https://github.com/franela/goreq) - Minimal and simple request library for Go language *(:star:435 | 9clm)*
  * [davecheney/**profile**](https://github.com/davecheney/profile) - A simple profiling support package for Go *(:star:417 | 0clm)*
  * [htcat/**htcat**](https://github.com/htcat/htcat) - Parallel and Pipelined HTTP GET Utility *(:star:415 | 0clm)*
  * [facebookgo/**httpcontrol**](https://github.com/facebookgo/httpcontrol) - Package httpcontrol allows for HTTP transport level control around timeouts and retries. *(:star:403 | 6clm)*
  * [afex/**hystrix-go**](https://github.com/afex/hystrix-go) - Netflix's Hystrix latency and fault tolerance library, for Go  *(:star:372 | 11clm)*
  * [bndr/**gopencils**](https://github.com/bndr/gopencils) - Easily consume REST APIs with Go (golang) *(:star:339 | 0clm)*
  * [tj/**go-debug**](https://github.com/tj/go-debug) - Conditional debug logging for Golang libraries & applications *(:star:306 | 0clm)*
  * [dghubble/**sling**](https://github.com/dghubble/sling) - A Go HTTP client library for creating and sending API requests *(:star:278 | 25clm)*
  * [go-resty/**resty**](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client *(:star:273 | 0clm)*
  * [joho/**godotenv**](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) *(:star:269 | 4clm)*
  * [bamzi/**jobrunner**](https://github.com/bamzi/jobrunner) - Framework for performing work asynchronously, outside of the request flow *(:star:260 | 0clm)*
  * [h2non/**gentleman**](https://github.com/h2non/gentleman) - Full-featured plugin-driven HTTP client toolkit for Go (͡° ͜ʖ ͡°) *(:star:253 | 0clm)*
  * [VividCortex/**godaemon**](https://github.com/VividCortex/godaemon) - Daemonize Go applications deviously. *(:star:243 | 0clm)*
  * [briandowns/**spinner**](https://github.com/briandowns/spinner) - Go (golang) package for providing a terminal spinner/progress indicator with options. *(:star:222 | 0clm)*
  * [cosiner/**gohper**](https://github.com/cosiner/gohper) - common libs here. *(:star:208 | 41clm)*
  * [minio/**mc**](https://github.com/minio/mc) - Minio Client is a replacement for ls, cp, mkdir, diff and rsync commands for filesystems and object storage. *(:star:198 | 514clm)*
  * [beefsack/**go-rate**](https://github.com/beefsack/go-rate) - A timed rate limiter for Go *(:star:184 | 0clm)*
  * [ungerik/**go-dry**](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go *(:star:143 | 0clm)*
  * [mozillazg/**request**](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™. *(:star:118 | 0clm)*
  * [StabbyCutyou/**moldova**](https://github.com/StabbyCutyou/moldova) - A lightweight templating system for generating random data *(:star:113 | 0clm)*
  * [rk/**go-cron**](https://github.com/rk/go-cron) - A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. *(:star:109 | 0clm)*
  * [imdario/**mergo**](https://github.com/imdario/mergo) - Mergo: merging Go structs and maps since 2013 *(:star:108 | 10clm)*
  * [VividCortex/**robustly**](https://github.com/VividCortex/robustly) - Run functions resiliently in Go, catching and restarting panics *(:star:102 | 0clm)*
  * [ivpusic/**rerun**](https://github.com/ivpusic/rerun) - Configurable recompiling and rerunning go apps when source changes *(:star:99 | 0clm)*
  * [ulule/**deepcopier**](https://github.com/ulule/deepcopier) - simple struct copying for golang *(:star:81 | 0clm)*
  * [carlescere/**scheduler**](https://github.com/carlescere/scheduler) - Job scheduling made easy. *(:star:79 | 1clm)*
  * [jaschaephraim/**lrserver**](https://github.com/jaschaephraim/lrserver) - LiveReload server for Go [golang] *(:star:63 | 0clm)*
  * [sethgrid/**pester**](https://github.com/sethgrid/pester) - Go (golang) http calls with retries and backoff  *(:star:60 | 0clm)*
  * [h2non/**filetype**](https://github.com/h2non/filetype) - Small Go package to infer the file type checking the magic numbers signature *(:star:49 | 0clm)*
  * [subosito/**gotenv**](https://github.com/subosito/gotenv) - Load environment variables dynamically in Go. *(:star:48 | 0clm)*
  * [ricardolonga/**jsongo**](https://github.com/ricardolonga/jsongo) - Fluent API to make it easier to create Json objects. *(:star:46 | 0clm)*
  * [tmrts/**boilr**](https://github.com/tmrts/boilr) - :zap: Blazingly fast CLI tool for creating projects from boilerplate templates *(:star:42 | 0clm)*
  * [sadlil/**go-trigger**](https://github.com/sadlil/go-trigger) - A Global event triggerer for golang. Defines functions as event with id string. Trigger the event anywhere from your project. *(:star:40 | 0clm)*
  * [VividCortex/**pm**](https://github.com/VividCortex/pm) - Processlist manager with TCP listener *(:star:39 | 0clm)*
  * [topfreegames/**apm**](https://github.com/topfreegames/apm) - APM is a process manager for Golang applications. *(:star:29 | 0clm)*
  * [e-dard/**netbug**](https://github.com/e-dard/netbug) - Package netbug provides a handler for registering profilers on your own ServeMux. *(:star:28 | 0clm)*
  * [miolini/**jsonf**](https://github.com/miolini/jsonf) - Console JSON formatter with query feature *(:star:23 | 0clm)*
  * [carlescere/**goback**](https://github.com/carlescere/goback) - Golang simple exponential backoff package. *(:star:23 | 0clm)*
  * [elgs/**gojq**](https://github.com/elgs/gojq) - JSON query in Golang *(:star:23 | 0clm)*
  * [vrecan/**death**](https://github.com/vrecan/death) - Managing go application shutdown with signals. *(:star:18 | 6clm)*
  * [msempere/**golarm**](https://github.com/msempere/golarm) - Fire alarms with system events *(:star:15 | 0clm)*
  * [mlimaloureiro/**golog**](https://github.com/mlimaloureiro/golog) - Easy and simple CLI time tracker for your tasks *(:star:13 | 0clm)*
  * [smallnest/**goreq**](https://github.com/smallnest/goreq) - A Simplified Golang Http Client *(:star:13 | 0clm)*
  * [xta/**okrun**](https://github.com/xta/okrun) - ok, run your gofile *(:star:11 | 0clm)*
  * [sanbornm/**mp**](https://github.com/sanbornm/mp) - Simple Email Parser *(:star:10 | 0clm)*
  * [michiwend/**goplaceholder**](https://github.com/michiwend/goplaceholder) - a small golang lib to generate placeholder images *(:star:10 | 0clm)*
  * [ikeikeikeike/**go-sitemap-generator**](https://github.com/ikeikeikeike/go-sitemap-generator) - go-sitemap-generator is the easiest way to generate Sitemaps in Go. *(:star:8 | 0clm)*
  * [alxrm/**ugo**](https://github.com/alxrm/ugo) - Simple and expressive toolbox written in Go *(:star:5 | 0clm)*
  * [bahlo/**abutil**](https://github.com/bahlo/abutil) - :ab: A collection of often-used Golang helpers *(:star:4 | 0clm)*
  * [digitalcrab/**fastlz**](https://github.com/digitalcrab/fastlz) - Wrap over FastLz for GoLang *(:star:4 | 0clm)*
  * [go-playground/**generate**](https://github.com/go-playground/generate) - :runner:runs go generate recursively on a specified path or environment variable and can filter by regex *(:star:1 | 0clm)*
  * [txgruppi/**command**](https://github.com/txgruppi/command) - Command pattern for Go with thread safe serial and parallel dispatcher *(:star:1 | 0clm)*

## Validation
*Libraries for validation.*

  * [asaskevich/**govalidator**](https://github.com/asaskevich/govalidator) - [Go] Package of validators and sanitizers for strings, numerics, slices and structs *(:star:974 | 0clm)*
  * [go-playground/**validator**](https://github.com/go-playground/validator) - :100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving *(:star:423 | 25clm)*

## Version Control
*Libraries for version control.*

  * [libgit2/**git2go**](https://github.com/libgit2/git2go) - Git to Go. Like McDonald's but tastier. *(:star:759 | 6clm)*
  * [sourcegraph/**go-vcs**](https://github.com/sourcegraph/go-vcs) - manipulate and inspect VCS repositories in Go *(:star:40 | 17clm)*
  * [rjeczalik/**gh**](https://github.com/rjeczalik/gh) - Scriptable server and net/http middleware for GitHub Webhooks. *(:star:36 | 13clm)*
  * [beyang/**hgo**](https://github.com/beyang/hgo) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories. *(:star:7 | 0clm)*

## Video
*Libraries for manipulating video.*

  * [3d0c/**gmf**](https://github.com/3d0c/gmf) - Go Media Framework *(:star:177 | 7clm)*
  * [nareix/**codec**](https://github.com/nareix/codec) - Golang libav codec bindings (h264,aac) *(:star:148 | 0clm)*
  * [giorgisio/**goav**](https://github.com/giorgisio/goav) - Golang bindings for FFmpeg *(:star:113 | 0clm)*
  * [ziutek/**gst**](https://github.com/ziutek/gst) - Go bindings for GStreamer *(:star:110 | 0clm)*

## Web Frameworks
*Full stack web frameworks.*

  * [astaxie/**beego**](https://github.com/astaxie/beego) - beego is an open-source, high-performance web framework for the Go programming language. *(:star:6560 | 72clm)*
  * [revel/**revel**](https://github.com/revel/revel) - A high productivity, full-stack web framework for the Go language. *(:star:6495 | 8clm)*
  * [gin-gonic/**gin**](https://github.com/gin-gonic/gin) - Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. *(:star:6000 | 55clm)*
  * [go-kit/**kit**](https://github.com/go-kit/kit) - A standard library for microservices. *(:star:3652 | 44clm)*
  * [labstack/**echo**](https://github.com/labstack/echo) - Echo is  a fast :rocket: and unfancy micro web framework for Go. *(:star:3134 | 99clm)*
  * [julienschmidt/**httprouter**](https://github.com/julienschmidt/httprouter) - A high performance HTTP request router that scales well *(:star:2936 | 3clm)*
  * [hoisie/**web**](https://github.com/hoisie/web) - The easiest way to create web applications with Go *(:star:2615 | 0clm)*
  * [ant0ine/**go-json-rest**](https://github.com/ant0ine/go-json-rest) - A quick and easy way to setup a RESTful JSON API *(:star:2235 | 10clm)*
  * [gorilla/**mux**](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for golang. *(:star:2153 | 0clm)*
  * [NYTimes/**gizmo**](https://github.com/NYTimes/gizmo) - A Microservice Toolkit from The New York Times *(:star:1303 | 0clm)*
  * [googollee/**go-socket.io**](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework. *(:star:1060 | 1clm)*
  * [go-macaron/**macaron**](https://github.com/go-macaron/macaron) - Package macaron is a high productive and modular web framework in Go. *(:star:1026 | 0clm)*
  * [gocraft/**web**](https://github.com/gocraft/web) - Go Router + Middleware. Your Contexts. *(:star:1005 | 0clm)*
  * [rcrowley/**go-tigertonic**](https://github.com/rcrowley/go-tigertonic) - A Go framework for building JSON web services inspired by Dropwizard *(:star:923 | 0clm)*
  * [bmizerany/**pat**](https://github.com/bmizerany/pat) -  *(:star:922 | 0clm)*
  * [go-zoo/**bone**](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer *(:star:885 | 1clm)*
  * [raphael/**goa**](https://github.com/raphael/goa) - Design-based microservices in Go *(:star:689 | 0clm)*
  * [pilu/**traffic**](https://github.com/pilu/traffic) - Sinatra inspired regexp/pattern mux and web framework for Go *(:star:466 | 0clm)*
  * [lunny/**tango**](https://github.com/lunny/tango) - Micro & pluggable web framework for Go *(:star:378 | 23clm)*
  * [VividCortex/**siesta**](https://github.com/VividCortex/siesta) - Composable framework for writing HTTP handlers in Go. *(:star:337 | 0clm)*
  * [rainycape/**gondola**](https://github.com/rainycape/gondola) - The web framework for writing faster sites, faster *(:star:296 | 7clm)*
  * [paulbellamy/**mango**](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. *(:star:295 | 0clm)*
  * [pressly/**chi**](https://github.com/pressly/chi) - Composable Go HTTP router built with net/context *(:star:274 | 0clm)*
  * [jcuga/**golongpoll**](https://github.com/jcuga/golongpoll) - golang long polling library.  Makes web pub-sub easy via HTTP long-poll server :smiley: :coffee: :computer: *(:star:274 | 0clm)*
  * [ivpusic/**neo**](https://github.com/ivpusic/neo) - Go Web Framework *(:star:247 | 1clm)*
  * [dimfeld/**httptreemux**](https://github.com/dimfeld/httptreemux) - High-speed, flexible tree-based HTTP router for Go. *(:star:138 | 5clm)*
  * [desertbit/**glue**](https://github.com/desertbit/glue) - Glue - Robust Go and Javascript Socket Library (Alternative to Socket.io) *(:star:137 | 0clm)*
  * [cosiner/**zerver**](https://github.com/cosiner/zerver) - a RESTful API framework *(:star:125 | 44clm)*
  * [codehack/**go-relax**](https://github.com/codehack/go-relax) - Framework for building RESTful API's in Go *(:star:119 | 0clm)*
  * [goji/**goji**](https://github.com/goji/goji) - Goji is a minimalistic and flexible HTTP request multiplexer for Go (golang) *(:star:117 | 0clm)*
  * [daryl/**zeus**](https://github.com/daryl/zeus) - Go HTTP router. *(:star:93 | 0clm)*
  * [ungerik/**go-rest**](https://github.com/ungerik/go-rest) - A small and evil REST framework for Go *(:star:88 | 0clm)*
  * [bahlo/**goat**](https://github.com/bahlo/goat) - :goat: A minimalistic JSON API server in Go *(:star:81 | 0clm)*
  * [volatile/**core**](https://github.com/volatile/core) - Pure handlers stack *(:star:81 | 0clm)*
  * [go-ozzo/**ozzo-routing**](https://github.com/go-ozzo/ozzo-routing) - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. *(:star:59 | 0clm)*
  * [dinever/**golf**](https://github.com/dinever/golf) - Golf: A micro web framework for Go. *(:star:48 | 0clm)*
  * [rs/**xmux**](https://github.com/rs/xmux) - xmux is a httprouter fork on top of xhandler (net/context aware) *(:star:46 | 0clm)*
  * [resoursea/**api**](https://github.com/resoursea/api) - A REST framework for quickly writing resource based services in Golang. *(:star:24 | 0clm)*
  * [ian-kent/**goose**](https://github.com/ian-kent/goose) - Server-Sent Events in Go *(:star:16 | 0clm)*
  * [imdario/**medeina**](https://github.com/imdario/medeina) - Go HTTP routing tree based on HttpRouter. Inspired by Roda and Cuba. *(:star:15 | 0clm)*
  * [goanywhere/**rex**](https://github.com/goanywhere/rex) - Pleasures for Web in Golang *(:star:10 | 0clm)*

### Middlewares
#### Actual middlewares

  * [didip/**tollbooth**](https://github.com/didip/tollbooth) - Simple middleware to rate-limit HTTP requests. *(:star:363 | 0clm)*
  * [rs/**cors**](https://github.com/rs/cors) - Go net/http configurable handler to handle CORS requests *(:star:315 | 6clm)*
  * [ulule/**limiter**](https://github.com/ulule/limiter) - Dead simple rate limit middleware for Go. *(:star:210 | 0clm)*
  * [sebest/**xff**](https://github.com/sebest/xff) - A Golang Middleware to handle X-Forwarded-For Header *(:star:42 | 0clm)*
  * [rs/**formjson**](https://github.com/rs/formjson) - Go net/http handler to transparently manage posted JSON *(:star:18 | 0clm)*

#### Libraries for creating HTTP middlewares

  * [codegangsta/**negroni**](https://github.com/codegangsta/negroni) - Idiomatic HTTP Middleware for Golang *(:star:3493 | 0clm)*
  * [justinas/**alice**](https://github.com/justinas/alice) - Painless middleware chaining for Go *(:star:861 | 1clm)*
  * [unrolled/**render**](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, binary data, and HTML templates responses. *(:star:644 | 2clm)*
  * [thoas/**stats**](https://github.com/thoas/stats) - A Go middleware that stores various information about your web application (response time, status code count, etc.) *(:star:329 | 4clm)*
  * [carbocation/**interpose**](https://github.com/carbocation/interpose) - Minimalist net/http middleware for golang *(:star:235 | 0clm)*
  * [stephens2424/**muxchain**](https://github.com/stephens2424/muxchain) - Lightweight Middleware for net/http *(:star:198 | 0clm)*
  * [go-on/**wrap**](https://github.com/go-on/wrap) - Go http.Hander based middleware stack with context sharing *(:star:50 | 0clm)*
  * [codemodus/**chain**](https://github.com/codemodus/chain) - Handler wrapper chaining with scoped data. *(:star:41 | 0clm)*
  * [alioygur/**gores**](https://github.com/alioygur/gores) - Go package that handles HTML, JSON, XML and etc. responses *(:star:38 | 0clm)*
  * [codemodus/**catena**](https://github.com/codemodus/catena) - http.Handler wrapper catenation. *(:star:2 | 0clm)*

# Tools
## Code Analysis

  * [golang/**lint**](https://github.com/golang/lint) - This is a linter for Go source code. *(:star:1280 | 3clm)*
  * [alecthomas/**gometalinter**](https://github.com/alecthomas/gometalinter) - Concurrently run Go lint tools and normalise their output *(:star:763 | 1clm)*
  * [kisielk/**errcheck**](https://github.com/kisielk/errcheck) - errcheck checks that you checked errors. *(:star:519 | 6clm)*
  * [davecheney/**gcvis**](https://github.com/davecheney/gcvis) - Visualise Go program GC trace data in real time *(:star:492 | 0clm)*
  * [mvdan/**interfacer**](https://github.com/mvdan/interfacer) - A linter that suggests interface types *(:star:223 | 0clm)*
  * [yuroyoro/**goast-viewer**](https://github.com/yuroyoro/goast-viewer) - Golang AST visualizer *(:star:137 | 0clm)*
  * [shurcooL/**gostatus**](https://github.com/shurcooL/gostatus) - A command line tool that shows the status of Go repositories. *(:star:123 | 1clm)*
  * [mccoyst/**validate**](https://github.com/mccoyst/validate) - A Go package to automatically validate fields with tags *(:star:56 | 0clm)*
  * [mibk/**dupl**](https://github.com/mibk/dupl) - a tool for code clone detection *(:star:44 | 23clm)*
  * [qiniu/**checkstyle**](https://github.com/qiniu/checkstyle) - checkstyle for go *(:star:26 | 0clm)*
  * [firstrow/**go-outdated**](https://github.com/firstrow/go-outdated) - Find outdated golang packages *(:star:19 | 0clm)*

## Editor Plugins

  * [fatih/**vim-go**](https://github.com/fatih/vim-go) - Go development plugin for Vim *(:star:4463 | 23clm)*
  * [go-lang-plugin-org/**go-lang-idea-plugin**](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) - Google Go language IDE built using the IntelliJ Platform *(:star:3139 | 118clm)*
  * [nsf/**gocode**](https://github.com/nsf/gocode) - An autocompletion daemon for the Go programming language *(:star:2998 | 4clm)*
  * [DisposaBoy/**GoSublime**](https://github.com/DisposaBoy/GoSublime) - A  Golang plugin collection for the text editor SublimeText 2 providing code completion and other IDE-like features. *(:star:2159 | 0clm)*
  * [joefitzgerald/**go-plus**](https://github.com/joefitzgerald/go-plus) - An Enhanced Go Experience For The Atom Editor *(:star:813 | 21clm)*
  * [GoClipse/**goclipse**](https://github.com/GoClipse/goclipse) - Eclipse IDE for the Go programming language *(:star:504 | 139clm)*
  * [eaburns/**Watch**](https://github.com/eaburns/Watch) - Watches for changes in a directory tree and reruns a command in an acme win or just on the terminal. *(:star:104 | 0clm)*
  * [rjohnsondev/**vim-compiler-go**](https://github.com/rjohnsondev/vim-compiler-go) - Vim compiler plugin for Go (golang) *(:star:63 | 0clm)*
  * [velour/**velour**](https://github.com/velour/velour) - An IRC client for acme — the project that started it all. *(:star:13 | 0clm)*

## Go Tools

  * [songgao/**colorgo**](https://github.com/songgao/colorgo) - Colorize (highlight) `go build` command output *(:star:61 | 0clm)*
  * [skelterjohn/**go-pkg-complete**](https://github.com/skelterjohn/go-pkg-complete) - bash completion for go and wgo *(:star:28 | 0clm)*

## Software Packages
### DevOps Tools

  * [kubernetes/**kubernetes**](https://github.com/kubernetes/kubernetes) - Container Cluster Manager from Google *(:star:13390 | 1644clm)*
  * [mitchellh/**packer**](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. *(:star:5128 | 24clm)*
  * [rakyll/**boom**](https://github.com/rakyll/boom) - HTTP(S) load generator, ApacheBench (ab) replacement, written in Go *(:star:4077 | 0clm)*
  * [tsenart/**vegeta**](https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000! *(:star:3059 | 20clm)*
  * [moovweb/**gvm**](https://github.com/moovweb/gvm) - Go Version Manager *(:star:1866 | 4clm)*
  * [bosun-monitor/**bosun**](https://github.com/bosun-monitor/bosun) - Time Series Alerting Framework *(:star:1719 | 102clm)*
  * [mitchellh/**gox**](https://github.com/mitchellh/gox) - A dead simple, no frills Go cross compile tool *(:star:1501 | 0clm)*
  * [laher/**goxc**](https://github.com/laher/goxc) - a build tool for Go, with a focus on cross-compiling, packaging and deployment *(:star:1248 | 0clm)*
  * [rcrowley/**go-metrics**](https://github.com/rcrowley/go-metrics) - Go port of Coda Hale's Metrics library *(:star:1137 | 1clm)*
  * [smira/**aptly**](https://github.com/smira/aptly) - aptly - Debian repository management tool *(:star:1013 | 38clm)*
  * [ajvb/**kala**](https://github.com/ajvb/kala) - Modern Job Scheduler *(:star:780 | 46clm)*
  * [heroku/**hk**](https://github.com/heroku/hk) - DEPRECATED: see *(:star:768 | 2clm)*
  * [rlmcpherson/**s3gof3r**](https://github.com/rlmcpherson/s3gof3r) - Fast, concurrent, streaming access to Amazon S3, including gof3r, a CLI. http://godoc.org/github.com/rlmcpherson/s3gof3r *(:star:634 | 1clm)*
  * [sanbornm/**go-selfupdate**](https://github.com/sanbornm/go-selfupdate) - Enable your Golang applications to self update *(:star:357 | 4clm)*
  * [adnanh/**webhook**](https://github.com/adnanh/webhook) - webhook is a lightweight configurable tool written in Go, that allows you to easily create HTTP endpoints (hooks) on your server, which you can use to execute configured commands. *(:star:353 | 8clm)*
  * [inconshreveable/**gonative**](https://github.com/inconshreveable/gonative) - Build Go Toolchains /w native libs for cross-compilation *(:star:224 | 0clm)*
  * [eleme/**banshee**](https://github.com/eleme/banshee) - Anomalies detection system for periodic metrics. *(:star:185 | 0clm)*
  * [sirnewton01/**godbg**](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application *(:star:185 | 0clm)*
  * [cryptojuice/**gobrew**](https://github.com/cryptojuice/gobrew) - Shell script to download and set GO environmental paths to allow multiple versions. *(:star:148 | 6clm)*
  * [emicklei/**mora**](https://github.com/emicklei/mora) - MongoDB generic REST server in Go *(:star:145 | 1clm)*
  * [scaleway/**scaleway-cli**](https://github.com/scaleway/scaleway-cli) - :computer: Manage BareMetal Servers from Command Line (as easily as with Docker) *(:star:109 | 9clm)*
  * [liudng/**dogo**](https://github.com/liudng/dogo) - Monitoring changes in the source file and automatically compile and run (restart). *(:star:88 | 2clm)*
  * [hypersleep/**easyssh**](https://github.com/hypersleep/easyssh) - Golang package for easy remote execution through SSH *(:star:85 | 0clm)*
  * [ostrost/**ostent**](https://github.com/ostrost/ostent) - ostent collects system metrics to display and relay *(:star:60 | 0clm)*
  * [alouche/**rodent**](https://github.com/alouche/rodent) - Manage Go Versions/Projects/Dependencies *(:star:31 | 2clm)*
  * [chrismckenzie/**dropship**](https://github.com/chrismckenzie/dropship) - instead of jumping ship... Dropship *(:star:21 | 0clm)*
  * [soniah/**awsenv**](https://github.com/soniah/awsenv) - AWS environment config loader *(:star:7 | 0clm)*

### Other Software

  * [coreos/**rkt**](https://github.com/coreos/rkt) - rkt is an App Container runtime for Linux *(:star:5334 | 245clm)*
  * [buger/**gor**](https://github.com/buger/gor) - Gor is an open-source tool for capturing and replaying live HTTP traffic into a test environment in order to continuously test your system with real data. It can be used to increase confidence in code deployments, configuration changes and infrastructure changes. *(:star:4868 | 1clm)*
  * [tylertreat/**Comcast**](https://github.com/tylertreat/Comcast) - Simulating shitty network connections so you can build better systems. *(:star:4165 | 0clm)*
  * [mozilla-services/**heka**](https://github.com/mozilla-services/heka) - Data collection and processing made easy. *(:star:2936 | 87clm)*
  * [visualfc/**liteide**](https://github.com/visualfc/liteide) - LiteIDE is a simple, open source, cross-platform Go IDE.  *(:star:2652 | 14clm)*
  * [kelseyhightower/**confd**](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul *(:star:2481 | 77clm)*
  * [chrislusf/**seaweedfs**](https://github.com/chrislusf/seaweedfs) - SeaweedFS is a simple and highly scalable distributed file system. There are two objectives:  to store billions of files! to serve the files fast! Instead of supporting full POSIX file system semantics, Seaweed-FS choose to implement only a key~file mapping. Similar to the word "NoSQL", you can call it as "NoFS". *(:star:2248 | 50clm)*
  * [fogleman/**nes**](https://github.com/fogleman/nes) - NES emulator written in Go. *(:star:2216 | 66clm)*
  * [coreos/**fleet**](https://github.com/coreos/fleet) - A Distributed init System *(:star:2119 | 39clm)*
  * [gocircuit/**circuit**](https://github.com/gocircuit/circuit) - Circuit: Dynamic cloud orchestration http://gocircuit.org *(:star:1245 | 1clm)*
  * [shopify/**toxiproxy**](https://github.com/shopify/toxiproxy) - :alarm_clock: :fire: A proxy to simulate network and system conditions *(:star:1185 | 20clm)*
  * [pressly/**sup**](https://github.com/pressly/sup) - Super simple deployment tool - just Unix - think of it like 'make' for a network of servers *(:star:901 | 1clm)*
  * [zachlatta/**postman**](https://github.com/zachlatta/postman) - Command-line utility for batch-sending email. *(:star:634 | 0clm)*
  * [intelsdi-x/**snap**](https://github.com/intelsdi-x/snap) - A powerful telemetry framework *(:star:571 | 74clm)*
  * [restic/**restic**](https://github.com/restic/restic) - restic backup program *(:star:551 | 44clm)*
  * [pointlander/**peg**](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. *(:star:280 | 0clm)*
  * [wellington/**wellington**](https://github.com/wellington/wellington) - Spriting that sass has been missing *(:star:171 | 106clm)*
  * [kyleterry/**tenyks**](https://github.com/kyleterry/tenyks) - The Tenyks IRC bot. *(:star:153 | 3clm)*
  * [dimiro1/**ipe**](https://github.com/dimiro1/ipe) - An open source Pusher server implementation compatible with Pusher client libraries written in GO *(:star:144 | 0clm)*
  * [noraesae/**orange-cat**](https://github.com/noraesae/orange-cat) - A Markdown previewer written in Go *(:star:139 | 0clm)*
  * [rafael-santiago/**cherry**](https://github.com/rafael-santiago/cherry) - A tiny webchat server in Go. *(:star:72 | 0clm)*
  * [tejo/**boxed**](https://github.com/tejo/boxed) - dropbox based blog engine, written in go. *(:star:46 | 9clm)*
  * [msoap/**shell2http**](https://github.com/msoap/shell2http) - Executing shell commands via http server (GoLang) *(:star:40 | 26clm)*
  * [goccmack/**gocc**](https://github.com/goccmack/gocc) - Parser / Scanner Generator *(:star:31 | 0clm)*
  * [ian-kent/**websysd**](https://github.com/ian-kent/websysd) - Like Marathon or Upstart - for your desktop! *(:star:18 | 0clm)*
  * [diankong/**GoDocTooltip**](https://github.com/diankong/GoDocTooltip) - A Chrome extension for golang users.When you're at golang's official doc site, it will show function's description as tooltip on function list *(:star:7 | 0clm)*
  * [unix4fun/**naclpipe**](https://github.com/unix4fun/naclpipe) - NaCL pipe *(:star:5 | 0clm)*
  * [blogcin/**ToTo**](https://github.com/blogcin/ToTo) - Proxy server written in Go language *(:star:4 | 0clm)*

# Resources
## Benchmarks

  * [julienschmidt/**go-http-routing-benchmark**](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router and web framework benchmark *(:star:651 | 17clm)*
  * [atemerev/**skynet**](https://github.com/atemerev/skynet) - Skynet 1M threads microbenchmark *(:star:558 | 0clm)*
  * [alecthomas/**go_serialization_benchmarks**](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods *(:star:285 | 0clm)*
  * [davecheney/**autobench**](https://github.com/davecheney/autobench) - Go benchmark harness.  *(:star:84 | 0clm)*
  * [fawick/**speedtest-resize**](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language *(:star:65 | 0clm)*
  * [feyeleanor/**GoSpeed**](https://github.com/feyeleanor/GoSpeed) - Go micro-benchmarks for calculating the speed of language constructs *(:star:48 | 0clm)*
  * [PuerkitoBio/**gocostmodel**](https://github.com/PuerkitoBio/gocostmodel) - Benchmarks of common basic operations for the Go language. *(:star:46 | 0clm)*
  * [tyler-smith/**golang-sql-benchmark**](https://github.com/tyler-smith/golang-sql-benchmark) - A benchmarking shootout of various db/SQL utilities for Go *(:star:20 | 0clm)*
  * [tylertreat/**go-benchmarks**](https://github.com/tylertreat/go-benchmarks) - A few miscellaneous Go microbenchmarks. *(:star:19 | 0clm)*
  * [jimrobinson/**kvbench**](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark *(:star:10 | 0clm)*
  * [amscanne/**golang-micro-benchmarks**](https://github.com/amscanne/golang-micro-benchmarks) - Tiny collection of micro benchmarks. *(:star:5 | 0clm)*
  * [hgfischer/**go-type-assertion-benchmark**](https://github.com/hgfischer/go-type-assertion-benchmark) - Naive performance test of two ways to do type assertion in Go. *(:star:3 | 0clm)*

## E-Books

  * [dariubs/**GoBooks**](https://github.com/dariubs/GoBooks) - List of Golang books *(:star:1973 | 0clm)*

## Websites

  * [bayandin/**awesome-awesomeness**](https://github.com/bayandin/awesome-awesomeness) - A curated list of awesome awesomeness *(:star:15739 | 20clm)*
  * [lukasz-madon/**awesome-remote-job**](https://github.com/lukasz-madon/awesome-remote-job) - A curated list of awesome remote jobs and resources. Inspired by https://github.com/vinta/awesome-python *(:star:5170 | 27clm)*
  * [mholt/**golang-graphics**](https://github.com/mholt/golang-graphics) - Community-contributed Go graphics files *(:star:77 | 0clm)*

### Tutorials

  * [mkaz/**working-with-go**](https://github.com/mkaz/working-with-go) - A set of example golang code to start learning Go *(:star:484 | 0clm)*

## Windows

  * [go-ole/**go-ole**](https://github.com/go-ole/go-ole) - win32 ole implementation for golang *(:star:186 | 0clm)*



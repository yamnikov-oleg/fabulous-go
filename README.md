# Fabulous Go

This application parses github repos from [Fabulous Go](https://github.com/avelino/awesome-go) and populates the lists with additional data, like stars count and commit activity.

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

  * [eaburns/**flac**](https://github.com/eaburns/flac) - A Free Lossless Audio Codec decoder in Go *(:star:43 | 0clm)*
  * [outrightmental/**go-atomix**](https://github.com/outrightmental/go-atomix) - Sequence-based Go-native audio mixer for Music apps *(:star:2 | 0clm)*
  * [krig/**go-sox**](https://github.com/krig/go-sox) - libsox bindings for go *(:star:33 | 1clm)*
  * [zhulik/**go_mediainfo**](https://github.com/zhulik/go_mediainfo) - Golang bindings for libmediainfo *(:star:1 | 0clm)*
  * [tcolgate/**mp3**](https://github.com/tcolgate/mp3) - golang mp3 frame parser *(:star:6 | 5clm)*
  * [gordonklaus/**portaudio**](https://github.com/gordonklaus/portaudio) - Go bindings for the PortAudio audio I/O library *(:star:21 | 0clm)*
  * [rakyll/**portmidi**](https://github.com/rakyll/portmidi) - Go bindings for libportmidi *(:star:73 | 0clm)*
  * [wtolson/**go-taglib**](https://github.com/wtolson/go-taglib) - Go wrapper for taglib *(:star:38 | 0clm)*
  * [mccoyst/**vorbis**](https://github.com/mccoyst/vorbis) - A "native" ogg vorbis decoder for Go (uses inline stb_vorbis) *(:star:9 | 0clm)*
  * [mdlayher/**waveform**](https://github.com/mdlayher/waveform) - Go package capable of generating waveform images from audio streams. MIT Licensed. *(:star:107 | 0clm)*

## Authentication & OAuth
*Libraries for implementing authentications schemes.*

  * [smartystreets/**go-aws-auth**](https://github.com/smartystreets/go-aws-auth) - Signs requests to Amazon Web Services (AWS) using IAM roles or signed signature versions 2, 3, and 4. Supports S3 and STS. *(:star:93 | 0clm)*
  * [square/**go-jose**](https://github.com/square/go-jose) - An implementation of JOSE standards in Golang. *(:star:451 | 18clm)*
  * [bradrydzewski/**go.auth**](https://github.com/bradrydzewski/go.auth) - authentication API for Go web applications *(:star:288 | 0clm)*
  * [dghubble/**gologin**](https://github.com/dghubble/gologin) - Chainable Go login handlers for authentication providers (oauth1, oauth2) *(:star:587 | 0clm)*
  * [mikespook/**gorbac**](https://github.com/mikespook/gorbac) - goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. *(:star:272 | 0clm)*
  * [markbates/**goth**](https://github.com/markbates/goth) - Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. *(:star:566 | 13clm)*
  * [goji/**httpauth**](https://github.com/goji/httpauth) - HTTP Authentication middlewares *(:star:85 | 0clm)*
  * [dgrijalva/**jwt-go**](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens (JWT) *(:star:906 | 0clm)*
  * [golang/**oauth2**](https://github.com/golang/oauth2) - Go OAuth2 *(:star:628 | 18clm)*
  * [RangelReale/**osin**](https://github.com/RangelReale/osin) - Golang OAuth2 server library *(:star:786 | 8clm)*
  * [xyproto/**permissions2**](https://github.com/xyproto/permissions2) -   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions *(:star:140 | 3clm)*
  * [GeertJohan/**yubigo**](https://github.com/GeertJohan/yubigo) -   Yubigo is a Yubikey client API library that provides an easy way to integrate the Yubico Yubikey into your existing Go-based user authentication infrastructure. *(:star:51 | 0clm)*

## Command Line
### Standard CLI
*Libraries for building standard or basic Command Line applications*

  * [tcnksm/**gcli**](https://github.com/tcnksm/gcli) - The easy way to build Golang command-line application. *(:star:435 | 1clm)*
  * [tucnak/**climax**](https://github.com/tucnak/climax) - Climax is an alternative CLI with "human face" *(:star:90 | 0clm)*
  * [spf13/**cobra**](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions *(:star:1690 | 9clm)*
  * [codegangsta/**cli**](https://github.com/codegangsta/cli) - A small package for building command line apps in Go *(:star:3728 | 6clm)*
  * [docopt/**docopt.go**](https://github.com/docopt/docopt.go) - A command-line arguments parser that will make you smile. *(:star:634 | 0clm)*
  * [jessevdk/**go-flags**](https://github.com/jessevdk/go-flags) - go command line option parser *(:star:526 | 2clm)*
  * [alecthomas/**kingpin**](https://github.com/alecthomas/kingpin) - A Go (golang) command line and flag parser *(:star:483 | 11clm)*
  * [peterh/**liner**](https://github.com/peterh/liner) - Pure Go line editor with history, inspired by linenoise *(:star:272 | 16clm)*
  * [mitchellh/**cli**](https://github.com/mitchellh/cli) - A Go library for implementing command-line interfaces. *(:star:372 | 0clm)*
  * [jawher/**mow.cli**](https://github.com/jawher/mow.cli) - A versatile library for building CLI applications in Go *(:star:320 | 4clm)*
  * [chzyer/**readline**](https://github.com/chzyer/readline) - Readline is a pure go(golang) implementation for GNU-Readline kind library *(:star:435 | 0clm)*
  * [ukautz/**clif**](https://github.com/ukautz/clif) - Another CLI framework for Go. It works on my machine. *(:star:60 | 0clm)*

### Advanced Console UIs
*Libraries for building Console Applications and Console User Interfaces*

  * [ttacon/**chalk**](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk *(:star:145 | 0clm)*
  * [fatih/**color**](https://github.com/fatih/color) - Color package for Go (golang) *(:star:692 | 0clm)*
  * [TreyBastian/**colourize**](https://github.com/TreyBastian/colourize) - An ANSI colour terminal package for Go *(:star:2 | 0clm)*
  * [daviddengcn/**go-colortext**](https://github.com/daviddengcn/go-colortext) - Change the color of console text. *(:star:138 | 0clm)*
  * [jroimartin/**gocui**](https://github.com/jroimartin/gocui) - Minimalist Go library aimed at creating Console User Interfaces. *(:star:358 | 32clm)*
  * [nsf/**termbox-go**](https://github.com/nsf/termbox-go) - Pure Go termbox implementation *(:star:1292 | 6clm)*
  * [apcera/**termtables**](https://github.com/apcera/termtables) - Go ASCII Table Generator, ported from the Ruby terminal-tables library *(:star:42 | 0clm)*
  * [gizak/**termui**](https://github.com/gizak/termui) - Golang terminal dashboard *(:star:4309 | 12clm)*
  * [gosuri/**uilive**](https://github.com/gosuri/uilive) - uilive is a go library for updating terminal output in realtime *(:star:407 | 0clm)*
  * [gosuri/**uiprogress**](https://github.com/gosuri/uiprogress) - A go library to render progress bars in terminal applications *(:star:761 | 0clm)*
  * [gosuri/**uitable**](https://github.com/gosuri/uitable) - A go library to improve readability in terminal apps using tabular data *(:star:269 | 0clm)*

## Configuration
*Libraries for configuration parsing*

  * [olebedev/**config**](https://github.com/olebedev/config) - JSON or YAML configuration wrapper with convenient access methods. *(:star:57 | 0clm)*
  * [paked/**configure**](https://github.com/paked/configure) - Configure is a Go package that gives you easy configuration of your project through redundancy *(:star:6 | 0clm)*
  * [caarlos0/**env**](https://github.com/caarlos0/env) - A KISS way to deal with environment variables in Go. *(:star:77 | 0clm)*
  * [tomazk/**envcfg**](https://github.com/tomazk/envcfg) - Un-marshaling environment variables to Go structs *(:star:75 | 0clm)*
  * [ian-kent/**envconf**](https://github.com/ian-kent/envconf) - Configure Go applications from the environment *(:star:2 | 0clm)*
  * [vrischmann/**envconfig**](https://github.com/vrischmann/envconfig) - Small library to read your configuration from environment variables *(:star:86 | 0clm)*
  * [go-gcfg/**gcfg**](https://github.com/go-gcfg/gcfg) - read INI-style configuration files into Go structs; supports user-defined types and subsections *(:star:23 | 0clm)*
  * [ian-kent/**gofigure**](https://github.com/ian-kent/gofigure) - Go configuration made easy! *(:star:28 | 0clm)*
  * [go-ini/**ini**](https://github.com/go-ini/ini) - Package ini provides INI file read and write functionality in Go. *(:star:203 | 3clm)*
  * [FogCreek/**mini**](https://github.com/FogCreek/mini) - A golang package for parsing ini-style configuration files *(:star:82 | 6clm)*
  * [tucnak/**store**](https://github.com/tucnak/store) - A dead simple configuration manager for Go applications *(:star:36 | 0clm)*
  * [spf13/**viper**](https://github.com/spf13/viper) - Go configuration with fangs *(:star:1242 | 5clm)*

## Continuous Integration
*Tools for help with continuous integration*

  * [drone/**drone**](https://github.com/drone/drone) - Drone is a Continuous Delivery platform built on Docker, written in Go *(:star:6133 | 73clm)*
  * [mattn/**goveralls**](https://github.com/mattn/goveralls) -  *(:star:218 | 2clm)*
  * [go-playground/**overalls**](https://github.com/go-playground/overalls) - :jeans:Multi-Package go project coverprofile for tools like goveralls *(:star:15 | 0clm)*

## CSS Preprocessors
*Libraries for preprocessing CSS files*

  * [c9s/**c6**](https://github.com/c9s/c6) - Compile SASS Faster ! C6 is a SASS-compatible compiler written in Go. *(:star:324 | 0clm)*
  * [yosssi/**gcss**](https://github.com/yosssi/gcss) - Pure Go CSS Preprocessor *(:star:333 | 0clm)*
  * [wellington/**go-libsass**](https://github.com/wellington/go-libsass) - Go wrapper for libsass, the only Sass 3.3 compiler for Go *(:star:16 | 0clm)*

## Data Structures
*Generic datastructures and algorithms in Go.*

  * [willf/**bitset**](https://github.com/willf/bitset) - Go package implementing bitsets *(:star:186 | 2clm)*
  * [zhenjl/**bloom**](https://github.com/zhenjl/bloom) - Bloom filters implemented in Go. *(:star:75 | 0clm)*
  * [tylertreat/**BoomFilters**](https://github.com/tylertreat/BoomFilters) - Probabilistic data structures for processing continuous, unbounded streams. *(:star:660 | 80clm)*
  * [seiflotfy/**cuckoofilter**](https://github.com/seiflotfy/cuckoofilter) - Cuckoo Filter: Practically Better Than Bloom *(:star:135 | 0clm)*
  * [zhenjl/**encoding**](https://github.com/zhenjl/encoding) - Integer Compression Libraries for Go *(:star:40 | 0clm)*
  * [Workiva/**go-datastructures**](https://github.com/Workiva/go-datastructures) -  *(:star:2177 | 174clm)*
  * [hailocab/**go-geoindex**](https://github.com/hailocab/go-geoindex) - Go native library for fast point tracking and K-Nearest queries *(:star:127 | 11clm)*
  * [deckarep/**golang-set**](https://github.com/deckarep/golang-set) - A simple set type for the Go language. *(:star:333 | 0clm)*
  * [ryszard/**goskiplist**](https://github.com/ryszard/goskiplist) - A skip list implementation in Go *(:star:97 | 0clm)*
  * [smartystreets/**mafsa**](https://github.com/smartystreets/mafsa) - Package mafsa implements Minimal Acyclic Finite State Automata in Go, essentially a high-speed, memory-efficient, Unicode-friendly set of strings. *(:star:202 | 0clm)*
  * [gansidui/**skiplist**](https://github.com/gansidui/skiplist) - skiplist for golang *(:star:25 | 0clm)*
  * [derekparker/**trie**](https://github.com/derekparker/trie) - Data structure and relevant algorithms for extremely fast prefix/fuzzy string searching. *(:star:104 | 5clm)*
  * [diegobernardes/**ttlcache**](https://github.com/diegobernardes/ttlcache) - An in-memory LRU string-interface{} map with expiration for golang *(:star:17 | 0clm)*
  * [willf/**bloom**](https://github.com/willf/bloom) - Go package implementing Bloom filters *(:star:218 | 6clm)*

## Database
*Databases implemented in Go.*

  * [boltdb/**bolt**](https://github.com/boltdb/bolt) - An embedded key/value database for Go. *(:star:3616 | 12clm)*
  * [muesli/**cache2go**](https://github.com/muesli/cache2go) - Simple go caching library with expiration capabilities and access counters *(:star:55 | 1clm)*
  * [cockroachdb/**cockroach**](https://github.com/cockroachdb/cockroach) - A Scalable, Survivable, Strongly-Consistent SQL Database *(:star:6109 | 125clm)*
  * [codingsince1985/**couchcache**](https://github.com/codingsince1985/couchcache) - A RESTful caching micro-service in Go backed by Couchbase server *(:star:14 | 0clm)*
  * [peterbourgon/**diskv**](https://github.com/peterbourgon/diskv) - A disk-backed key-value store. *(:star:327 | 0clm)*
  * [couchbase/**goforestdb**](https://github.com/couchbase/goforestdb) - Go bindings for ForestDB *(:star:18 | 4clm)*
  * [bluele/**gcache**](https://github.com/bluele/gcache) - Cache library for golang. It supports expirable Cache, LFU, LRU and ARC. *(:star:89 | 2clm)*
  * [pmylund/**go-cache**](https://github.com/pmylund/go-cache) - An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. *(:star:409 | 0clm)*
  * [syndtr/**goleveldb**](https://github.com/syndtr/goleveldb) - LevelDB key/value database in Go. *(:star:929 | 0clm)*
  * [golang/**groupcache**](https://github.com/golang/groupcache) - groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. *(:star:4166 | 0clm)*
  * [influxdb/**influxdb**](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics *(:star:7275 | 638clm)*
  * [siddontang/**ledisdb**](https://github.com/siddontang/ledisdb) - a high performance NoSQL powered by Go *(:star:1470 | 25clm)*
  * [jmhodges/**levigo**](https://github.com/jmhodges/levigo) - levigo is a Go wrapper for LevelDB *(:star:260 | 1clm)*
  * [prometheus/**prometheus**](https://github.com/prometheus/prometheus) - The Prometheus monitoring system and time series database. *(:star:3806 | 101clm)*
  * [pingcap/**tidb**](https://github.com/pingcap/tidb) - TiDB is a distributed NewSQL database compatible with MySQL protocol  *(:star:3359 | 0clm)*
  * [HouzuoGuo/**tiedot**](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang *(:star:1446 | 5clm)*

*Database tools.*

  * [siddontang/**go-mysql**](https://github.com/siddontang/go-mysql) - a powerful mysql toolset with Go *(:star:161 | 2clm)*
  * [siddontang/**go-mysql-elasticsearch**](https://github.com/siddontang/go-mysql-elasticsearch) - Sync MySQL data into elasticsearch  *(:star:200 | 0clm)*
  * [flike/**kingshard**](https://github.com/flike/kingshard) - A high-performance MySQL proxy *(:star:1467 | 0clm)*
  * [mattes/**migrate**](https://github.com/mattes/migrate) - Database migration handling in Golang *(:star:471 | 7clm)*
  * [2tvenom/**myreplication**](https://github.com/2tvenom/myreplication) - Golang MySql binary log replication listener *(:star:45 | 7clm)*
  * [outbrain/**orchestrator**](https://github.com/outbrain/orchestrator) - MySQL replication topology manager/visualizer *(:star:345 | 44clm)*
  * [sosedoff/**pgweb**](https://github.com/sosedoff/pgweb) - Cross-platform client for PostgreSQL databases *(:star:3195 | 10clm)*
  * [pravasan/**pravasan**](https://github.com/pravasan/pravasan) - Simple Migration Tool - written in Go *(:star:10 | 13clm)*
  * [rubenv/**sql-migrate**](https://github.com/rubenv/sql-migrate) - SQL schema migration tool for Go. *(:star:384 | 2clm)*
  * [youtube/**vitess**](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. *(:star:3201 | 190clm)*

*SQL query builder, libraries for building and using SQL.*

  * [mgutz/**dat**](https://github.com/mgutz/dat) - Go Postgres Data Access Toolkit *(:star:252 | 55clm)*
  * [gchaincl/**dotsql**](https://github.com/gchaincl/dotsql) - A Golang library for using SQL. *(:star:212 | 2clm)*
  * [doug-martin/**goqu**](https://github.com/doug-martin/goqu) - SQL builder and query library for golang *(:star:237 | 6clm)*
  * [go-ozzo/**ozzo-dbx**](https://github.com/go-ozzo/ozzo-dbx) - A Go package that enhances the standard database/sql package by providing powerful data retrieval methods as well as DB-agnostic query building capabilities. *(:star:53 | 0clm)*
  * [variadico/**scaneo**](https://github.com/variadico/scaneo) - Generate Go code to convert database rows into arbitrary structs. *(:star:53 | 0clm)*
  * [elgris/**sqrl**](https://github.com/elgris/sqrl) - Fluent SQL generation for golang *(:star:24 | 10clm)*
  * [Masterminds/**squirrel**](https://github.com/Masterminds/squirrel) - Fluent SQL generation for golang *(:star:732 | 10clm)*

## Database Drivers
*Libraries for connecting and operating databases.*
* Relational Databases

  * [nakagami/**firebirdsql**](https://github.com/nakagami/firebirdsql) - Firebird RDBMS sql driver for Go (golang) *(:star:40 | 12clm)*
  * [mattn/**go-adodb**](https://github.com/mattn/go-adodb) - Microsoft ActiveX Object DataBase driver for go that using exp/sql *(:star:30 | 0clm)*
  * [rounds/**go-bqstreamer**](https://github.com/rounds/go-bqstreamer) - Stream data into Google BigQuery concurrently using InsertAll() *(:star:72 | 0clm)*
  * [denisenkom/**go-mssqldb**](https://github.com/denisenkom/go-mssqldb) - Microsoft SQL server driver written in go language *(:star:284 | 1clm)*
  * [mattn/**go-oci8**](https://github.com/mattn/go-oci8) - oracle driver for go that using database/sql *(:star:115 | 16clm)*
  * [go-sql-driver/**mysql**](https://github.com/go-sql-driver/mysql) - Go-MySQL-Driver is a lightweight and fast MySQL-Driver for Go's (golang) database/sql package *(:star:2035 | 51clm)*
  * [mattn/**go-sqlite3**](https://github.com/mattn/go-sqlite3) - sqlite3 driver for go that using database/sql *(:star:1068 | 1clm)*
  * [minus5/**gofreetds**](https://github.com/minus5/gofreetds) - Go Sql Server database driver. *(:star:33 | 0clm)*
  * [jackc/**pgx**](https://github.com/jackc/pgx) - PostgreSQL client library for Go *(:star:481 | 4clm)*
  * [lib/**pq**](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql *(:star:1880 | 8clm)*

* NoSQL Databases

  * [aerospike/**aerospike-client-go**](https://github.com/aerospike/aerospike-client-go) - Aerospike Client Go  *(:star:151 | 8clm)*
  * [solher/**arangolite**](https://github.com/solher/arangolite) - Lightweight golang driver for ArangoDB *(:star:15 | 0clm)*
  * [google/**cayley**](https://github.com/google/cayley) - An open-source graph database *(:star:6849 | 38clm)*
  * [underarmour/**dynago**](https://github.com/underarmour/dynago) - A DynamoDB client for Go *(:star:11 | 0clm)*
  * [couchbase/**go-couchbase**](https://github.com/couchbase/go-couchbase) - Couchbase client in Go *(:star:198 | 8clm)*
  * [fjl/**go-couchdb**](https://github.com/fjl/go-couchdb) - Yet another CouchDB HTTP API wrapper for Go *(:star:29 | 0clm)*
  * [couchbase/**gocb**](https://github.com/couchbase/gocb) - The Couchbase Go SDK *(:star:152 | 16clm)*
  * [dancannon/**gorethink**](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB *(:star:748 | 46clm)*
  * [cihangir/**neo4j**](https://github.com/cihangir/neo4j) - Neo4j Rest API Client for Go lang *(:star:12 | 0clm)*
  * [davemeehan/**Neo4j-GO**](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang *(:star:60 | 0clm)*
  * [jmcvetta/**neoism**](https://github.com/jmcvetta/neoism) - Neo4j client for Golang *(:star:223 | 2clm)*
  * [garyburd/**redigo**](https://github.com/garyburd/redigo) - Go client for Redis *(:star:1922 | 0clm)*
  * [go-redis/**redis**](https://github.com/go-redis/redis) - Redis client for Golang. *(:star:642 | 12clm)*
  * [hoisie/**redis**](https://github.com/hoisie/redis) - A simple, powerful Redis client for Go *(:star:487 | 0clm)*
  * [bsm/**redeo**](https://github.com/bsm/redeo) - High-performance framework for building redis-protocol compatible TCP servers/services *(:star:32 | 7clm)*

* Search and Analytic Databases

  * [blevesearch/**bleve**](https://github.com/blevesearch/bleve) - A modern text indexing library for go *(:star:2295 | 19clm)*
  * [olivere/**elastic**](https://github.com/olivere/elastic) - Elasticsearch client for Go. *(:star:497 | 10clm)*
  * [mattbaird/**elastigo**](https://github.com/mattbaird/elastigo) - A Go (golang) based Elasticsearch client library. *(:star:681 | 7clm)*
  * [belogik/**goes**](https://github.com/belogik/goes) - A library to interact with Elasticsearch in Go! *(:star:72 | 5clm)*
  * [seiflotfy/**skizze**](https://github.com/seiflotfy/skizze) - A probabilistic data structure service and storage *(:star:7 | 0clm)*

## Date & Time
*Libraries for working with dates and times.*

  * [grsmv/**goweek**](https://github.com/grsmv/goweek) - ISO 8601 compatible library for working with week entities for Go *(:star:3 | 0clm)*
  * [jinzhu/**now**](https://github.com/jinzhu/now) - Now is a time toolkit for golang *(:star:544 | 2clm)*
  * [leekchan/**timeutil**](https://github.com/leekchan/timeutil) - timeutil - useful extensions (Timedelta, Strftime, ...) to the golang's time package *(:star:96 | 0clm)*

## Distributed Systems
*Packages that help with building Distributed Systems.*

  * [svcavallar/**celeriac.v1**](https://github.com/svcavallar/celeriac.v1) - Golang client library for adding support for interacting and monitoring Celery workers, tasks and events. *(:star:7 | 0clm)*
  * [vectaport/**flowgraph**](https://github.com/vectaport/flowgraph) - Ready-send coordination layer on top of goroutines. *(:star:6 | 0clm)*
  * [chrislusf/**glow**](https://github.com/chrislusf/glow) - Glow is an easy-to-use distributed computation system written in Go, similar to Hadoop Map Reduce, Spark, Flink, Samza, etc. Currently just started and not feature rich yet, but should be reliable to run most common cases. *(:star:596 | 0clm)*
  * [dgryski/**go-jump**](https://github.com/dgryski/go-jump) - go-jump: Jump consistent hashing *(:star:53 | 0clm)*
  * [valyala/**gorpc**](https://github.com/valyala/gorpc) - Simple, fast and scalable golang rpc library for high load *(:star:185 | 10clm)*
  * [grpc/**grpc-go**](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC *(:star:1429 | 149clm)*
  * [micro/**micro**](https://github.com/micro/micro) - A microservice toolkit *(:star:636 | 2clm)*
  * [hashicorp/**raft**](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol *(:star:451 | 4clm)*
  * [anacrolix/**torrent**](https://github.com/anacrolix/torrent) - Full-featured BitTorrent client package and utilities *(:star:1029 | 46clm)*
  * [Sioro-Neoku/**go-peerflix**](https://github.com/Sioro-Neoku/go-peerflix) - Go Peerflix *(:star:117 | 0clm)*

## Email
*Libraries that implement email creation and sending*

  * [aymerick/**douceur**](https://github.com/aymerick/douceur) - A simple CSS parser and inliner in Go *(:star:34 | 0clm)*
  * [jordan-wright/**email**](https://github.com/jordan-wright/email) - Robust and flexible email library for Go *(:star:643 | 0clm)*
  * [toorop/**go-dkim**](https://github.com/toorop/go-dkim) - DKIM package for golang *(:star:12 | 0clm)*
  * [hectane/**hectane**](https://github.com/hectane/hectane) - Lightweight SMTP client written in Go *(:star:34 | 0clm)*
  * [mailhog/**MailHog**](https://github.com/mailhog/MailHog) - Web and API based SMTP testing *(:star:591 | 6clm)*
  * [sendgrid/**sendgrid-go**](https://github.com/sendgrid/sendgrid-go) - SendGrid Library to Interface through Go *(:star:153 | 2clm)*
  * [mailhog/**smtp**](https://github.com/mailhog/smtp) - MailHog SMTP Protocol *(:star:18 | 0clm)*

## Embeddable Scripting Languages
*Embedding other languages inside your go code*

  * [PuerkitoBio/**agora**](https://github.com/PuerkitoBio/agora) - a dynamically typed, garbage collected, embeddable programming language built with Go *(:star:203 | 0clm)*
  * [mattn/**anko**](https://github.com/mattn/anko) - Scriptable interpreter written in golang *(:star:232 | 3clm)*
  * [jcla1/**gisp**](https://github.com/jcla1/gisp) - Simple LISP in Go *(:star:313 | 0clm)*
  * [olebedev/**go-duktape**](https://github.com/olebedev/go-duktape) - Duktape JavaScript engine bindings for Go *(:star:264 | 10clm)*
  * [Shopify/**go-lua**](https://github.com/Shopify/go-lua) - A Lua VM in Go *(:star:570 | 12clm)*
  * [deuill/**go-php**](https://github.com/deuill/go-php) - PHP bindings for the Go programming language (Golang) *(:star:89 | 0clm)*
  * [sbinet/**go-python**](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API *(:star:336 | 6clm)*
  * [aarzilli/**golua**](https://github.com/aarzilli/golua) - Go bindings for Lua C API - in progress *(:star:254 | 0clm)*
  * [yuin/**gopher-lua**](https://github.com/yuin/gopher-lua) - GopherLua: VM and compiler for Lua in Go *(:star:1073 | 26clm)*
  * [robertkrimen/**otto**](https://github.com/robertkrimen/otto) - A JavaScript interpreter in Go (golang) *(:star:1933 | 15clm)*
  * [ian-kent/**purl**](https://github.com/ian-kent/purl) - Perl, but fluffy like a cat! *(:star:8 | 0clm)*

## Financial
*Packages for accounting and finance*

  * [leekchan/**accounting**](https://github.com/leekchan/accounting) - money and currency formatting for golang *(:star:183 | 0clm)*
  * [shopspring/**decimal**](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers in go *(:star:222 | 6clm)*

## Forms
*Libraries for working with forms.*

  * [robfig/**bind**](https://github.com/robfig/bind) -  *(:star:9 | 0clm)*
  * [mholt/**binding**](https://github.com/mholt/binding) - Reflectionless data binding for Go's net/http (not yet a stable 1.0, but not likely to change much either) *(:star:424 | 0clm)*
  * [monoculum/**formam**](https://github.com/monoculum/formam) - a package for decode form's values into struct in Go *(:star:42 | 1clm)*
  * [albrow/**forms**](https://github.com/albrow/forms) - A lightweight go library for parsing form data or json from an http.Request. *(:star:51 | 5clm)*
  * [gorilla/**csrf**](https://github.com/gorilla/csrf) - gorilla/csrf provides Cross Site Request Forgery (CSRF) prevention middleware for Go web applications & services. *(:star:90 | 0clm)*
  * [justinas/**nosurf**](https://github.com/justinas/nosurf) - CSRF protection middleware for Go. *(:star:612 | 0clm)*

## Game Development
*Awesome game development libraries.*

  * [paked/**engi**](https://github.com/paked/engi) - A cross-platform game engine written in Go following my interpretation of the Entity Component System paradigm, A fork of ajhager/engi *(:star:74 | 1clm)*
  * [vova616/**GarageEngine**](https://github.com/vova616/GarageEngine) - Game engine written in Go (golang). *(:star:206 | 0clm)*
  * [runningwild/**glop**](https://github.com/runningwild/glop) - Bare-bones osx alternative to sdl *(:star:71 | 0clm)*
  * [beefsack/**go-astar**](https://github.com/beefsack/go-astar) - Go implementation of the A* search algorithm *(:star:135 | 0clm)*
  * [GlenKelley/**go-collada**](https://github.com/GlenKelley/go-collada) - Go package for working with the Collada file format. *(:star:8 | 0clm)*
  * [veandco/**go-sdl2**](https://github.com/veandco/go-sdl2) - SDL2 binding for Go *(:star:299 | 56clm)*
  * [ungerik/**go3d**](https://github.com/ungerik/go3d) - A performance oriented 2D/3D math package for Go *(:star:92 | 5clm)*
  * [xtaci/**gonet**](https://github.com/xtaci/gonet) - a game server skeleton in golang *(:star:606 | 0clm)*
  * [name5566/**leaf**](https://github.com/name5566/leaf) - A game server framework in Go (golang) *(:star:529 | 23clm)*
  * [JoelOtter/**termloop**](https://github.com/JoelOtter/termloop) - Terminal-based game engine for Go, built on top of Termbox *(:star:573 | 0clm)*

## Generation & Generics
*Tools to enhance the language with features like generics via code generation*

  * [clipperhouse/**gen**](https://github.com/clipperhouse/gen) - Type-driven code generation for Go *(:star:635 | 2clm)*
  * [ahmetalpbalkan/**go-linq**](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go *(:star:700 | 0clm)*
  * [rjeczalik/**interfaces**](https://github.com/rjeczalik/interfaces) - Code generation tools for Go. *(:star:60 | 0clm)*
  * [ungerik/**pkgreflect**](https://github.com/ungerik/pkgreflect) - A Go preprocessor for package scoped reflection *(:star:29 | 0clm)*

## Go Compilers
*Tools for compiling Go to other languages*

  * [gopherjs/**gopherjs**](https://github.com/gopherjs/gopherjs) - A compiler from Go to JavaScript for running Go code in a browser *(:star:3442 | 78clm)*
  * [go-llvm/**llgo**](https://github.com/go-llvm/llgo) - LLVM-based compiler for Go *(:star:717 | 0clm)*
  * [tardisgo/**tardisgo**](https://github.com/tardisgo/tardisgo) - Golang->Haxe->CPP/CSharp/Java/JavaScript transpiler   *(:star:301 | 28clm)*

## Goroutines
*Tools for managing and working with Goroutines*

  * [ivpusic/**grpool**](https://github.com/ivpusic/grpool) - Lightweight Goroutine pool *(:star:85 | 0clm)*
  * [go-playground/**pool**](https://github.com/go-playground/pool) - :speedboat: Go consumer goroutine pool for easy goroutine handling + time saving *(:star:34 | 0clm)*
  * [Jeffail/**tunny**](https://github.com/Jeffail/tunny) - A goroutine pool for golang *(:star:265 | 0clm)*

## GUI
*Libraries for building GUI Applications*

  * [go-qml/**qml**](https://github.com/go-qml/qml) - QML support for the Go language *(:star:1595 | 2clm)*
  * [visualfc/**goqt**](https://github.com/visualfc/goqt) - Golang bindings to the Qt cross-platform application framework. *(:star:847 | 0clm)*
  * [deckarep/**gosx-notifier**](https://github.com/deckarep/gosx-notifier) - gosx-notifier is a Go framework for sending desktop notifications to OSX 10.8 or higher *(:star:265 | 0clm)*
  * [gotk3/**gotk3**](https://github.com/gotk3/gotk3) - Go bindings for GTK3 *(:star:46 | 0clm)*
  * [oskca/**sciter**](https://github.com/oskca/sciter) - Golang bindings of Sciter: the Embeddable HTML/CSS/script engine for modern UI development *(:star:116 | 0clm)*
  * [getlantern/**systray**](https://github.com/getlantern/systray) - a cross platfrom Go library to place an icon and menu in the notification area *(:star:40 | 69clm)*
  * [shurcooL/**trayhost**](https://github.com/shurcooL/trayhost) - Cross-platform Go library to place an icon in the host operating system's taskbar. *(:star:28 | 0clm)*
  * [andlabs/**ui**](https://github.com/andlabs/ui) - Platform-native GUI library for Go. *(:star:2631 | 84clm)*
  * [lxn/**walk**](https://github.com/lxn/walk) - A Windows GUI toolkit for the Go Programming Language *(:star:1191 | 0clm)*

## Images
*Libraries for manipulating images.*

  * [h2non/**bimg**](https://github.com/h2non/bimg) - Small Go package for fast high-level image processing using libvips via C bindings *(:star:137 | 0clm)*
  * [pravj/**geopattern**](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string in golang. *(:star:832 | 0clm)*
  * [disintegration/**gift**](https://github.com/disintegration/gift) - Go Image Filtering Toolkit *(:star:684 | 5clm)*
  * [ungerik/**go-cairo**](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library *(:star:48 | 2clm)*
  * [bolknote/**go-gd**](https://github.com/bolknote/go-gd) - Go bingings for GD (http://www.boutell.com/gd/) *(:star:38 | 0clm)*
  * [koyachi/**go-nude**](https://github.com/koyachi/go-nude) - Nudity detection with Go. *(:star:164 | 0clm)*
  * [lazywei/**go-opencv**](https://github.com/lazywei/go-opencv) - Go bindings for OpenCV / 2.x API in gocv / 1.x API in opencv *(:star:337 | 39clm)*
  * [jyotiska/**go-webcolors**](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go *(:star:17 | 0clm)*
  * [gographics/**imagick**](https://github.com/gographics/imagick) - naive Go binding to ImageMagick's MagickWand C API *(:star:377 | 0clm)*
  * [h2non/**imaginary**](https://github.com/h2non/imaginary) - Fast HTTP microservice for high-level image processing. Perfectly fitted for Docker and Heroku. *(:star:655 | 0clm)*
  * [disintegration/**imaging**](https://github.com/disintegration/imaging) - Simple Go image processing package *(:star:741 | 2clm)*
  * [hawx/**img**](https://github.com/hawx/img) - A selection of image manipulation tools *(:star:72 | 1clm)*
  * [donatj/**mpo**](https://github.com/donatj/mpo) - Simple Go-lang JPEG MPO (Multi Picture Object Decoder) *(:star:4 | 0clm)*
  * [thoas/**picfit**](https://github.com/thoas/picfit) - An image resizing server written in Go *(:star:423 | 11clm)*
  * [nfnt/**resize**](https://github.com/nfnt/resize) - Pure golang image resizing  *(:star:969 | 0clm)*
  * [bamiaux/**rez**](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD *(:star:115 | 1clm)*
  * [muesli/**smartcrop**](https://github.com/muesli/smartcrop) - smartcrop implementation in Go *(:star:174 | 1clm)*
  * [ajstarks/**svgo**](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation *(:star:642 | 0clm)*
  * [ftrvxmtrx/**tga**](https://github.com/ftrvxmtrx/tga) - Go package for decoding and encoding TARGA image format *(:star:12 | 0clm)*

## Logging
*Libraries for generating and working with log files.*

  * [golang/**glog**](https://github.com/golang/glog) - Leveled execution logs for Go *(:star:907 | 0clm)*
  * [siddontang/**go-log**](https://github.com/siddontang/go-log) - a golang log lib supports level and multi handlers *(:star:13 | 0clm)*
  * [ian-kent/**go-log**](https://github.com/ian-kent/go-log) - A logger, for Go *(:star:16 | 0clm)*
  * [ventu-io/**go-log-interface**](https://github.com/ventu-io/go-log-interface) - Leveled log interface and the matching facade for the Go stdlib log.Logger *(:star:2 | 0clm)*
  * [apsdehal/**go-logger**](https://github.com/apsdehal/go-logger) - Simple logger for Go programs *(:star:149 | 0clm)*
  * [sadlil/**gologger**](https://github.com/sadlil/gologger) - Simple Logger for golang. Logs Into console, file or ElasticSearch. Simple, easy to use.  *(:star:16 | 0clm)*
  * [apex/**log**](https://github.com/apex/log) - Structured logging package for Go. *(:star:162 | 0clm)*
  * [firstrow/**logvoyage**](https://github.com/firstrow/logvoyage) - LogVoyage - logging SaaS written in GoLang *(:star:47 | 1clm)*
  * [inconshreveable/**log15**](https://github.com/inconshreveable/log15) - Powerful, composable logging for Go *(:star:383 | 1clm)*
  * [chzyer/**logex**](https://github.com/chzyer/logex) - An golang log lib, supports tracking and level, wrap by standard log lib *(:star:22 | 0clm)*
  * [azer/**logger**](https://github.com/azer/logger) - Minimalistic logging library for Go. *(:star:55 | 0clm)*
  * [Sirupsen/**logrus**](https://github.com/Sirupsen/logrus) - Structured, pluggable logging for Go. *(:star:2253 | 20clm)*
  * [sebest/**logrusly**](https://github.com/sebest/logrusly) - Loggly Hooks for GO Logrus logger *(:star:6 | 0clm)*
  * [hashicorp/**logutils**](https://github.com/hashicorp/logutils) - Utilities for slightly better logging in Go (Golang). *(:star:123 | 0clm)*
  * [mgutz/**logxi**](https://github.com/mgutz/logxi) - A 12-factor app logger built for performance and happy development *(:star:224 | 0clm)*
  * [natefinch/**lumberjack**](https://github.com/natefinch/lumberjack) - lumberjack is a rolling logger for Go *(:star:257 | 0clm)*
  * [jbrodriguez/**mlog**](https://github.com/jbrodriguez/mlog) -  *(:star:2 | 0clm)*
  * [cihub/**seelog**](https://github.com/cihub/seelog) - Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting. *(:star:591 | 3clm)*
  * [alexcesaro/**log**](https://github.com/alexcesaro/log) - Logging packages for Go *(:star:27 | 0clm)*
  * [hpcloud/**tail**](https://github.com/hpcloud/tail) - Go package for reading from continously updated files (tail -f) *(:star:416 | 0clm)*
  * [rs/**xlog**](https://github.com/rs/xlog) - xlog is a logger for net/context aware HTTP applications *(:star:55 | 0clm)*

## Machine Learning
*Libraries for Machine Learning.*

  * [jbrukh/**bayesian**](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang. *(:star:282 | 0clm)*
  * [ryanbressler/**CloudForest**](https://github.com/ryanbressler/CloudForest) - Ensembles of decision trees in go/golang. *(:star:321 | 0clm)*
  * [white-pony/**go-fann**](https://github.com/white-pony/go-fann) - Go bindings for FANN, library for artificial neural networks *(:star:68 | 1clm)*
  * [thoj/**go-galib**](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang *(:star:106 | 0clm)*
  * [daviddengcn/**go-pr**](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang. *(:star:33 | 0clm)*
  * [goml/**gobrain**](https://github.com/goml/gobrain) - Neural Networks written in go *(:star:68 | 0clm)*
  * [e-dard/**godist**](https://github.com/e-dard/godist) - Probability distributions and associated methods in Go *(:star:6 | 0clm)*
  * [tomcraven/**goga**](https://github.com/tomcraven/goga) - Golang Genetic Algorithm *(:star:21 | 0clm)*
  * [sjwhitworth/**golearn**](https://github.com/sjwhitworth/golearn) - Machine Learning for Go *(:star:2515 | 0clm)*
  * [cdipaolo/**goml**](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go (and so much more) *(:star:432 | 0clm)*
  * [timkaye11/**goRecommend**](https://github.com/timkaye11/goRecommend) - Collaborative Filtering (CF) Algorithms in Go!  *(:star:31 | 0clm)*
  * [datastream/**libsvm**](https://github.com/datastream/libsvm) - libsvm go version *(:star:33 | 0clm)*
  * [schuyler/**neural-go**](https://github.com/schuyler/neural-go) - A multilayer perceptron network implemented in Go, with training via backpropagation. *(:star:40 | 0clm)*
  * [muesli/**regommend**](https://github.com/muesli/regommend) - Recommendation engine for Go *(:star:86 | 0clm)*
  * [eaigner/**shield**](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go *(:star:74 | 0clm)*

## Messaging
*Libraries that implement messaging systems*

  * [godbus/**dbus**](https://github.com/godbus/dbus) - Native Go bindings for D-Bus *(:star:67 | 0clm)*
  * [olebedev/**emitter**](https://github.com/olebedev/emitter) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins *(:star:26 | 0clm)*
  * [asaskevich/**EventBus**](https://github.com/asaskevich/EventBus) - [Go] Lightweight eventbus with async compatibility for Go *(:star:121 | 0clm)*
  * [ventu-io/**go-longpoll**](https://github.com/ventu-io/go-longpoll) - PubSub queuing with long-polling subscribers for web and other distributed applications *(:star:3 | 0clm)*
  * [TheCreeper/**go-notify**](https://github.com/TheCreeper/go-notify) - Package notify provides an implementation of the Gnome DBus Notification Specification. *(:star:12 | 0clm)*
  * [nsqio/**go-nsq**](https://github.com/nsqio/go-nsq) - The official Go package for NSQ *(:star:525 | 27clm)*
  * [Terry-Mao/**gopush-cluster**](https://github.com/Terry-Mao/gopush-cluster) - Golang push server cluster *(:star:1102 | 0clm)*
  * [RichardKnop/**machinery**](https://github.com/RichardKnop/machinery) - Machinery is an asynchronous task queue/job queue based on distributed message passing. *(:star:691 | 0clm)*
  * [nats-io/**nats**](https://github.com/nats-io/nats) - Golang client for NATS, the cloud native messaging system. *(:star:572 | 5clm)*
  * [dailymotion/**oplog**](https://github.com/dailymotion/oplog) - A generic oplog/replication system for microservices *(:star:58 | 24clm)*
  * [tuxychandru/**pubsub**](https://github.com/tuxychandru/pubsub) - A simple pubsub package for go. *(:star:75 | 0clm)*
  * [Shopify/**sarama**](https://github.com/Shopify/sarama) - Sarama is a Go library for Apache Kafka 0.8 and 0.9  *(:star:817 | 112clm)*
  * [uniqush/**uniqush-push**](https://github.com/uniqush/uniqush-push) - Uniqush is a free and open source software which provides a unified push service for server-side notification to apps on mobile devices. *(:star:767 | 0clm)*
  * [pebbe/**zmq4**](https://github.com/pebbe/zmq4) - A Go interface to ZeroMQ version 4 *(:star:350 | 0clm)*

## Miscellaneous
*These libraries were placed here because none of the other categories seemed to fit*

  * [artyom/**autoflags**](https://github.com/artyom/autoflags) - Populate go command line app flags from config struct *(:star:12 | 0clm)*
  * [digitalcrab/**browscap_go**](https://github.com/digitalcrab/browscap_go) - GoLang Library for Browser Capabilities Project *(:star:12 | 0clm)*
  * [miolini/**datacounter**](https://github.com/miolini/datacounter) - Golang counters for readers/writers *(:star:4 | 0clm)*
  * [go-chat-bot/**bot**](https://github.com/go-chat-bot/bot) - IRC, Slack and Telegram bot written in go *(:star:23 | 0clm)*
  * [hashicorp/**go-multierror**](https://github.com/hashicorp/go-multierror) - A Go (golang) package for representing a list of errors as a single error. *(:star:178 | 0clm)*
  * [ventu-io/**go-shortid**](https://github.com/ventu-io/go-shortid) - Super short, unique, non sequential, URL friendly Ids for Go *(:star:32 | 0clm)*
  * [shirou/**gopsutil**](https://github.com/shirou/gopsutil) - psutil for golang *(:star:902 | 26clm)*
  * [haxpax/**gosms**](https://github.com/haxpax/gosms) - :mailbox_closed: Your own local SMS gateway in Go *(:star:897 | 3clm)*
  * [albrow/**jobs**](https://github.com/albrow/jobs) - A persistent and flexible background jobs library for go. *(:star:306 | 107clm)*
  * [zhulik/**margelet**](https://github.com/zhulik/margelet) - Telegram Bot Framework for Go *(:star:25 | 0clm)*
  * [rjeczalik/**notify**](https://github.com/rjeczalik/notify) - File system event notification library on steroids. *(:star:110 | 82clm)*
  * [go-playground/**stats**](https://github.com/go-playground/stats) - :chart_with_upwards_trend: Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... *(:star:10 | 0clm)*
  * [go-xkg/**xkg**](https://github.com/go-xkg/xkg) - X Keyboard Grabber *(:star:9 | 0clm)*
  * [huandu/**xstrings**](https://github.com/huandu/xstrings) - xstrings: A collection of useful string functions for Go. *(:star:299 | 1clm)*

## Natural Language Processing
*Libraries for working with human languages.*

  * [nuance/**go-nlp**](https://github.com/nuance/go-nlp) - Utilities for working with discrete probability distributions and other tools useful for doing NLP work *(:star:41 | 0clm)*
  * [agonopol/**go-stem**](https://github.com/agonopol/go-stem) - Word Stemming in Go *(:star:33 | 0clm)*
  * [danieldk/**go2vec**](https://github.com/danieldk/go2vec) - Read and use word2vec vectors in Go *(:star:5 | 1clm)*
  * [rjohnsondev/**golibstemmer**](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2 *(:star:10 | 0clm)*
  * [fiam/**gounidecode**](https://github.com/fiam/gounidecode) - Unicode transliterator for #golang *(:star:35 | 0clm)*
  * [goodsign/**icu**](https://github.com/goodsign/icu) - Cgo binding for icu4c library *(:star:15 | 0clm)*
  * [goodsign/**libtextcat**](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library *(:star:8 | 0clm)*
  * [awsong/**MMSEGO**](https://github.com/awsong/MMSEGO) - Chinese word splitting algorithm MMSEG in GO *(:star:44 | 0clm)*
  * [rookii/**paicehusk**](https://github.com/rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm *(:star:10 | 0clm)*
  * [a2800276/**porter**](https://github.com/a2800276/porter) - porter stemmer *(:star:2 | 0clm)*
  * [zhenjl/**porter2**](https://github.com/zhenjl/porter2) - High Performance Porter2 Stemmer *(:star:19 | 0clm)*
  * [blevesearch/**segment**](https://github.com/blevesearch/segment) - A Go library for performing Unicode Text Segmentation as described in Unicode Standard Annex #29 *(:star:14 | 0clm)*
  * [neurosnap/**sentences**](https://github.com/neurosnap/sentences) - A multilingual command line sentence tokenizer in Golang *(:star:103 | 0clm)*
  * [goodsign/**snowball**](https://github.com/goodsign/snowball) - Cgo binding for Snowball C library *(:star:13 | 0clm)*
  * [dchest/**stemmer**](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers. *(:star:25 | 0clm)*
  * [pebbe/**textcat**](https://github.com/pebbe/textcat) - A Go package for n-gram based text categorization, with support for utf-8 and raw text *(:star:39 | 0clm)*

## Networking
*Libraries for working with various layers of the network*

  * [mdlayher/**arp**](https://github.com/mdlayher/arp) - Package arp implements the ARP protocol, as described in RFC 826. MIT Licensed. *(:star:26 | 0clm)*
  * [stabbycutyou/**buffstreams**](https://github.com/stabbycutyou/buffstreams) - A library to simplify writing applications using TCP sockets to stream protobuff messages *(:star:122 | 0clm)*
  * [zubairhamed/**canopus**](https://github.com/zubairhamed/canopus) - CoAP Client/Server implementing RFC 7252 (For Go/Golang) *(:star:19 | 12clm)*
  * [mdlayher/**dhcp6**](https://github.com/mdlayher/dhcp6) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. MIT Licensed. *(:star:7 | 0clm)*
  * [miekg/**dns**](https://github.com/miekg/dns) - DNS library in Go *(:star:1320 | 42clm)*
  * [mdlayher/**ethernet**](https://github.com/mdlayher/ethernet) - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. *(:star:9 | 0clm)*
  * [valyala/**fasthttp**](https://github.com/valyala/fasthttp) - Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http *(:star:1635 | 0clm)*
  * [jlaffaye/**ftp**](https://github.com/jlaffaye/ftp) - FTP client package for Go *(:star:110 | 0clm)*
  * [hashicorp/**go-getter**](https://github.com/hashicorp/go-getter) - Package for downloading things from a string URL using a variety of protocols. *(:star:202 | 0clm)*
  * [ccding/**go-stun**](https://github.com/ccding/go-stun) - a go implementation of the STUN client (RFC 3489 and RFC 5389) *(:star:62 | 0clm)*
  * [sunwxg/**golibwireshark**](https://github.com/sunwxg/golibwireshark) -  *(:star:3 | 0clm)*
  * [google/**gopacket**](https://github.com/google/gopacket) - Provides packet processing capabilities for Go *(:star:472 | 0clm)*
  * [akrennmair/**gopcap**](https://github.com/akrennmair/gopcap) - A simple wrapper around libpcap for the Go programming language *(:star:197 | 2clm)*
  * [sunwxg/**goshark**](https://github.com/sunwxg/goshark) -  *(:star:2 | 0clm)*
  * [soniah/**gosnmp**](https://github.com/soniah/gosnmp) - An SNMP library written in GoLang. *(:star:97 | 3clm)*
  * [gansidui/**gotcp**](https://github.com/gansidui/gotcp) - A Go package for quickly building tcp servers *(:star:128 | 0clm)*
  * [koofr/**graval**](https://github.com/koofr/graval) - An experimental go FTP server framework *(:star:8 | 4clm)*
  * [ian-kent/**linkio**](https://github.com/ian-kent/linkio) - Simulate network link speed *(:star:15 | 0clm)*
  * [hashicorp/**mdns**](https://github.com/hashicorp/mdns) - Simple mDNS client/server library in Golang *(:star:206 | 1clm)*
  * [aybabtme/**portproxy**](https://github.com/aybabtme/portproxy) - TCP proxy, highjacks HTTP to allow CORS *(:star:17 | 0clm)*
  * [mdlayher/**raw**](https://github.com/mdlayher/raw) - Package raw enables reading and writing data at the device driver level for a network interface.  MIT Licensed. *(:star:14 | 0clm)*
  * [pkg/**sftp**](https://github.com/pkg/sftp) - SFTP support for the go.crypto/ssh package *(:star:220 | 1clm)*
  * [eduardonunesp/**sslb**](https://github.com/eduardonunesp/sslb) - Golang Super Simple Load Balance *(:star:56 | 0clm)*
  * [firstrow/**tcp_server**](https://github.com/firstrow/tcp_server) - GoLang simple TCP server *(:star:30 | 2clm)*
  * [anacrolix/**utp**](https://github.com/anacrolix/utp) - Implements uTP, the micro transport protocol as used with Bittorrent *(:star:42 | 9clm)*

## OpenGL
*Libraries for using OpenGL in Go.*

  * [go-gl/**gl**](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow) *(:star:154 | 3clm)*
  * [go-gl/**glfw**](https://github.com/go-gl/glfw) - Go bindings for GLFW 3 *(:star:212 | 6clm)*
  * [goxjs/**gl**](https://github.com/goxjs/gl) - Go cross-platform OpenGL bindings. *(:star:57 | 0clm)*
  * [goxjs/**glfw**](https://github.com/goxjs/glfw) - Go cross-platform glfw library for creating an OpenGL context and receiving events. *(:star:21 | 7clm)*
  * [go-gl/**mathgl**](https://github.com/go-gl/mathgl) - A pure Go 3D math library. *(:star:104 | 0clm)*

## ORM
*Libraries that implement Object-Relational Mapping or datamapping techniques.*

  * [astaxie/**beedb**](https://github.com/astaxie/beedb) - beedb is a go ORM,support database/sql interface，pq/mysql/sqlite *(:star:586 | 0clm)*
  * [gosuri/**go-store**](https://github.com/gosuri/go-store) - A simple and fast Redis backed key-value store library for Go *(:star:69 | 0clm)*
  * [cosiner/**gomodel**](https://github.com/cosiner/gomodel) - A lightweight, fast, orm-like library helps interactive with database *(:star:46 | 0clm)*
  * [jinzhu/**gorm**](https://github.com/jinzhu/gorm) - The fantastic ORM library for Golang, aims to be developer friendly *(:star:3202 | 69clm)*
  * [go-gorp/**gorp**](https://github.com/go-gorp/gorp) - Go Relational Persistence - an ORM-ish library for Go *(:star:1922 | 46clm)*
  * [coocood/**qbs**](https://github.com/coocood/qbs) - QBS stands for Query By Struct. A Go ORM. *(:star:373 | 0clm)*
  * [upper/**db**](https://github.com/upper/db) - Magic-free ORM-like package for Go. *(:star:304 | 12clm)*
  * [go-xorm/**xorm**](https://github.com/go-xorm/xorm) - Simple and Powerful ORM for Go, support mysql,postgres,tidb,sqlite3,mssql,oracle *(:star:1024 | 16clm)*
  * [albrow/**zoom**](https://github.com/albrow/zoom) - A blazing-fast datastore and querying engine for Go built on Redis. *(:star:119 | 0clm)*

## Package Management
*Libraries for package and dependency management.*

  * [LyricalSecurity/**gigo**](https://github.com/LyricalSecurity/gigo) - GIGO: PIP for GO *(:star:188 | 0clm)*
  * [Masterminds/**glide**](https://github.com/Masterminds/glide) - Vendor Package Management for Golang *(:star:930 | 0clm)*
  * [tools/**godep**](https://github.com/tools/godep) - dependency tool for go *(:star:3063 | 0clm)*
  * [mattn/**gom**](https://github.com/mattn/gom) - Go Manager - bundle for go *(:star:1009 | 0clm)*
  * [nitrous-io/**goop**](https://github.com/nitrous-io/goop) - A simple dependency manager for Go (golang), inspired by Bundler. *(:star:741 | 0clm)*
  * [gpmgo/**gopm**](https://github.com/gpmgo/gopm) - Go Package Manager (gopm) is a package manager and build tool for Go. *(:star:862 | 0clm)*
  * [pote/**gpm**](https://github.com/pote/gpm) - Barebones dependency manager for Go. *(:star:891 | 1clm)*
  * [VividCortex/**johnny-deps**](https://github.com/VividCortex/johnny-deps) - Barebones dependency manager for Go. *(:star:215 | 0clm)*
  * [jingweno/**nut**](https://github.com/jingweno/nut) - Vendor Go dependencies *(:star:243 | 14clm)*
  * [DamnWidget/**VenGO**](https://github.com/DamnWidget/VenGO) - Create and manage Isolated Virtual Environments for Go *(:star:84 | 3clm)*

## Query Language

  * [tmc/**graphql**](https://github.com/tmc/graphql) - graphql parser + utilities *(:star:30 | 0clm)*
  * [sevki/**graphql**](https://github.com/sevki/graphql) - GraphQL implementation in go *(:star:27 | 0clm)*
  * [chris-ramon/**graphql-go**](https://github.com/chris-ramon/graphql-go) - An implementation of GraphQL for Go / Golang *(:star:276 | 0clm)*
  * [elgs/**jsonql**](https://github.com/elgs/jsonql) - JSON query expression library in Golang. *(:star:24 | 0clm)*

## Resource Embedding

  * [UnnoTed/**fileb0x**](https://github.com/UnnoTed/fileb0x) - simple customizable tool to embed files in go *(:star:19 | 0clm)*
  * [jteeuwen/**go-bindata**](https://github.com/jteeuwen/go-bindata) - A small utility which generates Go code from any file. Useful for embedding binary data in a Go program. *(:star:1561 | 2clm)*
  * [pyros2097/**go-embed**](https://github.com/pyros2097/go-embed) - Generates go code to embed resource files into your library or executable *(:star:1 | 0clm)*
  * [omeid/**go-resources**](https://github.com/omeid/go-resources) - Unfancy resources embedding for Go with out of box http.FileSystem support. *(:star:95 | 3clm)*
  * [GeertJohan/**go.rice**](https://github.com/GeertJohan/go.rice) - go.rice is a Go package that makes working with resources such as html,js,css,images,templates, etc very easy. *(:star:663 | 11clm)*
  * [go-playground/**statics**](https://github.com/go-playground/statics) - :file_folder: Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks *(:star:30 | 0clm)*
  * [shurcooL/**vfsgen**](https://github.com/shurcooL/vfsgen) - Takes an input http.FileSystem (likely at go generate time) and generates Go code that statically implements it. *(:star:65 | 2clm)*

## Science and Data Analysis
*Libraries for scientific computing and data analyzing.*

  * [ziutek/**blas**](https://github.com/ziutek/blas) - Go implementation of BLAS (Basic Linear Algebra Subprograms) *(:star:88 | 0clm)*
  * [vdobler/**chart**](https://github.com/vdobler/chart) - Provide basic charts in go *(:star:276 | 0clm)*
  * [soniah/**evaler**](https://github.com/soniah/evaler) - Implements a simple floating point arithmetic expression evaluator in Go (golang). *(:star:20 | 0clm)*
  * [VividCortex/**ewma**](https://github.com/VividCortex/ewma) - Exponentially Weighted Moving Average algorithms for Go. *(:star:144 | 0clm)*
  * [skelterjohn/**geom**](https://github.com/skelterjohn/geom) - 2d geometry for golang *(:star:32 | 0clm)*
  * [mjibson/**go-dsp**](https://github.com/mjibson/go-dsp) - Digital Signal Processing for Go *(:star:367 | 0clm)*
  * [skelterjohn/**go.matrix**](https://github.com/skelterjohn/go.matrix) - linear algebra for go *(:star:233 | 0clm)*
  * [anschelsc/**gofrac**](https://github.com/anschelsc/gofrac) - A fractions library for go (http://golang.org) *(:star:2 | 0clm)*
  * [VividCortex/**gohistogram**](https://github.com/VividCortex/gohistogram) - Streaming approximate histograms in Go *(:star:61 | 0clm)*
  * [gonum/**matrix**](https://github.com/gonum/matrix) - Matrix packages for the Go language. *(:star:185 | 20clm)*
  * [gonum/**plot**](https://github.com/gonum/plot) - A repository for plotting and visualizing data *(:star:178 | 22clm)*
  * [gyuho/**goraph**](https://github.com/gyuho/goraph) - Package goraph implements graph, tree data structures and algorithms. *(:star:181 | 8clm)*
  * [alixaxel/**pagerank**](https://github.com/alixaxel/pagerank) - Weighted PageRank implementation in Go *(:star:10 | 0clm)*
  * [montanaflynn/**stats**](https://github.com/montanaflynn/stats) - A statistics package with common functions that are missing from the Golang standard library.  *(:star:557 | 0clm)*
  * [nytlabs/**streamtools**](https://github.com/nytlabs/streamtools) - tools for working with streams of data *(:star:1266 | 0clm)*
  * [spate/**vectormath**](https://github.com/spate/vectormath) - Vectormath for Go *(:star:53 | 0clm)*

## Security
*Libraries that are used to help make your application more secure.*

  * [hlandau/**acme**](https://github.com/hlandau/acme) - :lock: Automatic certificate acquisition tool for ACME (Let's Encrypt) *(:star:525 | 0clm)*
  * [jaredfolkins/**badactor**](https://github.com/jaredfolkins/badactor) - BadActor.org An in-memory application driven jailer written in Go *(:star:158 | 1clm)*
  * [hillu/**go-yara**](https://github.com/hillu/go-yara) - Go bindings for YARA *(:star:20 | 30clm)*
  * [xenolf/**lego**](https://github.com/xenolf/lego) - Let's Encrypt client and ACME library written in Go *(:star:890 | 0clm)*
  * [hlandau/**passlib**](https://github.com/hlandau/passlib) - :key: Idiotproof golang password validation library inspired by Python's passlib *(:star:74 | 0clm)*
  * [elithrar/**simple-scrypt**](https://github.com/elithrar/simple-scrypt) - A convenience library for generating, comparing and inspecting password hashes using the scrypt KDF in Go. *(:star:58 | 0clm)*

## Serialization
*Libraries and tools for binary serialization*

  * [glycerine/**go-capnproto**](https://github.com/glycerine/go-capnproto) - Cap'n Proto library and parser for go. This is go-capnproto-1.0, and does not have rpc. See https://github.com/zombiezen/go-capnproto2 for 2.0 which has rpc and capabilities. *(:star:239 | 12clm)*
  * [glycerine/**bambam**](https://github.com/glycerine/bambam) - auto-generate capnproto schema from your golang source files. Depends on go-capnproto-1.0 at https://github.com/glycerine/go-capnproto *(:star:51 | 9clm)*
  * [ugorji/**go**](https://github.com/ugorji/go) - idiomatic codec and rpc lib for msgpack, cbor, json, etc. msgpack.org[Go] *(:star:503 | 2clm)*
  * [gogo/**protobuf**](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets *(:star:327 | 3clm)*
  * [golang/**protobuf**](https://github.com/golang/protobuf) - Go support for Google's protocol buffers *(:star:717 | 8clm)*
  * [mitchellh/**mapstructure**](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures. *(:star:515 | 0clm)*
  * [yvasiyarov/**php_session_decoder**](https://github.com/yvasiyarov/php_session_decoder) - PHP session encoder/decoder written in Go *(:star:44 | 2clm)*
  * [tuvistavie/**structomap**](https://github.com/tuvistavie/structomap) - Easily and dynamically generate maps from Go static structures *(:star:26 | 0clm)*

## Server Applications

  * [xyproto/**algernon**](https://github.com/xyproto/algernon) - :diamond_shape_with_a_dot_inside: HTTP/2 web/application server with Lua support *(:star:201 | 0clm)*
  * [mholt/**caddy**](https://github.com/mholt/caddy) - Fast, cross-platform HTTP/2 web server with automatic HTTPS *(:star:3826 | 23clm)*
  * [cortesi/**devd**](https://github.com/cortesi/devd) -  A local webserver for developers *(:star:1883 | 0clm)*
  * [coreos/**etcd**](https://github.com/coreos/etcd) - Distributed reliable key-value store for the most critical data of a distributed system *(:star:8600 | 193clm)*
  * [minio/**minio**](https://github.com/minio/minio) - Minio is a distributed object storage server written in Golang. *(:star:500 | 213clm)*
  * [minio/**minio-xl**](https://github.com/minio/minio-xl) - Cloud Storage Server for Petascale. *(:star:14 | 0clm)*
  * [sci4me/**yakvs**](https://github.com/sci4me/yakvs) - YAKVS (Yet Another Key Value Store) *(:star:3 | 0clm)*

## Template Engines
*Libraries and tools for templating and lexing.*

  * [yosssi/**ace**](https://github.com/yosssi/ace) - HTML template engine for Go *(:star:409 | 0clm)*
  * [eknkc/**amber**](https://github.com/eknkc/amber) - Amber is an elegant templating engine for Go Programming Language, inspired from HAML and Jade *(:star:556 | 0clm)*
  * [dskinner/**damsel**](https://github.com/dskinner/damsel) - Package damsel provides html outlining via css-selectors and common template functionality. *(:star:16 | 0clm)*
  * [benbjohnson/**ego**](https://github.com/benbjohnson/ego) - An ERB-style templating language for Go. *(:star:235 | 1clm)*
  * [ziutek/**kasia.go**](https://github.com/ziutek/kasia.go) - Templating system for HTML and other text documents - go implementation *(:star:66 | 0clm)*
  * [hoisie/**mustache**](https://github.com/hoisie/mustache) - The mustache template language in Go *(:star:778 | 0clm)*
  * [flosch/**pongo2**](https://github.com/flosch/pongo2) - Django-syntax like template-engine for Go *(:star:644 | 5clm)*
  * [aymerick/**raymond**](https://github.com/aymerick/raymond) - Handlebars for golang *(:star:69 | 0clm)*
  * [sipin/**gorazor**](https://github.com/sipin/gorazor) - Razor view engine for Golang *(:star:473 | 0clm)*
  * [robfig/**soy**](https://github.com/robfig/soy) - Go implementation for Soy templates (Google Closure templates) *(:star:106 | 0clm)*

## Testing
*Libraries for testing codebases and generating test data.*
* Testing Frameworks

  * [go-playground/**assert**](https://github.com/go-playground/assert) - :exclamation:Basic Assertion Library used along side native go testing, with building blocks for custom assertions *(:star:5 | 0clm)*
  * [bmizerany/**assert**](https://github.com/bmizerany/assert) - Asserts to Go testing *(:star:119 | 0clm)*
  * [marioidival/**bro**](https://github.com/marioidival/bro) - bro watch files in directory and run tests for them *(:star:10 | 0clm)*
  * [verdverm/**frisby**](https://github.com/verdverm/frisby) - API testing framework inspired by frisby-js *(:star:120 | 0clm)*
  * [zimmski/**go-mutesting**](https://github.com/zimmski/go-mutesting) - Mutation testing for Go source code *(:star:33 | 4clm)*
  * [franela/**goblin**](https://github.com/franela/goblin) - Minimal and Beautiful Go testing framework *(:star:246 | 0clm)*
  * [DATA-DOG/**godog**](https://github.com/DATA-DOG/godog) - BDD Behat, Cucumber like gherkin feature runner for golang *(:star:35 | 0clm)*
  * [orfjackal/**gospec**](https://github.com/orfjackal/gospec) - Testing framework for Go. Allows writing self-documenting tests/specifications, and executes them concurrently and safely isolated. [UNMAINTAINED] *(:star:110 | 0clm)*
  * [stesla/**gospecify**](https://github.com/stesla/gospecify) - A BDD library for Go *(:star:50 | 0clm)*
  * [rdrdr/**hamcrest**](https://github.com/rdrdr/hamcrest) - Hamcrest matchers for the Go programming language *(:star:21 | 0clm)*
  * [yookoala/**restit**](https://github.com/yookoala/restit) - A Go micro framework to help writing RESTful API integration test *(:star:22 | 0clm)*
  * [stretchr/**testify**](https://github.com/stretchr/testify) - A sacred extension to the standard go testing package *(:star:1589 | 10clm)*

* Mock

  * [maxbrunsfeld/**counterfeiter**](https://github.com/maxbrunsfeld/counterfeiter) - A tool for generating self-contained, type-safe test doubles in go *(:star:49 | 0clm)*
  * [DATA-DOG/**go-sqlmock**](https://github.com/DATA-DOG/go-sqlmock) - Sql mock driver for golang to test database interactions *(:star:160 | 0clm)*
  * [DATA-DOG/**go-txdb**](https://github.com/DATA-DOG/go-txdb) - Single transaction sql.Driver for GO *(:star:1 | 0clm)*
  * [golang/**mock**](https://github.com/golang/mock) - GoMock is a mocking framework for the Go programming language. *(:star:172 | 0clm)*
  * [tv42/**mockhttp**](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter *(:star:17 | 0clm)*

* Fuzzing and delta-debugging/reducing/shrinking

  * [dvyukov/**go-fuzz**](https://github.com/dvyukov/go-fuzz) - Randomized testing for Go *(:star:1257 | 0clm)*
  * [google/**gofuzz**](https://github.com/google/gofuzz) - Fuzz testing for go. *(:star:205 | 0clm)*
  * [zimmski/**tavor**](https://github.com/zimmski/tavor) - A generic fuzzing and delta-debugging framework *(:star:137 | 51clm)*

## Text Processing
*Libraries for parsing and manipulating texts.*
* Specific Formats

  * [russross/**blackfriday**](https://github.com/russross/blackfriday) - Blackfriday: a markdown processor for Go *(:star:1680 | 6clm)*
  * [microcosm-cc/**bluemonday**](https://github.com/microcosm-cc/bluemonday) - bluemonday: a fast golang HTML sanitizer (inspired by the OWASP Java HTML Sanitizer) to scrub user generated content of XSS *(:star:305 | 0clm)*
  * [endeveit/**enca**](https://github.com/endeveit/enca) - Minimal cgo bindings for libenca *(:star:2 | 0clm)*
  * [alixaxel/**genex**](https://github.com/alixaxel/genex) - Genex package for Go *(:star:35 | 0clm)*
  * [dustin/**go-humanize**](https://github.com/dustin/go-humanize) - Go Humans! (formatters for units to human friendly sizes) *(:star:606 | 0clm)*
  * [adrianmo/**go-nmea**](https://github.com/adrianmo/go-nmea) - NMEA parser library for Golang *(:star:12 | 0clm)*
  * [jteeuwen/**go-pkg-rss**](https://github.com/jteeuwen/go-pkg-rss) - This package reads RSS and Atom feeds and provides a caching mechanism that adheres to the feed specs. *(:star:247 | 2clm)*
  * [jteeuwen/**go-pkg-xmlx**](https://github.com/jteeuwen/go-pkg-xmlx) - Extension to the standard Go XML package. Maintains a node tree that allows forward/backwards browsing and exposes some simple single/multi-node search functions. *(:star:100 | 2clm)*
  * [mattn/**go-runewidth**](https://github.com/mattn/go-runewidth) -  *(:star:51 | 2clm)*
  * [awalterschulze/**gographviz**](https://github.com/awalterschulze/gographviz) - Parses the Graphviz DOT language in golang *(:star:32 | 0clm)*
  * [polera/**gonameparts**](https://github.com/polera/gonameparts) - Takes a full name and splits it into individual name parts *(:star:15 | 0clm)*
  * [PuerkitoBio/**goquery**](https://github.com/PuerkitoBio/goquery) - A little like that j-thing, only in Go. *(:star:2332 | 3clm)*
  * [zach-klippenstein/**goregen**](https://github.com/zach-klippenstein/goregen) - randexp for Go. *(:star:14 | 0clm)*
  * [endeveit/**guesslanguage**](https://github.com/endeveit/guesslanguage) - Guess the natural language of a text in Go *(:star:16 | 0clm)*
  * [clbanning/**mxj**](https://github.com/clbanning/mxj) - Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages. *(:star:100 | 0clm)*
  * [gosimple/**slug**](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support. *(:star:65 | 0clm)*
  * [avelino/**slugify**](https://github.com/avelino/slugify) - A Go slugify application that handles string *(:star:9 | 0clm)*
  * [BurntSushi/**toml**](https://github.com/BurntSushi/toml) - TOML parser for Golang with reflection. *(:star:755 | 2clm)*

* Utility

  * [bndr/**gotabulate**](https://github.com/bndr/gotabulate) - Gotabulate - Easily pretty-print your tabular data with Go *(:star:111 | 0clm)*
  * [codemodus/**kace**](https://github.com/codemodus/kace) - Common case conversions covering common initialisms. *(:star:1 | 0clm)*
  * [codemodus/**parth**](https://github.com/codemodus/parth) - Path parsing / URL segmentation handling. *(:star:5 | 0clm)*
  * [mvdan/**xurls**](https://github.com/mvdan/xurls) - Extract urls from text *(:star:172 | 40clm)*

## Third-party APIs
*Libraries for accessing third party APIs.*

  * [ChimeraCoder/**anaconda**](https://github.com/ChimeraCoder/anaconda) - A Go client library for the Twitter 1.1 API *(:star:473 | 14clm)*
  * [aws/**aws-sdk-go**](https://github.com/aws/aws-sdk-go) - AWS SDK for the Go programming language. *(:star:2290 | 98clm)*
  * [naegelejd/**brewerydb**](https://github.com/naegelejd/brewerydb) - Go library for http://www.brewerydb.com/ API *(:star:10 | 0clm)*
  * [samuelcouch/**clarifai**](https://github.com/samuelcouch/clarifai) - Clarifai library for Go *(:star:27 | 0clm)*
  * [bwmarrin/**discordgo**](https://github.com/bwmarrin/discordgo) -  (Golang) Go bindings for Discord *(:star:31 | 0clm)*
  * [huandu/**facebook**](https://github.com/huandu/facebook) - A Facebook Graph API SDK Library For Golang *(:star:271 | 0clm)*
  * [emiddleton/**gads**](https://github.com/emiddleton/gads) - Google Adwords API for Go *(:star:11 | 1clm)*
  * [bit4bit/**gami**](https://github.com/bit4bit/gami) - GO - Asterisk AMI Interface *(:star:11 | 12clm)*
  * [Aorioli/**gcm**](https://github.com/Aorioli/gcm) - Google Cloud Messaging for application servers implemented using the Go programming language. *(:star:26 | 0clm)*
  * [codingsince1985/**geo-golang**](https://github.com/codingsince1985/geo-golang) - (reverse) geocoding service in Go *(:star:108 | 17clm)*
  * [neuegram/**ghost**](https://github.com/neuegram/ghost) - A Go library for Snapchat's API *(:star:13 | 3clm)*
  * [google/**go-github**](https://github.com/google/go-github) - Go library for accessing the GitHub API *(:star:1463 | 3clm)*
  * [gambol99/**go-marathon**](https://github.com/gambol99/go-marathon) - A GO API library for working with Marathon *(:star:62 | 64clm)*
  * [mitchellh/**goamz**](https://github.com/mitchellh/goamz) - Golang Amazon Library *(:star:610 | 11clm)*
  * [michiwend/**gomusicbrainz**](https://github.com/michiwend/gomusicbrainz) - a Go (Golang) MusicBrainz WS2 client library - work in progress *(:star:15 | 0clm)*
  * [google/**google-api-go-client**](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go *(:star:526 | 5clm)*
  * [chonthu/**go-google-analytics**](https://github.com/chonthu/go-google-analytics) - Simple Reporting for Google Analytics *(:star:4 | 0clm)*
  * [GoogleCloudPlatform/**gcloud-golang**](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library *(:star:250 | 8clm)*
  * [jsgilmore/**gostorm**](https://github.com/jsgilmore/gostorm) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. *(:star:61 | 1clm)*
  * [andybons/**hipchat**](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API. *(:star:94 | 0clm)*
  * [daneharrigan/**hipchat**](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP *(:star:87 | 0clm)*
  * [Medium/**medium-sdk-go**](https://github.com/Medium/medium-sdk-go) - A Golang SDK for Medium's OAuth2 API *(:star:34 | 0clm)*
  * [minio/**minio-go**](https://github.com/minio/minio-go) - Minio Go Library for Amazon S3 compatible cloud storage *(:star:71 | 0clm)*
  * [dukex/**mixpanel**](https://github.com/dukex/mixpanel) - Golang Mixpanel Client *(:star:12 | 0clm)*
  * [logpacker/**paypalsdk**](https://github.com/logpacker/paypalsdk) - Golang client for PayPal REST API *(:star:59 | 0clm)*
  * [playlyfe/**playlyfe-go-sdk**](https://github.com/playlyfe/playlyfe-go-sdk) - This is the official Playlyfe Golang Sdk *(:star:0 | 0clm)*
  * [gregdel/**pushover**](https://github.com/gregdel/pushover) - Go wrapper arround the Pushover API *(:star:10 | 7clm)*
  * [Omie/**rrdaclient**](https://github.com/Omie/rrdaclient) - Go bindings for RRDA https://github.com/fcambus/rrda *(:star:3 | 0clm)*
  * [rapito/**go-shopify**](https://github.com/rapito/go-shopify) - Simple Shopify API for the Go Programming Language *(:star:9 | 0clm)*
  * [nlopes/**slack**](https://github.com/nlopes/slack) - Slack API in Go *(:star:299 | 17clm)*
  * [sergiotapia/**smitego**](https://github.com/sergiotapia/smitego) - SmiteGo is an API wrapper for the Smite game from HiRez. It is written in Go! *(:star:6 | 0clm)*
  * [rapito/**go-spotify**](https://github.com/rapito/go-spotify) - Go library for the Spotify Web API *(:star:5 | 0clm)*
  * [sostronk/**go-steam**](https://github.com/sostronk/go-steam) - Go library for querying Source servers *(:star:7 | 6clm)*
  * [stripe/**stripe-go**](https://github.com/stripe/stripe-go) - Go client for the Stripe API *(:star:330 | 14clm)*
  * [tucnak/**telebot**](https://github.com/tucnak/telebot) - Telegram bot framework written in Go *(:star:103 | 0clm)*
  * [Syfaro/**telegram-bot-api**](https://github.com/Syfaro/telegram-bot-api) - Golang bindings for the Telegram Bot API *(:star:117 | 0clm)*
  * [jbrodriguez/**go-tmdb**](https://github.com/jbrodriguez/go-tmdb) -  *(:star:3 | 0clm)*
  * [poorny/**translate**](https://github.com/poorny/translate) - Go online translation package *(:star:9 | 0clm)*
  * [mattcunningham/**gumblr**](https://github.com/mattcunningham/gumblr) - A Go Wrapper for the Tumblr v2 API *(:star:0 | 0clm)*
  * [go-playground/**webhooks**](https://github.com/go-playground/webhooks) - :fishing_pole_and_fish: Webhook receiver for GitHub and Bitbucket *(:star:8 | 0clm)*

## Utilities
*General utilities and tools to make your life easier.*

  * [topfreegames/**apm**](https://github.com/topfreegames/apm) - APM is a process manager for Golang applications. *(:star:25 | 0clm)*
  * [tmrts/**boilr**](https://github.com/tmrts/boilr) - :zap: Blazingly fast CLI tool for creating projects from boilerplate templates *(:star:19 | 0clm)*
  * [rakyll/**coop**](https://github.com/rakyll/coop) - Cheat sheet for some of the common concurrent flows in Go *(:star:1127 | 0clm)*
  * [ulule/**deepcopier**](https://github.com/ulule/deepcopier) - simple struct copying for golang *(:star:79 | 0clm)*
  * [derekparker/**delve**](https://github.com/derekparker/delve) - Delve is a debugger for the Go programming language. *(:star:3012 | 21clm)*
  * [digitalcrab/**fastlz**](https://github.com/digitalcrab/fastlz) - Wrap over FastLz for GoLang *(:star:4 | 1clm)*
  * [h2non/**filetype**](https://github.com/h2non/filetype) - Small Go package to infer the file type checking the magic numbers signature *(:star:45 | 0clm)*
  * [junegunn/**fzf**](https://github.com/junegunn/fzf) - :cherry_blossom: A command-line fuzzy finder written in Go *(:star:3838 | 11clm)*
  * [go-playground/**generate**](https://github.com/go-playground/generate) - :runner:runs go generate recursively on a specified path or environment variable and can filter by regex *(:star:1 | 0clm)*
  * [rk/**go-cron**](https://github.com/rk/go-cron) - A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. *(:star:104 | 0clm)*
  * [tj/**go-debug**](https://github.com/tj/go-debug) - Conditional debug logging for Golang libraries & applications *(:star:296 | 0clm)*
  * [ungerik/**go-dry**](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go *(:star:139 | 0clm)*
  * [beefsack/**go-rate**](https://github.com/beefsack/go-rate) - A timed rate limiter for Go *(:star:184 | 0clm)*
  * [ikeikeikeike/**go-sitemap-generator**](https://github.com/ikeikeikeike/go-sitemap-generator) - go-sitemap-generator is the easiest way to generate Sitemaps in Go. *(:star:3 | 0clm)*
  * [sadlil/**go-trigger**](https://github.com/sadlil/go-trigger) - A Global event triggerer for golang. Defines functions as event with id string. Trigger the event anywhere from your project. *(:star:13 | 0clm)*
  * [tobyhede/**go-underscore**](https://github.com/tobyhede/go-underscore) -  Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness.  *(:star:824 | 0clm)*
  * [carlescere/**goback**](https://github.com/carlescere/goback) - Golang simple exponential backoff package. *(:star:21 | 0clm)*
  * [VividCortex/**godaemon**](https://github.com/VividCortex/godaemon) - Daemonize Go applications deviously. *(:star:232 | 0clm)*
  * [joho/**godotenv**](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) *(:star:247 | 0clm)*
  * [dropbox/**godropbox**](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications. *(:star:2953 | 18clm)*
  * [cosiner/**gohper**](https://github.com/cosiner/gohper) - common libs here. *(:star:201 | 22clm)*
  * [elgs/**gojq**](https://github.com/elgs/gojq) - JSON query in Golang *(:star:18 | 0clm)*
  * [msempere/**golarm**](https://github.com/msempere/golarm) - Fire alarms with system events *(:star:13 | 0clm)*
  * [mlimaloureiro/**golog**](https://github.com/mlimaloureiro/golog) - Easy and simple CLI time tracker for your tasks *(:star:11 | 0clm)*
  * [bndr/**gopencils**](https://github.com/bndr/gopencils) - Easily consume REST APIs with Go (golang) *(:star:329 | 0clm)*
  * [michiwend/**goplaceholder**](https://github.com/michiwend/goplaceholder) - a small golang lib to generate placeholder images *(:star:10 | 0clm)*
  * [franela/**goreq**](https://github.com/franela/goreq) - Minimal and simple request library for Go language *(:star:417 | 2clm)*
  * [smallnest/**goreq**](https://github.com/smallnest/goreq) - A Simplified Golang Http Client *(:star:11 | 0clm)*
  * [parnurzeal/**gorequest**](https://github.com/parnurzeal/gorequest) - GoRequest -- Simplified HTTP client ( inspired by nodejs SuperAgent ) *(:star:575 | 11clm)*
  * [subosito/**gotenv**](https://github.com/subosito/gotenv) - Load environment variables dynamically in Go. *(:star:48 | 0clm)*
  * [levigross/**grequests**](https://github.com/levigross/grequests) - A Go "clone" of the great and famous Requests library *(:star:687 | 0clm)*
  * [htcat/**htcat**](https://github.com/htcat/htcat) - Parallel and Pipelined HTTP GET Utility *(:star:415 | 0clm)*
  * [facebookgo/**httpcontrol**](https://github.com/facebookgo/httpcontrol) - Package httpcontrol allows for HTTP transport level control around timeouts and retries. *(:star:380 | 0clm)*
  * [afex/**hystrix-go**](https://github.com/afex/hystrix-go) - Netflix's Hystrix latency and fault tolerance library, for Go  *(:star:350 | 3clm)*
  * [bamzi/**jobrunner**](https://github.com/bamzi/jobrunner) - Framework for performing work asynchronously, outside of the request flow *(:star:255 | 0clm)*
  * [miolini/**jsonf**](https://github.com/miolini/jsonf) - Console JSON formatter with query feature *(:star:20 | 0clm)*
  * [ricardolonga/**jsongo**](https://github.com/ricardolonga/jsongo) - Fluent API to make it easier to create Json objects. *(:star:44 | 0clm)*
  * [jaschaephraim/**lrserver**](https://github.com/jaschaephraim/lrserver) - LiveReload server for Go [golang] *(:star:58 | 0clm)*
  * [minio/**mc**](https://github.com/minio/mc) - Minio Client is a replacement for ls, cp, mkdir, diff and rsync commands for filesystems and object storage. *(:star:170 | 66clm)*
  * [sanbornm/**mp**](https://github.com/sanbornm/mp) - Simple Email Parser *(:star:9 | 0clm)*
  * [VividCortex/**multitick**](https://github.com/VividCortex/multitick) - A multiplexor for aligned time.Time tickers in Go *(:star:43 | 0clm)*
  * [e-dard/**netbug**](https://github.com/e-dard/netbug) - Package netbug provides a handler for registering profilers on your own ServeMux. *(:star:25 | 0clm)*
  * [inconshreveable/**ngrok**](https://github.com/inconshreveable/ngrok) - Introspected tunnels to localhost *(:star:6445 | 0clm)*
  * [xta/**okrun**](https://github.com/xta/okrun) - ok, run your gofile *(:star:10 | 0clm)*
  * [maruel/**panicparse**](https://github.com/maruel/panicparse) - Crash your app in style (Golang) *(:star:1209 | 2clm)*
  * [peco/**peco**](https://github.com/peco/peco) - Simplistic interactive filtering tool *(:star:2811 | 9clm)*
  * [sethgrid/**pester**](https://github.com/sethgrid/pester) - Go (golang) http calls with retries and backoff  *(:star:50 | 0clm)*
  * [VividCortex/**pm**](https://github.com/VividCortex/pm) - Processlist manager with TCP listener *(:star:35 | 0clm)*
  * [davecheney/**profile**](https://github.com/davecheney/profile) - A simple profiling support package for Go *(:star:414 | 0clm)*
  * [mozillazg/**request**](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™. *(:star:106 | 7clm)*
  * [ivpusic/**rerun**](https://github.com/ivpusic/rerun) - Configurable recompiling and rerunning go apps when source changes *(:star:96 | 0clm)*
  * [go-resty/**resty**](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client *(:star:258 | 0clm)*
  * [VividCortex/**robustly**](https://github.com/VividCortex/robustly) - Run functions resiliently in Go, catching and restarting panics *(:star:101 | 0clm)*
  * [carlescere/**scheduler**](https://github.com/carlescere/scheduler) - Job scheduling made easy. *(:star:75 | 20clm)*
  * [dghubble/**sling**](https://github.com/dghubble/sling) - A Go HTTP client library for creating and sending API requests *(:star:258 | 0clm)*
  * [briandowns/**spinner**](https://github.com/briandowns/spinner) - Go (golang) package for providing a terminal spinner/progress indicator with options. *(:star:212 | 5clm)*
  * [jmoiron/**sqlx**](https://github.com/jmoiron/sqlx) - general purpose extensions to golang's database/sql *(:star:1515 | 1clm)*
  * [tealeg/**xlsx**](https://github.com/tealeg/xlsx) - Google Go (golang) library for reading and writing XLSX files. *(:star:819 | 13clm)*

## Validation
*Libraries for validation.*

  * [asaskevich/**govalidator**](https://github.com/asaskevich/govalidator) - [Go] Package of validators and sanitizers for strings, numerics, slices and structs *(:star:834 | 5clm)*
  * [go-playground/**validator**](https://github.com/go-playground/validator) - :100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving *(:star:384 | 63clm)*

## Version Control
*Libraries for version control.*

  * [rjeczalik/**gh**](https://github.com/rjeczalik/gh) - Scriptable server and net/http middleware for GitHub Webhooks. *(:star:31 | 0clm)*
  * [libgit2/**git2go**](https://github.com/libgit2/git2go) - Git to Go. Like McDonald's but tastier. *(:star:698 | 10clm)*
  * [sourcegraph/**go-vcs**](https://github.com/sourcegraph/go-vcs) - manipulate and inspect VCS repositories in Go *(:star:38 | 6clm)*
  * [beyang/**hgo**](https://github.com/beyang/hgo) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories. *(:star:7 | 0clm)*

## Video
*Libraries for manipulating video.*

  * [nareix/**codec**](https://github.com/nareix/codec) - Golang libav codec bindings (h264,aac) *(:star:139 | 0clm)*
  * [3d0c/**gmf**](https://github.com/3d0c/gmf) - Go Media Framework *(:star:171 | 0clm)*
  * [giorgisio/**goav**](https://github.com/giorgisio/goav) - Golang bindings for FFmpeg *(:star:90 | 0clm)*
  * [ziutek/**gst**](https://github.com/ziutek/gst) - Go bindings for GStreamer *(:star:109 | 0clm)*

## Web Frameworks
*Full stack web frameworks.*

  * [astaxie/**beego**](https://github.com/astaxie/beego) - beego is an open-source, high-performance web framework for the Go programming language. *(:star:6184 | 29clm)*
  * [go-zoo/**bone**](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer *(:star:854 | 0clm)*
  * [pressly/**chi**](https://github.com/pressly/chi) - small, fast and expressive router / mux for Go HTTP services built with net/context *(:star:247 | 0clm)*
  * [labstack/**echo**](https://github.com/labstack/echo) - Echo is  a fast :rocket: and unfancy micro web framework for Go. *(:star:2941 | 0clm)*
  * [gin-gonic/**gin**](https://github.com/gin-gonic/gin) - Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. *(:star:5366 | 33clm)*
  * [NYTimes/**gizmo**](https://github.com/NYTimes/gizmo) - A Microservice Toolkit from The New York Times *(:star:1068 | 0clm)*
  * [desertbit/**glue**](https://github.com/desertbit/glue) - Glue - Robust Go and Javascript Socket Library (Alternative to Socket.io) *(:star:124 | 0clm)*
  * [ant0ine/**go-json-rest**](https://github.com/ant0ine/go-json-rest) - A quick and easy way to setup a RESTful JSON API *(:star:2111 | 25clm)*
  * [go-kit/**kit**](https://github.com/go-kit/kit) - A standard library for microservices. *(:star:3019 | 34clm)*
  * [codehack/**go-relax**](https://github.com/codehack/go-relax) - Framework for building RESTful API's in Go *(:star:115 | 0clm)*
  * [ungerik/**go-rest**](https://github.com/ungerik/go-rest) - A small and evil REST framework for Go *(:star:84 | 0clm)*
  * [googollee/**go-socket.io**](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework. *(:star:994 | 2clm)*
  * [raphael/**goa**](https://github.com/raphael/goa) - Design-based HTTP microservices in Go *(:star:552 | 0clm)*
  * [bahlo/**goat**](https://github.com/bahlo/goat) - :goat: A minimalistic REST API server in Go *(:star:76 | 0clm)*
  * [gocraft/**web**](https://github.com/gocraft/web) - Go Router + Middleware. Your Contexts. *(:star:960 | 0clm)*
  * [goji/**goji**](https://github.com/goji/goji) - Goji is a minimalistic and flexible HTTP request multiplexer for Go (golang) *(:star:71 | 0clm)*
  * [jcuga/**golongpoll**](https://github.com/jcuga/golongpoll) - golang HTTP longpolling library.  Makes web pub-sub easy :smiley: :coffee: :computer: *(:star:244 | 0clm)*
  * [rainycape/**gondola**](https://github.com/rainycape/gondola) - The web framework for writing faster sites, faster *(:star:292 | 0clm)*
  * [ian-kent/**goose**](https://github.com/ian-kent/goose) - Server-Sent Events in Go *(:star:13 | 0clm)*
  * [julienschmidt/**httprouter**](https://github.com/julienschmidt/httprouter) - A high performance HTTP request router that scales well *(:star:2700 | 16clm)*
  * [dimfeld/**httptreemux**](https://github.com/dimfeld/httptreemux) - High-speed, flexible tree-based HTTP router for Go. *(:star:119 | 0clm)*
  * [Unknwon/**macaron**](https://github.com/Unknwon/macaron) - :exclamation::exclamation::exclamation: [deprecated] Moved to https://github.com/go-macaron/macaron *(:star:7 | 5clm)*
  * [paulbellamy/**mango**](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. *(:star:296 | 0clm)*
  * [imdario/**medeina**](https://github.com/imdario/medeina) - Go HTTP routing tree based on HttpRouter. Inspired by Roda and Cuba. *(:star:15 | 0clm)*
  * [gorilla/**mux**](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for golang. *(:star:1951 | 1clm)*
  * [ivpusic/**neo**](https://github.com/ivpusic/neo) - Go Web Framework *(:star:202 | 19clm)*
  * [bmizerany/**pat**](https://github.com/bmizerany/pat) -  *(:star:900 | 0clm)*
  * [resoursea/**api**](https://github.com/resoursea/api) - A REST framework for quickly writing resource based services in Golang. *(:star:24 | 8clm)*
  * [revel/**revel**](https://github.com/revel/revel) - A high productivity, full-stack web framework for the Go language. *(:star:6315 | 12clm)*
  * [goanywhere/**rex**](https://github.com/goanywhere/rex) - Pleasures for Web in Golang *(:star:10 | 0clm)*
  * [VividCortex/**siesta**](https://github.com/VividCortex/siesta) - Composable framework for writing HTTP handlers in Go. *(:star:325 | 13clm)*
  * [lunny/**tango**](https://github.com/lunny/tango) - Micro & pluggable web framework for Go *(:star:345 | 6clm)*
  * [rcrowley/**go-tigertonic**](https://github.com/rcrowley/go-tigertonic) - A Go framework for building JSON web services inspired by Dropwizard *(:star:914 | 1clm)*
  * [pilu/**traffic**](https://github.com/pilu/traffic) - Sinatra inspired regexp/pattern mux and web framework for Go *(:star:458 | 0clm)*
  * [volatile/**core**](https://github.com/volatile/core) - Pure handlers stack *(:star:78 | 0clm)*
  * [hoisie/**web**](https://github.com/hoisie/web) - The easiest way to create web applications with Go *(:star:2582 | 0clm)*
  * [rs/**xmux**](https://github.com/rs/xmux) - xmux is a httprouter fork on top of xhandler (net/context aware) *(:star:40 | 0clm)*
  * [cosiner/**zerver**](https://github.com/cosiner/zerver) - a RESTful API framework *(:star:119 | 14clm)*
  * [daryl/**zeus**](https://github.com/daryl/zeus) - Go HTTP router. *(:star:91 | 0clm)*

### Middlewares
#### Actual middlewares

  * [rs/**cors**](https://github.com/rs/cors) - Go net/http configurable handler to handle CORS requests *(:star:282 | 7clm)*
  * [rs/**formjson**](https://github.com/rs/formjson) - Go net/http handler to transparently manage posted JSON *(:star:17 | 0clm)*
  * [ulule/**limiter**](https://github.com/ulule/limiter) - Dead simple rate limit middleware for Go. *(:star:207 | 0clm)*
  * [didip/**tollbooth**](https://github.com/didip/tollbooth) - Simple middleware to rate-limit HTTP requests. *(:star:351 | 0clm)*
  * [sebest/**xff**](https://github.com/sebest/xff) - A Golang Middleware to handle X-Forwarded-For Header *(:star:41 | 10clm)*

#### Libraries for creating HTTP middlewares

  * [justinas/**alice**](https://github.com/justinas/alice) - Painless middleware chaining for Go *(:star:800 | 0clm)*
  * [codemodus/**catena**](https://github.com/codemodus/catena) - http.Handler wrapper catenation. *(:star:2 | 0clm)*
  * [codemodus/**chain**](https://github.com/codemodus/chain) - Handler wrapper chaining with scoped data. *(:star:37 | 0clm)*
  * [go-on/**wrap**](https://github.com/go-on/wrap) - Go http.Hander based middleware stack with context sharing *(:star:49 | 0clm)*
  * [carbocation/**interpose**](https://github.com/carbocation/interpose) - Minimalist net/http middleware for golang *(:star:219 | 4clm)*
  * [stephens2424/**muxchain**](https://github.com/stephens2424/muxchain) - Lightweight Middleware for net/http *(:star:192 | 0clm)*
  * [codegangsta/**negroni**](https://github.com/codegangsta/negroni) - Idiomatic HTTP Middleware for Golang *(:star:3328 | 0clm)*
  * [unrolled/**render**](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, binary data, and HTML templates responses. *(:star:612 | 1clm)*
  * [thoas/**stats**](https://github.com/thoas/stats) - A Go middleware that stores various information about your web application (response time, status code count, etc.) *(:star:305 | 0clm)*

# Tools
## Code Analysis

  * [mibk/**dupl**](https://github.com/mibk/dupl) - a tool for code clone detection *(:star:40 | 1clm)*
  * [kisielk/**errcheck**](https://github.com/kisielk/errcheck) - errcheck checks that you checked errors. *(:star:501 | 4clm)*
  * [davecheney/**gcvis**](https://github.com/davecheney/gcvis) - Visualise Go program GC trace data in real time *(:star:472 | 0clm)*
  * [alecthomas/**gometalinter**](https://github.com/alecthomas/gometalinter) - Concurrently run Go lint tools and normalise their output *(:star:684 | 16clm)*
  * [qiniu/**checkstyle**](https://github.com/qiniu/checkstyle) - checkstyle for go *(:star:22 | 0clm)*
  * [yuroyoro/**goast-viewer**](https://github.com/yuroyoro/goast-viewer) - Golang AST visualizer *(:star:130 | 0clm)*
  * [golang/**lint**](https://github.com/golang/lint) - This is a linter for Go source code. *(:star:1209 | 2clm)*
  * [shurcooL/**gostatus**](https://github.com/shurcooL/gostatus) - A command line tool that shows the status of Go repositories. *(:star:119 | 1clm)*
  * [mvdan/**interfacer**](https://github.com/mvdan/interfacer) - A linter that suggests interface types *(:star:203 | 0clm)*
  * [mccoyst/**validate**](https://github.com/mccoyst/validate) - A Go package to automatically validate fields with tags *(:star:55 | 0clm)*

## Editor Plugins

  * [go-lang-plugin-org/**go-lang-idea-plugin**](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) - Google Go language IDE built using the IntelliJ Platform *(:star:2904 | 256clm)*
  * [joefitzgerald/**go-plus**](https://github.com/joefitzgerald/go-plus) - An Enhanced Go Experience For The Atom Editor *(:star:722 | 28clm)*
  * [GoClipse/**goclipse**](https://github.com/GoClipse/goclipse) - Eclipse IDE for the Go programming language *(:star:473 | 76clm)*
  * [nsf/**gocode**](https://github.com/nsf/gocode) - An autocompletion daemon for the Go programming language *(:star:2912 | 2clm)*
  * [DisposaBoy/**GoSublime**](https://github.com/DisposaBoy/GoSublime) - A  Golang plugin collection for the text editor SublimeText 2 providing code completion and other IDE-like features. *(:star:2086 | 0clm)*
  * [velour/**velour**](https://github.com/velour/velour) - An IRC client for acme — the project that started it all. *(:star:12 | 0clm)*
  * [rjohnsondev/**vim-compiler-go**](https://github.com/rjohnsondev/vim-compiler-go) - Vim compiler plugin for Go (golang) *(:star:60 | 0clm)*
  * [fatih/**vim-go**](https://github.com/fatih/vim-go) - Go development plugin for Vim *(:star:4200 | 38clm)*
  * [eaburns/**Watch**](https://github.com/eaburns/Watch) - Watches for changes in a directory tree and reruns a command in an acme win or just on the terminal. *(:star:102 | 2clm)*

## Go Tools

  * [songgao/**colorgo**](https://github.com/songgao/colorgo) - Colorize (highlight) `go build` command output *(:star:57 | 0clm)*
  * [skelterjohn/**go-pkg-complete**](https://github.com/skelterjohn/go-pkg-complete) - bash completion for go and wgo *(:star:25 | 0clm)*

## Software Packages
### DevOps Tools

  * [smira/**aptly**](https://github.com/smira/aptly) - aptly - Debian repository management tool *(:star:971 | 74clm)*
  * [soniah/**awsenv**](https://github.com/soniah/awsenv) - AWS environment config loader *(:star:5 | 0clm)*
  * [rakyll/**boom**](https://github.com/rakyll/boom) - HTTP(S) load generator, ApacheBench (ab) replacement, written in Go *(:star:3871 | 0clm)*
  * [bosun-monitor/**bosun**](https://github.com/bosun-monitor/bosun) - Time Series Alerting Framework *(:star:1574 | 76clm)*
  * [liudng/**dogo**](https://github.com/liudng/dogo) - Monitoring changes in the source file and automatically compile and run (restart). *(:star:83 | 6clm)*
  * [chrismckenzie/**dropship**](https://github.com/chrismckenzie/dropship) - instead of jumping ship... Dropship *(:star:20 | 0clm)*
  * [hypersleep/**easyssh**](https://github.com/hypersleep/easyssh) - Golang package for easy remote execution through SSH *(:star:73 | 11clm)*
  * [rcrowley/**go-metrics**](https://github.com/rcrowley/go-metrics) - Go port of Coda Hale's Metrics library *(:star:1069 | 1clm)*
  * [sanbornm/**go-selfupdate**](https://github.com/sanbornm/go-selfupdate) - Enable your Golang applications to self update *(:star:323 | 3clm)*
  * [cryptojuice/**gobrew**](https://github.com/cryptojuice/gobrew) - Shell script to download and set GO environmental paths to allow multiple versions. *(:star:137 | 6clm)*
  * [sirnewton01/**godbg**](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application *(:star:180 | 0clm)*
  * [inconshreveable/**gonative**](https://github.com/inconshreveable/gonative) - Build Go Toolchains /w native libs for cross-compilation *(:star:219 | 1clm)*
  * [mitchellh/**gox**](https://github.com/mitchellh/gox) - A dead simple, no frills Go cross compile tool *(:star:1430 | 0clm)*
  * [laher/**goxc**](https://github.com/laher/goxc) - a build tool for Go, with a focus on cross-compiling, packaging and deployment *(:star:1208 | 14clm)*
  * [moovweb/**gvm**](https://github.com/moovweb/gvm) - Go Version Manager *(:star:1736 | 4clm)*
  * [heroku/**hk**](https://github.com/heroku/hk) - Fast Heroku command-line interface *(:star:766 | 2clm)*
  * [ajvb/**kala**](https://github.com/ajvb/kala) - Modern Job Scheduler *(:star:757 | 0clm)*
  * [kubernetes/**kubernetes**](https://github.com/kubernetes/kubernetes) - Container Cluster Manager from Google *(:star:12425 | 1217clm)*
  * [emicklei/**mora**](https://github.com/emicklei/mora) - MongoDB generic REST server in Go *(:star:133 | 0clm)*
  * [ostrost/**ostent**](https://github.com/ostrost/ostent) - ostent collects system metrics to display and relay *(:star:58 | 0clm)*
  * [mitchellh/**packer**](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. *(:star:4952 | 62clm)*
  * [alouche/**rodent**](https://github.com/alouche/rodent) - Manage Go Versions/Projects/Dependencies *(:star:31 | 0clm)*
  * [rlmcpherson/**s3gof3r**](https://github.com/rlmcpherson/s3gof3r) - Fast, concurrent, streaming access to Amazon S3, including gof3r, a CLI. http://godoc.org/github.com/rlmcpherson/s3gof3r *(:star:608 | 5clm)*
  * [tsenart/**vegeta**](https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000! *(:star:2909 | 5clm)*
  * [adnanh/**webhook**](https://github.com/adnanh/webhook) - webhook is a lightweight configurable tool written in Go, that allows you to easily create HTTP endpoints (hooks) on your server, which you can use to execute configured commands. *(:star:311 | 11clm)*

### Other Software

  * [tejo/**boxed**](https://github.com/tejo/boxed) - dropbox based blog engine, written in go. *(:star:41 | 36clm)*
  * [gocircuit/**circuit**](https://github.com/gocircuit/circuit) - Circuit: Dynamic cloud orchestration http://gocircuit.org *(:star:1209 | 52clm)*
  * [tylertreat/**Comcast**](https://github.com/tylertreat/Comcast) - Simulating shitty network connections so you can build better systems. *(:star:4099 | 9clm)*
  * [kelseyhightower/**confd**](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul *(:star:2314 | 10clm)*
  * [coreos/**fleet**](https://github.com/coreos/fleet) - A Distributed init System *(:star:2056 | 7clm)*
  * [goccmack/**gocc**](https://github.com/goccmack/gocc) - Parser / Scanner Generator *(:star:15 | 0clm)*
  * [diankong/**GoDocTooltip**](https://github.com/diankong/GoDocTooltip) - A Chrome extension for golang users.When you're at golang's official doc site, it will show function's description as tooltip on function list *(:star:7 | 0clm)*
  * [buger/**gor**](https://github.com/buger/gor) - Gor is an open-source tool for capturing and replaying live HTTP traffic into a test environment in order to continuously test your system with real data. It can be used to increase confidence in code deployments, configuration changes and infrastructure changes. *(:star:4588 | 4clm)*
  * [mozilla-services/**heka**](https://github.com/mozilla-services/heka) - Data collection and processing made easy. *(:star:2831 | 161clm)*
  * [dimiro1/**ipe**](https://github.com/dimiro1/ipe) - An open source Pusher server implementation compatible with Pusher client libraries written in GO *(:star:129 | 11clm)*
  * [visualfc/**liteide**](https://github.com/visualfc/liteide) - LiteIDE is a simple, open source, cross-platform Go IDE.  *(:star:2535 | 96clm)*
  * [unix4fun/**naclpipe**](https://github.com/unix4fun/naclpipe) - NaCL pipe *(:star:4 | 0clm)*
  * [fogleman/**nes**](https://github.com/fogleman/nes) - NES emulator written in Go. *(:star:2136 | 0clm)*
  * [noraesae/**orange-cat**](https://github.com/noraesae/orange-cat) - A Markdown previewer written in Go *(:star:134 | 0clm)*
  * [pointlander/**peg**](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. *(:star:267 | 1clm)*
  * [zachlatta/**postman**](https://github.com/zachlatta/postman) - Command-line utility for batch-sending email. *(:star:622 | 0clm)*
  * [restic/**restic**](https://github.com/restic/restic) - restic backup program *(:star:441 | 90clm)*
  * [coreos/**rkt**](https://github.com/coreos/rkt) - rkt is an App Container runtime for Linux *(:star:5048 | 146clm)*
  * [chrislusf/**seaweedfs**](https://github.com/chrislusf/seaweedfs) - SeaweedFS is a simple and highly scalable distributed file system. There are two objectives:  to store billions of files! to serve the files fast! Instead of supporting full POSIX file system semantics, Seaweed-FS choose to implement only a key~file mapping. Similar to the word "NoSQL", you can call it as "NoFS". *(:star:2111 | 39clm)*
  * [msoap/**shell2http**](https://github.com/msoap/shell2http) - Executing shell commands via http server (GoLang) *(:star:35 | 0clm)*
  * [intelsdi-x/**snap**](https://github.com/intelsdi-x/snap) - A powerful telemetry framework *(:star:518 | 131clm)*
  * [pressly/**sup**](https://github.com/pressly/sup) - Super simple deployment tool - just Unix - think of it like 'make' for a network of servers *(:star:804 | 8clm)*
  * [kyleterry/**tenyks**](https://github.com/kyleterry/tenyks) - The Tenyks IRC bot. *(:star:152 | 0clm)*
  * [shopify/**toxiproxy**](https://github.com/shopify/toxiproxy) - :alarm_clock: :fire: A proxy to simulate network and system conditions *(:star:1085 | 16clm)*
  * [ian-kent/**websysd**](https://github.com/ian-kent/websysd) - Like Marathon or Upstart - for your desktop! *(:star:17 | 0clm)*
  * [wellington/**wellington**](https://github.com/wellington/wellington) - Spriting that sass has been missing *(:star:160 | 43clm)*

# Resources
## Benchmarks

  * [davecheney/**autobench**](https://github.com/davecheney/autobench) - Go benchmark harness.  *(:star:83 | 0clm)*
  * [julienschmidt/**go-http-routing-benchmark**](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router and web framework benchmark *(:star:621 | 5clm)*
  * [hgfischer/**go-type-assertion-benchmark**](https://github.com/hgfischer/go-type-assertion-benchmark) - Naive performance test of two ways to do type assertion in Go. *(:star:2 | 0clm)*
  * [alecthomas/**go_serialization_benchmarks**](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods *(:star:260 | 2clm)*
  * [PuerkitoBio/**gocostmodel**](https://github.com/PuerkitoBio/gocostmodel) - Benchmarks of common basic operations for the Go language. *(:star:46 | 0clm)*
  * [amscanne/**golang-micro-benchmarks**](https://github.com/amscanne/golang-micro-benchmarks) - Tiny collection of micro benchmarks. *(:star:5 | 0clm)*
  * [tyler-smith/**golang-sql-benchmark**](https://github.com/tyler-smith/golang-sql-benchmark) - A benchmarking shootout of various db/SQL utilities for Go *(:star:18 | 0clm)*
  * [feyeleanor/**GoSpeed**](https://github.com/feyeleanor/GoSpeed) - Go micro-benchmarks for calculating the speed of language constructs *(:star:44 | 0clm)*
  * [jimrobinson/**kvbench**](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark *(:star:9 | 0clm)*
  * [fawick/**speedtest-resize**](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language *(:star:60 | 0clm)*

## E-Books

  * [dariubs/**GoBooks**](https://github.com/dariubs/GoBooks) - List of Golang books *(:star:1851 | 0clm)*

## Websites

  * [lukasz-madon/**awesome-remote-job**](https://github.com/lukasz-madon/awesome-remote-job) - A curated list of awesome remote jobs and resources. Inspired by https://github.com/vinta/awesome-python *(:star:4853 | 43clm)*
  * [bayandin/**awesome-awesomeness**](https://github.com/bayandin/awesome-awesomeness) - A curated list of awesome awesomeness *(:star:15222 | 0clm)*
  * [mholt/**golang-graphics**](https://github.com/mholt/golang-graphics) - Community-contributed Go graphics files *(:star:70 | 1clm)*

### Tutorials

  * [mkaz/**working-with-go**](https://github.com/mkaz/working-with-go) - A set of example golang code to start learning Go *(:star:471 | 0clm)*

## Windows

  * [go-ole/**go-ole**](https://github.com/go-ole/go-ole) - win32 ole implementation for golang *(:star:175 | 22clm)*



package conf

import (
	"gopkg.in/ini.v1"
	"strings"
)

var (
	Mode string
	Port string

	MysqlName     string
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDatabase string
	MysqlCharset  string

	RedisName     string
	RedisAddr     string
	RedisPassword string
	RedisDbName   string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host    string
	Product string
	Avatar  string

	ReadDB  string
	WriteDB string
)

func Init() {
	file, err := ini.Load("../conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMysql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPhotoPath(file)

	// 读写分离
	// mysql 读
	ReadDB = strings.Join([]string{MysqlUser, ":", MysqlPassword, "@tcp(", MysqlHost, ":", MysqlPort, ")/", MysqlDatabase, "?charset=", MysqlCharset, "&parseTime=true"}, "")
	// mysql 写
	WriteDB = strings.Join([]string{MysqlUser, ":", MysqlPassword, "@tcp(", MysqlHost, ":", MysqlPort, ")/", MysqlDatabase, "?charset=", MysqlCharset, "&parseTime=true"}, "")
}

// LoadServer 加载服务器配置
func LoadServer(file *ini.File) {
	Mode = file.Section("server").Key("mode").String()
	Port = file.Section("server").Key("port").String()
}

// LoadMysql 加载mysql配置
func LoadMysql(file *ini.File) {
	MysqlName = file.Section("mysql").Key("name").String()
	MysqlHost = file.Section("mysql").Key("host").String()
	MysqlPort = file.Section("mysql").Key("port").String()
	MysqlUser = file.Section("mysql").Key("user").String()
	MysqlPassword = file.Section("mysql").Key("password").String()
	MysqlDatabase = file.Section("mysql").Key("database").String()
	MysqlCharset = file.Section("mysql").Key("charset").String()
}

// LoadRedis 加载redis配置
func LoadRedis(file *ini.File) {
	RedisName = file.Section("redis").Key("name").String()
	RedisAddr = file.Section("redis").Key("addr").String()
	RedisPassword = file.Section("redis").Key("password").String()
	RedisDbName = file.Section("redis").Key("dbName").String()
}

// LoadEmail 加载邮件配置
func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("validEmail").String()
	SmtpHost = file.Section("email").Key("smtpHost").String()
	SmtpEmail = file.Section("email").Key("smtpEmail").String()
	SmtpPass = file.Section("email").Key("smtpPass").String()
}

// LoadPhotoPath 加载路径配置
func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("host").String()
	Product = file.Section("path").Key("product").String()
	Avatar = file.Section("path").Key("avatar").String()
}

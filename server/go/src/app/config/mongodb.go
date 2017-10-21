package config

import (
	"io"
	"io/ioutil"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/yaml.v1"
)

// Configs 環境ごとの設定情報をもつ
type Configs map[string]*Config

// Open 指定された環境についてDBに接続します。
func (cs Configs) MongodbOpen(env string) (*mgo.Session, error) {
	config, ok := cs[env]
	if !ok {
		return nil, nil
	}
	return config.MongodbOpen()
}

// Config sql-migrateの設定ファイルと同じ形式を想定している
type Config struct {
	Datasource string `yaml:"datasource"`
}

// DSN 設定されているDSNを返します
func (c *Config) DSN() string {
	return c.Datasource
}

// Open Configで指定されている接続先に接続する。
// MySQL固定
func (c *Config) MongodbOpen() (*mgo.Session, error) {
	return mgo.Dial(c.DSN())
}

// NewConfigsFromFile Configから設定を読み取る
func NewMongodbConfigsFromFile(path string) (Configs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close() // nolint: errcheck
	return NewMongodbConfigs(f)
}

// NewConfigs io.ReaderからDB用設定を読み取る
func NewMongodbConfigs(r io.Reader) (Configs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var configs Configs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

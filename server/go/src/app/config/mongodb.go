package config

import (
	"errors"
	"io"
	"io/ioutil"
	"os"

	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/yaml.v1"
)

// S3Configs は用途ごとのS3情報を持つ
type MongodbConfigs map[string]*Mongodb

// S3 S3の情報
type Mongodb struct {
	Hosts    string `yaml:"hosts"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Get は指定された用途のOauth認証情報を返す
func (sc MongodbConfigs) MongoOpen(purpose string) (*mgo.Session, error) {
	c, ok := sc[purpose]
	if !ok {
		return nil, errors.New("No such config")
	}
	return mgo.Dial(fmt.Sprintf("mongodb://%s", c.Hosts))
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	//fmt.Println(fmt.Sprintf("mongodb://%s:%s@%s/%s", c.Username, c.Password, c.Hosts, c.Database))
}

// NewS3ConfigsFromFile Configから設定を読み取る
func NewMongodbConfigsFromFile(path string) (MongodbConfigs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close() // nolint: errcheck
	return NewMongodbConfigs(f)
}

// NewS3Configs io.ReaderからS3設定を読み取る
func NewMongodbConfigs(r io.Reader) (MongodbConfigs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var configs MongodbConfigs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

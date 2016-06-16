// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	var testConfigData = []byte(
		`
addr : 0.0.0.0:9696
log_level : error
users:
-
  user: user_a
  password: pass_a

nodes :
- 
  name : node1
  down_after_noalive : 300
  max_conns_limit : 16
  master : 127.0.0.1:3306
  slave : 127.0.0.1:4306
  users:
  -
   user: user_b
   password: pass_b
- 
  name : node2
  master : 127.0.0.1:3307
  users:
  -
   user: user_c

- 
  name : node3
  down_after_noalive : 300
  max_conns_limit : 16
  master : 127.0.0.1:3308
  users:
  -
    user: user_d
    password: pass_d
  -
    user: user_e
    password: pass_e

schema :
  db : kingshard 
  nodes: [node1, node2, node3]
  default: node1
  shard:
    -   
      table: test_shard_hash
      key: id
      nodes: [node1, node2, node3]
      type: hash
      locations: [4,4,4]
    -   
      table: test_shard_range
      key: id
      type: range
      nodes: [node2, node3]
      locations: [4,4]
      table_row_limit: 10000
`)

	cfg, err := ParseConfigData(testConfigData)
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Nodes) != 3 {
		t.Fatal(len(cfg.Nodes))
	}

	testNode := NodeConfig{
		Name:             "node1",
		DownAfterNoAlive: 300,
		MaxConnNum:       16,
		Users: []AuthInfo{
			AuthInfo{
				User:     "user_b",
				Password: "pass_b",
			},
		},

		Master: "127.0.0.1:3306",
		Slave:  "127.0.0.1:4306",
	}

	if !reflect.DeepEqual(cfg.Nodes[0], testNode) {
		fmt.Printf("%v\n", cfg.Nodes[0])
		t.Fatal("node1 must equal")
	}

	testNode_2 := NodeConfig{
		Name:   "node2",
		Master: "127.0.0.1:3307",
		Users: []AuthInfo{
			AuthInfo{
				User: "user_c",
			},
		},
	}

	if !reflect.DeepEqual(cfg.Nodes[1], testNode_2) {
		t.Fatal("node2 must equal")
	}

	testShard_1 := ShardConfig{
		Table:         "test_shard_hash",
		Key:           "id",
		Nodes:         []string{"node1", "node2", "node3"},
		Locations:     []int{4, 4, 4},
		Type:          "hash",
		TableRowLimit: 0,
	}
	if !reflect.DeepEqual(cfg.Schema.ShardRule[0], testShard_1) {
		fmt.Printf("%v\n", cfg.Schema.ShardRule[0])
		t.Fatal("ShardConfig0 must equal")
	}

	testShard_2 := ShardConfig{
		Table:         "test_shard_range",
		Key:           "id",
		Nodes:         []string{"node2", "node3"},
		Type:          "range",
		Locations:     []int{4, 4},
		TableRowLimit: 10000,
	}
	if !reflect.DeepEqual(cfg.Schema.ShardRule[1], testShard_2) {
		fmt.Printf("%v\n", cfg.Schema.ShardRule[1])
		t.Fatal("ShardConfig1 must equal")
	}

	if 2 != len(cfg.Schema.ShardRule) {
		t.Fatal("ShardRule must 2")
	}

	testSchema := SchemaConfig{
		DB:        "kingshard",
		Nodes:     []string{"node1", "node2", "node3"},
		Default:   "node1",
		ShardRule: []ShardConfig{testShard_1, testShard_2},
	}

	if !reflect.DeepEqual(cfg.Schema, testSchema) {
		t.Fatal("schema must equal")
	}

	if cfg.Users[0].User != "user_a" || cfg.Users[0].Password != "pass_a" {
		t.Fatal("Top Config user / password not equal.")
	}

	if cfg.LogLevel != "error" || cfg.Addr != "0.0.0.0:9696" {
		t.Fatal("Top Config not equal.")
	}
}

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

package backend

import (
	"errors"
	"kingshard/core/golog"
	"sync"
)

type UserPool struct {
	sync.RWMutex
	pools map[string]*DB
}

func NewUserPool() *UserPool {

	pool := new(UserPool)
	pool.pools = make(map[string]*DB, 0)
	return pool
}

func (pool *UserPool) Open(addr string, user string, password string, dbName string, maxConnNum int) (*DB, error) {

	db, err := Open(addr, user, password, dbName, maxConnNum)

	if err != nil {
		golog.Error("UserPool", "Open", err.Error(), 0)
		return nil, err
	}
	pool.Lock()
	pool.pools[user] = db
	pool.Unlock()

	return db, err
}

func (pool *UserPool) GetDB(user string) (*DB, error) {

	pool.RLock()
	db, ok := pool.pools[user]
	pool.RUnlock()

	if !ok {
		return nil, errors.New("user not connect server")
	}

	return db, nil
}

//直接使用map range 的无序性来实现随机读
func (pool *UserPool) GetRandomDB() *DB {

	if len(pool.pools) < 1 {
		return nil
	}

	for _, item := range pool.pools {
		return item
	}

	return nil

}

func (pool *UserPool) GetPools() []*DB {

	pool.RLock()

	dbs := make([]*DB, 0)

	for _, item := range pool.pools {
		dbs = append(dbs, item)
	}
	pool.RUnlock()

	return dbs

}

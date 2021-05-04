// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strategystore

import (
	"errors"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/cmd/collector/app/sampling/model"
	ss "github.com/jaegertracing/jaeger/cmd/collector/app/sampling/strategystore"
	"github.com/jaegertracing/jaeger/pkg/distributedlock"
	"github.com/jaegertracing/jaeger/plugin"
	"github.com/jaegertracing/jaeger/storage/samplingstore"
)

func clearEnv() {
	os.Setenv(SamplingTypeEnvVar, "static")
}

var _ ss.Factory = new(Factory)
var _ plugin.Configurable = new(Factory)

func TestNewFactory(t *testing.T) {
	f, err := NewFactory(FactoryConfig{StrategyStoreType: "static"})
	require.NoError(t, err)
	assert.NotEmpty(t, f.factories)
	assert.NotEmpty(t, f.factories["static"])
	assert.Equal(t, Kind("static"), f.StrategyStoreType)

	mock := new(mockFactory)
	f.factories["static"] = mock

	lock := &mockLock{}
	store := &mockStore{}

	assert.NoError(t, f.Initialize(metrics.NullFactory, zap.NewNop(), lock, store))
	_, err = f.CreateStrategyStore()
	assert.NoError(t, err)

	// force the mock to return errors
	mock.retError = true
	assert.EqualError(t, f.Initialize(metrics.NullFactory, zap.NewNop(), lock, store), "error initializing store")
	_, err = f.CreateStrategyStore()
	assert.EqualError(t, err, "error creating store")

	f.StrategyStoreType = "nonsense"
	_, err = f.CreateStrategyStore()
	assert.EqualError(t, err, "no nonsense strategy store registered")

	_, err = NewFactory(FactoryConfig{StrategyStoreType: "nonsense"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown sampling strategy store type")
}

func TestConfigurable(t *testing.T) {
	clearEnv()
	defer clearEnv()

	f, err := NewFactory(FactoryConfig{StrategyStoreType: "static"})
	require.NoError(t, err)
	assert.NotEmpty(t, f.factories)
	assert.NotEmpty(t, f.factories["static"])

	mock := new(mockFactory)
	f.factories["static"] = mock

	fs := new(flag.FlagSet)
	v := viper.New()

	f.AddFlags(fs)
	f.InitFromViper(v)

	assert.Equal(t, fs, mock.flagSet)
	assert.Equal(t, v, mock.viper)
}

type mockFactory struct {
	flagSet  *flag.FlagSet
	viper    *viper.Viper
	retError bool
}

func (f *mockFactory) AddFlags(flagSet *flag.FlagSet) {
	f.flagSet = flagSet
}

func (f *mockFactory) InitFromViper(v *viper.Viper) {
	f.viper = v
}

func (f *mockFactory) CreateStrategyStore() (ss.StrategyStore, error) {
	if f.retError {
		return nil, errors.New("error creating store")
	}
	return nil, nil
}

func (f *mockFactory) Initialize(metricsFactory metrics.Factory, logger *zap.Logger, lock distributedlock.Lock, store samplingstore.Store) error {
	if f.retError {
		return errors.New("error initializing store")
	}
	return nil
}

func (f *mockFactory) RequiresLockAndSamplingStore() (bool, error) {
	if f.retError {
		return false, errors.New("error creating store")
	}
	return false, nil
}

type mockStore struct{}

func (m *mockStore) InsertThroughput(throughput []*model.Throughput) error {
	return nil
}
func (m *mockStore) InsertProbabilitiesAndQPS(hostname string, probabilities model.ServiceOperationProbabilities, qps model.ServiceOperationQPS) error {
	return nil
}
func (m *mockStore) GetThroughput(start, end time.Time) ([]*model.Throughput, error) {
	return nil, nil
}
func (m *mockStore) GetProbabilitiesAndQPS(start, end time.Time) (map[string][]model.ServiceOperationData, error) {
	return nil, nil
}
func (m *mockStore) GetLatestProbabilities() (model.ServiceOperationProbabilities, error) {
	return nil, nil
}

type mockLock struct{}

func (m *mockLock) Acquire(resource string, ttl time.Duration) (acquired bool, err error) {
	return true, nil
}

func (m *mockLock) Forfeit(resource string) (forfeited bool, err error) {
	return true, nil
}

// Copyright (c) 2020 The Jaeger Authors.
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

package defaultconfig

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/service/builder"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/cmd/opentelemetry/app"
	"github.com/jaegertracing/jaeger/cmd/opentelemetry/app/defaultcomponents"
	"github.com/jaegertracing/jaeger/cmd/opentelemetry/app/exporter/cassandraexporter"
	"github.com/jaegertracing/jaeger/cmd/opentelemetry/app/exporter/elasticsearchexporter"
	"github.com/jaegertracing/jaeger/cmd/opentelemetry/app/exporter/grpcpluginexporter"
	"github.com/jaegertracing/jaeger/cmd/opentelemetry/app/exporter/kafkaexporter"
	"github.com/jaegertracing/jaeger/cmd/opentelemetry/app/exporter/memoryexporter"
	"github.com/jaegertracing/jaeger/cmd/opentelemetry/app/receiver/kafkareceiver"
	jConfig "github.com/jaegertracing/jaeger/pkg/config"
)

func TestService(t *testing.T) {
	tests := []struct {
		service     configmodels.Service
		cfg         ComponentSettings
		err         string
		viperConfig map[string]interface{}
	}{
		{
			cfg: ComponentSettings{
				ComponentType: Agent,
			},
			service: configmodels.Service{
				Extensions: []string{"health_check"},
				Pipelines: configmodels.Pipelines{
					"traces": &configmodels.Pipeline{
						InputType:  configmodels.TracesDataType,
						Receivers:  []string{"otlp", "jaeger"},
						Processors: []string{"batch"},
						Exporters:  []string{"jaeger"},
					},
				},
			},
		},
		{
			viperConfig: map[string]interface{}{"resource.attributes": "foo=bar"},
			cfg: ComponentSettings{
				ComponentType: Collector,
				StorageType:   "elasticsearch,kafka,memory",
			},
			service: configmodels.Service{
				Extensions: []string{"health_check"},
				Pipelines: configmodels.Pipelines{
					"traces": &configmodels.Pipeline{
						InputType:  configmodels.TracesDataType,
						Receivers:  []string{"otlp", "jaeger"},
						Processors: []string{"resource", "batch"},
						Exporters:  []string{elasticsearchexporter.TypeStr, kafkaexporter.TypeStr, memoryexporter.TypeStr},
					},
				},
			},
		},
		{
			cfg: ComponentSettings{
				ComponentType: Ingester,
				StorageType:   "elasticsearch",
			},
			service: configmodels.Service{
				Extensions: []string{"health_check"},
				Pipelines: configmodels.Pipelines{
					"traces": &configmodels.Pipeline{
						InputType:  configmodels.TracesDataType,
						Receivers:  []string{kafkareceiver.TypeStr},
						Processors: []string{"batch"},
						Exporters:  []string{elasticsearchexporter.TypeStr},
					},
				},
			},
		},
		{
			cfg: ComponentSettings{
				ComponentType: Ingester,
				StorageType:   "cassandra,elasticsearch,grpc-plugin",
			},
			service: configmodels.Service{
				Extensions: []string{"health_check"},
				Pipelines: configmodels.Pipelines{
					"traces": &configmodels.Pipeline{
						InputType:  configmodels.TracesDataType,
						Receivers:  []string{kafkareceiver.TypeStr},
						Processors: []string{"batch"},
						Exporters:  []string{cassandraexporter.TypeStr, elasticsearchexporter.TypeStr, grpcpluginexporter.TypeStr},
					},
				},
			},
		},
		{
			viperConfig: map[string]interface{}{"collector.zipkin.host-port": "localhost:9411"},
			cfg: ComponentSettings{
				ComponentType: AllInOne,
				StorageType:   "elasticsearch",
			},
			service: configmodels.Service{
				Extensions: []string{"health_check"},
				Pipelines: configmodels.Pipelines{
					"traces": &configmodels.Pipeline{
						InputType:  configmodels.TracesDataType,
						Receivers:  []string{"otlp", "jaeger", "zipkin"},
						Processors: []string{"batch"},
						Exporters:  []string{elasticsearchexporter.TypeStr},
					},
				},
			},
		},
		{
			cfg: ComponentSettings{
				ComponentType: Collector,
				StorageType:   "floppy",
			},
			err: "unknown storage type: floppy",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v:%v", test.cfg.ComponentType, test.cfg.StorageType), func(t *testing.T) {
			v, _ := jConfig.Viperize(app.AddComponentFlags)
			for key, val := range test.viperConfig {
				v.Set(key, val)
			}
			factories := defaultcomponents.Components(v)
			test.cfg.Factories = factories
			cfg, err := test.cfg.createDefaultConfig()
			if test.err != "" {
				require.Nil(t, cfg)
				assert.Contains(t, err.Error(), test.err)
				return
			}
			sort.Strings(test.service.Pipelines["traces"].Exporters)
			sort.Strings(cfg.Service.Pipelines["traces"].Exporters)
			sort.Strings(test.service.Pipelines["traces"].Receivers)
			sort.Strings(cfg.Service.Pipelines["traces"].Receivers)
			require.NoError(t, err)
			require.NoError(t, config.ValidateConfig(cfg, zap.NewNop()))
			assert.Equal(t, test.service, cfg.Service)

			assert.Equal(t, len(test.service.Pipelines["traces"].Exporters), len(cfg.Exporters))
			types := []string{}
			for _, e := range cfg.Exporters {
				types = append(types, string(e.Type()))
			}
			sort.Strings(types)
			assert.Equal(t, test.service.Pipelines["traces"].Exporters, types)

			assert.Equal(t, len(test.service.Pipelines["traces"].Receivers), len(cfg.Receivers))
			types = []string{}
			for _, r := range cfg.Receivers {
				types = append(types, string(r.Type()))
			}
			sort.Strings(types)
			assert.Equal(t, test.service.Pipelines["traces"].Receivers, types)
		})
	}
}

func TestDumpOtelErrors(t *testing.T) {
	tests := []struct {
		name        string
		err         string
		viperConfig map[string]interface{}
	}{
		{
			name:        "don't dump otel",
			err:         "",
			viperConfig: nil,
		},
		{
			name:        "dump otel",
			err:         "otel config dump requested.",
			viperConfig: map[string]interface{}{"config.dump-otel": "true"},
		},
	}

	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			c := ComponentSettings{}
			jaegerViper := viper.New()
			for key, val := range test.viperConfig {
				jaegerViper.Set(key, val)
			}

			builder.Flags(flag.NewFlagSet("", flag.ContinueOnError))

			otelViper := viper.New()
			c.Factories = defaultcomponents.Components(otelViper)
			factory := c.DefaultConfigFactory(jaegerViper)

			// stdout shenanigans prevent factory() from actually dumping config during tests
			orig := os.Stdout
			os.Stdout = nil
			_, err := factory(otelViper, c.Factories)
			os.Stdout = orig

			if test.err == "" {
				assert.NoError(t, err)
			} else {
				assert.Contains(t, err.Error(), test.err)
			}
		})
	}
}

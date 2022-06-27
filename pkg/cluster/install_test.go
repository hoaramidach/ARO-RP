package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	configv1 "github.com/openshift/api/config/v1"
	operatorv1 "github.com/openshift/api/operator/v1"
	configfake "github.com/openshift/client-go/config/clientset/versioned/fake"
	operatorfake "github.com/openshift/client-go/operator/clientset/versioned/fake"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/util/steps"
	"github.com/Azure/ARO-RP/pkg/util/version"
	testdatabase "github.com/Azure/ARO-RP/test/database"
	testlog "github.com/Azure/ARO-RP/test/util/log"
)

func failingFunc(context.Context) error { return errors.New("oh no!") }

func normalFunc(context.Context) error {
	time.Sleep(1 * time.Second)
	return nil
}

type fakeMetricsEmitter struct {
	Topic      string
	IntallTime int64
}

func (e *fakeMetricsEmitter) EmitGauge(topic string, value int64, dims map[string]string) {
	e.Topic = topic
	e.IntallTime = value
}

func (e *fakeMetricsEmitter) EmitFloat(topic string, value float64, dims map[string]string) {}

var clusterOperator = &configv1.ClusterOperator{
	ObjectMeta: metav1.ObjectMeta{
		Name: "operator",
	},
}

var clusterVersion = &configv1.ClusterVersion{
	ObjectMeta: metav1.ObjectMeta{
		Name: "version",
	},
}

var node = &corev1.Node{
	ObjectMeta: metav1.ObjectMeta{
		Name: "node",
	},
}

var ingressController = &operatorv1.IngressController{
	ObjectMeta: metav1.ObjectMeta{
		Namespace: "openshift-ingress-operator",
		Name:      "ingress-controller",
	},
}

func TestStepRunnerWithInstaller(t *testing.T) {
	ctx := context.Background()

	for _, tt := range []struct {
		name          string
		steps         []steps.Step
		wantEntries   []map[string]types.GomegaMatcher
		wantErr       string
		kubernetescli *fake.Clientset
		configcli     *configfake.Clientset
		operatorcli   *operatorfake.Clientset
	}{
		{
			name: "Failed step run will log cluster version, cluster operator status, and ingress information if available",
			steps: []steps.Step{
				steps.Action(failingFunc),
			},
			wantErr: "oh no!",
			wantEntries: []map[string]types.GomegaMatcher{
				{
					"level": gomega.Equal(logrus.InfoLevel),
					"msg":   gomega.Equal(`running step [Action github.com/Azure/ARO-RP/pkg/cluster.failingFunc]`),
				},
				{
					"level": gomega.Equal(logrus.ErrorLevel),
					"msg":   gomega.Equal("step [Action github.com/Azure/ARO-RP/pkg/cluster.failingFunc] encountered error: oh no!"),
				},
				{
					"level": gomega.Equal(logrus.InfoLevel),
					"msg":   gomega.MatchRegexp(`(?s)github.com/Azure/ARO-RP/pkg/cluster.\(\*manager\).logClusterVersion\-fm:.*"name": "version"`),
				},
				{
					"level": gomega.Equal(logrus.InfoLevel),
					"msg":   gomega.MatchRegexp(`(?s)github.com/Azure/ARO-RP/pkg/cluster.\(\*manager\).logNodes\-fm:.*"name": "node"`),
				},
				{
					"level": gomega.Equal(logrus.InfoLevel),
					"msg":   gomega.MatchRegexp(`(?s)github.com/Azure/ARO-RP/pkg/cluster.\(\*manager\).logClusterOperators\-fm:.*"name": "operator"`),
				},
				{
					"level": gomega.Equal(logrus.InfoLevel),
					"msg":   gomega.MatchRegexp(`(?s)github.com/Azure/ARO-RP/pkg/cluster.\(\*manager\).logIngressControllers\-fm:.*"name": "ingress-controller"`),
				},
			},
			kubernetescli: fake.NewSimpleClientset(node),
			configcli:     configfake.NewSimpleClientset(clusterVersion, clusterOperator),
			operatorcli:   operatorfake.NewSimpleClientset(ingressController),
		},
		{
			name: "Failed step run will not crash if it cannot get the clusterversions, clusteroperators, ingresscontrollers",
			steps: []steps.Step{
				steps.Action(failingFunc),
			},
			wantErr: "oh no!",
			wantEntries: []map[string]types.GomegaMatcher{
				{
					"level": gomega.Equal(logrus.InfoLevel),
					"msg":   gomega.Equal(`running step [Action github.com/Azure/ARO-RP/pkg/cluster.failingFunc]`),
				},
				{
					"level": gomega.Equal(logrus.ErrorLevel),
					"msg":   gomega.Equal("step [Action github.com/Azure/ARO-RP/pkg/cluster.failingFunc] encountered error: oh no!"),
				},
				{
					"level": gomega.Equal(logrus.ErrorLevel),
					"msg":   gomega.Equal(`clusterversions.config.openshift.io "version" not found`),
				},
				{
					"level": gomega.Equal(logrus.InfoLevel),
					"msg":   gomega.Equal(`github.com/Azure/ARO-RP/pkg/cluster.(*manager).logNodes-fm: null`),
				},
				{
					"level": gomega.Equal(logrus.InfoLevel),
					"msg":   gomega.Equal(`github.com/Azure/ARO-RP/pkg/cluster.(*manager).logClusterOperators-fm: null`),
				},
				{
					"level": gomega.Equal(logrus.InfoLevel),
					"msg":   gomega.Equal(`github.com/Azure/ARO-RP/pkg/cluster.(*manager).logIngressControllers-fm: null`),
				},
			},
			kubernetescli: fake.NewSimpleClientset(),
			configcli:     configfake.NewSimpleClientset(),
			operatorcli:   operatorfake.NewSimpleClientset(),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			h, log := testlog.New()
			m := &manager{
				log:           log,
				kubernetescli: tt.kubernetescli,
				configcli:     tt.configcli,
				operatorcli:   tt.operatorcli,
			}

			err := m.runSteps(ctx, tt.steps)
			if err != nil && err.Error() != tt.wantErr ||
				err == nil && tt.wantErr != "" {
				t.Error(err)
			}

			err = testlog.AssertLoggingOutput(h, tt.wantEntries)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestUpdateProvisionedBy(t *testing.T) {
	ctx := context.Background()
	key := "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup/providers/Microsoft.RedHatOpenShift/openShiftClusters/resourceName1"

	openShiftClustersDatabase, _ := testdatabase.NewFakeOpenShiftClusters()
	fixture := testdatabase.NewFixture().WithOpenShiftClusters(openShiftClustersDatabase)
	fixture.AddOpenShiftClusterDocuments(&api.OpenShiftClusterDocument{
		Key: strings.ToLower(key),
		OpenShiftCluster: &api.OpenShiftCluster{
			ID: key,
			Properties: api.OpenShiftClusterProperties{
				ProvisioningState: api.ProvisioningStateCreating,
			},
		},
	})
	err := fixture.Create()
	if err != nil {
		t.Fatal(err)
	}

	clusterdoc, err := openShiftClustersDatabase.Dequeue(ctx)
	if err != nil {
		t.Fatal(err)
	}

	i := &manager{
		doc: clusterdoc,
		db:  openShiftClustersDatabase,
	}
	err = i.updateProvisionedBy(ctx)
	if err != nil {
		t.Error(err)
	}

	// Check it was set to the correct value in the database
	updatedClusterDoc, err := openShiftClustersDatabase.Get(ctx, strings.ToLower(key))
	if err != nil {
		t.Fatal(err)
	}
	if updatedClusterDoc.OpenShiftCluster.Properties.ProvisionedBy != version.GitCommit {
		t.Error("version was not added")
	}
}

func TestInstallationTimeMetrics(t *testing.T) {
	_, log := testlog.New()
	fm := &fakeMetricsEmitter{}

	for _, tt := range []struct {
		name  string
		steps []steps.Step
	}{
		{
			name: "Failed step run will not generate any install time metrics",
			steps: []steps.Step{
				steps.Action(normalFunc),
				steps.Action(failingFunc),
			},
		},
		{
			name: "Succeeded step run will generate a valid install time metrics",
			steps: []steps.Step{
				steps.Action(normalFunc),
				steps.Action(normalFunc),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			m := &manager{
				log: log,
				me:  fm,
			}
			err := m.runSteps(ctx, tt.steps)
			if err != nil {
				if fm.Topic != "" || fm.IntallTime != 0 {
					t.Error("fake metrics obj should be empty when run steps failed")
				}
			} else {
				if fm.Topic != "backend.openshiftcluster.installtime" {
					t.Error("wrong metrics topic")
				}
				if fm.IntallTime < 2 {
					t.Error("wrong metrics value")
				}
			}
		})
	}
}

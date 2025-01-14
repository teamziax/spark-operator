/*
Copyright 2024 The Kubeflow authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mutatingwebhookconfiguration

import (
	"context"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	logger = ctrl.Log.WithName("")
)

// Reconciler reconciles a webhook configuration object.
type Reconciler struct {
	client client.Client
	name   string
}

// MutatingWebhookConfigurationReconciler implements reconcile.Reconciler.
var _ reconcile.Reconciler = &Reconciler{}

// NewReconciler creates a new MutatingWebhookConfigurationReconciler instance.
func NewReconciler(client client.Client, name string) *Reconciler {
	return &Reconciler{
		client: client,
		name:   name,
	}
}

func (r *Reconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		Named("mutating-webhook-configuration-controller").
		Watches(
			&admissionregistrationv1.MutatingWebhookConfiguration{},
			NewEventHandler(),
			builder.WithPredicates(
				NewEventFilter(r.name),
			),
		).
		WithOptions(options).
		Complete(r)
}

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

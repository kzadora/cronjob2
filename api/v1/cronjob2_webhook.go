/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var cronjob2log = logf.Log.WithName("cronjob2-resource")

func (r *CronJob2) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-batch-tutorial-kubebuilder-io-v1-cronjob2,mutating=true,failurePolicy=fail,sideEffects=None,groups=batch.tutorial.kubebuilder.io,resources=cronjob2s,verbs=create;update,versions=v1,name=mcronjob2.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &CronJob2{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *CronJob2) Default() {
	cronjob2log.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-batch-tutorial-kubebuilder-io-v1-cronjob2,mutating=false,failurePolicy=fail,sideEffects=None,groups=batch.tutorial.kubebuilder.io,resources=cronjob2s,verbs=create;update,versions=v1,name=vcronjob2.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &CronJob2{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *CronJob2) ValidateCreate() error {
	cronjob2log.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *CronJob2) ValidateUpdate(old runtime.Object) error {
	cronjob2log.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *CronJob2) ValidateDelete() error {
	cronjob2log.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

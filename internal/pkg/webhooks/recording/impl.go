/*
Copyright 2021 The Kubernetes Authors.

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

package recording

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"sigs.k8s.io/security-profiles-operator/api/profilerecording/v1alpha1"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/config"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/webhooks/utils"
)

type defaultImpl struct {
	client  client.Client
	decoder *admission.Decoder
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . impl
type impl interface {
	ListProfileRecordings(context.Context, ...client.ListOption) (*v1alpha1.ProfileRecordingList, error)
	UpdateResource(context.Context, logr.Logger, client.Object, string) error
	UpdateResourceStatus(context.Context, logr.Logger, client.Object, string) error
	SetDecoder(*admission.Decoder)
	DecodePod(admission.Request) (*corev1.Pod, error)
	LabelSelectorAsSelector(*metav1.LabelSelector) (labels.Selector, error)
	GetOperatorNamespace() string
}

func (d *defaultImpl) ListProfileRecordings(
	ctx context.Context, opts ...client.ListOption,
) (*v1alpha1.ProfileRecordingList, error) {
	profileRecordings := &v1alpha1.ProfileRecordingList{}
	if err := d.client.List(ctx, profileRecordings, opts...); err != nil {
		return nil, fmt.Errorf("list profile recordings: %w", err)
	}
	return profileRecordings, nil
}

func (d *defaultImpl) UpdateResource(
	ctx context.Context,
	logger logr.Logger,
	object client.Object,
	name string,
) error {
	return utils.UpdateResource(ctx, logger, d.client, object, name)
}

func (d *defaultImpl) UpdateResourceStatus(
	ctx context.Context,
	logger logr.Logger,
	object client.Object,
	name string,
) error {
	return utils.UpdateResourceStatus(ctx, logger, d.client.Status(), object, name)
}

func (d *defaultImpl) SetDecoder(decoder *admission.Decoder) {
	d.decoder = decoder
}

func (d *defaultImpl) GetOperatorNamespace() string {
	return config.GetOperatorNamespace()
}

//nolint:gocritic
func (d *defaultImpl) DecodePod(req admission.Request) (*corev1.Pod, error) {
	pod := &corev1.Pod{}
	if err := d.decoder.Decode(req, pod); err != nil {
		return nil, fmt.Errorf("decode pod: %w", err)
	}
	return pod, nil
}

func (*defaultImpl) LabelSelectorAsSelector(
	ps *metav1.LabelSelector,
) (labels.Selector, error) {
	return metav1.LabelSelectorAsSelector(ps)
}

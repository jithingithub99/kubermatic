/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package resources

import (
	"encoding/base64"
	"fmt"

	"k8c.io/kubermatic/v2/pkg/resources/reconciling"

	corev1 "k8s.io/api/core/v1"
)

// ServiceAccountSecretCreator returns a creator function to create a Google Service Account.
func ServiceAccountSecretCreator(data CredentialsData) reconciling.NamedSecretCreatorGetter {
	return func() (string, reconciling.SecretCreator) {
		return GoogleServiceAccountSecretName, func(se *corev1.Secret) (*corev1.Secret, error) {
			credentials, err := GetCredentials(data)
			if err != nil {
				return nil, err
			}

			b, err := base64.StdEncoding.DecodeString(credentials.GCP.ServiceAccount)
			if err != nil {
				return nil, fmt.Errorf("error decoding service account: %w", err)
			}

			se.Type = corev1.SecretTypeOpaque

			if se.Data == nil {
				se.Data = map[string][]byte{}
			}
			se.Data["serviceAccount"] = b

			return se, nil
		}
	}
}

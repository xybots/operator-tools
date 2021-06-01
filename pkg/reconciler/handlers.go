// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reconciler

import (
	"strings"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	ottypes "github.com/banzaicloud/operator-tools/pkg/types"
)

func EnqueueByOwnerAnnotationMapper() handler.MapFunc {
	return func(a client.Object) []reconcile.Request {
		pieces := strings.SplitN(a.GetAnnotations()[ottypes.BanzaiCloudRelatedTo], string(types.Separator), 2)
		if len(pieces) != 2 {
			return []reconcile.Request{}
		}

		return []reconcile.Request{
			{NamespacedName: client.ObjectKey{
				Name:      pieces[1],
				Namespace: pieces[0],
			}},
		}
	}
}

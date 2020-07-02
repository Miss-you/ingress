/*


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

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	githubcomv1 "github.com/Miss-you/ingress/api/v1"
	extensions "k8s.io/api/extensions/v1beta1"
)

// ApisixRoutesReconciler reconciles a ApisixRoutes object
type ApisixRoutesReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=github.com.my.domain,resources=apisixroutes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=github.com.my.domain,resources=apisixroutes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=extensions,resources=ingress,verbs=get;list;watch;create;update;patch;delete

func (r *ApisixRoutesReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("apisixroutes", req.NamespacedName)

	// your logic here
	//get and print
	var ingressObj extensions.Ingress
	fmt.Println("req.NamespacedName: ", req.NamespacedName)
	if err := r.Get(ctx, req.NamespacedName, &ingressObj); err != nil {
		log.Error(err, "unable to fetch ingressObj")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	fmt.Println(ingressObj)

	var apisixRoute githubcomv1.ApisixRoutes
	if err := r.Get(ctx, req.NamespacedName, &apisixRoute); err != nil {
		log.Error(err, "unable to fetch apisixRoute")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	fmt.Println(apisixRoute)

	return ctrl.Result{}, nil
}

func (r *ApisixRoutesReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&githubcomv1.ApisixRoutes{}).
		Complete(r)
}

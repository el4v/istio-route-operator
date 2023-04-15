/*
Copyright 2023.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	routev1alpha1 "github.com/el4v/istio-route-operator/api/v1alpha1"
)

var (
	finalizer = "istio.el4v.com/finalizer"
)

// RouteReconciler reconciles a Route object
type RouteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=istio.el4v.com,resources=routes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=istio.el4v.com,resources=routes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=istio.el4v.com,resources=routes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Route object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *RouteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	recRoute := &routev1alpha1.Route{}

	err := r.Get(ctx, req.NamespacedName, recRoute)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			logger.Info("Route resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		logger.Error(err, "Failed to get Route.")

		return ctrl.Result{}, err
	}

	isPaasAppMarkedToBeDeleted := recRoute.GetDeletionTimestamp() != nil
	if isPaasAppMarkedToBeDeleted {
		// Finalize Route CR (after it was deleted)
		if controllerutil.ContainsFinalizer(recRoute, finalizer) {
			if err := r.finalizeRoute(ctx, recRoute); err != nil {
				return ctrl.Result{}, err
			}

			controllerutil.RemoveFinalizer(recRoute, finalizer)

			err := r.Update(ctx, recRoute)
			if err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(recRoute, finalizer) {
		controllerutil.AddFinalizer(recRoute, finalizer)

		err = r.Update(ctx, recRoute)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	// Get all Route CRs
	routeList := &routev1alpha1.RouteList{}
	routeListForProcessing := &routev1alpha1.RouteList{}

	err = r.List(ctx, routeList)
	if err != nil {
		logger.Error(err, "failed to fetch route list")
		return ctrl.Result{}, err
	}

	for _, r := range routeList.Items {
		if recRoute.Spec.TargetVSName == r.Spec.TargetVSName && recRoute.Spec.TargetVSNamespace == r.Spec.TargetVSNamespace && r.GetDeletionTimestamp() == nil {
			routeListForProcessing.Items = append(routeListForProcessing.Items, r)
		}
	}

	err = r.generateVirtualService(ctx, recRoute, routeListForProcessing)
	if err != nil {
		logger.Error(err, "failed to create virtual service")
		return ctrl.Result{}, err
	}

	err = r.generateEnvoyFilters(ctx, recRoute, routeListForProcessing)
	if err != nil {
		logger.Error(err, "failed to create envoy filter")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// the Route finalize logic
func (r *RouteReconciler) finalizeRoute(ctx context.Context, route *routev1alpha1.Route) error {
	_ = log.FromContext(ctx)

	// TODO: add logic for finalize route

	return nil
}

func (r *RouteReconciler) generateVirtualService(ctx context.Context, route *routev1alpha1.Route, routeList *routev1alpha1.RouteList) error {
	// TODO: add logic for generate target virtual service
	return nil
}

func (r *RouteReconciler) generateEnvoyFilters(ctx context.Context, route *routev1alpha1.Route, routeList *routev1alpha1.RouteList) error {
	// TODO: add logic for generate envoy filter for regex rewrite
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RouteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&routev1alpha1.Route{}).
		Complete(r)
}

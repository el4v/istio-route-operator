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

package v1alpha1

import (
	"time"

	istiov1beta1 "istio.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RouteSpec defines the desired state of Route

type RouteSpec struct {
	TargetVSName      string `json:"targetVSName,omitempty" yaml:"targetVSName"`
	TargetVSNamespace string `json:"targetVSNamespace,omitempty" yaml:"targetVSNamespace"`
	// +optional
	Hosts []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	// +optional
	Gateways []string `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	// +kubebuilder:validation:MinItems=1
	HTTPRules []HTTPRule `json:"httpRules" yaml:"httpRules"`
}

type HTTPRule struct {
	Name string `json:"name" yaml:"name"`
	// +optional
	Match []HTTPMatchRequest `json:"match,omitempty" yaml:"match,omitempty"`
	// +optional
	Route []HTTPRouteDestination `json:"route,omitempty" yaml:"route,omitempty"`
	// +optional
	Redirect HTTPRedirect `json:"redirect,omitempty" yaml:"redirect,omitempty"`
	// +optional
	DirectResponse HTTPDirectResponse `json:"directResponse,omitempty" yaml:"directResponse,omitempty"`
	// +optional
	Delegate Delegate `json:"delegate,omitempty" yaml:"delegate,omitempty"`
	// +optional
	Rewrite HTTPRewrite `json:"rewrite,omitempty" yaml:"rewrite,omitempty"`
	// +optional
	Timeout time.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	// +optional
	Retries HTTPRetry `json:"retries,omitempty" yaml:"retries,omitempty"`
	// +optional
	Fault HTTPFaultInjection `json:"fault,omitempty" yaml:"fault,omitempty"`
	// +optional
	Mirror Destination `json:"mirror,omitempty" yaml:"mirror,omitempty"`
	// +optional
	MirrorPercentage int `json:"mirrorPercentage,omitempty" yaml:"mirrorPercentage,omitempty"`
	// +optional
	CorsPolicy CorsPolicy `json:"corsPolicy,omitempty" yaml:"corsPolicy,omitempty"`
	// +optional
	Headers Headers `json:"headers,omitempty" yaml:"headers,omitempty"`
}

type HTTPRetry struct {
	Attempts int32 `yaml:"attempts,omitempty" json:"attempts,omitempty"`
	// +optional
	PerTryTimeout time.Duration `yaml:"perTryTimeout,omitempty" json:"perTryTimeout,omitempty"`
	// +optional
	RetryOn string `yaml:"retryOn,omitempty" json:"retryOn,omitempty"`
	// +optional
	RetryRemoteLocalities bool `yaml:"retryRemoteLocalities,omitempty" json:"retryRemoteLocalities,omitempty"`
}

type HTTPFaultInjection struct {
	// +optional
	Delay HTTPFaultInjectionDelay `yaml:"delay,omitempty" json:"delay,omitempty"`
	// +optional
	Abort HTTPFaultInjectionAbort `yaml:"abort,omitempty" json:"abort,omitempty"`
}

type HTTPFaultInjectionDelay struct {
	FixedDelay time.Duration `yaml:"fixedDelay,omitempty" json:"fixedDelay,omitempty"`
	// +optional
	Percentage Percent `yaml:"percentage,omitempty" json:"percentage,omitempty"`
	// +optional
	Percent int32 `yaml:"percent,omitempty" json:"percent,omitempty"`
}

type HTTPFaultInjectionAbort struct {
	HttpStatus int32 `yaml:"httpStatus,omitempty" json:"httpStatus,omitempty"`
	// +optional
	GrpcStatus string `yaml:"grpcStatus,omitempty" json:"grpcStatus,omitempty"`
	// +optional
	Percentage Percent `yaml:"percent,omitempty" json:"percent,omitempty"`
}

type Percent struct {
	// +optional
	Value float64 `yaml:"value,omitempty" json:"value,omitempty"`
}

type CorsPolicy struct {
	// deprecated field
	// +optional
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// +optional
	AllowOrigins ListStringMatch `json:"allowOrigins,omitempty" yaml:"allowOrigins,omitempty"`
	// +optional
	AllowMethods []string `json:"allowMethods,omitempty" yaml:"allowMethods,omitempty"`
	// +optional
	AllowHeaders []string `json:"allowHeaders,omitempty" yaml:"allowHeaders,omitempty"`
	// +optional
	ExposeHeaders []string `json:"exposeHeaders,omitempty" yaml:"exposeHeaders,omitempty"`
	// +optional
	MaxAge string `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`
	// +optional
	AllowCredentials *bool `json:"allowCredentials,omitempty" yaml:"allowCredentials,omitempty"`
}

type HTTPRewrite struct {
	// deprecated field
	// +optional
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// +optional
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"` //nolint:revive,stylecheck
	// +optional
	Authority string `json:"authority,omitempty" yaml:"authority,omitempty"`
}

type Delegate struct {
	// deprecated field
	// +optional
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}

type HTTPRedirect struct {
	// deprecated field
	// +optional
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// +optional
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"` //nolint:revive,stylecheck
	// +optional
	Authority string `json:"authority,omitempty" yaml:"authority,omitempty"`
	// +optional
	Port uint32 `json:"port,omitempty" yaml:"port,omitempty"`
	// +optional
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	// +optional
	RedirectCode uint32 `json:"redirectCode,omitempty" yaml:"redirectCode,omitempty"`
}

type HTTPDirectResponse struct {
	Status uint32 `json:"status,omitempty" yaml:"status"`
	// +optional
	Body HTTPBody `json:"body,omitempty" yaml:"body,omitempty"`
}

type HTTPBody struct {
	// +optional
	String string `json:"string,omitempty" yaml:"string,omitempty"`
	// +optional
	Bytes byte `json:"bytes,omitempty" yaml:"bytes,omitempty"`
}

type HTTPMatchRequest struct {
	// +optional
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// +optional
	Uri StringMatch `json:"uri,omitempty" yaml:"uri,omitempty"` //nolint:revive,stylecheck
	// +optional
	Scheme StringMatch `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	// +optional
	Method StringMatch `json:"method,omitempty" yaml:"method,omitempty"`
	// +optional
	Authority StringMatch `json:"authority,omitempty" yaml:"authority,omitempty"`
	// +optional
	Headers MapStringMatch `json:"headers,omitempty" yaml:"headers,omitempty"`
	// +optional
	Port uint32 `json:"port,omitempty" yaml:"port,omitempty"`
	// +optional
	SourceLabels map[string]string `json:"sourceLabels,omitempty" yaml:"sourceLabels,omitempty"`
	// +optional
	Gateways []string `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	// +optional
	QueryParams MapStringMatch `json:"queryParams,omitempty" yaml:"queryParams,omitempty"`
	// +optional
	IgnoreUriCase bool `json:"ignoreUriCase,omitempty" yaml:"ignoreUriCase,omitempty"` //nolint:revive,stylecheck
	// +optional
	WithoutHeaders MapStringMatch `json:"withoutHeaders,omitempty" yaml:"withoutHeaders,omitempty"`
	// +optional
	SourceNamespace string `json:"sourceNamespace,omitempty" yaml:"sourceNamespace,omitempty"`
}

type HTTPRouteDestination struct {
	// +optional
	Enabled     bool        `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Destination Destination `json:"destination" yaml:"destination"`
	// +optional
	Weight int32 `json:"weight,omitempty" yaml:"weight,omitempty"`
	// +optional
	Headers Headers `json:"headers,omitempty" yaml:"headers,omitempty"`
}

type Destination struct {
	Host string `json:"host" yaml:"host"`
	// +optional
	Subset string `json:"subset,omitempty" yaml:"subset,omitempty"`
	// +optional
	Port PortSettings `json:"port" yaml:"port"`
}

type PortSettings struct {
	Number uint32 `json:"number,omitempty" yaml:"number,omitempty"`
}

type Headers struct {
	// deprecated field
	// +optional
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// +optional
	Request HeaderOperations `json:"request,omitempty" yaml:"request,omitempty"`
	// +optional
	Response HeaderOperations `json:"response,omitempty" yaml:"response,omitempty"`
}

type HeaderOperations struct {
	// +optional
	Set map[string]string `json:"set,omitempty" yaml:"set,omitempty"`
	// +optional
	Add map[string]string `json:"add,omitempty" yaml:"add,omitempty"`
	// +optional
	Remove []string `json:"remove,omitempty" yaml:"remove,omitempty"`
}

type StringMatch struct {
	// +optional
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	// +optional
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`
}

func (in *StringMatch) String() string {
	if in.Exact != "" {
		return "exact"
	}

	if in.Prefix != "" {
		return "prefix"
	}

	if in.Regex != "" {
		return "regex"
	}

	return ""
}

func (in *StringMatch) ToStringMatch() *istiov1beta1.StringMatch {
	if in.Exact != "" {
		return &istiov1beta1.StringMatch{MatchType: &istiov1beta1.StringMatch_Exact{Exact: in.Exact}}
	}

	if in.Prefix != "" {
		return &istiov1beta1.StringMatch{MatchType: &istiov1beta1.StringMatch_Prefix{Prefix: in.Prefix}}
	}

	if in.Regex != "" {
		return &istiov1beta1.StringMatch{MatchType: &istiov1beta1.StringMatch_Regex{Regex: in.Regex}}
	}

	return nil
}

type MapStringMatch map[string]StringMatch

func (in MapStringMatch) ToStringMatch() map[string]*istiov1beta1.StringMatch {
	ret := make(map[string]*istiov1beta1.StringMatch)

	for k, v := range in {
		if v.Exact != "" {
			ret[k] = &istiov1beta1.StringMatch{MatchType: &istiov1beta1.StringMatch_Exact{Exact: v.Exact}}
			continue
		}

		if v.Prefix != "" {
			ret[k] = &istiov1beta1.StringMatch{MatchType: &istiov1beta1.StringMatch_Prefix{Prefix: v.Prefix}}
			continue
		}

		if v.Regex != "" {
			ret[k] = &istiov1beta1.StringMatch{MatchType: &istiov1beta1.StringMatch_Regex{Regex: v.Regex}}
			continue
		}
	}

	return ret
}

type ListStringMatch []StringMatch

func (in ListStringMatch) ToStringMatch() []*istiov1beta1.StringMatch {
	var ret []*istiov1beta1.StringMatch

	for _, v := range in {
		if v.Exact != "" {
			ret = append(ret, &istiov1beta1.StringMatch{MatchType: &istiov1beta1.StringMatch_Exact{Exact: v.Exact}})
			continue
		}

		if v.Prefix != "" {
			ret = append(ret, &istiov1beta1.StringMatch{MatchType: &istiov1beta1.StringMatch_Prefix{Prefix: v.Prefix}})
			continue
		}

		if v.Regex != "" {
			ret = append(ret, &istiov1beta1.StringMatch{MatchType: &istiov1beta1.StringMatch_Regex{Regex: v.Regex}})
			continue
		}
	}

	return ret
}

// RouteStatus defines the observed state of Route
type RouteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Route is the Schema for the routes API
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteSpec   `json:"spec,omitempty"`
	Status RouteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RouteList contains a list of Route
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}

package models

type ObjectMeta struct {
	Name                       string            `json:"name,omitempty"`
	GenerateName               string            `json:"generateName,omitempty"`
	Namespace                  string            `json:"namespace,omitempty"`
	SelfLink                   string            `json:"selfLink,omitempty"`
	UID                        string            `json:"uid,omitempty"`
	ResourceVersion            string            `json:"resourceVersion,omitempty"`
	Generation                 int64             `json:"generation,omitempty"`
	CreationTimestamp          string            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp          string            `json:"deletionTimestamp,omitempty"`
	DeletionGracePeriodSeconds int64             `json:"deletionGracePeriodSeconds,omitempty"`
	Labels                     map[string]string `json:"labels,omitempty"`
	Annotations                map[string]string `json:"annotations,omitempty"`
	OwnerReferences            []OwnerReference  `json:"ownerReferences,omitempty"`
	Finalizers                 []string          `json:"finalizers,omitempty"`
}
type OwnerReference struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Uid        string `json:"uid"`
	Controller bool   `json:"controller,omitempty"`
}
type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

type ListMeta struct {
	SelfLink        string `json:"selfLink,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
}
type Endpoints struct {
	TypeMeta `json:",inline"`
	Metadata ObjectMeta
	Subsets  []EndpointSubset
	Ages     string `json:ages`
}
type EndpointSubset struct {
	Addresses         []EndpointAddress
	NotReadyAddresses []EndpointAddress
	Ports             []EndpointPort
}

type EndpointAddress struct {
	Ip        string          `json:ip`
	Hostname  string          `json:hostname,omitempty`
	TargetRef ObjectReference `json:targetRef,omitempty`
}
type EndpointPort struct {
	Name     string `json:name,omitempty`
	Port     int32  `json:port`
	Protocol string `json:Protocol,omitempty`
}

type ObjectReference struct {
	Kind            string `json:"kind,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	Name            string `json:"name,omitempty"`
	Uid             string `json:"uid,omitempty"`
	APIVersion      string `json:"apiVersion,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	FieldPath       string `json:"fieldPath,omitempty"`
}

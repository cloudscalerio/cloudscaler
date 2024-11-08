// +kubebuilder:object:generate=true
package common

type ScalerPeriod struct {
	// +kubebuilder:validation:Enum=down;nominal;up;restore
	Type string     `json:"type"`
	Time TimePeriod `json:"time"`
	// Minimum replicas
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// Maximum replicas
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
}

type TimePeriod struct {
	Days []string `json:"days"`
	// +kubebuilder:validation:Pattern=`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$`
	StartTime string `json:"startTime"`
	// +kubebuilder:validation:Pattern=`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$`
	EndTime  string `json:"endTime"`
	Timezone string `json:"timezone,omitempty"`
	// Run once at StartTime
	Once bool `json:"once,omitempty"`
	// Grace period in seconds for deployments before scaling down
	GracePeriod int `json:"gracePeriod,omitempty"`
}

// ScalerStatus defines the observed state of Scaler
type ScalerStatus struct {
	CurrentPeriod *ScalerStatusPeriod `json:"currentPeriod,omitempty"`
	Comments      string              `json:"comments,omitempty"`
}

type ScalerStatusPeriod struct {
	Spec       *ScalerPeriod         `json:"spec"`
	SpecSHA    string                `json:"specSHA"`
	Successful []ScalerStatusSuccess `json:"success,omitempty"`
	Failed     []ScalerStatusFailed  `json:"failed,omitempty"`
}

type ScalerStatusSuccess struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type ScalerStatusFailed struct {
	Kind   string `json:"kind"`
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

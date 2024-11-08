//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package common

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalerPeriod) DeepCopyInto(out *ScalerPeriod) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalerPeriod.
func (in *ScalerPeriod) DeepCopy() *ScalerPeriod {
	if in == nil {
		return nil
	}
	out := new(ScalerPeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalerStatus) DeepCopyInto(out *ScalerStatus) {
	*out = *in
	if in.CurrentPeriod != nil {
		in, out := &in.CurrentPeriod, &out.CurrentPeriod
		*out = new(ScalerStatusPeriod)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalerStatus.
func (in *ScalerStatus) DeepCopy() *ScalerStatus {
	if in == nil {
		return nil
	}
	out := new(ScalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalerStatusFailed) DeepCopyInto(out *ScalerStatusFailed) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalerStatusFailed.
func (in *ScalerStatusFailed) DeepCopy() *ScalerStatusFailed {
	if in == nil {
		return nil
	}
	out := new(ScalerStatusFailed)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalerStatusPeriod) DeepCopyInto(out *ScalerStatusPeriod) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ScalerPeriod)
		(*in).DeepCopyInto(*out)
	}
	if in.Successful != nil {
		in, out := &in.Successful, &out.Successful
		*out = make([]ScalerStatusSuccess, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]ScalerStatusFailed, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalerStatusPeriod.
func (in *ScalerStatusPeriod) DeepCopy() *ScalerStatusPeriod {
	if in == nil {
		return nil
	}
	out := new(ScalerStatusPeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalerStatusSuccess) DeepCopyInto(out *ScalerStatusSuccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalerStatusSuccess.
func (in *ScalerStatusSuccess) DeepCopy() *ScalerStatusSuccess {
	if in == nil {
		return nil
	}
	out := new(ScalerStatusSuccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimePeriod) DeepCopyInto(out *TimePeriod) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimePeriod.
func (in *TimePeriod) DeepCopy() *TimePeriod {
	if in == nil {
		return nil
	}
	out := new(TimePeriod)
	in.DeepCopyInto(out)
	return out
}

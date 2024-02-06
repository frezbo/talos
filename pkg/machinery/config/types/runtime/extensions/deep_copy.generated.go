// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type ServiceConfigV1Alpha1 -pointer-receiver -header-file ../../../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package extensions

// DeepCopy generates a deep copy of *ServiceConfigV1Alpha1.
func (o *ServiceConfigV1Alpha1) DeepCopy() *ServiceConfigV1Alpha1 {
	var cp ServiceConfigV1Alpha1 = *o
	if o.ServiceConfigFiles != nil {
		cp.ServiceConfigFiles = make([]ConfigFile, len(o.ServiceConfigFiles))
		copy(cp.ServiceConfigFiles, o.ServiceConfigFiles)
	}
	if o.ServiceEnvironment != nil {
		cp.ServiceEnvironment = make([]string, len(o.ServiceEnvironment))
		copy(cp.ServiceEnvironment, o.ServiceEnvironment)
	}
	return &cp
}

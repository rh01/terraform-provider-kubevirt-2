/*
Copyright The Kubernetes Authors.

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

package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_CertificateSigningRequest = map[string]string{
	"":       "Describes a certificate signing request",
<<<<<<< HEAD
	"spec":   "spec contains the certificate request, and is immutable after creation. Only the request, signerName, expirationSeconds, and usages fields can be set on creation. Other fields are derived by Kubernetes and cannot be modified by users.",
=======
	"spec":   "The certificate request itself and any additional information.",
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
	"status": "Derived information about the request.",
}

func (CertificateSigningRequest) SwaggerDoc() map[string]string {
	return map_CertificateSigningRequest
}

var map_CertificateSigningRequestCondition = map[string]string{
	"type":               "type of the condition. Known conditions include \"Approved\", \"Denied\", and \"Failed\".",
	"status":             "Status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be \"False\" or \"Unknown\". Defaults to \"True\". If unset, should be treated as \"True\".",
	"reason":             "brief reason for the request state",
	"message":            "human readable message with details about the request state",
	"lastUpdateTime":     "timestamp for the last update to this condition",
	"lastTransitionTime": "lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.",
}

func (CertificateSigningRequestCondition) SwaggerDoc() map[string]string {
	return map_CertificateSigningRequestCondition
}

var map_CertificateSigningRequestSpec = map[string]string{
<<<<<<< HEAD
	"":                  "CertificateSigningRequestSpec contains the certificate request.",
	"request":           "Base64-encoded PKCS#10 CSR data",
	"signerName":        "Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be defaulted:\n 1. If it's a kubelet client certificate, it is assigned\n    \"kubernetes.io/kube-apiserver-client-kubelet\".\n 2. If it's a kubelet serving certificate, it is assigned\n    \"kubernetes.io/kubelet-serving\".\n 3. Otherwise, it is assigned \"kubernetes.io/legacy-unknown\".\nDistribution of trust for signers happens out of band. You can select on this field using `spec.signerName`.",
	"expirationSeconds": "expirationSeconds is the requested duration of validity of the issued certificate. The certificate signer may issue a certificate with a different validity duration so a client must check the delta between the notBefore and and notAfter fields in the issued certificate to determine the actual duration.\n\nThe v1.22+ in-tree implementations of the well-known Kubernetes signers will honor this field as long as the requested duration is not greater than the maximum duration they will honor per the --cluster-signing-duration CLI flag to the Kubernetes controller manager.\n\nCertificate signers may not honor this field for various reasons:\n\n  1. Old signer that is unaware of the field (such as the in-tree\n     implementations prior to v1.22)\n  2. Signer whose configured maximum is shorter than the requested duration\n  3. Signer whose configured minimum is longer than the requested duration\n\nThe minimum valid value for expirationSeconds is 600, i.e. 10 minutes.",
	"usages":            "allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3\n     https://tools.ietf.org/html/rfc5280#section-4.2.1.12\nValid values are:\n \"signing\",\n \"digital signature\",\n \"content commitment\",\n \"key encipherment\",\n \"key agreement\",\n \"data encipherment\",\n \"cert sign\",\n \"crl sign\",\n \"encipher only\",\n \"decipher only\",\n \"any\",\n \"server auth\",\n \"client auth\",\n \"code signing\",\n \"email protection\",\n \"s/mime\",\n \"ipsec end system\",\n \"ipsec tunnel\",\n \"ipsec user\",\n \"timestamping\",\n \"ocsp signing\",\n \"microsoft sgc\",\n \"netscape sgc\"",
	"username":          "Information about the requesting user. See user.Info interface for details.",
	"uid":               "UID information about the requesting user. See user.Info interface for details.",
	"groups":            "Group information about the requesting user. See user.Info interface for details.",
	"extra":             "Extra information about the requesting user. See user.Info interface for details.",
=======
	"":           "This information is immutable after the request is created. Only the Request and Usages fields can be set on creation, other fields are derived by Kubernetes and cannot be modified by users.",
	"request":    "Base64-encoded PKCS#10 CSR data",
	"signerName": "Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be defaulted:\n 1. If it's a kubelet client certificate, it is assigned\n    \"kubernetes.io/kube-apiserver-client-kubelet\".\n 2. If it's a kubelet serving certificate, it is assigned\n    \"kubernetes.io/kubelet-serving\".\n 3. Otherwise, it is assigned \"kubernetes.io/legacy-unknown\".\nDistribution of trust for signers happens out of band. You can select on this field using `spec.signerName`.",
	"usages":     "allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3\n     https://tools.ietf.org/html/rfc5280#section-4.2.1.12\nValid values are:\n \"signing\",\n \"digital signature\",\n \"content commitment\",\n \"key encipherment\",\n \"key agreement\",\n \"data encipherment\",\n \"cert sign\",\n \"crl sign\",\n \"encipher only\",\n \"decipher only\",\n \"any\",\n \"server auth\",\n \"client auth\",\n \"code signing\",\n \"email protection\",\n \"s/mime\",\n \"ipsec end system\",\n \"ipsec tunnel\",\n \"ipsec user\",\n \"timestamping\",\n \"ocsp signing\",\n \"microsoft sgc\",\n \"netscape sgc\"",
	"username":   "Information about the requesting user. See user.Info interface for details.",
	"uid":        "UID information about the requesting user. See user.Info interface for details.",
	"groups":     "Group information about the requesting user. See user.Info interface for details.",
	"extra":      "Extra information about the requesting user. See user.Info interface for details.",
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
}

func (CertificateSigningRequestSpec) SwaggerDoc() map[string]string {
	return map_CertificateSigningRequestSpec
}

var map_CertificateSigningRequestStatus = map[string]string{
	"conditions":  "Conditions applied to the request, such as approval or denial.",
	"certificate": "If request was approved, the controller will place the issued certificate here.",
}

func (CertificateSigningRequestStatus) SwaggerDoc() map[string]string {
	return map_CertificateSigningRequestStatus
}

// AUTO-GENERATED FUNCTIONS END HERE

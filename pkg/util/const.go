package util

import "time"

const (
	// Operand and operator run in the same namespace
	OperatorNamespace = "openshift-cluster-csi-drivers"
	OperatorName      = "ibm-vpc-block-csi-driver-operator"
	OperandName       = "ibm-vpc-block-csi-driver"

	InstanceName = "vpc.block.csi.ibm.io"
	Resync       = 20 * time.Minute

	// Name of the secret provided by cloud-credentials-operator
	CloudCredentialSecretName = "ibm-cloud-credentials"

	// Name of configmap with cluster info provided by cloud-credentials-operator
	ConfigMapNamespace = "openshift-cloud-controller-manager"
	ConfigMapName      = "cloud-conf"

	// Name of secret created in operand namespace
	IBMCSIDriverSecretName = "storage-secret-store"

	// Name of the configmap with the injected trusted CA bundle
	TrustedCAConfigMap = "ibm-vpc-block-csi-driver-trusted-ca-bundle"
)

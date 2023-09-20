package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/PatchExtension.json
func ExampleExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkubernetesconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExtensionsClient().BeginUpdate(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "ClusterMonitor", armkubernetesconfiguration.PatchExtension{
		Properties: &armkubernetesconfiguration.PatchExtensionProperties{
			AutoUpgradeMinorVersion: to.Ptr(true),
			ConfigurationProtectedSettings: map[string]*string{
				"omsagent.secret.key": to.Ptr("secretKeyValue01"),
			},
			ConfigurationSettings: map[string]*string{
				"omsagent.env.clusterName": to.Ptr("clusterName1"),
				"omsagent.secret.wsid":     to.Ptr("a38cef99-5a89-52ed-b6db-22095c23664b"),
			},
			ReleaseTrain: to.Ptr("Preview"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Extension = armkubernetesconfiguration.Extension{
	// 	Properties: &armkubernetesconfiguration.ExtensionProperties{
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		ConfigurationSettings: map[string]*string{
	// 			"omsagent.env.clusterName": to.Ptr("clusterName1"),
	// 			"omsagent.secret.wsid": to.Ptr("a38cef99-5a89-52ed-b6db-22095c23664b"),
	// 		},
	// 		ExtensionType: to.Ptr("azuremonitor-containers"),
	// 		ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateSucceeded),
	// 		ReleaseTrain: to.Ptr("Preview"),
	// 		Scope: &armkubernetesconfiguration.Scope{
	// 			Cluster: &armkubernetesconfiguration.ScopeCluster{
	// 				ReleaseNamespace: to.Ptr("kube-system"),
	// 			},
	// 		},
	// 		Statuses: []*armkubernetesconfiguration.ExtensionStatus{
	// 		},
	// 		Version: to.Ptr("0.1.4"),
	// 	},
	// }
}

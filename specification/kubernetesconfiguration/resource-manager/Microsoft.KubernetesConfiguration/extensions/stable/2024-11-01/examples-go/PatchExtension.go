package armextensions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armextensions"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba0c086df0ebe03a61579485c1c10de0d17804b2/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensions/stable/2024-11-01/examples/PatchExtension.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextensions.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginUpdate(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "ClusterMonitor", armextensions.PatchExtension{
		Properties: &armextensions.PatchExtensionProperties{
			AutoUpgradeMinorVersion: to.Ptr(true),
			ConfigurationProtectedSettings: map[string]*string{
				"omsagent.secret.key": to.Ptr("secretKeyValue01"),
			},
			ConfigurationSettings: map[string]*string{
				"omsagent.env.clusterName": to.Ptr("clusterName1"),
				"omsagent.secret.wsid":     to.Ptr("fakeTokenPlaceholder"),
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
	// res.Extension = armextensions.Extension{
	// 	Properties: &armextensions.ExtensionProperties{
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		ConfigurationSettings: map[string]*string{
	// 			"omsagent.env.clusterName": to.Ptr("clusterName1"),
	// 			"omsagent.secret.wsid": to.Ptr("fakeTokenPlaceholder"),
	// 		},
	// 		ExtensionType: to.Ptr("azuremonitor-containers"),
	// 		ProvisioningState: to.Ptr(armextensions.ProvisioningStateSucceeded),
	// 		ReleaseTrain: to.Ptr("Preview"),
	// 		Scope: &armextensions.Scope{
	// 			Cluster: &armextensions.ScopeCluster{
	// 				ReleaseNamespace: to.Ptr("kube-system"),
	// 			},
	// 		},
	// 		Statuses: []*armextensions.ExtensionStatus{
	// 		},
	// 		Version: to.Ptr("0.1.4"),
	// 	},
	// }
}

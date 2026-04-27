package armextensions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armextensions"
)

// Generated from example definition: 2025-03-01/PatchExtension.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextensions.NewClientFactory("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginUpdate(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "ClusterMonitor", armextensions.PatchExtension{
		Properties: &armextensions.PatchExtensionProperties{
			AutoUpgradeMode:         to.Ptr(armextensions.AutoUpgradeModeCompatible),
			AutoUpgradeMinorVersion: to.Ptr(true),
			ReleaseTrain:            to.Ptr("Preview"),
			ConfigurationSettings: map[string]*string{
				"omsagent.secret.wsid":     to.Ptr("fakeTokenPlaceholder"),
				"omsagent.env.clusterName": to.Ptr("clusterName1"),
			},
			ConfigurationProtectedSettings: map[string]*string{
				"omsagent.secret.key": to.Ptr("secretKeyValue01"),
			},
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
	// res = armextensions.ClientUpdateResponse{
	// 	Extension: &armextensions.Extension{
	// 		Properties: &armextensions.ExtensionProperties{
	// 			ExtensionType: to.Ptr("azuremonitor-containers"),
	// 			AutoUpgradeMode: to.Ptr(armextensions.AutoUpgradeModeCompatible),
	// 			AutoUpgradeMinorVersion: to.Ptr(true),
	// 			ReleaseTrain: to.Ptr("Preview"),
	// 			Version: to.Ptr("0.1.4"),
	// 			Scope: &armextensions.Scope{
	// 				Cluster: &armextensions.ScopeCluster{
	// 					ReleaseNamespace: to.Ptr("kube-system"),
	// 				},
	// 			},
	// 			ConfigurationSettings: map[string]*string{
	// 				"omsagent.secret.wsid": to.Ptr("fakeTokenPlaceholder"),
	// 				"omsagent.env.clusterName": to.Ptr("clusterName1"),
	// 			},
	// 			ProvisioningState: to.Ptr(armextensions.ProvisioningStateSucceeded),
	// 			Statuses: []*armextensions.ExtensionStatus{
	// 			},
	// 		},
	// 	},
	// }
}

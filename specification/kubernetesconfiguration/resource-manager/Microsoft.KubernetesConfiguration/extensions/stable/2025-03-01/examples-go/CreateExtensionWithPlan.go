package armextensions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armextensions"
)

// Generated from example definition: 2025-03-01/CreateExtensionWithPlan.json
func ExampleClient_BeginCreate_createExtensionWithPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextensions.NewClientFactory("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreate(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "azureVote", armextensions.Extension{
		Properties: &armextensions.ExtensionProperties{
			ExtensionType:           to.Ptr("azure-vote"),
			AutoUpgradeMode:         to.Ptr(armextensions.AutoUpgradeModeCompatible),
			AutoUpgradeMinorVersion: to.Ptr(true),
			ReleaseTrain:            to.Ptr("Preview"),
		},
		Plan: &armextensions.Plan{
			Name:      to.Ptr("azure-vote-standard"),
			Publisher: to.Ptr("Microsoft"),
			Product:   to.Ptr("azure-vote-standard-offer-id"),
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
	// res = armextensions.ClientCreateResponse{
	// 	Extension: &armextensions.Extension{
	// 		ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/extensions/azureVote"),
	// 		Type: to.Ptr("Microsoft.KubernetesConfiguration/extensions"),
	// 		Name: to.Ptr("azureVote"),
	// 		SystemData: &armextensions.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armextensions.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T05:10:57.027Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armextensions.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T05:10:57.027Z"); return t}()),
	// 		},
	// 		Plan: &armextensions.Plan{
	// 			Name: to.Ptr("azure-vote-standard"),
	// 			Publisher: to.Ptr("Microsoft"),
	// 			Product: to.Ptr("azure-vote-standard-offer-id"),
	// 		},
	// 		Properties: &armextensions.ExtensionProperties{
	// 			ExtensionType: to.Ptr("azure-vote"),
	// 			AutoUpgradeMode: to.Ptr(armextensions.AutoUpgradeModeCompatible),
	// 			AutoUpgradeMinorVersion: to.Ptr(true),
	// 			ReleaseTrain: to.Ptr("Preview"),
	// 			Version: to.Ptr("0.1.4"),
	// 			Scope: &armextensions.Scope{
	// 				Cluster: &armextensions.ScopeCluster{
	// 					ReleaseNamespace: to.Ptr("kube-system"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armextensions.ProvisioningStateCreating),
	// 			Statuses: []*armextensions.ExtensionStatus{
	// 			},
	// 			CurrentVersion: to.Ptr("0.1.4"),
	// 			IsSystemExtension: to.Ptr(false),
	// 		},
	// 	},
	// }
}

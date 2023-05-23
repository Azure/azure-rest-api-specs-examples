package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/CreateExtensionWithPlan.json
func ExampleExtensionsClient_BeginCreate_createExtensionWithPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkubernetesconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExtensionsClient().BeginCreate(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "azureVote", armkubernetesconfiguration.Extension{
		Plan: &armkubernetesconfiguration.Plan{
			Name:      to.Ptr("azure-vote-standard"),
			Product:   to.Ptr("azure-vote-standard-offer-id"),
			Publisher: to.Ptr("Microsoft"),
		},
		Properties: &armkubernetesconfiguration.ExtensionProperties{
			AutoUpgradeMinorVersion: to.Ptr(true),
			ExtensionType:           to.Ptr("azure-vote"),
			ReleaseTrain:            to.Ptr("Preview"),
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
	// 	Name: to.Ptr("azureVote"),
	// 	Type: to.Ptr("Microsoft.KubernetesConfiguration/extensions"),
	// 	ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/extensions/azureVote"),
	// 	Plan: &armkubernetesconfiguration.Plan{
	// 		Name: to.Ptr("azure-vote-standard"),
	// 		Product: to.Ptr("azure-vote-standard-offer-id"),
	// 		Publisher: to.Ptr("Microsoft"),
	// 	},
	// 	Properties: &armkubernetesconfiguration.ExtensionProperties{
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		CurrentVersion: to.Ptr("0.1.4"),
	// 		ExtensionType: to.Ptr("azure-vote"),
	// 		IsSystemExtension: to.Ptr(false),
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
	// 	SystemData: &armkubernetesconfiguration.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T05:10:57.027Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T05:10:57.027Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
	// 	},
	// }
}

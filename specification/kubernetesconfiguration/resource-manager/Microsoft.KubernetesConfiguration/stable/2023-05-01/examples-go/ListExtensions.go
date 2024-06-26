package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/ListExtensions.json
func ExampleExtensionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkubernetesconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExtensionsClient().NewListPager("rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ExtensionsList = armkubernetesconfiguration.ExtensionsList{
		// 	Value: []*armkubernetesconfiguration.Extension{
		// 		{
		// 			Name: to.Ptr("ClusterMonitor"),
		// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/extensions"),
		// 			ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/extensions/ClusterMonitor"),
		// 			Properties: &armkubernetesconfiguration.ExtensionProperties{
		// 				AutoUpgradeMinorVersion: to.Ptr(false),
		// 				ConfigurationSettings: map[string]*string{
		// 					"omsagent.env.clusterName": to.Ptr("clusterName1"),
		// 					"omsagent.secret.wsid": to.Ptr("a38cef99-5a89-52ed-b6db-22095c23664b"),
		// 				},
		// 				CurrentVersion: to.Ptr("0.1.4"),
		// 				ExtensionType: to.Ptr("azuremonitor-containers"),
		// 				IsSystemExtension: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateSucceeded),
		// 				ReleaseTrain: to.Ptr("Preview"),
		// 				Scope: &armkubernetesconfiguration.Scope{
		// 					Cluster: &armkubernetesconfiguration.ScopeCluster{
		// 						ReleaseNamespace: to.Ptr("kube-system"),
		// 					},
		// 				},
		// 				Statuses: []*armkubernetesconfiguration.ExtensionStatus{
		// 				},
		// 				Version: to.Ptr("0.1.4"),
		// 			},
		// 			SystemData: &armkubernetesconfiguration.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T05:10:57.027Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T05:10:57.027Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("App1Monitor"),
		// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/extensions"),
		// 			ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/extensions/BackupVault01"),
		// 			Properties: &armkubernetesconfiguration.ExtensionProperties{
		// 				AutoUpgradeMinorVersion: to.Ptr(true),
		// 				ConfigurationSettings: map[string]*string{
		// 				},
		// 				CurrentVersion: to.Ptr("1.0.1"),
		// 				ExtensionType: to.Ptr("Microsoft.RecoveryServices/recoveryVault"),
		// 				IsSystemExtension: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateSucceeded),
		// 				ReleaseTrain: to.Ptr("Stable"),
		// 				Scope: &armkubernetesconfiguration.Scope{
		// 					Cluster: &armkubernetesconfiguration.ScopeCluster{
		// 						ReleaseNamespace: to.Ptr("myKVNamespace"),
		// 					},
		// 				},
		// 				Statuses: []*armkubernetesconfiguration.ExtensionStatus{
		// 				},
		// 			},
		// 			SystemData: &armkubernetesconfiguration.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T04:09:23.011Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T04:09:23.011Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("azureVote"),
		// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/extensions"),
		// 			ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/extensions/azureVote"),
		// 			Plan: &armkubernetesconfiguration.Plan{
		// 				Name: to.Ptr("azure-vote-standard"),
		// 				Product: to.Ptr("azure-vote-standard-offer-id"),
		// 				Publisher: to.Ptr("Microsoft"),
		// 			},
		// 			Properties: &armkubernetesconfiguration.ExtensionProperties{
		// 				AutoUpgradeMinorVersion: to.Ptr(false),
		// 				ConfigurationSettings: map[string]*string{
		// 					"omsagent.env.clusterName": to.Ptr("clusterName1"),
		// 					"omsagent.secret.wsid": to.Ptr("a38cef99-5a89-52ed-b6db-22095c23664b"),
		// 				},
		// 				CurrentVersion: to.Ptr("0.1.4"),
		// 				ExtensionType: to.Ptr("azure-vote"),
		// 				IsSystemExtension: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armkubernetesconfiguration.ProvisioningStateSucceeded),
		// 				ReleaseTrain: to.Ptr("Preview"),
		// 				Scope: &armkubernetesconfiguration.Scope{
		// 					Cluster: &armkubernetesconfiguration.ScopeCluster{
		// 						ReleaseNamespace: to.Ptr("kube-system"),
		// 					},
		// 				},
		// 				Statuses: []*armkubernetesconfiguration.ExtensionStatus{
		// 				},
		// 				Version: to.Ptr("0.1.4"),
		// 			},
		// 			SystemData: &armkubernetesconfiguration.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T05:10:57.027Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-08T05:10:57.027Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armkubernetesconfiguration.CreatedByTypeApplication),
		// 			},
		// 	}},
		// }
	}
}

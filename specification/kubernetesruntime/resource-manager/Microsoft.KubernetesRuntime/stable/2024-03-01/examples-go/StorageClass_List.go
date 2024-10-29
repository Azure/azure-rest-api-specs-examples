package armcontainerorchestratorruntime_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerorchestratorruntime/armcontainerorchestratorruntime"
)

// Generated from example definition: 2024-03-01/StorageClass_List.json
func ExampleStorageClassClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerorchestratorruntime.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageClassClient().NewListPager("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1", nil)
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
		// page = armcontainerorchestratorruntime.StorageClassClientListResponse{
		// 	StorageClassResourceListResult: armcontainerorchestratorruntime.StorageClassResourceListResult{
		// 		Value: []*armcontainerorchestratorruntime.StorageClassResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1/providers/Microsoft.KubernetesRuntime/storageclasses/testrwx"),
		// 				Name: to.Ptr("testrwx"),
		// 				Type: to.Ptr("microsoft.kubernetesruntime/storageclass"),
		// 				Properties: &armcontainerorchestratorruntime.StorageClassProperties{
		// 					ProvisioningState: to.Ptr(armcontainerorchestratorruntime.ProvisioningStateSucceeded),
		// 					Performance: to.Ptr(armcontainerorchestratorruntime.PerformanceTierBasic),
		// 					TypeProperties: &armcontainerorchestratorruntime.RwxStorageClassTypeProperties{
		// 						Type: to.Ptr("RWX"),
		// 						BackingStorageClassName: to.Ptr("default"),
		// 					},
		// 					AccessModes: []*armcontainerorchestratorruntime.AccessMode{
		// 						to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteOnce),
		// 						to.Ptr(armcontainerorchestratorruntime.AccessModeReadWriteMany),
		// 					},
		// 					AllowVolumeExpansion: to.Ptr(armcontainerorchestratorruntime.VolumeExpansionAllow),
		// 					Provisioner: to.Ptr("rwx.csi.microsoft.com"),
		// 					VolumeBindingMode: to.Ptr(armcontainerorchestratorruntime.VolumeBindingModeImmediate),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

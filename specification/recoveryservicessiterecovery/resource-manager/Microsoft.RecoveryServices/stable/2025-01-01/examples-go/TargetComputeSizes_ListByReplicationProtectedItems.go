package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/TargetComputeSizes_ListByReplicationProtectedItems.json
func ExampleTargetComputeSizesClient_NewListByReplicationProtectedItemsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTargetComputeSizesClient().NewListByReplicationProtectedItemsPager("avraiMgDiskVaultRG", "avraiMgDiskVault", "asr-a2a-default-centraluseuap", "asr-a2a-default-centraluseuap-container", "468c912d-b1ab-4ea2-97eb-4b5095155db2", nil)
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
		// page.TargetComputeSizeCollection = armrecoveryservicessiterecovery.TargetComputeSizeCollection{
		// 	Value: []*armrecoveryservicessiterecovery.TargetComputeSize{
		// 		{
		// 			Name: to.Ptr("Basic_A0"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/targetComputeSizes"),
		// 			ID: to.Ptr("/Subscriptions/6808dbbc-98c7-431f-a1b1-9580902423b7/resourceGroups/avraiMgDiskVaultRG/providers/Microsoft.RecoveryServices/vaults/avraiMgDiskVault/replicationFabrics/asr-a2a-default-centraluseuap/replicationProtectionContainers/asr-a2a-default-centraluseuap-container/replicationProtectedItems/468c912d-b1ab-4ea2-97eb-4b5095155db2/targetComputeSizes/Basic_A0"),
		// 			Properties: &armrecoveryservicessiterecovery.TargetComputeSizeProperties{
		// 				Name: to.Ptr("Basic_A0"),
		// 				CPUCoresCount: to.Ptr[int32](1),
		// 				FriendlyName: to.Ptr("Basic_A0"),
		// 				HighIopsSupported: to.Ptr("NotSupported"),
		// 				MaxDataDiskCount: to.Ptr[int32](1),
		// 				MaxNicsCount: to.Ptr[int32](2),
		// 				MemoryInGB: to.Ptr[float64](0.75),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Standard_A0"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems/targetComputeSizes"),
		// 			ID: to.Ptr("/Subscriptions/6808dbbc-98c7-431f-a1b1-9580902423b7/resourceGroups/avraiMgDiskVaultRG/providers/Microsoft.RecoveryServices/vaults/avraiMgDiskVault/replicationFabrics/asr-a2a-default-centraluseuap/replicationProtectionContainers/asr-a2a-default-centraluseuap-container/replicationProtectedItems/468c912d-b1ab-4ea2-97eb-4b5095155db2/targetComputeSizes/Standard_A0"),
		// 			Properties: &armrecoveryservicessiterecovery.TargetComputeSizeProperties{
		// 				Name: to.Ptr("Standard_A0"),
		// 				CPUCoresCount: to.Ptr[int32](1),
		// 				FriendlyName: to.Ptr("Standard_A0"),
		// 				HighIopsSupported: to.Ptr("NotSupported"),
		// 				HyperVGenerations: []*string{
		// 					to.Ptr("V1")},
		// 					MaxDataDiskCount: to.Ptr[int32](1),
		// 					MaxNicsCount: to.Ptr[int32](2),
		// 					MemoryInGB: to.Ptr[float64](0.75),
		// 					VCPUsAvailable: to.Ptr[int32](1),
		// 				},
		// 		}},
		// 	}
	}
}

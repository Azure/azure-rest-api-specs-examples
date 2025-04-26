package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationRecoveryServicesProviders_ListByReplicationFabrics.json
func ExampleReplicationRecoveryServicesProvidersClient_NewListByReplicationFabricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationRecoveryServicesProvidersClient().NewListByReplicationFabricsPager("resourceGroupPS1", "vault1", "cloud1", nil)
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
		// page.RecoveryServicesProviderCollection = armrecoveryservicessiterecovery.RecoveryServicesProviderCollection{
		// 	Value: []*armrecoveryservicessiterecovery.RecoveryServicesProvider{
		// 		{
		// 			Name: to.Ptr("241641e6-ee7b-4ee4-8141-821fadda43fa"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationRecoveryServicesProviders"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationRecoveryServicesProviders/241641e6-ee7b-4ee4-8141-821fadda43fa"),
		// 			Properties: &armrecoveryservicessiterecovery.RecoveryServicesProviderProperties{
		// 				AllowedScenarios: []*string{
		// 					to.Ptr("Refresh")},
		// 					ConnectionStatus: to.Ptr("Connected"),
		// 					FabricFriendlyName: to.Ptr("cloud1"),
		// 					FabricType: to.Ptr("HyperVSite"),
		// 					FriendlyName: to.Ptr("CP-B3L40406-12.ntdev.corp.microsoft.com"),
		// 					LastHeartBeat: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-27T09:06:38.272Z"); return t}()),
		// 					ProtectedItemCount: to.Ptr[int32](2),
		// 					ProviderVersion: to.Ptr("5.1.2250.0"),
		// 					ProviderVersionState: to.Ptr("Latest"),
		// 					ServerVersion: to.Ptr("3.2.7510.0"),
		// 				},
		// 		}},
		// 	}
	}
}

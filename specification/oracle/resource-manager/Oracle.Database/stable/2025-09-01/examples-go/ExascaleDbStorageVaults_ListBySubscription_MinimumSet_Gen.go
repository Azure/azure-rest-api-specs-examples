package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/ExascaleDbStorageVaults_ListBySubscription_MinimumSet_Gen.json
func ExampleExascaleDbStorageVaultsClient_NewListBySubscriptionPager_exascaleDbStorageVaultsListBySubscriptionMaximumSetGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExascaleDbStorageVaultsClient().NewListBySubscriptionPager(nil)
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
		// page = armoracledatabase.ExascaleDbStorageVaultsClientListBySubscriptionResponse{
		// 	ExascaleDbStorageVaultListResult: armoracledatabase.ExascaleDbStorageVaultListResult{
		// 		Value: []*armoracledatabase.ExascaleDbStorageVault{
		// 			{
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/exascaleDbStorageVaults/storagevault1"),
		// 			},
		// 		},
		// 	},
		// }
	}
}

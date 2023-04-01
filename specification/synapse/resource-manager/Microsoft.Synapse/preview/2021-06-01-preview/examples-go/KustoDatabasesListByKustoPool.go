package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoDatabasesListByKustoPool.json
func ExampleKustoPoolDatabasesClient_NewListByKustoPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKustoPoolDatabasesClient().NewListByKustoPoolPager("kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", nil)
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
		// page.DatabaseListResult = armsynapse.DatabaseListResult{
		// 	Value: []armsynapse.DatabaseClassification{
		// 		&armsynapse.ReadWriteDatabase{
		// 			Name: to.Ptr("KustoClusterRPTest4/KustoDatabase8"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/kustopools/Databases"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustopools/KustoClusterRPTest4/Databases/KustoDatabase8"),
		// 			Kind: to.Ptr(armsynapse.KindReadWrite),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armsynapse.ReadWriteDatabaseProperties{
		// 				ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
		// 				SoftDeletePeriod: to.Ptr("P1D"),
		// 			},
		// 		},
		// 		&armsynapse.ReadWriteDatabase{
		// 			Name: to.Ptr("KustoClusterRPTest4/KustoDatabase9"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/kustopools/Databases"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustopools/KustoClusterRPTest4/Databases/KustoDatabase9"),
		// 			Kind: to.Ptr(armsynapse.KindReadWrite),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armsynapse.ReadWriteDatabaseProperties{
		// 				ProvisioningState: to.Ptr(armsynapse.ResourceProvisioningStateSucceeded),
		// 				SoftDeletePeriod: to.Ptr("P1D"),
		// 			},
		// 	}},
		// }
	}
}

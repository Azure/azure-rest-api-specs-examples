package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasesCreateOrUpdate.json
func ExampleDatabasesClient_BeginCreateOrUpdate_kustoReadWriteDatabaseCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginCreateOrUpdate(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", &armkusto.ReadWriteDatabase{
		Kind:     to.Ptr(armkusto.KindReadWrite),
		Location: to.Ptr("westus"),
		Properties: &armkusto.ReadWriteDatabaseProperties{
			SoftDeletePeriod: to.Ptr("P1D"),
		},
	}, &armkusto.DatabasesClientBeginCreateOrUpdateOptions{CallerRole: to.Ptr(armkusto.CallerRoleAdmin)})
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
	// res = armkusto.DatabasesClientCreateOrUpdateResponse{
	// 	                            DatabaseClassification: &armkusto.ReadWriteDatabase{
	// 		Name: to.Ptr("kustoCluster/KustoDatabase8"),
	// 		Type: to.Ptr("Microsoft.Kusto/Clusters/Databases"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8"),
	// 		Kind: to.Ptr(armkusto.KindReadWrite),
	// 		Location: to.Ptr("westus"),
	// 	},
	// 	                        }
}

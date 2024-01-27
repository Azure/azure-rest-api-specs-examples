package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6f8faf5da91b5b9af5f3512fe609e22e99383d41/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/throughputPool/CosmosDBThroughputPoolAccountCreate.json
func ExampleThroughputPoolAccountClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewThroughputPoolAccountClient().BeginCreate(ctx, "rg1", "tp1", "db1", armcosmos.ThroughputPoolAccountResource{
		Properties: &armcosmos.ThroughputPoolAccountProperties{
			AccountLocation:           to.Ptr("West US"),
			AccountResourceIdentifier: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/providers/Microsoft.DocumentDB/resourceGroup/rg1/databaseAccounts/db1/"),
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
	// res.ThroughputPoolAccountResource = armcosmos.ThroughputPoolAccountResource{
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/throughputPools/tp1/throughputPoolAccounts/db1"),
	// 	Properties: &armcosmos.ThroughputPoolAccountProperties{
	// 		AccountInstanceID: to.Ptr("{instance-id1}"),
	// 		AccountLocation: to.Ptr("west us"),
	// 		AccountResourceIdentifier: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/db1"),
	// 		ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
	// 	},
	// }
}

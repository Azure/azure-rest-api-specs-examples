package armresourcemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Delete.json
func ExampleMoveResourcesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveResourcesClient().BeginDelete(ctx, "rg1", "movecollection1", "moveresourcename1", nil)
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
	// res.OperationStatus = armresourcemover.OperationStatus{
	// 	Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	EndTime: to.Ptr("6/17/2020 6:45:56 AM"),
	// 	ID: to.Ptr("/subscriptions/e80eb9fa-c996-4435-aa32-5af6f3d3077c/resourceGroups/RegionMoveRG-southcentralus-southeastasia/providers/Microsoft.Migrate/MoveCollections/MoveCollection-southcentralus-southeastasia/operations/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("6/17/2020 6:45:55 AM"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

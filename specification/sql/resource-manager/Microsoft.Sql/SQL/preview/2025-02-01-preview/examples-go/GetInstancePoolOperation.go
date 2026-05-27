package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/GetInstancePoolOperation.json
func ExampleInstancePoolOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInstancePoolOperationsClient().Get(ctx, "resource-group", "test-instance-pool", "c218773b-203f-4c7a-b174-6bd71fe20f72", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.InstancePoolOperationsClientGetResponse{
	// 	InstancePoolOperation: armsql.InstancePoolOperation{
	// 		Name: to.Ptr("c218773b-203f-4c7a-b174-6bd71fe20f72"),
	// 		Type: to.Ptr("Microsoft.Sql/instancePools/operations"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resource-group/providers/Microsoft.Sql/instancePools/test-instance-pool"),
	// 		Properties: &armsql.InstancePoolOperationProperties{
	// 			InstancePoolName: to.Ptr("test-instance-pool"),
	// 			IsCancellable: to.Ptr(true),
	// 			Operation: to.Ptr("UpsertInstancePoolAsync"),
	// 			OperationFriendlyName: to.Ptr("UPDATE INSTANCE POOL"),
	// 			PercentComplete: to.Ptr[int32](100),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-22T14:46:00.423Z"); return t}()),
	// 			State: to.Ptr(armsql.ManagementOperationStateSucceeded),
	// 		},
	// 	},
	// }
}

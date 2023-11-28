package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/GetTaskSubscriptionLocation_example.json
func ExampleTasksClient_GetSubscriptionLevelTask() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTasksClient().GetSubscriptionLevelTask(ctx, "westeurope", "62609ee7-d0a5-8616-9fe4-1df5cca7758d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Task = armsecurity.Task{
	// 	Name: to.Ptr("62609ee7-d0a5-8616-9fe4-1df5cca7758d"),
	// 	Type: to.Ptr("Microsoft.Security/locations/tasks"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/locations/westeurope/tasks/62609ee7-d0a5-8616-9fe4-1df5cca7758d"),
	// 	Properties: &armsecurity.TaskProperties{
	// 		CreationTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-05T10:42:03.993Z"); return t}()),
	// 		LastStateChangeTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-05T10:42:03.993Z"); return t}()),
	// 		SecurityTaskParameters: &armsecurity.TaskParameters{
	// 			AdditionalProperties: map[string]any{
	// 				"location": "uksouth",
	// 				"resourceGroup": "myRg",
	// 				"resourceId": "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/default",
	// 				"resourceName": "default",
	// 				"resourceParent": "vnet1",
	// 				"resourceType": "Subnet",
	// 				"uniqueKey": "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/default",
	// 			},
	// 			Name: to.Ptr("NetworkSecurityGroupMissingOnSubnet"),
	// 		},
	// 		State: to.Ptr("Active"),
	// 		SubState: to.Ptr("NA"),
	// 	},
	// }
}

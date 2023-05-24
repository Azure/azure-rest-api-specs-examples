package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a7af0df86022e5e6cc6e8f40ca1981c4557a4bc/specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ConnectedEnvironmentsStorages_List.json
func ExampleConnectedEnvironmentsStoragesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsStoragesClient().List(ctx, "examplerg", "managedEnv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedEnvironmentStoragesCollection = armappcontainers.ConnectedEnvironmentStoragesCollection{
	// 	Value: []*armappcontainers.ConnectedEnvironmentStorage{
	// 		{
	// 			Name: to.Ptr("jlaw-demo1"),
	// 			Type: to.Ptr("Microsoft.App/connectedEnvironments/storages"),
	// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/managedEnv/storages/jlaw-demo1"),
	// 			Properties: &armappcontainers.ConnectedEnvironmentStorageProperties{
	// 				AzureFile: &armappcontainers.AzureFileProperties{
	// 					AccessMode: to.Ptr(armappcontainers.AccessModeReadOnly),
	// 					AccountName: to.Ptr("account1"),
	// 					ShareName: to.Ptr("share1"),
	// 				},
	// 			},
	// 	}},
	// }
}

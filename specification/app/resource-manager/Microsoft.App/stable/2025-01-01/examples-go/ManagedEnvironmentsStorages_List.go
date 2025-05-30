package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ManagedEnvironmentsStorages_List.json
func ExampleManagedEnvironmentsStoragesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedEnvironmentsStoragesClient().List(ctx, "examplerg", "managedEnv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedEnvironmentStoragesCollection = armappcontainers.ManagedEnvironmentStoragesCollection{
	// 	Value: []*armappcontainers.ManagedEnvironmentStorage{
	// 		{
	// 			Name: to.Ptr("jlaw-demo1"),
	// 			Type: to.Ptr("Microsoft.App/managedEnvironments/storages"),
	// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/managedEnv/storages/jlaw-demo1"),
	// 			Properties: &armappcontainers.ManagedEnvironmentStorageProperties{
	// 				AzureFile: &armappcontainers.AzureFileProperties{
	// 					AccessMode: to.Ptr(armappcontainers.AccessModeReadOnly),
	// 					AccountName: to.Ptr("account1"),
	// 					ShareName: to.Ptr("share1"),
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("jlaw-demo2"),
	// 			Type: to.Ptr("Microsoft.App/managedEnvironments/storages"),
	// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/managedEnv/storages/jlaw-demo2"),
	// 			Properties: &armappcontainers.ManagedEnvironmentStorageProperties{
	// 				NfsAzureFile: &armappcontainers.NfsAzureFileProperties{
	// 					AccessMode: to.Ptr(armappcontainers.AccessModeReadOnly),
	// 					Server: to.Ptr("server1"),
	// 					ShareName: to.Ptr("share1"),
	// 				},
	// 			},
	// 	}},
	// }
}

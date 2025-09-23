package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/ManagedEnvironmentsStorages_List.json
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
	// 					AccountKeyVaultProperties: &armappcontainers.SecretKeyVaultProperties{
	// 						Identity: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity"),
	// 						KeyVaultURL: to.Ptr("https://myvault.vault.azure.net/secrets/mysecret"),
	// 					},
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

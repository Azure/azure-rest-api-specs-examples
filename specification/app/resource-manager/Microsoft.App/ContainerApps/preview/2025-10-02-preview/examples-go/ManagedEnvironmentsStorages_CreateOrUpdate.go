package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/ManagedEnvironmentsStorages_CreateOrUpdate.json
func ExampleManagedEnvironmentsStoragesClient_CreateOrUpdate_createOrUpdateEnvironmentsStorage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("8efdecc5-919e-44eb-b179-915dca89ebf9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedEnvironmentsStoragesClient().CreateOrUpdate(ctx, "examplerg", "managedEnv", "jlaw-demo1", armappcontainers.ManagedEnvironmentStorage{
		Properties: &armappcontainers.ManagedEnvironmentStorageProperties{
			AzureFile: &armappcontainers.AzureFileProperties{
				AccessMode:  to.Ptr(armappcontainers.AccessModeReadOnly),
				AccountKey:  to.Ptr("key"),
				AccountName: to.Ptr("account1"),
				ShareName:   to.Ptr("share1"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.ManagedEnvironmentsStoragesClientCreateOrUpdateResponse{
	// 	ManagedEnvironmentStorage: armappcontainers.ManagedEnvironmentStorage{
	// 	},
	// }
}

package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironmentsStorages_CreateOrUpdate.json
func ExampleConnectedEnvironmentsStoragesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappcontainers.NewConnectedEnvironmentsStoragesClient("8efdecc5-919e-44eb-b179-915dca89ebf9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "examplerg", "env", "jlaw-demo1", armappcontainers.ConnectedEnvironmentStorage{
		Properties: &armappcontainers.ConnectedEnvironmentStorageProperties{
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
	// TODO: use response item
	_ = res
}

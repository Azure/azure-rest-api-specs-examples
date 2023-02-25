package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3d7a3848106b831a4a7f46976fe38aa605c4f44d/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/Projects_CreateOrUpdate.json
func ExampleProjectsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewProjectsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "examples-rg", "examples-storageMoverName", "examples-projectName", armstoragemover.Project{
		Properties: &armstoragemover.ProjectProperties{
			Description: to.Ptr("Example Project Description"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Project = armstoragemover.Project{
	// 	Name: to.Ptr("examples-projectName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/projects"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-projectName"),
	// 	Properties: &armstoragemover.ProjectProperties{
	// 		Description: to.Ptr("Example Project Description"),
	// 	},
	// }
}

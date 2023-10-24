package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/Projects_CreateOrUpdate.json
func ExampleProjectsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().CreateOrUpdate(ctx, "examples-rg", "examples-storageMoverName", "examples-projectName", armstoragemover.Project{
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
	// 	ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projects/examples-projectName"),
	// 	Properties: &armstoragemover.ProjectProperties{
	// 		Description: to.Ptr("Example Project Description"),
	// 	},
	// }
}

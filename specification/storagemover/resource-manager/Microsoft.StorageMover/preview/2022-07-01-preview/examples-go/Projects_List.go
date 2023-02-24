package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3d7a3848106b831a4a7f46976fe38aa605c4f44d/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/Projects_List.json
func ExampleProjectsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewProjectsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("examples-rg", "examples-storageMoverName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ProjectList = armstoragemover.ProjectList{
		// 	Value: []*armstoragemover.Project{
		// 		{
		// 			Name: to.Ptr("examples-projectName1"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/projects"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projects/examples-projectName1"),
		// 			Properties: &armstoragemover.ProjectProperties{
		// 				Description: to.Ptr("Example Project 1 Description"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-projectName2"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/projects"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projects/examples-projectName2"),
		// 			Properties: &armstoragemover.ProjectProperties{
		// 				Description: to.Ptr("Example Project 2 Description"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-projectName3"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/projects"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projects/examples-projectName2"),
		// 			Properties: &armstoragemover.ProjectProperties{
		// 				Description: to.Ptr("Example Project 3 Description"),
		// 			},
		// 	}},
		// }
	}
}

package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/JobDefinitions_List.json
func ExampleJobDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobDefinitionsClient().NewListPager("examples-rg", "examples-storageMoverName", "examples-projectName", nil)
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
		// page.JobDefinitionList = armstoragemover.JobDefinitionList{
		// 	Value: []*armstoragemover.JobDefinition{
		// 		{
		// 			Name: to.Ptr("examples-jobDefinitionName1"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/jobDefinitions"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/jobDefinitions/examples-jobDefinitionName1"),
		// 			Properties: &armstoragemover.JobDefinitionProperties{
		// 				Description: to.Ptr("Example Job Definition 1 Description"),
		// 				AgentName: to.Ptr("migration-agent"),
		// 				AgentResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
		// 				CopyMode: to.Ptr(armstoragemover.CopyModeAdditive),
		// 				SourceName: to.Ptr("examples-sourceEndpointName1"),
		// 				SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-sourceEndpointName1"),
		// 				SourceSubpath: to.Ptr("/"),
		// 				TargetName: to.Ptr("examples-targetEndpointName1"),
		// 				TargetResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-targetEndpointName1"),
		// 				TargetSubpath: to.Ptr("/"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-jobDefinitionName2"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/jobDefinitions"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/jobDefinitions/examples-jobDefinitionName2"),
		// 			SystemData: &armstoragemover.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T01:01:01.1075056Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:01:01.1075056Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
		// 			},
		// 			Properties: &armstoragemover.JobDefinitionProperties{
		// 				Description: to.Ptr("Example Job Definition 2 Description"),
		// 				AgentName: to.Ptr("migration-agent"),
		// 				AgentResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
		// 				CopyMode: to.Ptr(armstoragemover.CopyModeAdditive),
		// 				SourceName: to.Ptr("examples-sourceEndpointName2"),
		// 				SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-sourceEndpointName2"),
		// 				SourceSubpath: to.Ptr("/"),
		// 				TargetName: to.Ptr("examples-targetEndpointName2"),
		// 				TargetResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-targetEndpointName2"),
		// 				TargetSubpath: to.Ptr("/"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-jobDefinitionName3"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/jobDefinitions"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/jobDefinitions/examples-jobDefinitionName3"),
		// 			Properties: &armstoragemover.JobDefinitionProperties{
		// 				Description: to.Ptr("Example Job Definition 3 Description"),
		// 				AgentName: to.Ptr("migration-agent"),
		// 				AgentResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
		// 				CopyMode: to.Ptr(armstoragemover.CopyModeMirror),
		// 				SourceName: to.Ptr("examples-sourceEndpointName3"),
		// 				SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-sourceEndpointName3"),
		// 				SourceSubpath: to.Ptr("/"),
		// 				TargetName: to.Ptr("examples-targetEndpointName3"),
		// 				TargetResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-targetEndpointName3"),
		// 				TargetSubpath: to.Ptr("/"),
		// 			},
		// 	}},
		// }
	}
}

package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: 2025-07-01/JobDefinitions_CreateOrUpdate_CloudToCloud.json
func ExampleJobDefinitionsClient_CreateOrUpdate_jobDefinitionsCreateOrUpdateCloudToCloud() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("60bcfc77-6589-4da2-b7fd-f9ec9322cf95", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobDefinitionsClient().CreateOrUpdate(ctx, "examples-rg", "examples-storageMoverName", "examples-projectName", "examples-jobDefinitionName", armstoragemover.JobDefinition{
		Properties: &armstoragemover.JobDefinitionProperties{
			Description:   to.Ptr("Example Job Definition Description"),
			CopyMode:      to.Ptr(armstoragemover.CopyModeAdditive),
			JobType:       to.Ptr(armstoragemover.JobTypeCloudToCloud),
			SourceName:    to.Ptr("examples-sourceEndpointName"),
			SourceSubpath: to.Ptr("/"),
			TargetName:    to.Ptr("examples-targetEndpointName"),
			TargetSubpath: to.Ptr("/"),
			AgentName:     to.Ptr("dummy-agent"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragemover.JobDefinitionsClientCreateOrUpdateResponse{
	// 	JobDefinition: &armstoragemover.JobDefinition{
	// 		ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projects/examples-projectName/jobDefinitions/examples-jobDefinitionName"),
	// 		Name: to.Ptr("examples-jobDefinitionName"),
	// 		Type: to.Ptr("Microsoft.StorageMover/storageMovers/projectName/jobDefinitionName"),
	// 		Properties: &armstoragemover.JobDefinitionProperties{
	// 			Description: to.Ptr("Example Job Definition Description"),
	// 			CopyMode: to.Ptr(armstoragemover.CopyModeAdditive),
	// 			JobType: to.Ptr(armstoragemover.JobTypeCloudToCloud),
	// 			SourceName: to.Ptr("examples-sourceEndpointName"),
	// 			SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-sourceEndpointName"),
	// 			SourceSubpath: to.Ptr("/"),
	// 			TargetName: to.Ptr("examples-targetEndpointName"),
	// 			TargetResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-targetEndpointName"),
	// 			TargetSubpath: to.Ptr("/"),
	// 			AgentName: to.Ptr("migration-agent"),
	// 			AgentResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
	// 		},
	// 	},
	// }
}

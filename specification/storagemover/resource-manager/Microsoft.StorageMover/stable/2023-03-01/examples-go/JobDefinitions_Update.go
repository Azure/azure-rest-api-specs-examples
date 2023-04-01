package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/241e964afe675a7be98aa6a2e171a3c5f830816c/specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/JobDefinitions_Update.json
func ExampleJobDefinitionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobDefinitionsClient().Update(ctx, "examples-rg", "examples-storageMoverName", "examples-projectName", "examples-jobDefinitionName", armstoragemover.JobDefinitionUpdateParameters{
		Properties: &armstoragemover.JobDefinitionUpdateProperties{
			Description: to.Ptr("Updated Job Definition Description"),
			AgentName:   to.Ptr("updatedAgentName"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobDefinition = armstoragemover.JobDefinition{
	// 	Name: to.Ptr("examples-jobDefinitionName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/projectName/jobDefinitionName"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projects/examples-projectName/jobDefinitions/examples-jobDefinitionName"),
	// 	Properties: &armstoragemover.JobDefinitionProperties{
	// 		Description: to.Ptr("Updated Job Definition Description"),
	// 		AgentName: to.Ptr("updatedAgentName"),
	// 		AgentResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
	// 		CopyMode: to.Ptr(armstoragemover.CopyModeAdditive),
	// 		SourceName: to.Ptr("updatedSource"),
	// 		SourceResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/updatedSource"),
	// 		SourceSubpath: to.Ptr("/"),
	// 		TargetName: to.Ptr("updatedTarget"),
	// 		TargetResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/updatedTarget"),
	// 		TargetSubpath: to.Ptr("/"),
	// 	},
	// }
}

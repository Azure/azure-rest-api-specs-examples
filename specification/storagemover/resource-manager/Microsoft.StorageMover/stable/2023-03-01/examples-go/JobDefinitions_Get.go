package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/241e964afe675a7be98aa6a2e171a3c5f830816c/specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/JobDefinitions_Get.json
func ExampleJobDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobDefinitionsClient().Get(ctx, "examples-rg", "examples-storageMoverName", "examples-projectName", "examples-jobDefinitionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobDefinition = armstoragemover.JobDefinition{
	// 	Name: to.Ptr("examples-jobDefinitionName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/jobDefinitions"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/jobDefinitions/examples-jobDefinitionName"),
	// 	Properties: &armstoragemover.JobDefinitionProperties{
	// 		Description: to.Ptr("Example Job Definition Description"),
	// 		AgentName: to.Ptr("migration-agent"),
	// 		AgentResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
	// 		CopyMode: to.Ptr(armstoragemover.CopyModeAdditive),
	// 		SourceName: to.Ptr("examples-sourceEndpointName"),
	// 		SourceResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-sourceEndpointName"),
	// 		SourceSubpath: to.Ptr("/"),
	// 		TargetName: to.Ptr("examples-targetEndpointName"),
	// 		TargetResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-targetEndpointName"),
	// 		TargetSubpath: to.Ptr("/"),
	// 	},
	// }
}

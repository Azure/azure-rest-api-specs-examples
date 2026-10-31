package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: 2026-05-01/JobDefinitions_Get_CrossTenant_Mirror.json
func ExampleJobDefinitionsClient_Get_jobDefinitionsGetCrossTenantMirror() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("0a2b3c4d-5e6f-7081-92a3-b4c5d6e7f809", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobDefinitionsClient().Get(ctx, "partner-rg", "partner-storageMover", "partner-projectName", "examples-jobDefinitionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragemover.JobDefinitionsClientGetResponse{
	// 	JobDefinition: armstoragemover.JobDefinition{
	// 		Name: to.Ptr("examples-jobDefinitionName"),
	// 		Type: to.Ptr("Microsoft.StorageMover/storageMovers/jobDefinitions"),
	// 		ID: to.Ptr("/subscriptions/0a2b3c4d-5e6f-7081-92a3-b4c5d6e7f809/resourceGroups/partner-rg/providers/Microsoft.StorageMover/storageMovers/partner-storageMover/projects/partner-projectName/jobDefinitions/examples-jobDefinitionName"),
	// 		Properties: &armstoragemover.JobDefinitionProperties{
	// 			Description: to.Ptr("Cross-tenant Blob-to-Blob copy job (guest-tenant mirror). This read-only mirror is created by the resource provider in the target (guest) tenant alongside the local target endpoint; the source endpoint lives in the partner (host) tenant."),
	// 			CopyMode: to.Ptr(armstoragemover.CopyModeMirror),
	// 			JobType: to.Ptr(armstoragemover.JobTypeCloudToCloud),
	// 			SourceName: to.Ptr("examples-sourceEndpointName"),
	// 			SourceResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-sourceEndpointName"),
	// 			SourceSubpath: to.Ptr("/"),
	// 			TargetName: to.Ptr("partner-targetEndpoint"),
	// 			TargetResourceID: to.Ptr("/subscriptions/0a2b3c4d-5e6f-7081-92a3-b4c5d6e7f809/resourceGroups/partner-rg/providers/Microsoft.StorageMover/storageMovers/partner-storageMover/endpoints/partner-targetEndpoint"),
	// 			TargetSubpath: to.Ptr("/"),
	// 			IsCrossTenantJob: to.Ptr(true),
	// 			CrossTenantEndpointTenantID: to.Ptr("99999999-8888-7777-6666-555555555555"),
	// 			CrossTenantEndpointResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-sourceEndpointName"),
	// 			ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

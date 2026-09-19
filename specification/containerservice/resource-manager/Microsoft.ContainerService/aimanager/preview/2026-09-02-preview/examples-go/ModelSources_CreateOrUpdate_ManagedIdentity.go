package armcontainerserviceaimanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerserviceaimanager/armcontainerserviceaimanager"
)

// Generated from example definition: 2026-09-02-preview/ModelSources_CreateOrUpdate_ManagedIdentity.json
func ExampleModelSourcesClient_BeginCreateOrUpdate_modelSourcesCreateOrUpdateManagedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerserviceaimanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewModelSourcesClient().BeginCreateOrUpdate(ctx, "rgaimanagers", "aimanager1", "foundry", armcontainerserviceaimanager.ModelSource{
		Properties: &armcontainerserviceaimanager.ModelSourceProperties{
			SourceType:  to.Ptr(armcontainerserviceaimanager.ModelSourceTypeMicrosoftFoundry),
			Description: to.Ptr("Foundry model source"),
			MicrosoftFoundry: &armcontainerserviceaimanager.MicrosoftFoundrySource{
				ProjectResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.CognitiveServices/accounts/test-account/projects/test-model-project"),
			},
			Credential: &armcontainerserviceaimanager.CredentialValue{
				ManagedIdentity: &armcontainerserviceaimanager.ManagedIdentityCredential{
					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/mytestidentity"),
				},
			},
		},
	}, &armcontainerserviceaimanager.ModelSourcesClientBeginCreateOrUpdateOptions{
		IfMatch:     to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
		IfNoneMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerserviceaimanager.ModelSourcesClientCreateOrUpdateResponse{
	// 	ModelSource: armcontainerserviceaimanager.ModelSource{
	// 		Properties: &armcontainerserviceaimanager.ModelSourceProperties{
	// 			ProvisioningState: to.Ptr(armcontainerserviceaimanager.ResourceProvisioningStateSucceeded),
	// 			SourceType: to.Ptr(armcontainerserviceaimanager.ModelSourceTypeMicrosoftFoundry),
	// 			Description: to.Ptr("Foundry model source"),
	// 			MicrosoftFoundry: &armcontainerserviceaimanager.MicrosoftFoundrySource{
	// 				ProjectResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.CognitiveServices/accounts/test-account/projects/test-model-project"),
	// 			},
	// 			Credential: &armcontainerserviceaimanager.CredentialValue{
	// 				ManagedIdentity: &armcontainerserviceaimanager.ManagedIdentityCredential{
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/mytestidentity"),
	// 				},
	// 			},
	// 		},
	// 		ETag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgaimanagers/providers/Microsoft.ContainerService/aiManagers/aimanager1/modelSources/foundry"),
	// 		Name: to.Ptr("foundry"),
	// 		Type: to.Ptr("Microsoft.ContainerService/aiManagers/modelSources"),
	// 		SystemData: &armcontainerserviceaimanager.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}

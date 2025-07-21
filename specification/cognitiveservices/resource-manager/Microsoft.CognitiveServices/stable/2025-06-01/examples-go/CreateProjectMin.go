package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/CreateProjectMin.json
func ExampleProjectsClient_BeginCreate_createProjectMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProjectsClient().BeginCreate(ctx, "myResourceGroup", "testCreate1", "testProject1", armcognitiveservices.Project{
		Identity: &armcognitiveservices.Identity{
			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
		},
		Location:   to.Ptr("West US"),
		Properties: &armcognitiveservices.ProjectProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Project = armcognitiveservices.Project{
	// 	Name: to.Ptr("testProject1"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/testCreate1/projects/testProject1"),
	// 	Etag: to.Ptr("W/\"datetime'2017-04-10T08%3A00%3A05.445595Z'\""),
	// 	Identity: &armcognitiveservices.Identity{
	// 		Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("b5cf119e-a5c2-42c7-802f-592e0efb169f"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcognitiveservices.ProjectProperties{
	// 		Endpoints: map[string]*string{
	// 			"OpenAI Dall-E API": to.Ptr("https://sub-donmain-name.openai.azure.com/"),
	// 			"OpenAI Language Model Instance API": to.Ptr("https://sub-donmain-name.openai.azure.com/"),
	// 			"OpenAI Sora API": to.Ptr("https://sub-donmain-name.openai.azure.com/"),
	// 		},
	// 		IsDefault: to.Ptr(true),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 	},
	// }
}

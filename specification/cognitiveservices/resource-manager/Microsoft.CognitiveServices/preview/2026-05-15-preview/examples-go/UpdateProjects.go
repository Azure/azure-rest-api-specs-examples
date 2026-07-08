package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/UpdateProjects.json
func ExampleProjectsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProjectsClient().BeginUpdate(ctx, "bvttest", "bingSearch", "projectName", armcognitiveservices.Project{
		Location: to.Ptr("global"),
		Properties: &armcognitiveservices.ProjectProperties{
			Description: to.Ptr("new description."),
		},
	}, nil)
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
	// res = armcognitiveservices.ProjectsClientUpdateResponse{
	// 	Project: armcognitiveservices.Project{
	// 		Name: to.Ptr("projectName"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects"),
	// 		Etag: to.Ptr("W/\"datetime'2017-04-10T07%3A46%3A21.5618831Z'\""),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/bingSearch/projects/projectName"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armcognitiveservices.ProjectProperties{
	// 			Description: to.Ptr("new description."),
	// 			DisplayName: to.Ptr("myProject"),
	// 			Endpoints: map[string]*string{
	// 				"OpenAI Dall-E API": to.Ptr("https://customSubDomainName.openai.azure.com/"),
	// 				"OpenAI Language Model Instance API": to.Ptr("https://customSubDomainName.openai.azure.com/"),
	// 			},
	// 			IsDefault: to.Ptr(false),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

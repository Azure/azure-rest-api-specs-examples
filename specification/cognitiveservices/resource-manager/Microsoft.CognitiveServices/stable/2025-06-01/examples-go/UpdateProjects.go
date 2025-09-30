package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1004eed4202d64b48157c084fe2830760f8190f4/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/UpdateProjects.json
func ExampleProjectsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
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
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Project = armcognitiveservices.Project{
	// 	Name: to.Ptr("projectName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/bvttest/providers/Microsoft.CognitiveServices/accounts/bingSearch/projects/projectName"),
	// 	Etag: to.Ptr("W/\"datetime'2017-04-10T07%3A46%3A21.5618831Z'\""),
	// 	Location: to.Ptr("global"),
	// 	Properties: &armcognitiveservices.ProjectProperties{
	// 		Description: to.Ptr("new description."),
	// 		DisplayName: to.Ptr("myProject"),
	// 		Endpoints: map[string]*string{
	// 			"OpenAI Dall-E API": to.Ptr("https://customSubDomainName.openai.azure.com/"),
	// 			"OpenAI Language Model Instance API": to.Ptr("https://customSubDomainName.openai.azure.com/"),
	// 		},
	// 		IsDefault: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 	},
	// }
}

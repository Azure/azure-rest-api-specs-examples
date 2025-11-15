package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/GetProject.json
func ExampleProjectsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().Get(ctx, "myResourceGroup", "myAccount", "myProject", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Project = armcognitiveservices.Project{
	// 	Name: to.Ptr("myProject"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myAccount/projects/myProject"),
	// 	Etag: to.Ptr("W/\"datetime'2017-04-10T04%3A42%3A19.7067387Z'\""),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcognitiveservices.ProjectProperties{
	// 		Description: to.Ptr("This is my project"),
	// 		DisplayName: to.Ptr("myProject"),
	// 		Endpoints: map[string]*string{
	// 			"OpenAI Dall-E API": to.Ptr("https://customSubDomainName.openai.azure.com/"),
	// 			"OpenAI Language Model Instance API": to.Ptr("https://customSubDomainName.openai.azure.com/"),
	// 		},
	// 		IsDefault: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
	// 	},
	// 	Tags: map[string]*string{
	// 		"ExpiredDate": to.Ptr("2017/09/01"),
	// 		"Owner": to.Ptr("felixwa"),
	// 	},
	// }
}

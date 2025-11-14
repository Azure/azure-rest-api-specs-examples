package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/ListProjects.json
func ExampleProjectsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectsClient().NewListPager("myResourceGroup", "myAccount", nil)
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
		// page.ProjectListResult = armcognitiveservices.ProjectListResult{
		// 	Value: []*armcognitiveservices.Project{
		// 		{
		// 			Name: to.Ptr("myProject"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myAccount/projects/myProject"),
		// 			Etag: to.Ptr("W/\"datetime'2017-04-10T04%3A42%3A19.7067387Z'\""),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.ProjectProperties{
		// 				Description: to.Ptr("This is my project 1"),
		// 				DisplayName: to.Ptr("myProject1"),
		// 				Endpoints: map[string]*string{
		// 					"OpenAI Dall-E API": to.Ptr("https://customSubDomainName.openai.azure.com/"),
		// 					"OpenAI Language Model Instance API": to.Ptr("https://customSubDomainName.openai.azure.com/"),
		// 				},
		// 				IsDefault: to.Ptr(true),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			Tags: map[string]*string{
		// 				"ExpiredDate": to.Ptr("2017/09/01"),
		// 				"Owner": to.Ptr("felixwa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myProject-2"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.CognitiveServices/accounts/myAccount/projects/myProject-2"),
		// 			Etag: to.Ptr("W/\"datetime'2017-04-07T04%3A32%3A38.9187216Z'\""),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcognitiveservices.ProjectProperties{
		// 				Description: to.Ptr("This is my project 2"),
		// 				DisplayName: to.Ptr("myProject2"),
		// 				Endpoints: map[string]*string{
		// 					"OpenAI Dall-E API": to.Ptr("https://customSubDomainName2.openai.azure.com/"),
		// 					"OpenAI Language Model Instance API": to.Ptr("https://customSubDomainName2.openai.azure.com/"),
		// 				},
		// 				IsDefault: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.ProvisioningStateSucceeded),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 	}},
		// }
	}
}

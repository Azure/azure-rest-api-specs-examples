package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiTagDescription.json
func ExampleAPITagDescriptionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPITagDescriptionClient().CreateOrUpdate(ctx, "rg1", "apimService1", "5931a75ae4bbd512a88c680b", "tagId1", armapimanagement.TagDescriptionCreateParameters{
		Properties: &armapimanagement.TagDescriptionBaseProperties{
			Description:             to.Ptr("Some description that will be displayed for operation's tag if the tag is assigned to operation of the API"),
			ExternalDocsDescription: to.Ptr("Description of the external docs resource"),
			ExternalDocsURL:         to.Ptr("http://some.url/additionaldoc"),
		},
	}, &armapimanagement.APITagDescriptionClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TagDescriptionContract = armapimanagement.TagDescriptionContract{
	// 	Name: to.Ptr("tagId1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/tagDescriptions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/5931a75ae4bbd512a88c680b/tagDescriptions/tagId1"),
	// 	Properties: &armapimanagement.TagDescriptionContractProperties{
	// 		Description: to.Ptr("Some description that will be displayed for operation's tag if the tag is assigned to operation of the API"),
	// 		ExternalDocsDescription: to.Ptr("some additional info"),
	// 		ExternalDocsURL: to.Ptr("http://some_url.com"),
	// 		DisplayName: to.Ptr("tag1"),
	// 		TagID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tags/tagId1"),
	// 	},
	// }
}

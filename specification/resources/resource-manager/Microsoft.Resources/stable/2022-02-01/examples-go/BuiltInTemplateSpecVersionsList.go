package armtemplatespecs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/BuiltInTemplateSpecVersionsList.json
func ExampleTemplateSpecVersionsClient_NewListBuiltInsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTemplateSpecVersionsClient().NewListBuiltInsPager("nameOfTheBuiltIn", nil)
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
		// page.TemplateSpecVersionsListResult = armtemplatespecs.TemplateSpecVersionsListResult{
		// 	Value: []*armtemplatespecs.TemplateSpecVersion{
		// 		{
		// 			Name: to.Ptr("v1.0"),
		// 			Type: to.Ptr("Microsoft.Resources/builtInTemplateSpecs/versions"),
		// 			ID: to.Ptr("/providers/Microsoft.Resources/builtInTemplateSpecs/nameOfTheBuiltIn/versions/v1.0"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-01T01:01:01.107Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-01T01:01:01.107Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armtemplatespecs.TemplateSpecVersionProperties{
		// 				Description: to.Ptr("This is version v1.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("v2.0"),
		// 			Type: to.Ptr("Microsoft.Resources/builtInTemplateSpecs/versions"),
		// 			ID: to.Ptr("/providers/Microsoft.Resources/builtInTemplateSpecs/nameOfTheBuiltIn/versions/v2.0"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-02T02:03:01.197Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-02T02:03:01.197Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armtemplatespecs.TemplateSpecVersionProperties{
		// 				Description: to.Ptr("This is another version (v2.0)"),
		// 			},
		// 	}},
		// }
	}
}

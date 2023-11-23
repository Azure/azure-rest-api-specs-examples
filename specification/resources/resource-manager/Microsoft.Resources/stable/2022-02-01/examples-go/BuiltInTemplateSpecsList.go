package armtemplatespecs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/BuiltInTemplateSpecsList.json
func ExampleClient_NewListBuiltInsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListBuiltInsPager(&armtemplatespecs.ClientListBuiltInsOptions{Expand: nil})
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
		// page.ListResult = armtemplatespecs.ListResult{
		// 	Value: []*armtemplatespecs.TemplateSpec{
		// 		{
		// 			Name: to.Ptr("nameOfBuiltIn"),
		// 			Type: to.Ptr("Microsoft.Resources/builtInTemplateSpecs"),
		// 			ID: to.Ptr("/providers/Microsoft.Resources/builtInTemplateSpecs/nameOfBuiltIn"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-01T01:01:01.107Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-02T02:03:01.197Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armtemplatespecs.TemplateSpecProperties{
		// 				Description: to.Ptr("The description of the built-in template spec"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("nameOfAnotherBuiltIn"),
		// 			Type: to.Ptr("Microsoft.Resources/builtInTemplateSpecs"),
		// 			ID: to.Ptr("/providers/Microsoft.Resources/builtInTemplateSpecs/nameOfAnotherBuiltIn"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-03T03:01:01.128Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-03T05:22:06.197Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armtemplatespecs.TemplateSpecProperties{
		// 				Description: to.Ptr("The description of another built-in template spec"),
		// 			},
		// 	}},
		// }
	}
}

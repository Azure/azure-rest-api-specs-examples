package armtemplatespecs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/TemplateSpecsListBySubscription.json
func ExampleClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListBySubscriptionPager(&armtemplatespecs.ClientListBySubscriptionOptions{Expand: nil})
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
		// 			Name: to.Ptr("simpleTemplateSpec"),
		// 			Type: to.Ptr("Microsoft.Resources/templateSpecs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/templateSpecRG/providers/Microsoft.Resources/templateSpecs/simpleTemplateSpec"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.1974346Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armtemplatespecs.TemplateSpecProperties{
		// 				Description: to.Ptr("A very simple Template Spec"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("anotherSimpleTemplateSpec"),
		// 			Type: to.Ptr("Microsoft.Resources/templateSpecs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/templateSpecRG2/providers/Microsoft.Resources/templateSpecs/anotherSimpleTemplateSpec"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T01:01:01.1075056Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T02:03:01.1974346Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armtemplatespecs.TemplateSpecProperties{
		// 				Description: to.Ptr("Another very simple Template Spec"),
		// 			},
		// 	}},
		// }
	}
}

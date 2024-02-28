package armapicenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Apis_List.json
func ExampleApisClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApisClient().NewListPager("contoso-resources", "contoso", "default", &armapicenter.ApisClientListOptions{Filter: nil})
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
		// page.APIListResult = armapicenter.APIListResult{
		// 	Value: []*armapicenter.API{
		// 		{
		// 			Name: to.Ptr("echo-api"),
		// 			Type: to.Ptr("Microsoft.ApiCenter/services/environments"),
		// 			ID: to.Ptr("/subscriptions/a200340d-6b82-494d-9dbf-687ba6e33f9e/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/apis/echo-api"),
		// 			SystemData: &armapicenter.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
		// 			},
		// 			Properties: &armapicenter.APIProperties{
		// 				Description: to.Ptr("A simple HTTP request/response service."),
		// 				CustomProperties: map[string]any{
		// 					"author": "John Doe",
		// 				},
		// 				ExternalDocumentation: []*armapicenter.ExternalDocumentation{
		// 					{
		// 						Title: to.Ptr("Onboarding docs"),
		// 						URL: to.Ptr("https://docs.contoso.com"),
		// 				}},
		// 				Kind: to.Ptr(armapicenter.APIKindRest),
		// 				License: &armapicenter.License{
		// 					URL: to.Ptr("https://contoso.com/license"),
		// 				},
		// 				LifecycleStage: to.Ptr(armapicenter.LifecycleStageDesign),
		// 				TermsOfService: &armapicenter.TermsOfService{
		// 					URL: to.Ptr("https://contoso.com/terms-of-service"),
		// 				},
		// 				Title: to.Ptr("Echo API"),
		// 			},
		// 	}},
		// }
	}
}

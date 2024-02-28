package armapicenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Apis_Get.json
func ExampleApisClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApisClient().Get(ctx, "contoso-resources", "contoso", "default", "echo-api", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.API = armapicenter.API{
	// 	Name: to.Ptr("public"),
	// 	Type: to.Ptr("Microsoft.ApiCenter/services/apis"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/apis/echo-api"),
	// 	SystemData: &armapicenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
	// 	},
	// 	Properties: &armapicenter.APIProperties{
	// 		Description: to.Ptr("A simple HTTP request/response service."),
	// 		CustomProperties: map[string]any{
	// 			"author": "John Doe",
	// 		},
	// 		ExternalDocumentation: []*armapicenter.ExternalDocumentation{
	// 			{
	// 				Title: to.Ptr("Onboarding docs"),
	// 				URL: to.Ptr("https://docs.contoso.com"),
	// 		}},
	// 		Kind: to.Ptr(armapicenter.APIKindRest),
	// 		License: &armapicenter.License{
	// 			URL: to.Ptr("https://contoso.com/license"),
	// 		},
	// 		LifecycleStage: to.Ptr(armapicenter.LifecycleStageDesign),
	// 		TermsOfService: &armapicenter.TermsOfService{
	// 			URL: to.Ptr("https://contoso.com/terms-of-service"),
	// 		},
	// 		Title: to.Ptr("Echo API"),
	// 	},
	// }
}

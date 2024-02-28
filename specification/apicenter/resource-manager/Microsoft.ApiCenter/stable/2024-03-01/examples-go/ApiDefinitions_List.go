package armapicenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiDefinitions_List.json
func ExampleAPIDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPIDefinitionsClient().NewListPager("contoso-resources", "contoso", "default", "echo-api", "2023-01-01", &armapicenter.APIDefinitionsClientListOptions{Filter: nil})
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
		// page.APIDefinitionListResult = armapicenter.APIDefinitionListResult{
		// 	Value: []*armapicenter.APIDefinition{
		// 		{
		// 			Name: to.Ptr("openapi"),
		// 			Type: to.Ptr("Microsoft.ApiCenter/services/apis/versions/definitions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/apis/echo-api/versions/2023-01-01/definitions/openapi"),
		// 			SystemData: &armapicenter.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
		// 			},
		// 			Properties: &armapicenter.APIDefinitionProperties{
		// 				Description: to.Ptr("Default spec"),
		// 				Specification: &armapicenter.APIDefinitionPropertiesSpecification{
		// 					Name: to.Ptr("openapi"),
		// 					Version: to.Ptr("3.0.6"),
		// 				},
		// 				Title: to.Ptr("OpenAPI"),
		// 			},
		// 	}},
		// }
	}
}

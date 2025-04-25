package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateGlobalSchema2.json
func ExampleGlobalSchemaClient_BeginCreateOrUpdate_apiManagementCreateSchema2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGlobalSchemaClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "schema1", armapimanagement.GlobalSchemaContract{
		Properties: &armapimanagement.GlobalSchemaContractProperties{
			Description: to.Ptr("sample schema description"),
			Document: map[string]any{
				"type":    "object",
				"$id":     "https://example.com/person.schema.json",
				"$schema": "https://json-schema.org/draft/2020-12/schema",
				"properties": map[string]any{
					"age": map[string]any{
						"type":        "integer",
						"description": "Age in years which must be equal to or greater than zero.",
						"minimum":     float64(0),
					},
					"firstName": map[string]any{
						"type":        "string",
						"description": "The person's first name.",
					},
					"lastName": map[string]any{
						"type":        "string",
						"description": "The person's last name.",
					},
				},
				"title": "Person",
			},
			SchemaType: to.Ptr(armapimanagement.SchemaTypeJSON),
		},
	}, &armapimanagement.GlobalSchemaClientBeginCreateOrUpdateOptions{IfMatch: nil})
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
	// res.GlobalSchemaContract = armapimanagement.GlobalSchemaContract{
	// 	Name: to.Ptr("schema1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/schemas"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/schemas/schema1"),
	// 	Properties: &armapimanagement.GlobalSchemaContractProperties{
	// 		Description: to.Ptr("sample schema description"),
	// 		Document: map[string]any{
	// 			"type": "object",
	// 			"$id": "https://example.com/person.schema.json",
	// 			"$schema": "https://json-schema.org/draft/2020-12/schema",
	// 			"properties":map[string]any{
	// 				"age":map[string]any{
	// 					"type": "integer",
	// 					"description": "Age in years which must be equal to or greater than zero.",
	// 					"minimum": float64(0),
	// 				},
	// 				"firstName":map[string]any{
	// 					"type": "string",
	// 					"description": "The person's first name.",
	// 				},
	// 				"lastName":map[string]any{
	// 					"type": "string",
	// 					"description": "The person's last name.",
	// 				},
	// 			},
	// 			"title": "Person",
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SchemaType: to.Ptr(armapimanagement.SchemaTypeJSON),
	// 	},
	// }
}

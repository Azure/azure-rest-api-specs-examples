package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/PutApplication_example.json
func ExampleApplicationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationClient().CreateOrUpdate(ctx, "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", armsecurity.Application{
		Properties: &armsecurity.ApplicationProperties{
			Description: to.Ptr("An application on critical recommendations"),
			ConditionSets: []any{
				map[string]any{
					"conditions": []any{
						map[string]any{
							"operator": "contains",
							"property": "$.Id",
							"value":    "-bil-",
						},
					},
				}},
			DisplayName:        to.Ptr("Admin's application"),
			SourceResourceType: to.Ptr(armsecurity.ApplicationSourceResourceTypeAssessments),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armsecurity.Application{
	// 	Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 	Type: to.Ptr("Microsoft.Security/applications"),
	// 	ID: to.Ptr("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/applications/ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
	// 	Properties: &armsecurity.ApplicationProperties{
	// 		Description: to.Ptr("An application on critical recommendations"),
	// 		ConditionSets: []any{
	// 			map[string]any{
	// 				"conditions":[]any{
	// 					map[string]any{
	// 						"operator": "contains",
	// 						"property": "$.Id",
	// 						"value": "-dev-",
	// 					},
	// 				},
	// 		}},
	// 		DisplayName: to.Ptr("Admin's application"),
	// 		SourceResourceType: to.Ptr(armsecurity.ApplicationSourceResourceTypeAssessments),
	// 	},
	// }
}

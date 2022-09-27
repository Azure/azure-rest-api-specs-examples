package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/PutSecurityConnectorApplication_example.json
func ExampleConnectorApplicationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewConnectorApplicationClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "gcpResourceGroup", "gcpconnector", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", armsecurity.Application{
		Properties: &armsecurity.ApplicationProperties{
			Description: to.Ptr("An application on critical GCP recommendations"),
			ConditionSets: []interface{}{
				map[string]interface{}{
					"conditions": []interface{}{
						map[string]interface{}{
							"operator": "contains",
							"property": "$.Id",
							"value":    "-prod-",
						},
					},
				}},
			DisplayName:        to.Ptr("GCP Admin's application"),
			SourceResourceType: to.Ptr(armsecurity.ApplicationSourceResourceTypeAssessments),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

package armdeployments_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeployments"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/18609d68cf243ee3ce35d7c005ff3c7dd2cd9477/specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/CalculateTemplateHash.json
func ExampleDeploymentsClient_CalculateTemplateHash() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeployments.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeploymentsClient().CalculateTemplateHash(ctx, map[string]any{
		"$schema":        "http://schemas.management.azure.com/deploymentTemplate?api-version=2014-04-01-preview",
		"contentVersion": "1.0.0.0",
		"outputs": map[string]any{
			"string": map[string]any{
				"type":  "string",
				"value": "myvalue",
			},
		},
		"parameters": map[string]any{
			"string": map[string]any{
				"type": "string",
			},
		},
		"resources": []any{},
		"variables": map[string]any{
			"array": []any{
				float64(1),
				float64(2),
				float64(3),
				float64(4),
			},
			"bool": true,
			"int":  float64(42),
			"object": map[string]any{
				"object": map[string]any{
					"location": "West US",
					"vmSize":   "Large",
				},
			},
			"string": "string",
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TemplateHashResult = armdeployments.TemplateHashResult{
	// 	MinifiedTemplate: to.Ptr("{\"$SCHEMA\":\"HTTP://SCHEMAS.MANAGEMENT.AZURE.COM/DEPLOYMENTTEMPLATE?API-VERSION=2014-04-01-PREVIEW\",\"CONTENTVERSION\":\"1.0.0.0\",\"PARAMETERS\":{\"STRING\":{\"TYPE\":\"STRING\"}},\"VARIABLES\":{\"STRING\":\"STRING\",\"INT\":42,\"BOOL\":TRUE,\"ARRAY\":[1,2,3,4],\"OBJECT\":{\"OBJECT\":{\"VMSIZE\":\"LARGE\",\"LOCATION\":\"WEST US\"}}},\"RESOURCES\":[],\"OUTPUTS\":{\"STRING\":{\"TYPE\":\"STRING\",\"VALUE\":\"MYVALUE\"}}}"),
	// 	TemplateHash: to.Ptr("695440707931307747"),
	// }
}

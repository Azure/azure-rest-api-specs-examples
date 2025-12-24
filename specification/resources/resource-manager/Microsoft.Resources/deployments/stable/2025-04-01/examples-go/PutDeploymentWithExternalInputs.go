package armdeployments_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeployments"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05809299e74fc16df7badbab3fb29e25da59f7d2/specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/PutDeploymentWithExternalInputs.json
func ExampleDeploymentsClient_BeginCreateOrUpdate_createDeploymentUsingExternalInputs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeployments.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginCreateOrUpdate(ctx, "my-resource-group", "my-deployment", armdeployments.Deployment{
		Properties: &armdeployments.DeploymentProperties{
			ExternalInputDefinitions: map[string]*armdeployments.DeploymentExternalInputDefinition{
				"fooValue": {
					Config: "FOO_VALUE",
					Kind:   to.Ptr("sys.envVar"),
				},
			},
			ExternalInputs: map[string]*armdeployments.DeploymentExternalInput{
				"fooValue": {
					Value: "baz",
				},
			},
			Mode: to.Ptr(armdeployments.DeploymentModeIncremental),
			Parameters: map[string]*armdeployments.DeploymentParameter{
				"inputObj": {
					Expression: to.Ptr("[createObject('foo', externalInputs('fooValue'))]"),
				},
			},
			Template: map[string]any{
				"$schema":        "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
				"contentVersion": "1.0.0.0",
				"outputs": map[string]any{
					"inputObj": map[string]any{
						"type":  "object",
						"value": "[parameters('inputObj')]",
					},
				},
				"parameters": map[string]any{
					"inputObj": map[string]any{
						"type": "object",
					},
				},
				"resources": []any{},
			},
		},
	}, nil)
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
	// res.DeploymentExtended = armdeployments.DeploymentExtended{
	// 	Name: to.Ptr("my-deployment"),
	// 	Type: to.Ptr("Microsoft.Resources/deployments"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.Resources/deployments/my-deployment"),
	// 	Properties: &armdeployments.DeploymentPropertiesExtended{
	// 		CorrelationID: to.Ptr("ef613b6c-f76e-48fd-9da7-28884243c5e5"),
	// 		Dependencies: []*armdeployments.Dependency{
	// 		},
	// 		Mode: to.Ptr(armdeployments.DeploymentModeIncremental),
	// 		OutputResources: []*armdeployments.ResourceReference{
	// 		},
	// 		Outputs: map[string]any{
	// 			"inputObj":map[string]any{
	// 				"type": "Object",
	// 				"value":map[string]any{
	// 					"foo": "baz",
	// 				},
	// 			},
	// 		},
	// 		Parameters: map[string]any{
	// 			"inputObj":map[string]any{
	// 				"type": "Object",
	// 				"value":map[string]any{
	// 					"foo": "baz",
	// 				},
	// 			},
	// 		},
	// 		Providers: []*armdeployments.Provider{
	// 		},
	// 		ProvisioningState: to.Ptr(armdeployments.ProvisioningStateSucceeded),
	// 		TemplateHash: to.Ptr("17686481789412793580"),
	// 		Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-09T14:36:48.204Z"); return t}()),
	// 	},
	// }
}

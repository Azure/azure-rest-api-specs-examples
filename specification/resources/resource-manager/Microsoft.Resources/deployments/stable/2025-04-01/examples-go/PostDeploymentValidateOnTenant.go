package armdeployments_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeployments"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05809299e74fc16df7badbab3fb29e25da59f7d2/specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/PostDeploymentValidateOnTenant.json
func ExampleDeploymentsClient_BeginValidateAtTenantScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeployments.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginValidateAtTenantScope(ctx, "my-deployment", armdeployments.ScopedDeployment{
		Location: to.Ptr("eastus"),
		Properties: &armdeployments.DeploymentProperties{
			Mode:       to.Ptr(armdeployments.DeploymentModeIncremental),
			Parameters: map[string]*armdeployments.DeploymentParameter{},
			TemplateLink: &armdeployments.TemplateLink{
				URI: to.Ptr("https://example.com/exampleTemplate.json"),
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
	// res.DeploymentValidateResult = armdeployments.DeploymentValidateResult{
	// 	Name: to.Ptr("my-deployment"),
	// 	Type: to.Ptr("Microsoft.Resources/deployments"),
	// 	ID: to.Ptr("/providers/Microsoft.Resources/deployments/my-deployment"),
	// 	Properties: &armdeployments.DeploymentPropertiesExtended{
	// 		CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Dependencies: []*armdeployments.Dependency{
	// 		},
	// 		Diagnostics: []*armdeployments.DeploymentDiagnosticsDefinition{
	// 			{
	// 				Code: to.Ptr("NestedDeploymentShortCircuited"),
	// 				Level: to.Ptr(armdeployments.LevelWarning),
	// 				Message: to.Ptr("A nested deployment got short-circuited and all its resources got skipped from validation. This is due to a nested template having a parameter that was not fully evaluated (e.g. contains a reference() function)."),
	// 			},
	// 			{
	// 				Code: to.Ptr("NestedDeploymentSkippedFromInternalExpansion"),
	// 				Level: to.Ptr(armdeployments.LevelWarning),
	// 				Message: to.Ptr("When nested deployments are expanded, all its inner resources are retrieved for further validation. This process is performed in batch of: '10' at a time. Nested deployments exceeding this batch count are skipped from expansion."),
	// 			},
	// 			{
	// 				Code: to.Ptr("NestedDeploymentExpansionLimitReached"),
	// 				Level: to.Ptr(armdeployments.LevelWarning),
	// 				Message: to.Ptr("Nested deployments are expanded up to: '50' in total. Nested deployments exceeding this count are skipped from expansion."),
	// 		}},
	// 		Duration: to.Ptr("PT22.8356799S"),
	// 		Mode: to.Ptr(armdeployments.DeploymentModeIncremental),
	// 		Parameters: map[string]any{
	// 		},
	// 		Providers: []*armdeployments.Provider{
	// 			{
	// 				Namespace: to.Ptr("Microsoft.Authorization"),
	// 				ResourceTypes: []*armdeployments.ProviderResourceType{
	// 					{
	// 						Locations: []*string{
	// 							nil},
	// 							ResourceType: to.Ptr("policyDefinitions"),
	// 					}},
	// 			}},
	// 			ProvisioningState: to.Ptr(armdeployments.ProvisioningStateSucceeded),
	// 			TemplateHash: to.Ptr("0000000000000000000"),
	// 			TemplateLink: &armdeployments.TemplateLink{
	// 				ContentVersion: to.Ptr("1.0.0.0"),
	// 				URI: to.Ptr("https://example.com/exampleTemplate.json"),
	// 			},
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-05T01:20:01.723Z"); return t}()),
	// 			ValidatedResources: []*armdeployments.ResourceReference{
	// 				{
	// 					ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/policy-definition-name"),
	// 			}},
	// 		},
	// 	}
}

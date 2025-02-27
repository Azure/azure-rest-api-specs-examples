package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/91bfc0d02eaed75e6a3bfb5b9b150c84c79400ed/specification/resources/resource-manager/Microsoft.Resources/stable/2024-11-01/examples/PostDeploymentValidateOnManagementGroup.json
func ExampleDeploymentsClient_BeginValidateAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginValidateAtManagementGroupScope(ctx, "my-management-group-id", "my-deployment", armresources.ScopedDeployment{
		Location: to.Ptr("eastus"),
		Properties: &armresources.DeploymentProperties{
			Mode:       to.Ptr(armresources.DeploymentModeIncremental),
			Parameters: map[string]*armresources.DeploymentParameter{},
			TemplateLink: &armresources.TemplateLink{
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
	// res.DeploymentValidateResult = armresources.DeploymentValidateResult{
	// 	Name: to.Ptr("my-deployment"),
	// 	Type: to.Ptr("Microsoft.Resources/deployments"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Resources/deployments/my-deployment"),
	// 	Properties: &armresources.DeploymentPropertiesExtended{
	// 		CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Dependencies: []*armresources.Dependency{
	// 		},
	// 		Diagnostics: []*armresources.DeploymentDiagnosticsDefinition{
	// 			{
	// 				Code: to.Ptr("NestedDeploymentShortCircuited"),
	// 				Level: to.Ptr(armresources.LevelWarning),
	// 				Message: to.Ptr("A nested deployment got short-circuited and all its resources got skipped from validation. This is due to a nested template having a parameter that was not fully evaluated (e.g. contains a reference() function)."),
	// 			},
	// 			{
	// 				Code: to.Ptr("NestedDeploymentSkippedFromInternalExpansion"),
	// 				Level: to.Ptr(armresources.LevelWarning),
	// 				Message: to.Ptr("When nested deployments are expanded, all its inner resources are retrieved for further validation. This process is performed in batch of: '10' at a time. Nested deployments exceeding this batch count are skipped from expansion."),
	// 			},
	// 			{
	// 				Code: to.Ptr("NestedDeploymentExpansionLimitReached"),
	// 				Level: to.Ptr(armresources.LevelWarning),
	// 				Message: to.Ptr("Nested deployments are expanded up to: '50' in total. Nested deployments exceeding this count are skipped from expansion."),
	// 		}},
	// 		Duration: to.Ptr("PT22.8356799S"),
	// 		Mode: to.Ptr(armresources.DeploymentModeIncremental),
	// 		Parameters: map[string]any{
	// 		},
	// 		Providers: []*armresources.Provider{
	// 			{
	// 				Namespace: to.Ptr("Microsoft.Authorization"),
	// 				ResourceTypes: []*armresources.ProviderResourceType{
	// 					{
	// 						Locations: []*string{
	// 							nil},
	// 							ResourceType: to.Ptr("policyAssignments"),
	// 					}},
	// 			}},
	// 			ProvisioningState: to.Ptr(armresources.ProvisioningStateSucceeded),
	// 			TemplateHash: to.Ptr("0000000000000000000"),
	// 			TemplateLink: &armresources.TemplateLink{
	// 				ContentVersion: to.Ptr("1.0.0.0"),
	// 				URI: to.Ptr("https://example.com/exampleTemplate.json"),
	// 			},
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-05T01:20:01.723Z"); return t}()),
	// 			ValidatedResources: []*armresources.ResourceReference{
	// 				{
	// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment"),
	// 			}},
	// 		},
	// 	}
}

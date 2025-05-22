package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ada20fa9d6dcc6453f04f32d20fe3cf97293aa88/specification/resources/resource-manager/Microsoft.Resources/stable/2025-03-01/examples/PutDeploymentAtManagementGroup.json
func ExampleDeploymentsClient_BeginCreateOrUpdateAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginCreateOrUpdateAtManagementGroupScope(ctx, "my-management-group-id", "my-deployment", armresources.ScopedDeployment{
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
	// res.DeploymentExtended = armresources.DeploymentExtended{
	// 	Name: to.Ptr("my-deployment"),
	// 	Type: to.Ptr("Microsoft.Resources/deployments"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Resources/deployments/my-deployment"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armresources.DeploymentPropertiesExtended{
	// 		CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Dependencies: []*armresources.Dependency{
	// 			{
	// 				DependsOn: []*armresources.BasicDependency{
	// 					{
	// 						ID: to.Ptr("/providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Authorization/policyDefinitions/policy-definition-name"),
	// 						ResourceName: to.Ptr("policy-definition-name"),
	// 						ResourceType: to.Ptr("Microsoft.Authorization/policyDefinitions"),
	// 				}},
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/my-management-group-id/providers/Microsoft.Authorization/policyAssignments/location-lock"),
	// 				ResourceName: to.Ptr("location-lock"),
	// 				ResourceType: to.Ptr("Microsoft.Authorization/policyAssignments"),
	// 		}},
	// 		Duration: to.Ptr("PT1.2970875S"),
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
	// 							ResourceType: to.Ptr("policyDefinitions"),
	// 						},
	// 						{
	// 							Locations: []*string{
	// 								nil},
	// 								ResourceType: to.Ptr("policyAssignments"),
	// 						}},
	// 					},
	// 					{
	// 						Namespace: to.Ptr("Microsoft.Resources"),
	// 						ResourceTypes: []*armresources.ProviderResourceType{
	// 							{
	// 								Locations: []*string{
	// 									to.Ptr("eastus")},
	// 									ResourceType: to.Ptr("deployments"),
	// 							}},
	// 					}},
	// 					ProvisioningState: to.Ptr(armresources.ProvisioningStateSucceeded),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-24T22:52:38.789Z"); return t}()),
	// 				},
	// 			}
}

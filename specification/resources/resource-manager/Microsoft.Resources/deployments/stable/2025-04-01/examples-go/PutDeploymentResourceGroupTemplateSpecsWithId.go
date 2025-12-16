package armdeployments_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeployments"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/18609d68cf243ee3ce35d7c005ff3c7dd2cd9477/specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/PutDeploymentResourceGroupTemplateSpecsWithId.json
func ExampleDeploymentsClient_BeginCreateOrUpdate_createADeploymentThatWillDeployATemplateSpecWithTheGivenResourceId() {
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
			Mode:       to.Ptr(armdeployments.DeploymentModeIncremental),
			Parameters: map[string]*armdeployments.DeploymentParameter{},
			TemplateLink: &armdeployments.TemplateLink{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.Resources/TemplateSpecs/TemplateSpec-Name/versions/v1"),
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
	// 		CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Dependencies: []*armdeployments.Dependency{
	// 		},
	// 		Duration: to.Ptr("PT22.8356799S"),
	// 		Mode: to.Ptr(armdeployments.DeploymentModeIncremental),
	// 		OutputResources: []*armdeployments.ResourceReference{
	// 			{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.Storage/storageAccounts/my-storage-account"),
	// 		}},
	// 		Parameters: map[string]any{
	// 		},
	// 		Providers: []*armdeployments.Provider{
	// 			{
	// 				Namespace: to.Ptr("Microsoft.Storage"),
	// 				ResourceTypes: []*armdeployments.ProviderResourceType{
	// 					{
	// 						Locations: []*string{
	// 							to.Ptr("eastus")},
	// 							ResourceType: to.Ptr("storageAccounts"),
	// 					}},
	// 			}},
	// 			ProvisioningState: to.Ptr(armdeployments.ProvisioningStateSucceeded),
	// 			TemplateHash: to.Ptr("0000000000000000000"),
	// 			TemplateLink: &armdeployments.TemplateLink{
	// 				ContentVersion: to.Ptr("1.0.0.0"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.Resources/TemplateSpecs/TemplateSpec-Name/versions/v1"),
	// 			},
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-05T01:20:01.723Z"); return t}()),
	// 		},
	// 	}
}

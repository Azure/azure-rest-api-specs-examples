package armdeployments_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeployments"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/18609d68cf243ee3ce35d7c005ff3c7dd2cd9477/specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/PutDeploymentWithOnErrorDeploymentLastSuccessful.json
func ExampleDeploymentsClient_BeginCreateOrUpdate_createADeploymentThatWillRedeployTheLastSuccessfulDeploymentOnFailure() {
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
			Mode: to.Ptr(armdeployments.DeploymentModeComplete),
			OnErrorDeployment: &armdeployments.OnErrorDeployment{
				Type: to.Ptr(armdeployments.OnErrorDeploymentTypeLastSuccessful),
			},
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
	// res.DeploymentExtended = armdeployments.DeploymentExtended{
	// 	Name: to.Ptr("my-deployment"),
	// 	Type: to.Ptr("Microsoft.Resources/deployments"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/my-resource-group/providers/Microsoft.Resources/deployments/my-deployment"),
	// 	Properties: &armdeployments.DeploymentPropertiesExtended{
	// 		CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Dependencies: []*armdeployments.Dependency{
	// 			{
	// 				DependsOn: []*armdeployments.BasicDependency{
	// 					{
	// 						ID: to.Ptr("{resourceid}"),
	// 						ResourceName: to.Ptr("VNet1"),
	// 						ResourceType: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 				}},
	// 				ID: to.Ptr("{resourceid}"),
	// 				ResourceName: to.Ptr("VNet1/Subnet1"),
	// 				ResourceType: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 			},
	// 			{
	// 				DependsOn: []*armdeployments.BasicDependency{
	// 					{
	// 						ID: to.Ptr("{resourceid}"),
	// 						ResourceName: to.Ptr("VNet1"),
	// 						ResourceType: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 					},
	// 					{
	// 						ID: to.Ptr("{resourceid}"),
	// 						ResourceName: to.Ptr("VNet1/Subnet1"),
	// 						ResourceType: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 				}},
	// 				ID: to.Ptr("{resourceid}"),
	// 				ResourceName: to.Ptr("VNet1/Subnet2"),
	// 				ResourceType: to.Ptr("Microsoft.Network/virtualNetworks/subnets"),
	// 		}},
	// 		Duration: to.Ptr("PT0.8204881S"),
	// 		Mode: to.Ptr(armdeployments.DeploymentModeComplete),
	// 		OnErrorDeployment: &armdeployments.OnErrorDeploymentExtended{
	// 			Type: to.Ptr(armdeployments.OnErrorDeploymentTypeLastSuccessful),
	// 			DeploymentName: to.Ptr("{nameOfLastSuccesfulDeployment}"),
	// 		},
	// 		Parameters: map[string]any{
	// 		},
	// 		Providers: []*armdeployments.Provider{
	// 			{
	// 				Namespace: to.Ptr("Microsoft.Network"),
	// 				ResourceTypes: []*armdeployments.ProviderResourceType{
	// 					{
	// 						Locations: []*string{
	// 							to.Ptr("centralus")},
	// 							ResourceType: to.Ptr("virtualNetworks"),
	// 						},
	// 						{
	// 							Locations: []*string{
	// 								to.Ptr("centralus")},
	// 								ResourceType: to.Ptr("virtualNetworks/subnets"),
	// 						}},
	// 				}},
	// 				ProvisioningState: to.Ptr(armdeployments.ProvisioningStateSucceeded),
	// 				TemplateLink: &armdeployments.TemplateLink{
	// 					ContentVersion: to.Ptr("1.0.0.0"),
	// 					URI: to.Ptr("https://example.com/exampleTemplate.json"),
	// 				},
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T00:00:00.000Z"); return t}()),
	// 			},
	// 		}
}

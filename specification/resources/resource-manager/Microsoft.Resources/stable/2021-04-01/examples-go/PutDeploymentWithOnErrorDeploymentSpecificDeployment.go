package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fd842fb73656039ec94ce367bcedee25a57bd18/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentWithOnErrorDeploymentSpecificDeployment.json
func ExampleDeploymentsClient_BeginCreateOrUpdate_createADeploymentThatWillRedeployAnotherDeploymentOnFailure() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginCreateOrUpdate(ctx, "my-resource-group", "my-deployment", armresources.Deployment{
		Properties: &armresources.DeploymentProperties{
			Mode: to.Ptr(armresources.DeploymentModeComplete),
			OnErrorDeployment: &armresources.OnErrorDeployment{
				Type:           to.Ptr(armresources.OnErrorDeploymentTypeSpecificDeployment),
				DeploymentName: to.Ptr("name-of-deployment-to-use"),
			},
			Parameters: map[string]any{},
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
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/my-resource-group/providers/Microsoft.Resources/deployments/my-deployment"),
	// 	Properties: &armresources.DeploymentPropertiesExtended{
	// 		CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Dependencies: []*armresources.Dependency{
	// 			{
	// 				DependsOn: []*armresources.BasicDependency{
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
	// 				DependsOn: []*armresources.BasicDependency{
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
	// 		Mode: to.Ptr(armresources.DeploymentModeComplete),
	// 		OnErrorDeployment: &armresources.OnErrorDeploymentExtended{
	// 			Type: to.Ptr(armresources.OnErrorDeploymentTypeSpecificDeployment),
	// 			DeploymentName: to.Ptr("name-of-deployment-to-use"),
	// 		},
	// 		Parameters: map[string]any{
	// 		},
	// 		Providers: []*armresources.Provider{
	// 			{
	// 				Namespace: to.Ptr("Microsoft.Network"),
	// 				ResourceTypes: []*armresources.ProviderResourceType{
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
	// 				ProvisioningState: to.Ptr(armresources.ProvisioningStateSucceeded),
	// 				TemplateLink: &armresources.TemplateLink{
	// 					ContentVersion: to.Ptr("1.0.0.0"),
	// 					URI: to.Ptr("https://example.com/exampleTemplate.json"),
	// 				},
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-01T00:00:00.000Z"); return t}()),
	// 			},
	// 		}
}

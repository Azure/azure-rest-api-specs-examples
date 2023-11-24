package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fd842fb73656039ec94ce367bcedee25a57bd18/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentSubscriptionTemplateSpecsWithId.json
func ExampleDeploymentsClient_BeginCreateOrUpdateAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginCreateOrUpdateAtSubscriptionScope(ctx, "my-deployment", armresources.Deployment{
		Location: to.Ptr("eastus"),
		Properties: &armresources.DeploymentProperties{
			Mode:       to.Ptr(armresources.DeploymentModeIncremental),
			Parameters: map[string]any{},
			TemplateLink: &armresources.TemplateLink{
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
	// res.DeploymentExtended = armresources.DeploymentExtended{
	// 	Name: to.Ptr("my-deployment"),
	// 	Type: to.Ptr("Microsoft.Resources/deployments"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/providers/Microsoft.Resources/deployments/my-deployment"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armresources.DeploymentPropertiesExtended{
	// 		CorrelationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Dependencies: []*armresources.Dependency{
	// 		},
	// 		Duration: to.Ptr("PT1.2637681S"),
	// 		Mode: to.Ptr(armresources.DeploymentModeIncremental),
	// 		OutputResources: []*armresources.ResourceReference{
	// 		},
	// 		Parameters: map[string]any{
	// 		},
	// 		Providers: []*armresources.Provider{
	// 		},
	// 		ProvisioningState: to.Ptr(armresources.ProvisioningStateSucceeded),
	// 		TemplateHash: to.Ptr("0000000000000000000"),
	// 		TemplateLink: &armresources.TemplateLink{
	// 			ContentVersion: to.Ptr("1.0.0.0"),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group/providers/Microsoft.Resources/TemplateSpecs/TemplateSpec-Name/versions/v1"),
	// 		},
	// 		Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-05T01:51:58.628Z"); return t}()),
	// 	},
	// }
}

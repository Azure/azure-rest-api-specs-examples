package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_CreateOrUpdate.json
func ExampleDeploymentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewDeploymentsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myservice",
		"myapp",
		"mydeployment",
		armappplatform.DeploymentResource{
			Properties: &armappplatform.DeploymentResourceProperties{
				DeploymentSettings: &armappplatform.DeploymentSettings{
					AddonConfigs: map[string]map[string]interface{}{
						"ApplicationConfigurationService": {
							"patterns": []interface{}{
								"mypattern",
							},
						},
					},
					EnvironmentVariables: map[string]*string{
						"env": to.Ptr("test"),
					},
					ResourceRequests: &armappplatform.ResourceRequests{
						CPU:    to.Ptr("1000m"),
						Memory: to.Ptr("3Gi"),
					},
				},
				Source: &armappplatform.SourceUploadedUserSourceInfo{
					Type:             to.Ptr("Source"),
					Version:          to.Ptr("1.0"),
					RelativePath:     to.Ptr("resources/a172cedcae47474b615c54d510a5d84a8dea3032e958587430b413538be3f333-2019082605-e3095339-1723-44b7-8b5e-31b1003978bc"),
					ArtifactSelector: to.Ptr("sub-module-1"),
				},
			},
			SKU: &armappplatform.SKU{
				Name:     to.Ptr("S0"),
				Capacity: to.Ptr[int32](1),
				Tier:     to.Ptr("Standard"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

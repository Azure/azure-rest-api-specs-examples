```go
package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-05-01-preview/examples/Deployments_CreateOrUpdate.json
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
					LivenessProbe: &armappplatform.Probe{
						DisableProbe:        to.Ptr(false),
						FailureThreshold:    to.Ptr[int32](3),
						InitialDelaySeconds: to.Ptr[int32](30),
						PeriodSeconds:       to.Ptr[int32](10),
						ProbeAction: &armappplatform.HTTPGetAction{
							Type:   to.Ptr(armappplatform.ProbeActionTypeHTTPGetAction),
							Path:   to.Ptr("/health"),
							Scheme: to.Ptr(armappplatform.HTTPSchemeTypeHTTP),
						},
					},
					ReadinessProbe: &armappplatform.Probe{
						DisableProbe:        to.Ptr(false),
						FailureThreshold:    to.Ptr[int32](3),
						InitialDelaySeconds: to.Ptr[int32](30),
						PeriodSeconds:       to.Ptr[int32](10),
						ProbeAction: &armappplatform.HTTPGetAction{
							Type:   to.Ptr(armappplatform.ProbeActionTypeHTTPGetAction),
							Path:   to.Ptr("/health"),
							Scheme: to.Ptr(armappplatform.HTTPSchemeTypeHTTP),
						},
					},
					ResourceRequests: &armappplatform.ResourceRequests{
						CPU:    to.Ptr("1000m"),
						Memory: to.Ptr("3Gi"),
					},
					TerminationGracePeriodSeconds: to.Ptr[int32](30),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappplatform%2Farmappplatform%2Fv1.1.0-beta.1/sdk/resourcemanager/appplatform/armappplatform/README.md) on how to add the SDK to your project and authenticate.

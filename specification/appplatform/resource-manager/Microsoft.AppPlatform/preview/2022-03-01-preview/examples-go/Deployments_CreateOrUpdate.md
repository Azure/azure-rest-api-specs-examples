Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappplatform%2Farmappplatform%2Fv0.5.0/sdk/resourcemanager/appplatform/armappplatform/README.md) on how to add the SDK to your project and authenticate.

```go
package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/Deployments_CreateOrUpdate.json
func ExampleDeploymentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappplatform.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		"<deployment-name>",
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
						CPU:    to.Ptr("<cpu>"),
						Memory: to.Ptr("<memory>"),
					},
				},
				Source: &armappplatform.SourceUploadedUserSourceInfo{
					Type:             to.Ptr("<type>"),
					Version:          to.Ptr("<version>"),
					RelativePath:     to.Ptr("<relative-path>"),
					ArtifactSelector: to.Ptr("<artifact-selector>"),
				},
			},
			SKU: &armappplatform.SKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int32](1),
				Tier:     to.Ptr("<tier>"),
			},
		},
		&armappplatform.DeploymentsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

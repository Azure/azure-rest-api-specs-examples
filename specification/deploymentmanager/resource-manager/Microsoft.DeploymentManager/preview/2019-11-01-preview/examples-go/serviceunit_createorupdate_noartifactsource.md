Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeploymentmanager%2Farmdeploymentmanager%2Fv0.1.0/sdk/resourcemanager/deploymentmanager/armdeploymentmanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeploymentmanager_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate_noartifactsource.json
func ExampleServiceUnitsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewServiceUnitsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-topology-name>",
		"<service-name>",
		"<service-unit-name>",
		armdeploymentmanager.ServiceUnitResource{
			TrackedResource: armdeploymentmanager.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Properties: &armdeploymentmanager.ServiceUnitResourceProperties{
				ServiceUnitProperties: armdeploymentmanager.ServiceUnitProperties{
					Artifacts: &armdeploymentmanager.ServiceUnitArtifacts{
						ParametersURI: to.StringPtr("<parameters-uri>"),
						TemplateURI:   to.StringPtr("<template-uri>"),
					},
					DeploymentMode:      armdeploymentmanager.DeploymentModeIncremental.ToPtr(),
					TargetResourceGroup: to.StringPtr("<target-resource-group>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ServiceUnitResource.ID: %s\n", *res.ID)
}
```

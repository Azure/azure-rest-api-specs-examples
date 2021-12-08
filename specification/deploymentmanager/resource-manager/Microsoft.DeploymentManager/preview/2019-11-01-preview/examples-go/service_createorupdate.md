Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeploymentmanager%2Farmdeploymentmanager%2Fv0.1.0/sdk/resourcemanager/deploymentmanager/armdeploymentmanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/service_createorupdate.json
func ExampleServicesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentmanager.NewServicesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-topology-name>",
		"<service-name>",
		armdeploymentmanager.ServiceResource{
			TrackedResource: armdeploymentmanager.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Properties: &armdeploymentmanager.ServiceResourceProperties{
				ServiceProperties: armdeploymentmanager.ServiceProperties{
					TargetLocation:       to.StringPtr("<target-location>"),
					TargetSubscriptionID: to.StringPtr("<target-subscription-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ServiceResource.ID: %s\n", *res.ID)
}
```

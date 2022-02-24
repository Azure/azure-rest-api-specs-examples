Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiotcentral%2Farmiotcentral%2Fv0.2.1/sdk/resourcemanager/iotcentral/armiotcentral/README.md) on how to add the SDK to your project and authenticate.

```go
package armiotcentral_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"
)

// x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_CreateOrUpdate.json
func ExampleAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotcentral.NewAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armiotcentral.App{
			Location: to.StringPtr("<location>"),
			Identity: &armiotcentral.SystemAssignedServiceIdentity{
				Type: armiotcentral.SystemAssignedServiceIdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armiotcentral.AppProperties{
				DisplayName: to.StringPtr("<display-name>"),
				Subdomain:   to.StringPtr("<subdomain>"),
				Template:    to.StringPtr("<template>"),
			},
			SKU: &armiotcentral.AppSKUInfo{
				Name: armiotcentral.AppSKU("ST2").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.AppsClientCreateOrUpdateResult)
}
```

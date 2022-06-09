```go
package armiotcentral_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_CreateOrUpdate.json
func ExampleAppsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiotcentral.NewAppsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resRg",
		"myIoTCentralApp",
		armiotcentral.App{
			Location: to.Ptr("westus"),
			Identity: &armiotcentral.SystemAssignedServiceIdentity{
				Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
			},
			Properties: &armiotcentral.AppProperties{
				DisplayName: to.Ptr("My IoT Central App"),
				Subdomain:   to.Ptr("my-iot-central-app"),
				Template:    to.Ptr("iotc-pnp-preview@1.0.0"),
			},
			SKU: &armiotcentral.AppSKUInfo{
				Name: to.Ptr(armiotcentral.AppSKUST2),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiotcentral%2Farmiotcentral%2Fv1.0.0/sdk/resourcemanager/iotcentral/armiotcentral/README.md) on how to add the SDK to your project and authenticate.

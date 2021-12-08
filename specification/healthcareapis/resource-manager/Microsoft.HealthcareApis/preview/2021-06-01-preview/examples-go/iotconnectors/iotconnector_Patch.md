Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthcareapis%2Farmhealthcareapis%2Fv0.1.0/sdk/resourcemanager/healthcareapis/armhealthcareapis/README.md) on how to add the SDK to your project and authenticate.

```go
package armhealthcareapis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/iotconnectors/iotconnector_Patch.json
func ExampleIotConnectorsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewIotConnectorsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<iot-connector-name>",
		"<workspace-name>",
		armhealthcareapis.IotConnectorPatchResource{
			ResourceTags: armhealthcareapis.ResourceTags{
				Tags: map[string]*string{
					"additionalProp1": to.StringPtr("string"),
					"additionalProp2": to.StringPtr("string"),
					"additionalProp3": to.StringPtr("string"),
				},
			},
			ServiceManagedIdentity: armhealthcareapis.ServiceManagedIdentity{
				Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
					Type: armhealthcareapis.ManagedServiceIdentityTypeSystemAssigned.ToPtr(),
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
	log.Printf("IotConnector.ID: %s\n", *res.ID)
}
```

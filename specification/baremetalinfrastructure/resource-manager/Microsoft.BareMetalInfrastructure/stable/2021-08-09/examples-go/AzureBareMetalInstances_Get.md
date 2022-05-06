Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbaremetalinfrastructure%2Farmbaremetalinfrastructure%2Fv0.4.0/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure/README.md) on how to add the SDK to your project and authenticate.

```go
package armbaremetalinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalInstances_Get.json
func ExampleAzureBareMetalInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armbaremetalinfrastructure.NewAzureBareMetalInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<azure-bare-metal-instance-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

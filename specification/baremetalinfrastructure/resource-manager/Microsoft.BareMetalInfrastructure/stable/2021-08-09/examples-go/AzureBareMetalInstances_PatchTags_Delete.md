Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbaremetalinfrastructure%2Farmbaremetalinfrastructure%2Fv1.0.0/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure/README.md) on how to add the SDK to your project and authenticate.

```go
package armbaremetalinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalInstances_PatchTags_Delete.json
func ExampleAzureBareMetalInstancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbaremetalinfrastructure.NewAzureBareMetalInstancesClient("f0f4887f-d13c-4943-a8ba-d7da28d2a3fd", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"myResourceGroup",
		"myABMInstance",
		armbaremetalinfrastructure.Tags{
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

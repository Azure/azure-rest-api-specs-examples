Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerservice%2Farmcontainerservice%2Fv0.2.1/sdk/resourcemanager/containerservice/armcontainerservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/MaintenanceConfigurationsGet.json
func ExampleMaintenanceConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewMaintenanceConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<config-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("MaintenanceConfiguration.ID: %s\n", *res.ID)
}
```

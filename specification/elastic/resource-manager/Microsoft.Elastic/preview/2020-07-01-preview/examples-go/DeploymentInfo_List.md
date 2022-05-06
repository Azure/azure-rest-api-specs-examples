Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Felastic%2Farmelastic%2Fv0.4.0/sdk/resourcemanager/elastic/armelastic/README.md) on how to add the SDK to your project and authenticate.

```go
package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/elastic/resource-manager/Microsoft.Elastic/preview/2020-07-01-preview/examples/DeploymentInfo_List.json
func ExampleDeploymentInfoClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armelastic.NewDeploymentInfoClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

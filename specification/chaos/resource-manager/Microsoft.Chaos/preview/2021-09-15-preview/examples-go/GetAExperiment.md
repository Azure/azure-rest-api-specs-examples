Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fchaos%2Farmchaos%2Fv0.1.0/sdk/resourcemanager/chaos/armchaos/README.md) on how to add the SDK to your project and authenticate.

```go
package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2021-09-15-preview/examples/GetAExperiment.json
func ExampleExperimentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armchaos.NewExperimentsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<experiment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Experiment.ID: %s\n", *res.ID)
}
```

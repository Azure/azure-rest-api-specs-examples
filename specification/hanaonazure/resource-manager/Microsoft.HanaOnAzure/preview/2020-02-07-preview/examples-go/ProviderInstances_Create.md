Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhanaonazure%2Farmhanaonazure%2Fv0.1.0/sdk/resourcemanager/hanaonazure/armhanaonazure/README.md) on how to add the SDK to your project and authenticate.

```go
package armhanaonazure_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
)

// x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/ProviderInstances_Create.json
func ExampleProviderInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhanaonazure.NewProviderInstancesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<sap-monitor-name>",
		"<provider-instance-name>",
		armhanaonazure.ProviderInstance{
			Properties: &armhanaonazure.ProviderInstanceProperties{
				Type:       to.StringPtr("<type>"),
				Metadata:   to.StringPtr("<metadata>"),
				Properties: to.StringPtr("<properties>"),
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
	log.Printf("ProviderInstance.ID: %s\n", *res.ID)
}
```

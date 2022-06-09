```go
package armservicelinker_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker"
)

// x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2021-11-01-preview/examples/ValidateLinkFailure.json
func ExampleLinkerClient_BeginValidate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicelinker.NewLinkerClient(cred, nil)
	poller, err := client.BeginValidate(ctx,
		"<resource-uri>",
		"<linker-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LinkerClientValidateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicelinker%2Farmservicelinker%2Fv0.2.1/sdk/resourcemanager/servicelinker/armservicelinker/README.md) on how to add the SDK to your project and authenticate.

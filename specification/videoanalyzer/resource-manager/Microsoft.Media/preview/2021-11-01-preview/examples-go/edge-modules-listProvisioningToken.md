Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.2.1/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```go
package armvideoanalyzer_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/edge-modules-listProvisioningToken.json
func ExampleEdgeModulesClient_ListProvisioningToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewEdgeModulesClient("<subscription-id>", cred, nil)
	res, err := client.ListProvisioningToken(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<edge-module-name>",
		armvideoanalyzer.ListProvisioningTokenInput{
			ExpirationDate: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-23T11:04:49.0526841-08:00"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EdgeModulesClientListProvisioningTokenResult)
}
```

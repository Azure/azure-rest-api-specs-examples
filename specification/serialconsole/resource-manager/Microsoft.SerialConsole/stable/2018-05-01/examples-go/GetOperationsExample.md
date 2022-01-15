Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fserialconsole%2Farmserialconsole%2Fv0.2.0/sdk/resourcemanager/serialconsole/armserialconsole/README.md) on how to add the SDK to your project and authenticate.

```go
package armserialconsole_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole"
)

// x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/GetOperationsExample.json
func ExampleMicrosoftSerialConsoleClient_ListOperations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armserialconsole.NewMicrosoftSerialConsoleClient("<subscription-id>", cred, nil)
	res, err := client.ListOperations(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MicrosoftSerialConsoleClientListOperationsResult)
}
```

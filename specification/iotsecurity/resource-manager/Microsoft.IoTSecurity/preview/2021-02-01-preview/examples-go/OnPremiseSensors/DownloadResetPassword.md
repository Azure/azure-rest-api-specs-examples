Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiotsecurity%2Farmiotsecurity%2Fv0.1.0/sdk/resourcemanager/iotsecurity/armiotsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/OnPremiseSensors/DownloadResetPassword.json
func ExampleOnPremiseSensorsClient_DownloadResetPassword() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewOnPremiseSensorsClient("<subscription-id>", cred, nil)
	_, err = client.DownloadResetPassword(ctx,
		"<on-premise-sensor-name>",
		armiotsecurity.ResetPasswordInput{
			ApplianceID: to.StringPtr("<appliance-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

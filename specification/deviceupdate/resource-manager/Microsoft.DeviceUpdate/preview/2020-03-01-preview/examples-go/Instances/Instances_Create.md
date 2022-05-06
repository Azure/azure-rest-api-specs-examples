Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeviceupdate%2Farmdeviceupdate%2Fv0.4.0/sdk/resourcemanager/deviceupdate/armdeviceupdate/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeviceupdate_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2020-03-01-preview/examples/Instances/Instances_Create.json
func ExampleInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdeviceupdate.NewInstancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<instance-name>",
		armdeviceupdate.Instance{
			Location: to.Ptr("<location>"),
			Properties: &armdeviceupdate.InstanceProperties{
				DiagnosticStorageProperties: &armdeviceupdate.DiagnosticStorageProperties{
					AuthenticationType: to.Ptr(armdeviceupdate.AuthenticationTypeKeyBased),
					ConnectionString:   to.Ptr("<connection-string>"),
					ResourceID:         to.Ptr("<resource-id>"),
				},
				EnableDiagnostics: to.Ptr(false),
				IotHubs: []*armdeviceupdate.IotHubSettings{
					{
						EventHubConnectionString: to.Ptr("<event-hub-connection-string>"),
						IoTHubConnectionString:   to.Ptr("<io-thub-connection-string>"),
						ResourceID:               to.Ptr("<resource-id>"),
					}},
			},
		},
		&armdeviceupdate.InstancesClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```

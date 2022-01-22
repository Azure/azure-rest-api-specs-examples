Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv0.2.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.

```go
package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/UpdateDiagnosticRemoteSupportSettings.json
func ExampleDiagnosticSettingsClient_BeginUpdateDiagnosticRemoteSupportSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDiagnosticSettingsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateDiagnosticRemoteSupportSettings(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.DiagnosticRemoteSupportSettings{
			Properties: &armdataboxedge.DiagnosticRemoteSupportSettingsProperties{
				RemoteSupportSettingsList: []*armdataboxedge.RemoteSupportSettings{
					{
						AccessLevel:              armdataboxedge.AccessLevel("ReadWrite").ToPtr(),
						ExpirationTimeStampInUTC: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-07T00:00:00+00:00"); return t }()),
						RemoteApplicationType:    armdataboxedge.RemoteApplicationType("Powershell").ToPtr(),
					}},
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
	log.Printf("Response result: %#v\n", res.DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResult)
}
```

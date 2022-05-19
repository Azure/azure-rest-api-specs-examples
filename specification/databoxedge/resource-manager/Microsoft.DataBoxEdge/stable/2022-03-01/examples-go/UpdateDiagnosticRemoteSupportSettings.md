Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv1.0.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/UpdateDiagnosticRemoteSupportSettings.json
func ExampleDiagnosticSettingsClient_BeginUpdateDiagnosticRemoteSupportSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataboxedge.NewDiagnosticSettingsClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateDiagnosticRemoteSupportSettings(ctx,
		"testedgedevice",
		"GroupForEdgeAutomation",
		armdataboxedge.DiagnosticRemoteSupportSettings{
			Properties: &armdataboxedge.DiagnosticRemoteSupportSettingsProperties{
				RemoteSupportSettingsList: []*armdataboxedge.RemoteSupportSettings{
					{
						AccessLevel:              to.Ptr(armdataboxedge.AccessLevelReadWrite),
						ExpirationTimeStampInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-07T00:00:00+00:00"); return t }()),
						RemoteApplicationType:    to.Ptr(armdataboxedge.RemoteApplicationTypePowershell),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

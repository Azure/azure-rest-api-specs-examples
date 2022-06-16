package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/UpdateDiagnosticProactiveLogCollectionSettings.json
func ExampleDiagnosticSettingsClient_BeginUpdateDiagnosticProactiveLogCollectionSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDiagnosticSettingsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateDiagnosticProactiveLogCollectionSettings(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.DiagnosticProactiveLogCollectionSettings{
			Properties: &armdataboxedge.ProactiveLogCollectionSettingsProperties{
				UserConsent: armdataboxedge.ProactiveDiagnosticsConsent("Enabled").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResult)
}

package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/MonitoringSettings_UpdatePatch.json
func ExampleMonitoringSettingsClient_BeginUpdatePatch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappplatform.NewMonitoringSettingsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdatePatch(ctx,
		"<resource-group-name>",
		"<service-name>",
		armappplatform.MonitoringSettingResource{
			Properties: &armappplatform.MonitoringSettingProperties{
				AppInsightsInstrumentationKey: to.Ptr("<app-insights-instrumentation-key>"),
				AppInsightsSamplingRate:       to.Ptr[float64](10),
				TraceEnabled:                  to.Ptr(true),
			},
		},
		&armappplatform.MonitoringSettingsClientBeginUpdatePatchOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

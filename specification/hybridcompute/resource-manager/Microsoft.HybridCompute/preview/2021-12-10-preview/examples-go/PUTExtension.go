package armhybridcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-12-10-preview/examples/PUTExtension.json
func ExampleMachineExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhybridcompute.NewMachineExtensionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<machine-name>",
		"<extension-name>",
		armhybridcompute.MachineExtension{
			Location: to.Ptr("<location>"),
			Properties: &armhybridcompute.MachineExtensionProperties{
				Type:      to.Ptr("<type>"),
				Publisher: to.Ptr("<publisher>"),
				Settings: map[string]interface{}{
					"commandToExecute": "powershell.exe -c \"Get-Process | Where-Object { $_.CPU -gt 10000 }\"",
				},
				TypeHandlerVersion: to.Ptr("<type-handler-version>"),
			},
		},
		&armhybridcompute.MachineExtensionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

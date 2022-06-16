package armazurestackhci_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-03-01/examples/PatchExtension.json
func ExampleExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armazurestackhci.NewExtensionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<arc-setting-name>",
		"<extension-name>",
		armazurestackhci.Extension{
			Properties: &armazurestackhci.ExtensionProperties{
				ExtensionParameters: &armazurestackhci.ExtensionParameters{
					Type:      to.Ptr("<type>"),
					Publisher: to.Ptr("<publisher>"),
					Settings: map[string]interface{}{
						"workspaceId": "xx",
					},
					TypeHandlerVersion: to.Ptr("<type-handler-version>"),
				},
			},
		},
		&armazurestackhci.ExtensionsClientBeginUpdateOptions{ResumeToken: ""})
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

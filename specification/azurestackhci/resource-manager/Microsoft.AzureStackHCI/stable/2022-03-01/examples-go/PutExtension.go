package armazurestackhci_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-03-01/examples/PutExtension.json
func ExampleExtensionsClient_BeginCreate() {
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
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<arc-setting-name>",
		"<extension-name>",
		armazurestackhci.Extension{
			Properties: &armazurestackhci.ExtensionProperties{
				ExtensionParameters: &armazurestackhci.ExtensionParameters{
					Type: to.Ptr("<type>"),
					ProtectedSettings: map[string]interface{}{
						"workspaceKey": "xx",
					},
					Publisher: to.Ptr("<publisher>"),
					Settings: map[string]interface{}{
						"workspaceId": "xx",
					},
					TypeHandlerVersion: to.Ptr("<type-handler-version>"),
				},
			},
		},
		&armazurestackhci.ExtensionsClientBeginCreateOptions{ResumeToken: ""})
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

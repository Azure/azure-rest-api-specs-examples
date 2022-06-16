package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2021-12-01-preview/examples/ConnectedRegistryUpdate.json
func ExampleConnectedRegistriesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewConnectedRegistriesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<connected-registry-name>",
		armcontainerregistry.ConnectedRegistryUpdateParameters{
			Properties: &armcontainerregistry.ConnectedRegistryUpdateProperties{
				ClientTokenIDs: []*string{
					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"),
					to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token")},
				Logging: &armcontainerregistry.LoggingProperties{
					AuditLogStatus: to.Ptr(armcontainerregistry.AuditLogStatusEnabled),
					LogLevel:       to.Ptr(armcontainerregistry.LogLevelDebug),
				},
				NotificationsList: []*string{
					to.Ptr("hello-world:*:*"),
					to.Ptr("sample/repo/*:1.0:*")},
				SyncProperties: &armcontainerregistry.SyncUpdateProperties{
					MessageTTL: to.Ptr("<message-ttl>"),
					Schedule:   to.Ptr("<schedule>"),
					SyncWindow: to.Ptr("<sync-window>"),
				},
			},
		},
		&armcontainerregistry.ConnectedRegistriesClientBeginUpdateOptions{ResumeToken: ""})
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

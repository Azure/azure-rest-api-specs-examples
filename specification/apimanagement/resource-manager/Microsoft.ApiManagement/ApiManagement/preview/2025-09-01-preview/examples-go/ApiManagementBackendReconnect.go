package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementBackendReconnect.json
func ExampleBackendClient_Reconnect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBackendClient().Reconnect(ctx, "rg1", "apimService1", "proxybackend", &armapimanagement.BackendClientReconnectOptions{
		Parameters: &armapimanagement.BackendReconnectContract{
			Properties: &armapimanagement.BackendReconnectProperties{
				After: to.Ptr("PT3S"),
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

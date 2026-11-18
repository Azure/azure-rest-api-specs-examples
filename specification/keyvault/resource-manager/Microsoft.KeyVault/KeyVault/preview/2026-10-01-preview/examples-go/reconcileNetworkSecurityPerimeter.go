package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2026-10-01-preview/reconcileNetworkSecurityPerimeter.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_ReconcileConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().ReconcileConfiguration(ctx, "sample-resource-group", "sample-vault", "00000000-0000-0000-0000-000000000000.sampleAssociation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

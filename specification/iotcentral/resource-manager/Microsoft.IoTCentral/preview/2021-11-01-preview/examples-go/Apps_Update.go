package armiotcentral_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/96e52e2b911d533f95a0ad8e324c828d556c5f2b/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Update.json
func ExampleAppsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppsClient().BeginUpdate(ctx, "resRg", "myIoTCentralApp", armiotcentral.AppPatch{
		Identity: &armiotcentral.SystemAssignedServiceIdentity{
			Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armiotcentral.AppProperties{
			DisplayName: to.Ptr("My IoT Central App 2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

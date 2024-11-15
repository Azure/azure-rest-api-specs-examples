package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/98d74b2db60e46ceb7e3b75755e51519cd500485/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/StorageAppliances_EnableRemoteVendorManagement.json
func ExampleStorageAppliancesClient_BeginEnableRemoteVendorManagement() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageAppliancesClient().BeginEnableRemoteVendorManagement(ctx, "resourceGroupName", "storageApplianceName", &armnetworkcloud.StorageAppliancesClientBeginEnableRemoteVendorManagementOptions{StorageApplianceEnableRemoteVendorManagementParameters: &armnetworkcloud.StorageApplianceEnableRemoteVendorManagementParameters{
		SupportEndpoints: []*string{
			to.Ptr("10.0.0.0/24")},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

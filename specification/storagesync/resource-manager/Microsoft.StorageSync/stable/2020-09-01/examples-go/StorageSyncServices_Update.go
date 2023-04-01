package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/StorageSyncServices_Update.json
func ExampleServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginUpdate(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", &armstoragesync.ServicesClientBeginUpdateOptions{Parameters: &armstoragesync.ServiceUpdateParameters{
		Properties: &armstoragesync.ServiceUpdateProperties{
			IncomingTrafficPolicy: to.Ptr(armstoragesync.IncomingTrafficPolicyAllowAllTraffic),
		},
		Tags: map[string]*string{
			"Dept":        to.Ptr("IT"),
			"Environment": to.Ptr("Test"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Service = armstoragesync.Service{
	// 	Name: to.Ptr("SampleStorageSyncService_1"),
	// 	Type: to.Ptr("Microsoft.StorageSync/storageSyncServices"),
	// 	ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1"),
	// 	Location: to.Ptr("WestUS"),
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("IT"),
	// 		"Environment": to.Ptr("Test"),
	// 	},
	// 	Properties: &armstoragesync.ServiceProperties{
	// 		IncomingTrafficPolicy: to.Ptr(armstoragesync.IncomingTrafficPolicyAllowAllTraffic),
	// 	},
	// }
}

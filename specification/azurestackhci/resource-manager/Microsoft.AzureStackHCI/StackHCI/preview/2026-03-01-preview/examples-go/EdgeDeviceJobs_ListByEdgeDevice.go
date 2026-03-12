package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/EdgeDeviceJobs_ListByEdgeDevice.json
func ExampleEdgeDeviceJobsClient_NewListByEdgeDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEdgeDeviceJobsClient().NewListByEdgeDevicePager("subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1", "YE-855IEIN585-", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armazurestackhci.EdgeDeviceJobsClientListByEdgeDeviceResponse{
		// 	EdgeDeviceJobListResult: armazurestackhci.EdgeDeviceJobListResult{
		// 		NextLink: to.Ptr("https://microsoft.com/alqov"),
		// 		Value: []armazurestackhci.EdgeDeviceJobClassification{
		// 			&armazurestackhci.HciEdgeDeviceJob{
		// 				Name: to.Ptr("collectLog"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/edgeDevices/jobs"),
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/providers/Microsoft.AzureStackHCI/edgeDevices/default/jobs/collectLog"),
		// 				Kind: to.Ptr(armazurestackhci.EdgeDeviceKindHCI),
		// 				Properties: &armazurestackhci.HciRemoteSupportJobProperties{
		// 					Type: to.Ptr(armazurestackhci.RemoteSupportTypeEnable),
		// 					AccessLevel: to.Ptr(armazurestackhci.RemoteSupportAccessLevelDiagnostics),
		// 					ExpirationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t}()),
		// 					JobType: to.Ptr(armazurestackhci.HciEdgeDeviceJobTypeRemoteSupport),
		// 				},
		// 				SystemData: &armazurestackhci.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
		// 					CreatedBy: to.Ptr("zsnvvvbj"),
		// 					CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("qxlrx"),
		// 					LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

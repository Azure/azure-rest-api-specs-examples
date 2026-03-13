package armazurestackhci_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/EdgeDeviceJobs_CreateOrUpdate_CollectLog.json
func ExampleEdgeDeviceJobsClient_BeginCreateOrUpdate_edgeDeviceJobsCreateOrUpdateCollectLog() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEdgeDeviceJobsClient().BeginCreateOrUpdate(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1", "default", "collectLog", &armazurestackhci.HciEdgeDeviceJob{
		Kind: to.Ptr(armazurestackhci.EdgeDeviceKindHCI),
		Properties: &armazurestackhci.HciCollectLogJobProperties{
			FromDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t }()),
			JobType:  to.Ptr(armazurestackhci.HciEdgeDeviceJobTypeCollectLog),
			ToDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t }()),
		},
	}, nil)
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
	// res = armazurestackhci.EdgeDeviceJobsClientCreateOrUpdateResponse{
	// 	HciEdgeDeviceJob: &armazurestackhci.HciEdgeDeviceJob{
	// 		Name: to.Ptr("collectLog"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeDevices/jobs"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/providers/Microsoft.AzureStackHCI/edgeDevices/default/jobs/collectLog"),
	// 		Kind: to.Ptr(armazurestackhci.EdgeDeviceKindHCI),
	// 		Properties: &armazurestackhci.HciCollectLogJobProperties{
	// 			FromDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t}()),
	// 			JobType: to.Ptr(armazurestackhci.HciEdgeDeviceJobTypeCollectLog),
	// 			ToDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T10:43:27.9471574Z"); return t}()),
	// 		},
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
	// 			CreatedBy: to.Ptr("zsnvvvbj"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T17:02:21.168Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("qxlrx"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}

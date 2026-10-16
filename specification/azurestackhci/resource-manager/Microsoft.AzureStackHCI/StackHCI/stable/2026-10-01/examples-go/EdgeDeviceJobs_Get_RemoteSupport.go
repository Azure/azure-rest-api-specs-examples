package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeDeviceJobs_Get_RemoteSupport.json
func ExampleEdgeDeviceJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEdgeDeviceJobsClient().Get(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1", "-5M1G7G10OZ-o5b-HS3-c72", "-oUxs", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.EdgeDeviceJobsClientGetResponse{
	// 	EdgeDeviceJobClassification: &armazurestackhci.HciEdgeDeviceJob{
	// 		Name: to.Ptr("collectLog"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeDevices/jobs"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/providers/Microsoft.AzureStackHCI/edgeDevices/default/jobs/RemoteSupport"),
	// 		Kind: to.Ptr(armazurestackhci.EdgeDeviceKindHCI),
	// 		Properties: &armazurestackhci.HciRemoteSupportJobProperties{
	// 			Type: to.Ptr(armazurestackhci.RemoteSupportTypeEnable),
	// 			AccessLevel: to.Ptr(armazurestackhci.RemoteSupportAccessLevelDiagnostics),
	// 			ExpirationTimestamp: to.Ptr(time.Date(2024, time.January, 29, 10, 43, 27, 947157400, time.UTC)),
	// 			JobType: to.Ptr(armazurestackhci.HciEdgeDeviceJobTypeRemoteSupport),
	// 		},
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2024, time.July, 1, 17, 2, 21, 168000000, time.UTC)),
	// 			CreatedBy: to.Ptr("zsnvvvbj"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.July, 1, 17, 2, 21, 168000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("qxlrx"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}

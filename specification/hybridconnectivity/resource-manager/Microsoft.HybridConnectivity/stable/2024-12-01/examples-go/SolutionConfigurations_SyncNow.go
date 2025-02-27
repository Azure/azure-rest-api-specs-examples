package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/SolutionConfigurations_SyncNow.json
func ExampleSolutionConfigurationsClient_BeginSyncNow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSolutionConfigurationsClient().BeginSyncNow(ctx, "ymuj", "tks", nil)
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
	// res = armhybridconnectivity.SolutionConfigurationsClientSyncNowResponse{
	// 	OperationStatusResult: &armhybridconnectivity.OperationStatusResult{
	// 		ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		ResourceID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		Name: to.Ptr("svqtraeuwvyvblujlvqilypwpdrt"),
	// 		Status: to.Ptr("bevmrejij"),
	// 		PercentComplete: to.Ptr[float64](15),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T18:38:19.143Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-02T18:38:19.143Z"); return t}()),
	// 		Operations: []*armhybridconnectivity.OperationStatusResult{
	// 		},
	// 		Error: &armhybridconnectivity.ErrorDetail{
	// 			Code: to.Ptr("ykzvluyqiftfsumgvwzdh"),
	// 			Message: to.Ptr("krbjgtqkjgiux"),
	// 			Target: to.Ptr("nsaucxt"),
	// 			Details: []*armhybridconnectivity.ErrorDetail{
	// 			},
	// 			AdditionalInfo: []*armhybridconnectivity.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr("qivvrewsjvcildjgwwytgimwklh"),
	// 					Info: &armhybridconnectivity.ErrorAdditionalInfoInfo{
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

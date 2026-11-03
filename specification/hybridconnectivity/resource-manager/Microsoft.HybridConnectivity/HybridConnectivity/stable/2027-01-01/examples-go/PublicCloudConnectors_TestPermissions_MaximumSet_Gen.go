package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2027-01-01/PublicCloudConnectors_TestPermissions_MaximumSet_Gen.json
func ExamplePublicCloudConnectorsClient_BeginTestPermissions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").BeginTestPermissions(ctx, "rghybridconnectivity", "abc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.PublicCloudConnectorsClientTestPermissionsResponse{
	// 	OperationStatusResult: armhybridconnectivity.OperationStatusResult{
	// 		ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		ResourceID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
	// 		Name: to.Ptr("svqtraeuwvyvblujlvqilypwpdrt"),
	// 		Status: to.Ptr("bevmrejij"),
	// 		PercentComplete: to.Ptr[float64](15),
	// 		StartTime: to.Ptr(time.Date(2024, time.October, 2, 18, 38, 19, 143000000, time.UTC)),
	// 		EndTime: to.Ptr(time.Date(2024, time.October, 2, 18, 38, 19, 143000000, time.UTC)),
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
	// 					Info: map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

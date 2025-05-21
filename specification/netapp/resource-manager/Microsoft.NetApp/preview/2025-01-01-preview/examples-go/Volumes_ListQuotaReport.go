package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb3217991ff57b5760525aeba1a0670bfe0880fa/specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/Volumes_ListQuotaReport.json
func ExampleVolumesClient_BeginListQuotaReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumesClient().BeginListQuotaReport(ctx, "myRG", "account1", "pool1", "volume1", nil)
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
	// res.ListQuotaReportResponse = armnetapp.ListQuotaReportResponse{
	// 	Value: []*armnetapp.QuotaReport{
	// 		{
	// 			IsDerivedQuota: to.Ptr(false),
	// 			PercentageUsed: to.Ptr[float32](5),
	// 			QuotaLimitTotalInKiBs: to.Ptr[int64](204914688),
	// 			QuotaLimitUsedInKiBs: to.Ptr[int64](8192),
	// 			QuotaTarget: to.Ptr("1013"),
	// 			QuotaType: to.Ptr(armnetapp.TypeIndividualUserQuota),
	// 		},
	// 		{
	// 			IsDerivedQuota: to.Ptr(false),
	// 			PercentageUsed: to.Ptr[float32](5),
	// 			QuotaLimitTotalInKiBs: to.Ptr[int64](204914688),
	// 			QuotaLimitUsedInKiBs: to.Ptr[int64](8192),
	// 			QuotaTarget: to.Ptr("1012"),
	// 			QuotaType: to.Ptr(armnetapp.TypeIndividualGroupQuota),
	// 	}},
	// }
}

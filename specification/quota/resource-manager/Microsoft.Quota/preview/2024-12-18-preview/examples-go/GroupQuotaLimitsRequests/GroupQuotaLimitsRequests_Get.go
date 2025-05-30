package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/55c5a0cd6da80b2700333c01e9a9c6067de9cef0/specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/GroupQuotaLimitsRequests/GroupQuotaLimitsRequests_Get.json
func ExampleGroupQuotaLimitsRequestClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGroupQuotaLimitsRequestClient().Get(ctx, "E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", "requestId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubmittedResourceRequestStatus = armquota.SubmittedResourceRequestStatus{
	// 	Name: to.Ptr("requestId"),
	// 	Type: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaLimitsRequests"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1/resourceProviders/Microsoft.Compute/groupQuotaLimitsRequests/requestId"),
	// 	Properties: &armquota.SubmittedResourceRequestStatusProperties{
	// 		FaultCode: to.Ptr("ResourceNotAvailableForOffer"),
	// 		ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
	// 		RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-20T05:29:34.144Z"); return t}()),
	// 		RequestedResource: &armquota.GroupQuotaRequestBase{
	// 			Properties: &armquota.GroupQuotaRequestBaseProperties{
	// 				Name: &armquota.GroupQuotaRequestBasePropertiesName{
	// 					LocalizedValue: to.Ptr("standard DDv4 Family vCPUs"),
	// 					Value: to.Ptr("standardddv4family"),
	// 				},
	// 				Comments: to.Ptr(""),
	// 				Limit: to.Ptr[int64](100),
	// 				Region: to.Ptr("westus"),
	// 			},
	// 		},
	// 	},
	// }
}

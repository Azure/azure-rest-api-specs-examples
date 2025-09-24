package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota/v2"
)

// Generated from example definition: 2025-09-01/GroupQuotaLimitsRequests/GroupQuotaLimitsRequests_List.json
func ExampleGroupQuotaLimitsRequestClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupQuotaLimitsRequestClient().NewListPager("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", "Microsoft.Compute", "location eq westus", nil)
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
		// page = armquota.GroupQuotaLimitsRequestClientListResponse{
		// 	SubmittedResourceRequestStatusList: armquota.SubmittedResourceRequestStatusList{
		// 		NextLink: to.Ptr("https://yourLinkHere.com"),
		// 		Value: []*armquota.SubmittedResourceRequestStatus{
		// 			{
		// 				Name: to.Ptr("requestId1"),
		// 				Type: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaLimitsRequests"),
		// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1/resourceProviders/Microsoft.Compute/groupQuotaLimitsRequests/requestId1"),
		// 				Properties: &armquota.SubmittedResourceRequestStatusProperties{
		// 					FaultCode: to.Ptr("ResourceNotAvailableForOffer"),
		// 					ProvisioningState: to.Ptr(armquota.RequestStateFailed),
		// 					RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-18T00:12:34.004Z"); return t}()),
		// 					RequestedResource: &armquota.GroupQuotaRequestBase{
		// 						Properties: &armquota.GroupQuotaRequestBaseProperties{
		// 							Name: &armquota.GroupQuotaRequestBasePropertiesName{
		// 								LocalizedValue: to.Ptr("standard DDv4 Family vCPUs"),
		// 								Value: to.Ptr("standardddv4family"),
		// 							},
		// 							Comments: to.Ptr("Contoso requires more quota."),
		// 							Limit: to.Ptr[int64](100),
		// 							Region: to.Ptr("westus"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("requestId2"),
		// 				Type: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaLimitsRequests"),
		// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1/resourceProviders/Microsoft.Compute/groupQuotaLimitsRequests/requestId2"),
		// 				Properties: &armquota.SubmittedResourceRequestStatusProperties{
		// 					FaultCode: to.Ptr(""),
		// 					ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
		// 					RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-17T00:57:43.410Z"); return t}()),
		// 					RequestedResource: &armquota.GroupQuotaRequestBase{
		// 						Properties: &armquota.GroupQuotaRequestBaseProperties{
		// 							Name: &armquota.GroupQuotaRequestBasePropertiesName{
		// 								LocalizedValue: to.Ptr("standardav2family"),
		// 								Value: to.Ptr("standardav2family"),
		// 							},
		// 							Comments: to.Ptr("Contoso requires more quota."),
		// 							Limit: to.Ptr[int64](150),
		// 							Region: to.Ptr("westus"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimitsRequests/PatchGroupQuotaLimitsRequests-Compute.json
func ExampleGroupQuotaLimitsRequestClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGroupQuotaLimitsRequestClient().BeginUpdate(ctx, "E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", "Microsoft.Compute", "standardav2family", &armquota.GroupQuotaLimitsRequestClientBeginUpdateOptions{GroupQuotaRequest: nil})
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
	// res.SubmittedResourceRequestStatus = armquota.SubmittedResourceRequestStatus{
	// 	Name: to.Ptr("requestId1"),
	// 	Type: to.Ptr("Microsoft.Quota/groupQuotas/groupQuotaLimitsRequests"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/E7EC67B3-7657-4966-BFFC-41EFD36BAA09/providers/Microsoft.Quota/groupQuotas/groupquota1/resourceProviders/Microsoft.Compute/groupQuotaLimitsRequests/requestId1"),
	// 	Properties: &armquota.SubmittedResourceRequestStatusProperties{
	// 		ProvisioningState: to.Ptr(armquota.RequestStateSucceeded),
	// 		RequestSubmitTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-08T12:09:27.978Z"); return t}()),
	// 		RequestedResource: &armquota.GroupQuotaRequestBase{
	// 			Properties: &armquota.GroupQuotaRequestBaseProperties{
	// 				Name: &armquota.GroupQuotaRequestBasePropertiesName{
	// 					LocalizedValue: to.Ptr("Standard AV2 Family vCPUs"),
	// 					Value: to.Ptr("standardav2family"),
	// 				},
	// 				Comments: to.Ptr("Contoso requires more quota."),
	// 				Limit: to.Ptr[int64](100),
	// 				Region: to.Ptr("westus"),
	// 			},
	// 		},
	// 	},
	// }
}

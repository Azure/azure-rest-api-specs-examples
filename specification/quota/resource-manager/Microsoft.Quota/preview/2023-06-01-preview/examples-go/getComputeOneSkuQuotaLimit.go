package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/106483d9f698ac3b6c0d481ab0c5fab14152e21f/specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/getComputeOneSkuQuotaLimit.json
func ExampleClient_Get_quotasGetRequestForCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "standardNDSFamily", "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CurrentQuotaLimitBase = armquota.CurrentQuotaLimitBase{
	// 	Name: to.Ptr("standardNDSFamily"),
	// 	Type: to.Ptr("Microsoft.Quota/Quotas"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.Quota/Quotas/standardNDSFamily"),
	// 	Properties: &armquota.Properties{
	// 		Name: &armquota.ResourceName{
	// 			LocalizedValue: to.Ptr("Standard NDS Family vCPUs"),
	// 			Value: to.Ptr("standardNDSFamily"),
	// 		},
	// 		IsQuotaApplicable: to.Ptr(true),
	// 		Limit: &armquota.LimitObject{
	// 			LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
	// 			LimitType: to.Ptr(armquota.QuotaLimitTypesIndependent),
	// 			Value: to.Ptr[int32](100),
	// 		},
	// 		Unit: to.Ptr("Count"),
	// 	},
	// }
}

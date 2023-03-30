package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b9b91929c304f8fb44002267b6c98d9fb9dde014/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GrantIncludeAllocatedBudget.json
func ExampleGrantsClient_Get_grantIncludeAllocatedBudget() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeducation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGrantsClient().Get(ctx, "{billingAccountName}", "{billingProfileName}", &armeducation.GrantsClientGetOptions{IncludeAllocatedBudget: to.Ptr(false)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GrantDetails = armeducation.GrantDetails{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Education/Grants"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/providers/Microsoft.Education/grants/default"),
	// 	Properties: &armeducation.GrantDetailProperties{
	// 		AllocatedBudget: &armeducation.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](0),
	// 		},
	// 		EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-09T09:08:05.505Z"); return t}()),
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-09T09:08:05.505Z"); return t}()),
	// 		OfferCap: &armeducation.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](100),
	// 		},
	// 		OfferType: to.Ptr(armeducation.GrantTypeStudent),
	// 		Status: to.Ptr(armeducation.GrantStatusActive),
	// 	},
	// }
}

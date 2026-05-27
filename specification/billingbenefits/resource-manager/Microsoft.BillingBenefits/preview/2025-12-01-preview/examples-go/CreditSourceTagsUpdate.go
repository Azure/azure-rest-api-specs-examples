package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/CreditSourceTagsUpdate.json
func ExampleSourcesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSourcesClient("97ee05f2-07d5-494d-908c-081a197f4277").Update(ctx, "resource_group_name_01", "credit_20231212", "source_20231212", armbillingbenefits.CreditSourcePatchRequest{
		Tags: map[string]*string{
			"key1": to.Ptr("value4"),
			"key2": to.Ptr("value5"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbillingbenefits.SourcesClientUpdateResponse{
	// 	CreditSource: armbillingbenefits.CreditSource{
	// 		Name: to.Ptr("source_20231212"),
	// 		Type: to.Ptr("Microsoft.BillingBenefits/credits/sources"),
	// 		ID: to.Ptr("/subscriptions/97ee05f2-07d5-494d-908c-081a197f4277/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/credits/credit_20231212/sources/source_20231212"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armbillingbenefits.CreditSourceProperties{
	// 			Credit: &armbillingbenefits.Commitment{
	// 				Amount: to.Ptr[float64](20000),
	// 				CurrencyCode: to.Ptr("USD"),
	// 				Grain: to.Ptr(armbillingbenefits.CommitmentGrainFullTerm),
	// 			},
	// 			ImpactedBillingPeriod: to.Ptr("202304"),
	// 			SourceResourceID: to.Ptr("/subscriptions/{subId}"),
	// 			Status: to.Ptr(armbillingbenefits.CreditStatusPending),
	// 		},
	// 		SystemData: &armbillingbenefits.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-12T01:01:01.1075056Z"); return t}()),
	// 			CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-12T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value4"),
	// 			"key2": to.Ptr("value5"),
	// 		},
	// 	},
	// }
}

package armbillingbenefits_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/ContributorCreate.json
func ExampleMaccsClient_BeginCreate_contributorCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMaccsClient("20000000-0000-0000-0000-000000000000").BeginCreate(ctx, "resource_group_name_01", "macc_contributor_20230614", armbillingbenefits.Macc{
		Location: to.Ptr("global"),
		Properties: &armbillingbenefits.MaccModelProperties{
			EndAt:             to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00Z"); return t }()),
			EntityType:        to.Ptr(armbillingbenefits.MaccEntityTypeContributor),
			PrimaryResourceID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/resource_group_name_02/providers/Microsoft.BillingBenefits/maccs/macc_20230614"),
			ProductCode:       to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
			StartAt:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t }()),
			SystemID:          to.Ptr("13810867107109237"),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
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
	// res = armbillingbenefits.MaccsClientCreateResponse{
	// 	Macc: armbillingbenefits.Macc{
	// 		Name: to.Ptr("macc_contributor_20230614"),
	// 		Type: to.Ptr("Microsoft.BillingBenefits/maccs"),
	// 		ID: to.Ptr("/subscriptions/20000000-0000-0000-0000-000000000000/resourceGroups/resource_group_name_02/providers/Microsoft.BillingBenefits/maccs/macc_contributor_20230614"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armbillingbenefits.MaccModelProperties{
	// 			AutomaticShortfall: to.Ptr(armbillingbenefits.EnablementModeEnabled),
	// 			DisplayName: to.Ptr("macc contributor 20230614"),
	// 			EndAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-01T00:00:00Z"); return t}()),
	// 			EntityType: to.Ptr(armbillingbenefits.MaccEntityTypeContributor),
	// 			PrimaryResourceID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/resource_group_name_02/providers/Microsoft.BillingBenefits/maccs/macc_20230614"),
	// 			ProductCode: to.Ptr("0001d726-0000-0160-330f-a0b98cdbbdc4"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T00:00:00Z"); return t}()),
	// 			Status: to.Ptr(armbillingbenefits.MaccStatusScheduled),
	// 			SystemID: to.Ptr("13810867107109237"),
	// 		},
	// 		SystemData: &armbillingbenefits.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T01:01:01.1075056Z"); return t}()),
	// 			CreatedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedByType: to.Ptr(armbillingbenefits.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}

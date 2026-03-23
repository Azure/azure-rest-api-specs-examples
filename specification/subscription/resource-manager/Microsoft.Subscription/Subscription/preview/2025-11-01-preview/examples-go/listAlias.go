package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription/v2"
)

// Generated from example definition: 2025-11-01-preview/listAlias.json
func ExampleAliasClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAliasClient().NewListPager(nil)
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
		// page = armsubscription.AliasClientListResponse{
		// 	AliasListResult: armsubscription.AliasListResult{
		// 		Value: []*armsubscription.AliasResponse{
		// 			{
		// 				Name: to.Ptr("string"),
		// 				Type: to.Ptr("string"),
		// 				ID: to.Ptr("/providers/Microsoft.Subscription/aliases/dummyalias"),
		// 				Properties: &armsubscription.AliasResponseProperties{
		// 					AcceptOwnershipState: to.Ptr(armsubscription.AcceptOwnershipPending),
		// 					AcceptOwnershipURL: to.Ptr("/providers/Microsoft.Subscription/e2283d0f-acad-4904-b803-627dd74cc072/acceptOwnership"),
		// 					BillingScope: to.Ptr("/billingAccounts/af6231a7-7f8d-4fcc-a993-dd8466108d07:c663dac6-a9a5-405a-8938-cd903e12ab5b_2019_05_31/billingProfiles/QWDQ-QWHI-AUW-SJDO-DJH/invoiceSections/FEUF-EUHE-ISJ-SKDW-DJH"),
		// 					DisplayName: to.Ptr("Test Subscription"),
		// 					ProvisioningState: to.Ptr(armsubscription.ProvisioningStateAccepted),
		// 					SubscriptionID: to.Ptr("e2283d0f-acad-4904-b803-627dd74cc072"),
		// 					SubscriptionOwnerID: to.Ptr("f09b39eb-c496-482c-9ab9-afd799572f4c"),
		// 					Tags: map[string]*string{
		// 						"tag1": to.Ptr("Messi"),
		// 						"tag2": to.Ptr("Ronaldo"),
		// 						"tag3": to.Ptr("Lebron"),
		// 					},
		// 					Workload: to.Ptr(armsubscription.WorkloadProduction),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("string"),
		// 				Type: to.Ptr("string"),
		// 				ID: to.Ptr("/providers/Microsoft.Subscription/aliases/dummyalias2"),
		// 				Properties: &armsubscription.AliasResponseProperties{
		// 					AcceptOwnershipState: to.Ptr(armsubscription.AcceptOwnershipPending),
		// 					AcceptOwnershipURL: to.Ptr("/providers/Microsoft.Subscription/091c6e56-a835-4422-a082-0427308ca9ee/acceptOwnership"),
		// 					BillingScope: to.Ptr("/billingAccounts/af6231a7-7f8d-4fcc-a993-dd8466108d07:c663dac6-a9a5-405a-8938-cd903e12ab5b_2019_05_31/billingProfiles/QWDQ-QWHI-AUW-SJDO-DJH/invoiceSections/FEUF-EUHE-ISJ-SKDW-DJH"),
		// 					DisplayName: to.Ptr("Test Subscription 2"),
		// 					ProvisioningState: to.Ptr(armsubscription.ProvisioningStateAccepted),
		// 					SubscriptionID: to.Ptr("091c6e56-a835-4422-a082-0427308ca9ee"),
		// 					SubscriptionOwnerID: to.Ptr("f09b39eb-c496-482c-9ab9-afd799572f4c"),
		// 					Tags: map[string]*string{
		// 						"tag1": to.Ptr("Messi2"),
		// 						"tag2": to.Ptr("Ronaldo2"),
		// 						"tag3": to.Ptr("Lebron2"),
		// 					},
		// 					Workload: to.Ptr(armsubscription.WorkloadProduction),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

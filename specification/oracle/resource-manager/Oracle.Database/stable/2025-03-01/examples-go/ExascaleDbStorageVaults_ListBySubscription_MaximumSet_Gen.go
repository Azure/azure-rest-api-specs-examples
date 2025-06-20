package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/ExascaleDbStorageVaults_ListBySubscription_MaximumSet_Gen.json
func ExampleExascaleDbStorageVaultsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExascaleDbStorageVaultsClient().NewListBySubscriptionPager(nil)
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
		// page = armoracledatabase.ExascaleDbStorageVaultsClientListBySubscriptionResponse{
		// 	ExascaleDbStorageVaultListResult: armoracledatabase.ExascaleDbStorageVaultListResult{
		// 		Value: []*armoracledatabase.ExascaleDbStorageVault{
		// 			{
		// 				Properties: &armoracledatabase.ExascaleDbStorageVaultProperties{
		// 					AdditionalFlashCacheInPercent: to.Ptr[int32](0),
		// 					Description: to.Ptr("dmnvnnduldfmrmkkvvsdtuvmsmruxzzpsfdydgytlckutfozephjygjetrauvbdfcwmti"),
		// 					DisplayName: to.Ptr("hbsybtelyvhpalemszcvartlhwvskrnpiveqfblvkdihoytqaotdgsgauvgivzaftfgeiwlyeqzssicwrrnlxtsmeakbcsxabjlt"),
		// 					HighCapacityDatabaseStorageInput: &armoracledatabase.ExascaleDbStorageInputDetails{
		// 						TotalSizeInGbs: to.Ptr[int32](21),
		// 					},
		// 					HighCapacityDatabaseStorage: &armoracledatabase.ExascaleDbStorageDetails{
		// 						AvailableSizeInGbs: to.Ptr[int32](28),
		// 						TotalSizeInGbs: to.Ptr[int32](16),
		// 					},
		// 					TimeZone: to.Ptr("ltrbozwxjunncicrtzjrpqnqrcjgghohztrdlbfjrbkpenopyldwolslwgrgumjfkyovvkzcuxjujuxtjjzubvqvnhrswnbdgcbslopeofmtepbrrlymqwwszvsglmyuvlcuejshtpokirwklnwpcykhyinjmlqvxtyixlthtdishhmtipbygsayvgqzfrprgppylydlcskbmvwctxifdltippfvsxiughqbojqpqrekxsotnqsk"),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 					LifecycleState: to.Ptr(armoracledatabase.ExascaleDbStorageVaultLifecycleStateProvisioning),
		// 					LifecycleDetails: to.Ptr("iiav"),
		// 					VMClusterCount: to.Ptr[int32](26),
		// 					Ocid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					OciURL: to.Ptr("https://microsoft.com/a"),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("qk"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key7827": to.Ptr("xqi"),
		// 				},
		// 				Location: to.Ptr("ltguhzffucaytqg"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/exascaleDbStorageVaults/storageVaultName"),
		// 				Name: to.Ptr("s"),
		// 				Type: to.Ptr("efxpfdvhkaiqveccmdmaghlxxnmpwb"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("ilrpjodjmvzhybazxipoplnql"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("lhjbxchqkaia"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

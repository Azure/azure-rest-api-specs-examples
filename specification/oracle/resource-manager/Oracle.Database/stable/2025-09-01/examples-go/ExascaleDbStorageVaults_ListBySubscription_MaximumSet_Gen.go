package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/ExascaleDbStorageVaults_ListBySubscription_MaximumSet_Gen.json
func ExampleExascaleDbStorageVaultsClient_NewListBySubscriptionPager_exascaleDbStorageVaultsListBySubscriptionMaximumSet() {
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
		// 					Description: to.Ptr("kgqvxvtegzwyppegpvqxnlslvetbjlgveofcpjddenhbpocyzwtswaeaetqkipcxyhedsymuljalirryldlbviuvidhssyiwodacajjnbpkbvbvzwzsjctsidchalyjkievnivikwnnypaojcvhmokddstxwiqxmbfmbvglfimseguwyvibwzllggjtwejdfgezoeuvjjbsyfozswihydzuscjrqnklewongumiljeordhqlsclwlmftzdoey"),
		// 					DisplayName: to.Ptr("storagevault1"),
		// 					HighCapacityDatabaseStorageInput: &armoracledatabase.ExascaleDbStorageInputDetails{
		// 						TotalSizeInGbs: to.Ptr[int32](1),
		// 					},
		// 					HighCapacityDatabaseStorage: &armoracledatabase.ExascaleDbStorageDetails{
		// 						AvailableSizeInGbs: to.Ptr[int32](4),
		// 						TotalSizeInGbs: to.Ptr[int32](12),
		// 					},
		// 					TimeZone: to.Ptr("hyjcftlal"),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 					LifecycleState: to.Ptr(armoracledatabase.ExascaleDbStorageVaultLifecycleStateProvisioning),
		// 					LifecycleDetails: to.Ptr("mvikacxnfgannekl"),
		// 					VMClusterCount: to.Ptr[int32](10),
		// 					Ocid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					OciURL: to.Ptr("https://microsoft.com/a"),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("zone1"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key4521": to.Ptr("rrgotvwzckepkhgkbz"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/exascaleDbStorageVaults/storagevault1"),
		// 				Name: to.Ptr("storagevault1"),
		// 				Type: to.Ptr("pnac"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("sqehacivpuim"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("axrqfdkqylvjv"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

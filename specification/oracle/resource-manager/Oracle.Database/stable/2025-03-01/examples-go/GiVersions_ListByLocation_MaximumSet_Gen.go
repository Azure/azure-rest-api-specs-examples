package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/GiVersions_ListByLocation_MaximumSet_Gen.json
func ExampleGiVersionsClient_NewListByLocationPager_listGiVersionsByLocationGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGiVersionsClient().NewListByLocationPager("eastus", &armoracledatabase.GiVersionsClientListByLocationOptions{
		Shape: to.Ptr(armoracledatabase.SystemShapes("osixsklyaauhoqnkxvnvsqeqenhzogntqnpubldrrfvqncwetdtwqwjjcvspwhgecbimdlulwcubikebrdzmidrucgtsuqvytkqutmbyrvvyioxpocpmuwiivyanjzucaegihztluuvpznzaoakfsselumhhsvrtrbzwpjhcihsvyouonlxdluwhqfxoqvgthkaxppbydtqjntscgzbivfdcaobbkthrbdjwpejirqmbly"))})
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
		// page = armoracledatabase.GiVersionsClientListByLocationResponse{
		// 	GiVersionListResult: armoracledatabase.GiVersionListResult{
		// 		Value: []*armoracledatabase.GiVersion{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/giVersions/19.0.0.0"),
		// 				Type: to.Ptr("Oracle.Database/locations/giVersions"),
		// 				Properties: &armoracledatabase.GiVersionProperties{
		// 					Version: to.Ptr("19.0.0.0"),
		// 				},
		// 				Name: to.Ptr("lmktkddwvpv"),
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

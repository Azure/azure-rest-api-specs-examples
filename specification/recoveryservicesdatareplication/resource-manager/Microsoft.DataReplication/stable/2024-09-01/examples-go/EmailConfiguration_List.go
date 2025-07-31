package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/EmailConfiguration_List.json
func ExampleEmailConfigurationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEmailConfigurationClient().NewListPager("rgswagger_2024-09-01", "4", nil)
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
		// page = armrecoveryservicesdatareplication.EmailConfigurationClientListResponse{
		// 	EmailConfigurationModelListResult: armrecoveryservicesdatareplication.EmailConfigurationModelListResult{
		// 		Value: []*armrecoveryservicesdatareplication.EmailConfigurationModel{
		// 			{
		// 				Properties: &armrecoveryservicesdatareplication.EmailConfigurationModelProperties{
		// 					SendToOwners: to.Ptr(true),
		// 					CustomEmailAddresses: []*string{
		// 						to.Ptr("ketvbducyailcny"),
		// 					},
		// 					Locale: to.Ptr("vpnjxjvdqtebnucyxiyrjiko"),
		// 					ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateCanceled),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/alertSettings/emailConfiguration1"),
		// 				Name: to.Ptr("ywjplnjzaeu"),
		// 				Type: to.Ptr("bkaq"),
		// 				SystemData: &armrecoveryservicesdatareplication.SystemData{
		// 					CreatedBy: to.Ptr("ewufpudzcjrljhmmzhfnxoqdqwnya"),
		// 					CreatedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("zioqm")),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("rx"),
		// 					LastModifiedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("tqbvuqoakaaqij")),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

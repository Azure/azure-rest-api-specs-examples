package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/EmailConfiguration_Get.json
func ExampleEmailConfigurationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEmailConfigurationClient().Get(ctx, "rgswagger_2024-09-01", "4", "0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesdatareplication.EmailConfigurationClientGetResponse{
	// 	EmailConfigurationModel: &armrecoveryservicesdatareplication.EmailConfigurationModel{
	// 		Properties: &armrecoveryservicesdatareplication.EmailConfigurationModelProperties{
	// 			SendToOwners: to.Ptr(true),
	// 			CustomEmailAddresses: []*string{
	// 				to.Ptr("ketvbducyailcny"),
	// 			},
	// 			Locale: to.Ptr("vpnjxjvdqtebnucyxiyrjiko"),
	// 			ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateCanceled),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/alertSettings/emailConfiguration1"),
	// 		Name: to.Ptr("ywjplnjzaeu"),
	// 		Type: to.Ptr("bkaq"),
	// 		SystemData: &armrecoveryservicesdatareplication.SystemData{
	// 			CreatedBy: to.Ptr("ewufpudzcjrljhmmzhfnxoqdqwnya"),
	// 			CreatedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("zioqm")),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("rx"),
	// 			LastModifiedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("tqbvuqoakaaqij")),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
	// 		},
	// 	},
	// }
}

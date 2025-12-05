package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/VolumeQuotaRules_List.json
func ExampleVolumeQuotaRulesClient_NewListByVolumePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("5275316f-a498-48d6-b324-2cbfdc4311b9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeQuotaRulesClient().NewListByVolumePager("myRG", "account-9957", "pool-5210", "volume-6387", nil)
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
		// page = armnetapp.VolumeQuotaRulesClientListByVolumeResponse{
		// 	VolumeQuotaRulesList: armnetapp.VolumeQuotaRulesList{
		// 		Value: []*armnetapp.VolumeQuotaRule{
		// 			{
		// 				Name: to.Ptr("account-9957/pool-5210/volume-6387/rule-0004"),
		// 				Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/volumeQuotaRules"),
		// 				ID: to.Ptr("/subscriptions/5275316f-a498-48d6-b324-2cbfdc4311b9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account-9957/capacityPools/pool-5210/volumes/volume-6387/volumeQuotaRules/rule-0004"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetapp.VolumeQuotaRulesProperties{
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					QuotaSizeInKiBs: to.Ptr[int64](100005),
		// 					QuotaTarget: to.Ptr("1821"),
		// 					QuotaType: to.Ptr(armnetapp.TypeIndividualUserQuota),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/VolumeQuotaRules_Get.json
func ExampleVolumeQuotaRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeQuotaRulesClient().Get(ctx, "myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VolumeQuotaRule = armnetapp.VolumeQuotaRule{
	// 	Name: to.Ptr("account-9957/pool-5210/volume-6387/rule-0004"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/volumeQuotaRules"),
	// 	ID: to.Ptr("/subscriptions/5275316f-a498-48d6-b324-2cbfdc4311b9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account-9957/capacityPools/pool-5210/volumes/volume-6387/volumeQuotaRules/rule-0004"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetapp.VolumeQuotaRulesProperties{
	// 		ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 		QuotaSizeInKiBs: to.Ptr[int64](100005),
	// 		QuotaTarget: to.Ptr("1821"),
	// 		QuotaType: to.Ptr(armnetapp.TypeIndividualUserQuota),
	// 	},
	// }
}

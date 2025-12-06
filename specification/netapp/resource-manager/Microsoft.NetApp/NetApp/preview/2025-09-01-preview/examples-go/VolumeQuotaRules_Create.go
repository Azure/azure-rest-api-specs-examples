package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/VolumeQuotaRules_Create.json
func ExampleVolumeQuotaRulesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("5275316f-a498-48d6-b324-2cbfdc4311b9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeQuotaRulesClient().BeginCreate(ctx, "myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004", armnetapp.VolumeQuotaRule{
		Location: to.Ptr("westus"),
		Properties: &armnetapp.VolumeQuotaRulesProperties{
			QuotaSizeInKiBs: to.Ptr[int64](100005),
			QuotaTarget:     to.Ptr("1821"),
			QuotaType:       to.Ptr(armnetapp.TypeIndividualUserQuota),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.VolumeQuotaRulesClientCreateResponse{
	// 	VolumeQuotaRule: &armnetapp.VolumeQuotaRule{
	// 		Name: to.Ptr("account-9957/pool-5210/volume-6387/rule-0004"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/volumeQuotaRules"),
	// 		ID: to.Ptr("/subscriptions/5275316f-a498-48d6-b324-2cbfdc4311b9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account-9957/capacityPools/pool-5210/volumes/volume-6387/volumeQuotaRules/rule-0004"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetapp.VolumeQuotaRulesProperties{
	// 			ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 			QuotaSizeInKiBs: to.Ptr[int64](100005),
	// 			QuotaTarget: to.Ptr("1821"),
	// 			QuotaType: to.Ptr(armnetapp.TypeIndividualUserQuota),
	// 		},
	// 	},
	// }
}

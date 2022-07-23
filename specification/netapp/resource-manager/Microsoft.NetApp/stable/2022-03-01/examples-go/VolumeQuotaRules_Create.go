package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-03-01/examples/VolumeQuotaRules_Create.json
func ExampleVolumeQuotaRulesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewVolumeQuotaRulesClient("5275316f-a498-48d6-b324-2cbfdc4311b9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"myRG",
		"account-9957",
		"pool-5210",
		"volume-6387",
		"rule-0004",
		armnetapp.VolumeQuotaRule{
			Location: to.Ptr("westus"),
			Properties: &armnetapp.VolumeQuotaRulesProperties{
				QuotaSizeInKiBs: to.Ptr[int64](100005),
				QuotaTarget:     to.Ptr("1821"),
				QuotaType:       to.Ptr(armnetapp.TypeIndividualUserQuota),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/changeTenantPolicy.json
func ExamplePolicyClient_AddUpdatePolicyForTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsubscription.NewPolicyClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.AddUpdatePolicyForTenant(ctx,
		armsubscription.PutTenantPolicyRequestProperties{
			BlockSubscriptionsIntoTenant:    to.Ptr(true),
			BlockSubscriptionsLeavingTenant: to.Ptr(true),
			ExemptedPrincipals: []*string{
				to.Ptr("e879cf0f-2b4d-5431-109a-f72fc9868693"),
				to.Ptr("9792da87-c97b-410d-a97d-27021ba09ce6")},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

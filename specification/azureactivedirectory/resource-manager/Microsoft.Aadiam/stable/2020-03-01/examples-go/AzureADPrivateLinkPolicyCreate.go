package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateLinkPolicyCreate.json
func ExamplePrivateLinkForAzureAdClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateLinkForAzureAdClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"ddb1",
		armaad.PrivateLinkPolicy{
			Name:           to.Ptr("myOrgPrivateLinkPolicy"),
			AllTenants:     to.Ptr(false),
			OwnerTenantID:  to.Ptr("950f8bca-bf4d-4a41-ad10-034e792a243d"),
			ResourceGroup:  to.Ptr("myOrgVnetRG"),
			ResourceName:   to.Ptr("myOrgVnetPrivateLink"),
			SubscriptionID: to.Ptr("57849194-ea1f-470b-abda-d195b25634c1"),
			Tenants: []*string{
				to.Ptr("3616657d-1c80-41ae-9d83-2a2776f2c9be"),
				to.Ptr("727b6ef1-18ab-4627-ac95-3f9cd945ed87")},
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

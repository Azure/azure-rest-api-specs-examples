package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateLinkPolicyMinCreate.json
func ExamplePrivateLinkForAzureAdClient_BeginCreate_privateLinkPolicyMinCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armaad.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkForAzureAdClient().BeginCreate(ctx, "rg1", "ddb1", armaad.PrivateLinkPolicy{
		Name:           to.Ptr("myOrgPrivateLinkPolicy"),
		AllTenants:     to.Ptr(false),
		OwnerTenantID:  to.Ptr("950f8bca-bf4d-4a41-ad10-034e792a243d"),
		ResourceGroup:  to.Ptr("myOrgVnetRG"),
		ResourceName:   to.Ptr("myOrgVnetPrivateLink"),
		SubscriptionID: to.Ptr("57849194-ea1f-470b-abda-d195b25634c1"),
		Tenants: []*string{
			to.Ptr("3616657d-1c80-41ae-9d83-2a2776f2c9be"),
			to.Ptr("727b6ef1-18ab-4627-ac95-3f9cd945ed87")},
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
	// res.PrivateLinkPolicy = armaad.PrivateLinkPolicy{
	// 	Name: to.Ptr("myOrgPrivateLinkPolicy"),
	// 	AllTenants: to.Ptr(false),
	// 	OwnerTenantID: to.Ptr("950f8bca-bf4d-4a41-ad10-034e792a243d"),
	// 	ResourceGroup: to.Ptr("myOrgVnetRG"),
	// 	ResourceName: to.Ptr("myOrgVnetPrivateLink"),
	// 	SubscriptionID: to.Ptr("57849194-ea1f-470b-abda-d195b25634c1"),
	// 	Tenants: []*string{
	// 		to.Ptr("3616657d-1c80-41ae-9d83-2a2776f2c9be"),
	// 		to.Ptr("727b6ef1-18ab-4627-ac95-3f9cd945ed87")},
	// 	}
}

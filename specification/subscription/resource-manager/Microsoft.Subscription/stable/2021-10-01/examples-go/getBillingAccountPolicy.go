package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/getBillingAccountPolicy.json
func ExampleBillingAccountClient_GetPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBillingAccountClient().GetPolicy(ctx, "testBillingAccountId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BillingAccountPoliciesResponse = armsubscription.BillingAccountPoliciesResponse{
	// 	Name: to.Ptr("testBillingAccountId"),
	// 	Type: to.Ptr("Microsoft.Subscription/policies"),
	// 	ID: to.Ptr("/providers/Microsoft.Subscription/Policies/policyForBillingAccount"),
	// 	Properties: &armsubscription.BillingAccountPoliciesResponseProperties{
	// 		AllowTransfers: to.Ptr(true),
	// 		ServiceTenants: []*armsubscription.ServiceTenantResponse{
	// 			{
	// 				TenantID: to.Ptr("b8ed2088-c458-4e77-bd61-9e048d96a1c0"),
	// 				TenantName: to.Ptr("testServiceTenant"),
	// 		}},
	// 	},
	// }
}

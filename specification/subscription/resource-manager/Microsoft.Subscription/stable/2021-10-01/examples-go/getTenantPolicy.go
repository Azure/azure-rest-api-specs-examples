package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/getTenantPolicy.json
func ExamplePolicyClient_GetPolicyForTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyClient().GetPolicyForTenant(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GetTenantPolicyResponse = armsubscription.GetTenantPolicyResponse{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("providers/Microsoft.Subscription/policies"),
	// 	ID: to.Ptr("providers/Microsoft.Subscription/policies/default"),
	// 	Properties: &armsubscription.TenantPolicy{
	// 		BlockSubscriptionsIntoTenant: to.Ptr(true),
	// 		BlockSubscriptionsLeavingTenant: to.Ptr(true),
	// 		ExemptedPrincipals: []*string{
	// 			to.Ptr("e879cf0f-2b4d-5431-109a-f72fc9868693"),
	// 			to.Ptr("9792da87-c97b-410d-a97d-27021ba09ce6")},
	// 			PolicyID: to.Ptr("291bba3f-e0a5-47bc-a099-3bdcb2a50a05"),
	// 		},
	// 	}
}

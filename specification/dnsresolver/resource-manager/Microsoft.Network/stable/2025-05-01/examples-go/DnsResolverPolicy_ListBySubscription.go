package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d0d3a9b4fe0fce880fded7a617e71f84406bacbd/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolverPolicy_ListBySubscription.json
func ExamplePoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPoliciesClient().NewListPager(&armdnsresolver.PoliciesClientListOptions{Top: nil})
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
		// page.PolicyListResult = armdnsresolver.PolicyListResult{
		// 	Value: []*armdnsresolver.Policy{
		// 		{
		// 			Name: to.Ptr("sampleDnsResolverPolicy1"),
		// 			Type: to.Ptr("Microsoft.Network/dnsResolverPolicies"),
		// 			ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolverPolicies/sampleDnsResolverPolicy1"),
		// 			SystemData: &armdnsresolver.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.107Z"); return t}()),
		// 				CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.197Z"); return t}()),
		// 				LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armdnsresolver.PolicyProperties{
		// 				ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("ad9c8da4-3bb2-4821-a878-c2cb07b01fb6"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleDnsResolverPolicy2"),
		// 			Type: to.Ptr("Microsoft.Network/dnsResolverPolicies"),
		// 			ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolverPolicies/sampleDnsResolverPolicy2"),
		// 			SystemData: &armdnsresolver.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.107Z"); return t}()),
		// 				CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.197Z"); return t}()),
		// 				LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armdnsresolver.PolicyProperties{
		// 				ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("b6b2d964-8588-4e3a-a7fe-8a5b7fe8eca5"),
		// 			},
		// 	}},
		// }
	}
}

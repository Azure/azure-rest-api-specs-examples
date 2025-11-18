package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver/v2"
)

// Generated from example definition: 2025-10-01-preview/DnsResolver_ListByResourceGroup.json
func ExampleDNSResolversClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDNSResolversClient().NewListByResourceGroupPager("sampleResourceGroup", nil)
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
		// page = armdnsresolver.DNSResolversClientListByResourceGroupResponse{
		// 	ListResult: armdnsresolver.ListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/api/mresolver/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers?$skipToken=skipToken&api-version=2025-10-01-preview"),
		// 		Value: []*armdnsresolver.DNSResolver{
		// 			{
		// 				Name: to.Ptr("sampleDnsResolver1"),
		// 				Type: to.Ptr("Microsoft.Network/dnsResolvers"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver1"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armdnsresolver.Properties{
		// 					DNSResolverState: to.Ptr(armdnsresolver.DNSResolverStateConnected),
		// 					ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("ad9c8da4-3bb2-4821-a878-c2cb07b01fb6"),
		// 					VirtualNetwork: &armdnsresolver.SubResource{
		// 						ID: to.Ptr("/subscriptions/cbb1387e-4b03-44f2-ad41-58d4677b9873/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork1"),
		// 					},
		// 				},
		// 				SystemData: &armdnsresolver.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.1075056Z"); return t}()),
		// 					CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.1974346Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sampleDnsResolver2"),
		// 				Type: to.Ptr("Microsoft.Network/dnsResolvers"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver2"),
		// 				Location: to.Ptr("westus2"),
		// 				Properties: &armdnsresolver.Properties{
		// 					DNSResolverState: to.Ptr(armdnsresolver.DNSResolverStateConnected),
		// 					ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("b6b2d964-8588-4e3a-a7fe-8a5b7fe8eca5"),
		// 					VirtualNetwork: &armdnsresolver.SubResource{
		// 						ID: to.Ptr("/subscriptions/cbb1387e-4b03-44f2-ad41-58d4677b9873/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork2"),
		// 					},
		// 				},
		// 				SystemData: &armdnsresolver.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-03T01:01:01.1075056Z"); return t}()),
		// 					CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-04T02:03:01.1974346Z"); return t}()),
		// 					LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

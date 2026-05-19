package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/CustomDomains_ListByEndpoint.json
func ExampleCustomDomainsClient_NewListByEndpointPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomDomainsClient().NewListByEndpointPager("RG", "profile1", "endpoint1", nil)
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
		// page = armcdn.CustomDomainsClientListByEndpointResponse{
		// 	CustomDomainListResult: armcdn.CustomDomainListResult{
		// 		Value: []*armcdn.CustomDomain{
		// 			{
		// 				Name: to.Ptr("www-someDomain-net"),
		// 				Type: to.Ptr("Microsoft.Cdn/profiles/endpoints/customdomains"),
		// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customdomains/www-someDomain-net"),
		// 				Properties: &armcdn.CustomDomainProperties{
		// 					CustomHTTPSProvisioningState: to.Ptr(armcdn.CustomHTTPSProvisioningStateDisabled),
		// 					CustomHTTPSProvisioningSubstate: to.Ptr(armcdn.CustomHTTPSProvisioningSubstate("None")),
		// 					HostName: to.Ptr("www.someDomain.net"),
		// 					ProvisioningState: to.Ptr(armcdn.CustomHTTPSProvisioningState("Succeeded")),
		// 					ResourceState: to.Ptr(armcdn.CustomDomainResourceStateActive),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

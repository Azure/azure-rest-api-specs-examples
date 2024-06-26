package armsubscriptions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4f4073bdb028bc84bc3e6405c1cbaf8e89b83caf/specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/GetTenants.json
func ExampleTenantsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscriptions.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTenantsClient().NewListPager(nil)
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
		// page.TenantListResult = armsubscriptions.TenantListResult{
		// 	Value: []*armsubscriptions.TenantIDDescription{
		// 		{
		// 			CountryCode: to.Ptr("US"),
		// 			DefaultDomain: to.Ptr("aad50.ccsctp.net"),
		// 			DisplayName: to.Ptr("Test_Test_aad50"),
		// 			Domains: []*string{
		// 				to.Ptr("aad50.ccsctp.net")},
		// 				ID: to.Ptr("/tenants/a70a1586-9c4a-4373-b907-1d310660dbd1"),
		// 				TenantCategory: to.Ptr(armsubscriptions.TenantCategoryManagedBy),
		// 				TenantID: to.Ptr("a70a1586-9c4a-4373-b907-1d310660dbd1"),
		// 				TenantType: to.Ptr("AAD"),
		// 			},
		// 			{
		// 				CountryCode: to.Ptr("US"),
		// 				DefaultDomain: to.Ptr("auxteststagemanual.ccsctp.net"),
		// 				DisplayName: to.Ptr("Contoso Corp."),
		// 				Domains: []*string{
		// 					to.Ptr("auxteststagemanual.ccsctp.net")},
		// 					ID: to.Ptr("/tenants/83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
		// 					TenantCategory: to.Ptr(armsubscriptions.TenantCategoryHome),
		// 					TenantID: to.Ptr("83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
		// 					TenantType: to.Ptr("AAD"),
		// 				},
		// 				{
		// 					CountryCode: to.Ptr("US"),
		// 					DefaultDomain: to.Ptr("rdvmohoro.ccsctp.net"),
		// 					DisplayName: to.Ptr("TEST_TEST_RDV"),
		// 					Domains: []*string{
		// 						to.Ptr("rdvmohoro.ccsctp.net"),
		// 						to.Ptr("rdvmohoro.mail.ccsctp.net"),
		// 						to.Ptr("rdvmohoro.com")},
		// 						ID: to.Ptr("/tenants/daea2e9b-847b-4c93-850d-2aa6f2d7af33"),
		// 						TenantBrandingLogoURL: to.Ptr("logo1.logo.rdvmohoro.com"),
		// 						TenantCategory: to.Ptr(armsubscriptions.TenantCategoryProjectedBy),
		// 						TenantID: to.Ptr("daea2e9b-847b-4c93-850d-2aa6f2d7af33"),
		// 						TenantType: to.Ptr("AAD"),
		// 				}},
		// 			}
	}
}

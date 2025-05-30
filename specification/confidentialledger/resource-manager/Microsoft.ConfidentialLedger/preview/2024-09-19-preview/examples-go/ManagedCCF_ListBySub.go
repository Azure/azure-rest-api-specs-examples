package armconfidentialledger_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/40c12a2dd76da721d480d55a44e6ec666045d508/specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2024-09-19-preview/examples/ManagedCCF_ListBySub.json
func ExampleManagedCCFClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfidentialledger.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedCCFClient().NewListBySubscriptionPager(&armconfidentialledger.ManagedCCFClientListBySubscriptionOptions{Filter: nil})
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
		// page.ManagedCCFList = armconfidentialledger.ManagedCCFList{
		// 	Value: []*armconfidentialledger.ManagedCCF{
		// 		{
		// 			Name: to.Ptr("DummyLedgerName"),
		// 			Type: to.Ptr("Microsoft.ConfidentialLedger/ManagedCCFs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/providers/Microsoft.ConfidentialLedger/ManagedCCFs/DummyLedgerName"),
		// 			SystemData: &armconfidentialledger.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T00:00:00.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("ledgerAdmin@contoso.com"),
		// 				CreatedByType: to.Ptr(armconfidentialledger.CreatedByType("Admin1")),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-02T00:00:00.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("ledgerAdmin2@outlook.com"),
		// 				LastModifiedByType: to.Ptr(armconfidentialledger.CreatedByType("Admin2")),
		// 			},
		// 			Location: to.Ptr("EastUS"),
		// 			Tags: map[string]*string{
		// 				"additionalProps1": to.Ptr("additional properties"),
		// 			},
		// 			Properties: &armconfidentialledger.ManagedCCFProperties{
		// 				AppName: to.Ptr("DummyMccfAppName"),
		// 				AppURI: to.Ptr("https://dummy.accledger.domain.com/DummyLedgerName"),
		// 				DeploymentType: &armconfidentialledger.DeploymentType{
		// 					AppSourceURI: to.Ptr("https://myaccount.blob.core.windows.net/storage/mccfsource?sv=2022-02-11%st=2022-03-11"),
		// 					LanguageRuntime: to.Ptr(armconfidentialledger.LanguageRuntimeCPP),
		// 				},
		// 				IdentityServiceURI: to.Ptr("https://dummy.accledger.identity.com/DummyLedgerName"),
		// 				MemberIdentityCertificates: []*armconfidentialledger.MemberIdentityCertificate{
		// 					{
		// 						Certificate: to.Ptr("-----BEGIN CERTIFICATE-----MIIBsjCCATigAwIBAgIUZWIbyG79TniQLd2UxJuU74tqrKcwCgYIKoZIzj0EAwMwEDEOMAwGA1UEAwwFdXNlcjAwHhcNMjEwMzE2MTgwNjExWhcNMjIwMzE2MTgwNjExWjAQMQ4wDAYDVQQDDAV1c2VyMDB2MBAGByqGSM49AgEGBSuBBAAiA2IABBiWSo/j8EFit7aUMm5lF+lUmCu+IgfnpFD+7QMgLKtxRJ3aGSqgS/GpqcYVGddnODtSarNE/HyGKUFUolLPQ5ybHcouUk0kyfA7XMeSoUA4lBz63Wha8wmXo+NdBRo39qNTMFEwHQYDVR0OBBYEFPtuhrwgGjDFHeUUT4nGsXaZn69KMB8GA1UdIwQYMBaAFPtuhrwgGjDFHeUUT4nGsXaZn69KMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwMDaAAwZQIxAOnozm2CyqRwSSQLls5r+mUHRGRyXHXwYtM4Dcst/VEZdmS9fqvHRCHbjUlO/+HNfgIwMWZ4FmsjD3wnPxONOm9YdVn/PRD7SsPRPbOjwBiE4EBGaHDsLjYAGDSGi7NJnSkA-----END CERTIFICATE-----"),
		// 						Encryptionkey: to.Ptr("ledgerencryptionkey"),
		// 						Tags: map[string]any{
		// 							"additionalProps1": "additional properties",
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armconfidentialledger.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

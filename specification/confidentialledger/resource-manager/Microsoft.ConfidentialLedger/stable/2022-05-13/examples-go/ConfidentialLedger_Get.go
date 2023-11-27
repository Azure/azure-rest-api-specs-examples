package armconfidentialledger_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/792db17291c758b2bfdbbc0d35d0e2f5b5a1bd05/specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/ConfidentialLedger_Get.json
func ExampleLedgerClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfidentialledger.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLedgerClient().Get(ctx, "DummyResourceGroupName", "DummyLedgerName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfidentialLedger = armconfidentialledger.ConfidentialLedger{
	// 	Name: to.Ptr("DummyLedgerName"),
	// 	Type: to.Ptr("Microsoft.ConfidentialLedger/ledgers"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001/providers/Microsoft.ConfidentialLedger/ledgers/DummyLedgerName"),
	// 	SystemData: &armconfidentialledger.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T00:00:00.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("ledgerAdmin@contoso.com"),
	// 		CreatedByType: to.Ptr(armconfidentialledger.CreatedByType("Admin1")),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-02T00:00:00.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("ledgerAdmin2@outlook.com"),
	// 		LastModifiedByType: to.Ptr(armconfidentialledger.CreatedByType("Admin2")),
	// 	},
	// 	Location: to.Ptr("EastUS"),
	// 	Tags: map[string]*string{
	// 		"additionalProps1": to.Ptr("additional properties"),
	// 	},
	// 	Properties: &armconfidentialledger.LedgerProperties{
	// 		AADBasedSecurityPrincipals: []*armconfidentialledger.AADBasedSecurityPrincipal{
	// 			{
	// 				LedgerRoleName: to.Ptr(armconfidentialledger.LedgerRoleNameAdministrator),
	// 				PrincipalID: to.Ptr("34621747-6fc8-4771-a2eb-72f31c461f2e"),
	// 				TenantID: to.Ptr("bce123b9-2b7b-4975-8360-5ca0b9b1cd08"),
	// 		}},
	// 		IdentityServiceURI: to.Ptr("https://dummy.accledger.identity.com/DummyLedgerName"),
	// 		LedgerInternalNamespace: to.Ptr("dummyNamespace"),
	// 		LedgerName: to.Ptr("DummyLedgerName"),
	// 		LedgerType: to.Ptr(armconfidentialledger.LedgerTypePublic),
	// 		LedgerURI: to.Ptr("https://dummy.accledger.domain.com/DummyLedgerName"),
	// 		ProvisioningState: to.Ptr(armconfidentialledger.ProvisioningStateSucceeded),
	// 	},
	// }
}

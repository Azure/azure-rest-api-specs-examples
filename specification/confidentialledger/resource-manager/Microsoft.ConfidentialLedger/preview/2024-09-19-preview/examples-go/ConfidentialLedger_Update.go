package armconfidentialledger_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/40c12a2dd76da721d480d55a44e6ec666045d508/specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2024-09-19-preview/examples/ConfidentialLedger_Update.json
func ExampleLedgerClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfidentialledger.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLedgerClient().BeginUpdate(ctx, "DummyResourceGroupName", "DummyLedgerName", armconfidentialledger.ConfidentialLedger{
		Location: to.Ptr("EastUS"),
		Tags: map[string]*string{
			"additionProps2":   to.Ptr("additional property value"),
			"additionalProps1": to.Ptr("additional properties"),
		},
		Properties: &armconfidentialledger.LedgerProperties{
			AADBasedSecurityPrincipals: []*armconfidentialledger.AADBasedSecurityPrincipal{
				{
					LedgerRoleName: to.Ptr(armconfidentialledger.LedgerRoleNameAdministrator),
					PrincipalID:    to.Ptr("34621747-6fc8-4771-a2eb-72f31c461f2e"),
					TenantID:       to.Ptr("bce123b9-2b7b-4975-8360-5ca0b9b1cd08"),
				}},
			CertBasedSecurityPrincipals: []*armconfidentialledger.CertBasedSecurityPrincipal{
				{
					Cert:           to.Ptr("-----BEGIN CERTIFICATE-----\nMIIDUjCCAjqgAwIBAgIQJ2IrDBawSkiAbkBYmiAopDANBgkqhkiG9w0BAQsFADAmMSQwIgYDVQQDExtTeW50aGV0aWNzIExlZGdlciBVc2VyIENlcnQwHhcNMjAwOTIzMjIxODQ2WhcNMjEwOTIzMjIyODQ2WjAmMSQwIgYDVQQDExtTeW50aGV0aWNzIExlZGdlciBVc2VyIENlcnQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCX2s/Eu4q/eQ63N+Ugeg5oAciZua/YCJr41c/696szvSY7Zg1SNJlW88/nbz70+QpO55OmqlEE3QCU+T0Vl/h0Gf//n1PYcoBbTGUnYEmV+fTTHict6rFiEwrGJ62tvcpYgwapInSLyEeUzjki0zhOLJ1OfRnYd1eGnFVMpE5aVjiS8Q5dmTEUyd51EIprGE8RYAW9aeWSwTH7gjHUsRlJnHKcdhaK/v5QKJnNu5bzPFUcpC0ZBcizoMPAtroLAD4B68Jl0z3op18MgZe6lRrVoWuxfqnk5GojuB/Vu8ohAZKoFhQ6NB6r+LL2AUs+Zr7Bt26IkEdR178n9JMEA4gHAgMBAAGjfDB6MA4GA1UdDwEB/wQEAwIFoDAJBgNVHRMEAjAAMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBS/a7PU9iOfOKEyZCp11Oen5VSuuDAdBgNVHQ4EFgQUv2uz1PYjnzihMmQqddTnp+VUrrgwDQYJKoZIhvcNAQELBQADggEBAF5q2fDwnse8egXhfaJCqqM969E9gSacqFmASpoDJPRPEX7gqoO7v1ww7nqRtRDoRiBvo/yNk7jlSAkRN3nRRnZLZZ3MYQdmCr4FGyIqRg4Y94+nja+Du9pDD761rxRktMVPSOaAVM/E5DQvscDlPvlPYe9mkcrLCE4DXYpiMmLT8Tm55LJJq5m07dVDgzAIR1L/hmEcbK0pnLgzciMtMLxGO2udnyyW/UW9WxnjvrrD2JluTHH9mVbb+XQP1oFtlRBfH7aui1ZgWfKvxrdP4zdK9QoWSUvRux3TLsGmHRBjBMtqYDY3y5mB+aNjLelvWpeVb0m2aOSVXynrLwNCAVA=\n-----END CERTIFICATE-----"),
					LedgerRoleName: to.Ptr(armconfidentialledger.LedgerRoleNameReader),
				}},
			LedgerType: to.Ptr(armconfidentialledger.LedgerTypePublic),
		},
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
	// 		"additionProps2": to.Ptr("additional property value"),
	// 		"additionalProps1": to.Ptr("additional properties"),
	// 	},
	// 	Properties: &armconfidentialledger.LedgerProperties{
	// 		AADBasedSecurityPrincipals: []*armconfidentialledger.AADBasedSecurityPrincipal{
	// 			{
	// 				LedgerRoleName: to.Ptr(armconfidentialledger.LedgerRoleNameAdministrator),
	// 				PrincipalID: to.Ptr("34621747-6fc8-4771-a2eb-72f31c461f2e"),
	// 				TenantID: to.Ptr("bce123b9-2b7b-4975-8360-5ca0b9b1cd08"),
	// 		}},
	// 		CertBasedSecurityPrincipals: []*armconfidentialledger.CertBasedSecurityPrincipal{
	// 			{
	// 				Cert: to.Ptr("-----BEGIN CERTIFICATE-----\nMIIDUjCCAjqgAwIBAgIQJ2IrDBawSkiAbkBYmiAopDANBgkqhkiG9w0BAQsFADAmMSQwIgYDVQQDExtTeW50aGV0aWNzIExlZGdlciBVc2VyIENlcnQwHhcNMjAwOTIzMjIxODQ2WhcNMjEwOTIzMjIyODQ2WjAmMSQwIgYDVQQDExtTeW50aGV0aWNzIExlZGdlciBVc2VyIENlcnQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCX2s/Eu4q/eQ63N+Ugeg5oAciZua/YCJr41c/696szvSY7Zg1SNJlW88/nbz70+QpO55OmqlEE3QCU+T0Vl/h0Gf//n1PYcoBbTGUnYEmV+fTTHict6rFiEwrGJ62tvcpYgwapInSLyEeUzjki0zhOLJ1OfRnYd1eGnFVMpE5aVjiS8Q5dmTEUyd51EIprGE8RYAW9aeWSwTH7gjHUsRlJnHKcdhaK/v5QKJnNu5bzPFUcpC0ZBcizoMPAtroLAD4B68Jl0z3op18MgZe6lRrVoWuxfqnk5GojuB/Vu8ohAZKoFhQ6NB6r+LL2AUs+Zr7Bt26IkEdR178n9JMEA4gHAgMBAAGjfDB6MA4GA1UdDwEB/wQEAwIFoDAJBgNVHRMEAjAAMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBS/a7PU9iOfOKEyZCp11Oen5VSuuDAdBgNVHQ4EFgQUv2uz1PYjnzihMmQqddTnp+VUrrgwDQYJKoZIhvcNAQELBQADggEBAF5q2fDwnse8egXhfaJCqqM969E9gSacqFmASpoDJPRPEX7gqoO7v1ww7nqRtRDoRiBvo/yNk7jlSAkRN3nRRnZLZZ3MYQdmCr4FGyIqRg4Y94+nja+Du9pDD761rxRktMVPSOaAVM/E5DQvscDlPvlPYe9mkcrLCE4DXYpiMmLT8Tm55LJJq5m07dVDgzAIR1L/hmEcbK0pnLgzciMtMLxGO2udnyyW/UW9WxnjvrrD2JluTHH9mVbb+XQP1oFtlRBfH7aui1ZgWfKvxrdP4zdK9QoWSUvRux3TLsGmHRBjBMtqYDY3y5mB+aNjLelvWpeVb0m2aOSVXynrLwNCAVA=\n-----END CERTIFICATE-----"),
	// 				LedgerRoleName: to.Ptr(armconfidentialledger.LedgerRoleNameReader),
	// 		}},
	// 		IdentityServiceURI: to.Ptr("https://identity.confidential-ledger.core.azure.com/ledgerIdentity/dummyledgername"),
	// 		LedgerInternalNamespace: to.Ptr("dummyNamespace"),
	// 		LedgerName: to.Ptr("DummyLedgerName"),
	// 		LedgerSKU: to.Ptr(armconfidentialledger.LedgerSKUStandard),
	// 		LedgerType: to.Ptr(armconfidentialledger.LedgerTypePublic),
	// 		LedgerURI: to.Ptr("https://dummyledgername.confidential-ledger.azure.com"),
	// 		ProvisioningState: to.Ptr(armconfidentialledger.ProvisioningStateSucceeded),
	// 		RunningState: to.Ptr(armconfidentialledger.RunningStateActive),
	// 	},
	// }
}

package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/AFDCustomDomains_ListByProfile.json
func ExampleAFDCustomDomainsClient_NewListByProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAFDCustomDomainsClient().NewListByProfilePager("RG", "profile1", nil)
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
		// page = armcdn.AFDCustomDomainsClientListByProfileResponse{
		// 	AFDDomainListResult: armcdn.AFDDomainListResult{
		// 		Value: []*armcdn.AFDDomain{
		// 			{
		// 				Name: to.Ptr("domain1"),
		// 				Type: to.Ptr("Microsoft.Cdn/profiles/customdomains"),
		// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/domain1"),
		// 				Properties: &armcdn.AFDDomainProperties{
		// 					AzureDNSZone: &armcdn.ResourceReference{
		// 						ID: to.Ptr(""),
		// 					},
		// 					DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 					DomainValidationState: to.Ptr(armcdn.DomainValidationStatePending),
		// 					HostName: to.Ptr("www.contoso.com"),
		// 					PreValidatedCustomDomainResourceID: &armcdn.ResourceReference{
		// 						ID: to.Ptr(""),
		// 					},
		// 					ProfileName: to.Ptr("profile1"),
		// 					ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateUpdating),
		// 					TLSSettings: &armcdn.AFDDomainHTTPSParameters{
		// 						CertificateType: to.Ptr(armcdn.AfdCertificateTypeManagedCertificate),
		// 						CipherSuiteSetType: to.Ptr(armcdn.AfdCipherSuiteSetTypeCustomized),
		// 						CustomizedCipherSuiteSet: &armcdn.AFDDomainHTTPSCustomizedCipherSuiteSet{
		// 							CipherSuiteSetForTls12: []*armcdn.AfdCustomizedCipherSuiteForTls12{
		// 								to.Ptr(armcdn.AfdCustomizedCipherSuiteForTls12ECDHERSAAES128GCMSHA256),
		// 							},
		// 							CipherSuiteSetForTls13: []*armcdn.AfdCustomizedCipherSuiteForTls13{
		// 								to.Ptr(armcdn.AfdCustomizedCipherSuiteForTls13TLSAES128GCMSHA256),
		// 								to.Ptr(armcdn.AfdCustomizedCipherSuiteForTls13TLSAES256GCMSHA384),
		// 							},
		// 						},
		// 						MinimumTLSVersion: to.Ptr(armcdn.AfdMinimumTLSVersionTLS12),
		// 						Secret: &armcdn.ResourceReference{
		// 							ID: to.Ptr(""),
		// 						},
		// 					},
		// 					ValidationProperties: &armcdn.DomainValidationProperties{
		// 						ExpirationDate: to.Ptr("2009-06-15T13:45:43.0000000Z"),
		// 						ValidationToken: to.Ptr("8c9912db-c615-4eeb-8465"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

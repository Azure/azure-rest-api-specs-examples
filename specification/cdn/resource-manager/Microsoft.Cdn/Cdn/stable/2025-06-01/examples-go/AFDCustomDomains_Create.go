package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/AFDCustomDomains_Create.json
func ExampleAFDCustomDomainsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAFDCustomDomainsClient().BeginCreate(ctx, "RG", "profile1", "domain1", armcdn.AFDDomain{
		Properties: &armcdn.AFDDomainProperties{
			AzureDNSZone: &armcdn.ResourceReference{
				ID: to.Ptr(""),
			},
			HostName: to.Ptr("www.someDomain.net"),
			TLSSettings: &armcdn.AFDDomainHTTPSParameters{
				CertificateType:    to.Ptr(armcdn.AfdCertificateTypeManagedCertificate),
				CipherSuiteSetType: to.Ptr(armcdn.AfdCipherSuiteSetTypeCustomized),
				CustomizedCipherSuiteSet: &armcdn.AFDDomainHTTPSCustomizedCipherSuiteSet{
					CipherSuiteSetForTls12: []*armcdn.AfdCustomizedCipherSuiteForTls12{
						to.Ptr(armcdn.AfdCustomizedCipherSuiteForTls12ECDHERSAAES128GCMSHA256),
					},
					CipherSuiteSetForTls13: []*armcdn.AfdCustomizedCipherSuiteForTls13{
						to.Ptr(armcdn.AfdCustomizedCipherSuiteForTls13TLSAES128GCMSHA256),
						to.Ptr(armcdn.AfdCustomizedCipherSuiteForTls13TLSAES256GCMSHA384),
					},
				},
				MinimumTLSVersion: to.Ptr(armcdn.AfdMinimumTLSVersionTLS12),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcdn.AFDCustomDomainsClientCreateResponse{
	// 	AFDDomain: armcdn.AFDDomain{
	// 		Name: to.Ptr("domain1"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/customdomains"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/domain1"),
	// 		Properties: &armcdn.AFDDomainProperties{
	// 			AzureDNSZone: &armcdn.ResourceReference{
	// 				ID: to.Ptr(""),
	// 			},
	// 			DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 			DomainValidationState: to.Ptr(armcdn.DomainValidationStateSubmitting),
	// 			HostName: to.Ptr("www.contoso.com"),
	// 			PreValidatedCustomDomainResourceID: &armcdn.ResourceReference{
	// 				ID: to.Ptr(""),
	// 			},
	// 			ProfileName: to.Ptr("profile1"),
	// 			ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 			TLSSettings: &armcdn.AFDDomainHTTPSParameters{
	// 				CertificateType: to.Ptr(armcdn.AfdCertificateTypeManagedCertificate),
	// 				CipherSuiteSetType: to.Ptr(armcdn.AfdCipherSuiteSetTypeCustomized),
	// 				CustomizedCipherSuiteSet: &armcdn.AFDDomainHTTPSCustomizedCipherSuiteSet{
	// 					CipherSuiteSetForTls12: []*armcdn.AfdCustomizedCipherSuiteForTls12{
	// 						to.Ptr(armcdn.AfdCustomizedCipherSuiteForTls12ECDHERSAAES128GCMSHA256),
	// 					},
	// 					CipherSuiteSetForTls13: []*armcdn.AfdCustomizedCipherSuiteForTls13{
	// 						to.Ptr(armcdn.AfdCustomizedCipherSuiteForTls13TLSAES128GCMSHA256),
	// 						to.Ptr(armcdn.AfdCustomizedCipherSuiteForTls13TLSAES256GCMSHA384),
	// 					},
	// 				},
	// 				MinimumTLSVersion: to.Ptr(armcdn.AfdMinimumTLSVersionTLS12),
	// 				Secret: &armcdn.ResourceReference{
	// 					ID: to.Ptr(""),
	// 				},
	// 			},
	// 			ValidationProperties: &armcdn.DomainValidationProperties{
	// 				ExpirationDate: to.Ptr(""),
	// 				ValidationToken: to.Ptr(""),
	// 			},
	// 		},
	// 	},
	// }
}

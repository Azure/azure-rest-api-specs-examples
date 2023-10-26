package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDCustomDomains_Get.json
func ExampleAFDCustomDomainsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAFDCustomDomainsClient().Get(ctx, "RG", "profile1", "domain1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AFDDomain = armcdn.AFDDomain{
	// 	Name: to.Ptr("domain1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/customdomains"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/domain1"),
	// 	Properties: &armcdn.AFDDomainProperties{
	// 		AzureDNSZone: &armcdn.ResourceReference{
	// 			ID: to.Ptr(""),
	// 		},
	// 		PreValidatedCustomDomainResourceID: &armcdn.ResourceReference{
	// 			ID: to.Ptr(""),
	// 		},
	// 		ProfileName: to.Ptr("profile1"),
	// 		TLSSettings: &armcdn.AFDDomainHTTPSParameters{
	// 			CertificateType: to.Ptr(armcdn.AfdCertificateTypeManagedCertificate),
	// 			MinimumTLSVersion: to.Ptr(armcdn.AfdMinimumTLSVersionTLS12),
	// 			Secret: &armcdn.ResourceReference{
	// 				ID: to.Ptr(""),
	// 			},
	// 		},
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		DomainValidationState: to.Ptr(armcdn.DomainValidationStatePending),
	// 		HostName: to.Ptr("www.contoso.com"),
	// 		ValidationProperties: &armcdn.DomainValidationProperties{
	// 			ExpirationDate: to.Ptr("2009-06-15T13:45:43.0000000Z"),
	// 			ValidationToken: to.Ptr("8c9912db-c615-4eeb-8465"),
	// 		},
	// 	},
	// }
}

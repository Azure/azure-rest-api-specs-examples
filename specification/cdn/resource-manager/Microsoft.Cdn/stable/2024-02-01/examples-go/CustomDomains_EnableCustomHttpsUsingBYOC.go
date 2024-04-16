package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/CustomDomains_EnableCustomHttpsUsingBYOC.json
func ExampleCustomDomainsClient_BeginEnableCustomHTTPS_customDomainsEnableCustomHttpsUsingYourOwnCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomDomainsClient().BeginEnableCustomHTTPS(ctx, "RG", "profile1", "endpoint1", "www-someDomain-net", &armcdn.CustomDomainsClientBeginEnableCustomHTTPSOptions{CustomDomainHTTPSParameters: nil})
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
	// res.CustomDomain = armcdn.CustomDomain{
	// 	Name: to.Ptr("www-someDomain-net"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/endpoints/customdomains"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customdomains/www-someDomain-net"),
	// 	Properties: &armcdn.CustomDomainProperties{
	// 		CustomHTTPSProvisioningState: to.Ptr(armcdn.CustomHTTPSProvisioningStateEnabled),
	// 		CustomHTTPSProvisioningSubstate: to.Ptr(armcdn.CustomHTTPSProvisioningSubstateCertificateDeployed),
	// 		HostName: to.Ptr("www.someDomain.net"),
	// 		ProvisioningState: to.Ptr(armcdn.CustomHTTPSProvisioningState("Succeeded")),
	// 		ResourceState: to.Ptr(armcdn.CustomDomainResourceStateActive),
	// 		ValidationData: to.Ptr("validationdata"),
	// 	},
	// }
}

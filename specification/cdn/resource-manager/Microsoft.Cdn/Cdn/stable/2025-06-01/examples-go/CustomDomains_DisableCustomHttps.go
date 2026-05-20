package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/CustomDomains_DisableCustomHttps.json
func ExampleCustomDomainsClient_BeginDisableCustomHTTPS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomDomainsClient().BeginDisableCustomHTTPS(ctx, "RG", "profile1", "endpoint1", "www-someDomain-net", nil)
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
	// res = armcdn.CustomDomainsClientDisableCustomHTTPSResponse{
	// 	CustomDomain: armcdn.CustomDomain{
	// 		Name: to.Ptr("www-someDomain-net"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/endpoints/customdomains"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customdomains/www-someDomain-net"),
	// 		Properties: &armcdn.CustomDomainProperties{
	// 			CustomHTTPSProvisioningState: to.Ptr(armcdn.CustomHTTPSProvisioningStateDisabled),
	// 			CustomHTTPSProvisioningSubstate: to.Ptr(armcdn.CustomHTTPSProvisioningSubstateCertificateDeleted),
	// 			HostName: to.Ptr("www.someDomain.net"),
	// 			ProvisioningState: to.Ptr(armcdn.CustomHTTPSProvisioningState("Succeeded")),
	// 			ResourceState: to.Ptr(armcdn.CustomDomainResourceStateActive),
	// 		},
	// 	},
	// }
}

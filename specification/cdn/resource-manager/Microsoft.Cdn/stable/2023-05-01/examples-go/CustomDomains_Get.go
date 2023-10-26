package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7b551033155a63739b6d28f79b9c07569f6179b8/specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/CustomDomains_Get.json
func ExampleCustomDomainsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomDomainsClient().Get(ctx, "RG", "profile1", "endpoint1", "www-someDomain-net", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomDomain = armcdn.CustomDomain{
	// 	Name: to.Ptr("www-someDomain-net"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/endpoints/customdomains"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customdomains/www-someDomain-net"),
	// 	Properties: &armcdn.CustomDomainProperties{
	// 		CustomHTTPSProvisioningState: to.Ptr(armcdn.CustomHTTPSProvisioningStateDisabled),
	// 		CustomHTTPSProvisioningSubstate: to.Ptr(armcdn.CustomHTTPSProvisioningSubstate("None")),
	// 		HostName: to.Ptr("www.someDomain.net"),
	// 		ProvisioningState: to.Ptr(armcdn.CustomHTTPSProvisioningState("Succeeded")),
	// 		ResourceState: to.Ptr(armcdn.CustomDomainResourceStateActive),
	// 	},
	// }
}

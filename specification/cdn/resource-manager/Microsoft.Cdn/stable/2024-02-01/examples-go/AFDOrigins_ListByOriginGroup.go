package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOrigins_ListByOriginGroup.json
func ExampleAFDOriginsClient_NewListByOriginGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAFDOriginsClient().NewListByOriginGroupPager("RG", "profile1", "origingroup1", nil)
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
		// page.AFDOriginListResult = armcdn.AFDOriginListResult{
		// 	Value: []*armcdn.AFDOrigin{
		// 		{
		// 			Name: to.Ptr("origin1"),
		// 			Type: to.Ptr("Microsoft.Cdn/profiles/origingroups/origins"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1/origins/origin1"),
		// 			Properties: &armcdn.AFDOriginProperties{
		// 				EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
		// 				EnforceCertificateNameCheck: to.Ptr(true),
		// 				HostName: to.Ptr("host1.blob.core.windows.net"),
		// 				HTTPPort: to.Ptr[int32](80),
		// 				HTTPSPort: to.Ptr[int32](443),
		// 				OriginGroupName: to.Ptr("origingroup1"),
		// 				OriginHostHeader: to.Ptr("host1.foo.com"),
		// 				DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
		// 				ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

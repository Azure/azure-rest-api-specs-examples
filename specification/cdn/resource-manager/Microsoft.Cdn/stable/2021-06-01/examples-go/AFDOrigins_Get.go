package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOrigins_Get.json
func ExampleAFDOriginsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAFDOriginsClient().Get(ctx, "RG", "profile1", "origingroup1", "origin1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AFDOrigin = armcdn.AFDOrigin{
	// 	Name: to.Ptr("origin1"),
	// 	Type: to.Ptr("Microsoft.Cdn/profiles/origingroups/origins"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1/origins/origin1"),
	// 	Properties: &armcdn.AFDOriginProperties{
	// 		EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
	// 		EnforceCertificateNameCheck: to.Ptr(true),
	// 		HostName: to.Ptr("host1.blob.core.windows.net"),
	// 		HTTPPort: to.Ptr[int32](80),
	// 		HTTPSPort: to.Ptr[int32](443),
	// 		OriginGroupName: to.Ptr("origingroup1"),
	// 		OriginHostHeader: to.Ptr("host1.foo.com"),
	// 		DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 		ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 	},
	// }
}

package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/AFDOrigins_Update.json
func ExampleAFDOriginsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAFDOriginsClient().BeginUpdate(ctx, "RG", "profile1", "origingroup1", "origin1", armcdn.AFDOriginUpdateParameters{
		Properties: &armcdn.AFDOriginUpdatePropertiesParameters{
			EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
			HostName:     to.Ptr("host1.blob.core.windows.net"),
			HTTPPort:     to.Ptr[int32](80),
			HTTPSPort:    to.Ptr[int32](443),
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
	// res = armcdn.AFDOriginsClientUpdateResponse{
	// 	AFDOrigin: armcdn.AFDOrigin{
	// 		Name: to.Ptr("origin1"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/origingroups/origins"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/origingroups/origingroup1/origins/origin1"),
	// 		Properties: &armcdn.AFDOriginProperties{
	// 			DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 			EnabledState: to.Ptr(armcdn.EnabledStateEnabled),
	// 			EnforceCertificateNameCheck: to.Ptr(true),
	// 			HostName: to.Ptr("host1.blob.core.windows.net"),
	// 			HTTPPort: to.Ptr[int32](80),
	// 			HTTPSPort: to.Ptr[int32](443),
	// 			OriginGroupName: to.Ptr("origingroup1"),
	// 			OriginHostHeader: to.Ptr("host1.foo.com"),
	// 			ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

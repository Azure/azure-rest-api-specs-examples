Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.7.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/CreateCloudServiceWithMultiRole.json
func ExampleCloudServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewCloudServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cloud-service-name>",
		&armcompute.CloudServicesClientBeginCreateOrUpdateOptions{Parameters: &armcompute.CloudService{
			Location: to.Ptr("<location>"),
			Properties: &armcompute.CloudServiceProperties{
				Configuration: to.Ptr("<configuration>"),
				NetworkProfile: &armcompute.CloudServiceNetworkProfile{
					LoadBalancerConfigurations: []*armcompute.LoadBalancerConfiguration{
						{
							Name: to.Ptr("<name>"),
							Properties: &armcompute.LoadBalancerConfigurationProperties{
								FrontendIPConfigurations: []*armcompute.LoadBalancerFrontendIPConfiguration{
									{
										Name: to.Ptr("<name>"),
										Properties: &armcompute.LoadBalancerFrontendIPConfigurationProperties{
											PublicIPAddress: &armcompute.SubResource{
												ID: to.Ptr("<id>"),
											},
										},
									}},
							},
						}},
				},
				PackageURL: to.Ptr("<package-url>"),
				RoleProfile: &armcompute.CloudServiceRoleProfile{
					Roles: []*armcompute.CloudServiceRoleProfileProperties{
						{
							Name: to.Ptr("<name>"),
							SKU: &armcompute.CloudServiceRoleSKU{
								Name:     to.Ptr("<name>"),
								Capacity: to.Ptr[int64](1),
								Tier:     to.Ptr("<tier>"),
							},
						},
						{
							Name: to.Ptr("<name>"),
							SKU: &armcompute.CloudServiceRoleSKU{
								Name:     to.Ptr("<name>"),
								Capacity: to.Ptr[int64](1),
								Tier:     to.Ptr("<tier>"),
							},
						}},
				},
				UpgradeMode: to.Ptr(armcompute.CloudServiceUpgradeModeAuto),
			},
		},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

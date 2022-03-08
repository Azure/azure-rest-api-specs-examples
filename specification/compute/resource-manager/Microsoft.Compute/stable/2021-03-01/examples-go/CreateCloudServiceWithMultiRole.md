Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.5.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/CreateCloudServiceWithMultiRole.json
func ExampleCloudServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewCloudServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cloud-service-name>",
		&armcompute.CloudServicesClientBeginCreateOrUpdateOptions{Parameters: &armcompute.CloudService{
			Location: to.StringPtr("<location>"),
			Properties: &armcompute.CloudServiceProperties{
				Configuration: to.StringPtr("<configuration>"),
				NetworkProfile: &armcompute.CloudServiceNetworkProfile{
					LoadBalancerConfigurations: []*armcompute.LoadBalancerConfiguration{
						{
							Name: to.StringPtr("<name>"),
							Properties: &armcompute.LoadBalancerConfigurationProperties{
								FrontendIPConfigurations: []*armcompute.LoadBalancerFrontendIPConfiguration{
									{
										Name: to.StringPtr("<name>"),
										Properties: &armcompute.LoadBalancerFrontendIPConfigurationProperties{
											PublicIPAddress: &armcompute.SubResource{
												ID: to.StringPtr("<id>"),
											},
										},
									}},
							},
						}},
				},
				PackageURL: to.StringPtr("<package-url>"),
				RoleProfile: &armcompute.CloudServiceRoleProfile{
					Roles: []*armcompute.CloudServiceRoleProfileProperties{
						{
							Name: to.StringPtr("<name>"),
							SKU: &armcompute.CloudServiceRoleSKU{
								Name:     to.StringPtr("<name>"),
								Capacity: to.Int64Ptr(1),
								Tier:     to.StringPtr("<tier>"),
							},
						},
						{
							Name: to.StringPtr("<name>"),
							SKU: &armcompute.CloudServiceRoleSKU{
								Name:     to.StringPtr("<name>"),
								Capacity: to.Int64Ptr(1),
								Tier:     to.StringPtr("<tier>"),
							},
						}},
				},
				UpgradeMode: armcompute.CloudServiceUpgradeMode("Auto").ToPtr(),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CloudServicesClientCreateOrUpdateResult)
}
```

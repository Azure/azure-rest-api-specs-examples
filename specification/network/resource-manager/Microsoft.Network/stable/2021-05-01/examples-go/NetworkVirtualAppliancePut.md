Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkVirtualAppliancePut.json
func ExampleVirtualAppliancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualAppliancesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<network-virtual-appliance-name>",
		armnetwork.VirtualAppliance{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Identity: &armnetwork.ManagedServiceIdentity{
				Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
				UserAssignedIdentities: map[string]*armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
				},
			},
			Properties: &armnetwork.VirtualAppliancePropertiesFormat{
				BootStrapConfigurationBlobs: []*string{
					to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig")},
				CloudInitConfigurationBlobs: []*string{
					to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig")},
				NvaSKU: &armnetwork.VirtualApplianceSKUProperties{
					BundledScaleUnit:   to.Ptr("<bundled-scale-unit>"),
					MarketPlaceVersion: to.Ptr("<market-place-version>"),
					Vendor:             to.Ptr("<vendor>"),
				},
				VirtualApplianceAsn: to.Ptr[int64](10000),
				VirtualHub: &armnetwork.SubResource{
					ID: to.Ptr("<id>"),
				},
			},
		},
		&armnetwork.VirtualAppliancesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

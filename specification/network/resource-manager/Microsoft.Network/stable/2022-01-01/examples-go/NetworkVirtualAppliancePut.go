package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkVirtualAppliancePut.json
func ExampleVirtualAppliancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualAppliancesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "nva", armnetwork.VirtualAppliance{
		Location: to.Ptr("West US"),
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
				BundledScaleUnit:   to.Ptr("1"),
				MarketPlaceVersion: to.Ptr("12.1"),
				Vendor:             to.Ptr("Cisco SDWAN"),
			},
			VirtualApplianceAsn: to.Ptr[int64](10000),
			VirtualHub: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

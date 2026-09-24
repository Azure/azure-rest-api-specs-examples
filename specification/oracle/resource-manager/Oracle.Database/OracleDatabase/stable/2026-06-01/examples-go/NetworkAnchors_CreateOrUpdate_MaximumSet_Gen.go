package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v3"
)

// Generated from example definition: 2026-06-01/NetworkAnchors_CreateOrUpdate_MaximumSet_Gen.json
func ExampleNetworkAnchorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkAnchorsClient().BeginCreateOrUpdate(ctx, "rgopenapi", "resource1", armoracledatabase.NetworkAnchor{
		Properties: &armoracledatabase.NetworkAnchorProperties{
			ResourceAnchorID:                     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resourceAnchors/anchor1"),
			VnetID:                               to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
			SubnetID:                             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
			OciVcnID:                             to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
			OciVcnDNSLabel:                       to.Ptr("example"),
			OciSubnetID:                          to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
			OciBackupCidrBlock:                   to.Ptr("example"),
			IsOracleToAzureDNSZoneSyncEnabled:    to.Ptr(true),
			IsOracleDNSListeningEndpointEnabled:  to.Ptr(true),
			IsOracleDNSForwardingEndpointEnabled: to.Ptr(true),
			DNSForwardingRules: []*armoracledatabase.DNSForwardingRule{
				{
					DomainNames:         to.Ptr("ghs"),
					ForwardingIPAddress: to.Ptr("example"),
				},
			},
			DNSListeningEndpointAllowedCidrs: to.Ptr("toqgyp"),
			ProximityPlacementGroup: &armoracledatabase.ProximityPlacementGroup{
				ProximityPlacementGroupID: to.Ptr("example"),
				ProximityAnchorID:         to.Ptr("example"),
				EntityTypeIntendedToUse:   to.Ptr(armoracledatabase.ProximityPlacementGroupEntityTypeCloudExadataInfrastructure),
			},
		},
		Zones: []*string{
			to.Ptr("zznbkklaih"),
		},
		Tags: map[string]*string{
			"key6589": to.Ptr("mcg"),
		},
		Location: to.Ptr("eastus"),
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
	// res = armoracledatabase.NetworkAnchorsClientCreateOrUpdateResponse{
	// 	NetworkAnchor: armoracledatabase.NetworkAnchor{
	// 		Properties: &armoracledatabase.NetworkAnchorProperties{
	// 			ResourceAnchorID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resourceAnchors/anchor1"),
	// 			ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 			VnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 			CidrBlock: to.Ptr("ehm"),
	// 			OciVcnID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			OciVcnDNSLabel: to.Ptr("example"),
	// 			OciSubnetID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
	// 			OciBackupCidrBlock: to.Ptr("example"),
	// 			IsOracleToAzureDNSZoneSyncEnabled: to.Ptr(true),
	// 			IsOracleDNSListeningEndpointEnabled: to.Ptr(true),
	// 			IsOracleDNSForwardingEndpointEnabled: to.Ptr(true),
	// 			DNSListeningEndpointIPAddress: to.Ptr("example"),
	// 			DNSForwardingEndpointIPAddress: to.Ptr("example"),
	// 			DNSForwardingRulesURL: to.Ptr("https://example.com"),
	// 			DNSListeningEndpointNsgRulesURL: to.Ptr("wmjxzjkoktq"),
	// 			DNSForwardingEndpointNsgRulesURL: to.Ptr("https://example.com"),
	// 			ProximityPlacementGroup: &armoracledatabase.ProximityPlacementGroup{
	// 				ProximityPlacementGroupID: to.Ptr("example"),
	// 				ProximityAnchorID: to.Ptr("example"),
	// 				EntityTypeIntendedToUse: to.Ptr(armoracledatabase.ProximityPlacementGroupEntityTypeCloudExadataInfrastructure),
	// 			},
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("zznbkklaih"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key6589": to.Ptr("mcg"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resources/resource1"),
	// 		Name: to.Ptr("adg"),
	// 		Type: to.Ptr("Oracle.Database/resource"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("ns"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("example"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
	// 		},
	// 	},
	// }
}

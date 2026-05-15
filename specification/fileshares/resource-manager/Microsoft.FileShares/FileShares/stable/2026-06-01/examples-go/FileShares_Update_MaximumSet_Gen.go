package armfileshares_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fileshares/armfileshares"
)

// Generated from example definition: 2026-06-01/FileShares_Update_MaximumSet_Gen.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfileshares.NewClientFactory("0681745E-3F9F-4966-80E6-69624A3B29F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginUpdate(ctx, "rgfileshares", "fileshare", armfileshares.FileShareUpdate{
		Tags: map[string]*string{
			"key173": to.Ptr("uyf"),
		},
		Properties: &armfileshares.FileShareUpdateProperties{
			ProvisionedStorageGiB:          to.Ptr[int32](7),
			ProvisionedIOPerSec:            to.Ptr[int32](1),
			ProvisionedThroughputMiBPerSec: to.Ptr[int32](29),
			NfsProtocolProperties: &armfileshares.NfsProtocolProperties{
				RootSquash:                  to.Ptr(armfileshares.ShareRootSquashNoRootSquash),
				EncryptionInTransitRequired: to.Ptr(armfileshares.EncryptionInTransitRequiredEnabled),
			},
			PublicAccessProperties: &armfileshares.PublicAccessProperties{
				AllowedSubnets: []*string{
					to.Ptr("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
				},
			},
			PublicNetworkAccess: to.Ptr(armfileshares.PublicNetworkAccessEnabled),
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
	// res = armfileshares.ClientUpdateResponse{
	// 	FileShare: armfileshares.FileShare{
	// 		Properties: &armfileshares.FileShareProperties{
	// 			MountName: to.Ptr("fileshare"),
	// 			HostName: to.Ptr("venxryqlzdhnb"),
	// 			MediaTier: to.Ptr(armfileshares.MediaTierSSD),
	// 			Redundancy: to.Ptr(armfileshares.RedundancyLocal),
	// 			Protocol: to.Ptr(armfileshares.ProtocolNFS),
	// 			ProvisionedStorageGiB: to.Ptr[int32](20),
	// 			ProvisionedStorageNextAllowedDowngrade: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.954Z"); return t}()),
	// 			ProvisionedIOPerSec: to.Ptr[int32](9),
	// 			ProvisionedIOPerSecNextAllowedDowngrade: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.954Z"); return t}()),
	// 			ProvisionedThroughputMiBPerSec: to.Ptr[int32](5),
	// 			ProvisionedThroughputNextAllowedDowngrade: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.954Z"); return t}()),
	// 			IncludedBurstIOPerSec: to.Ptr[int32](24523),
	// 			MaxBurstIOPerSecCredits: to.Ptr[int64](0),
	// 			NfsProtocolProperties: &armfileshares.NfsProtocolProperties{
	// 				RootSquash: to.Ptr(armfileshares.ShareRootSquashNoRootSquash),
	// 				EncryptionInTransitRequired: to.Ptr(armfileshares.EncryptionInTransitRequiredEnabled),
	// 			},
	// 			PublicAccessProperties: &armfileshares.PublicAccessProperties{
	// 				AllowedSubnets: []*string{
	// 					to.Ptr("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armfileshares.FileShareProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armfileshares.PublicNetworkAccessEnabled),
	// 			PrivateEndpointConnections: []*armfileshares.PrivateEndpointConnection{
	// 				{
	// 					Properties: &armfileshares.PrivateEndpointConnectionProperties{
	// 						GroupIDs: []*string{
	// 							to.Ptr("ndoii"),
	// 						},
	// 						PrivateEndpoint: &armfileshares.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.Network/privateEndpoints/testresourcename"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armfileshares.PrivateLinkServiceConnectionState{
	// 							Status: to.Ptr(armfileshares.PrivateEndpointServiceConnectionStatusPending),
	// 							Description: to.Ptr("a"),
	// 							ActionsRequired: to.Ptr("pvlmopquhexpdjpcaruvj"),
	// 						},
	// 						ProvisioningState: to.Ptr(armfileshares.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 					},
	// 					ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.FileShares/fileShares/testresourcename/privateEndpointConnections/privateEndpointConnectionName"),
	// 					Name: to.Ptr("kjavgkroj"),
	// 					Type: to.Ptr("Microsoft.FileShares/fileShares/privateEndpointConnections"),
	// 					SystemData: &armfileshares.SystemData{
	// 						CreatedBy: to.Ptr("moat"),
	// 						CreatedByType: to.Ptr(armfileshares.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.955Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("dxukcadzapywjcgnxsqchaa"),
	// 						LastModifiedByType: to.Ptr(armfileshares.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.955Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key9647": to.Ptr("xwokdvyoae"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		ID: to.Ptr("/subscriptions/0590b445-524e-4caa-9870-2e2461d23434/resourceGroups/fiddler-mfs-rg/providers/Microsoft.FileShares/fileShares/fileshare"),
	// 		Name: to.Ptr("fileshare"),
	// 		Type: to.Ptr("Microsoft.FileShares/fileShares"),
	// 		SystemData: &armfileshares.SystemData{
	// 			CreatedBy: to.Ptr("moat"),
	// 			CreatedByType: to.Ptr(armfileshares.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.955Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("dxukcadzapywjcgnxsqchaa"),
	// 			LastModifiedByType: to.Ptr(armfileshares.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-23T12:37:53.955Z"); return t}()),
	// 		},
	// 	},
	// }
}

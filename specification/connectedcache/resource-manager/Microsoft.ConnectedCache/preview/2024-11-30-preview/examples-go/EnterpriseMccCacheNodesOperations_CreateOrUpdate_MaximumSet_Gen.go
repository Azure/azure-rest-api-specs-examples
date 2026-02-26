package armconnectedcache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
)

// Generated from example definition: 2024-11-30-preview/EnterpriseMccCacheNodesOperations_CreateOrUpdate_MaximumSet_Gen.json
func ExampleEnterpriseMccCacheNodesOperationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEnterpriseMccCacheNodesOperationsClient().BeginCreateOrUpdate(ctx, "rgConnectedCache", "nhdkvstdrrtsxxuz", "fgduqdovidpemlnmhelomffuafdrbgaasqznvrdkbvspfzsnrhncdtqquhijhdpwyzwleukqldpceyxqhqlftqrr", armconnectedcache.EnterpriseMccCacheNodeResource{
		Location: to.Ptr("westus"),
		Properties: &armconnectedcache.CacheNodeProperty{
			CacheNode: &armconnectedcache.CacheNodeEntity{
				FullyQualifiedResourceID: to.Ptr("yeinlleavzbehelhsucb"),
				CustomerName:             to.Ptr("zsklcocjfjhkcpmzyefzkwamdzc"),
				IPAddress:                to.Ptr("gbfkdhloyphnpnhemwrcrxlk"),
				CustomerIndex:            to.Ptr("vafvezmelfgmjsrccjukrhppljvipg"),
				CacheNodeID:              to.Ptr("fmrjefyddfn"),
				CacheNodeName:            to.Ptr("qppvqxliajjfoalzjbgmxggr"),
				CustomerAsn:              to.Ptr[int32](25),
				IsEnabled:                to.Ptr(true),
				MaxAllowableEgressInMbps: to.Ptr[int32](27),
				IsEnterpriseManaged:      to.Ptr(true),
				CidrCSV: []*string{
					to.Ptr("nlqlvrthafvvljuupcbcw"),
				},
				ShouldMigrate:     to.Ptr(true),
				CidrSelectionType: to.Ptr[int32](11),
			},
			StatusCode:    to.Ptr("1"),
			StatusText:    to.Ptr("Success"),
			StatusDetails: to.Ptr("lgljxmyyoakw"),
			AdditionalCacheNodeProperties: &armconnectedcache.AdditionalCacheNodeProperties{
				CacheNodePropertiesDetailsIssuesList: []*string{
					to.Ptr("ennbzfpuszgalzpawmwicaofqcwcj"),
				},
				DriveConfiguration: []*armconnectedcache.CacheNodeDriveConfiguration{
					{
						PhysicalPath: to.Ptr("pcbkezoofhamkycot"),
						SizeInGb:     to.Ptr[int32](14),
						CacheNumber:  to.Ptr[int32](11),
						NginxMapping: to.Ptr("cirlpkpuxg"),
					},
				},
				BgpConfiguration: &armconnectedcache.BgpConfiguration{
					AsnToIPAddressMapping: to.Ptr("fjbggfvumrn"),
				},
				ProxyURLConfiguration: &armconnectedcache.ProxyURLConfiguration{
					ProxyURL: to.Ptr("hplstyg"),
				},
				OptionalProperty1: to.Ptr("ph"),
				OptionalProperty2: to.Ptr("soqqgpgcbhb"),
				OptionalProperty3: to.Ptr("fpnycrbagptsujiotnjfuhlm"),
				OptionalProperty4: to.Ptr("gesqugrxvhxlxxyvatgrautxwlmxsf"),
				OptionalProperty5: to.Ptr("zknjgzpaqtvdqjydd"),
			},
			Error: &armconnectedcache.ErrorDetail{},
		},
		Tags: map[string]*string{
			"key259": to.Ptr("qbkixjuyjkoj"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconnectedcache.EnterpriseMccCacheNodesOperationsClientCreateOrUpdateResponse{
	// 	EnterpriseMccCacheNodeResource: &armconnectedcache.EnterpriseMccCacheNodeResource{
	// 		Location: to.Ptr("westus"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/DOTest/providers/Microsoft.ConnectedCache/enterpriseMccCustomers/MccRPTest1/enterpriseMccCacheNodes/MCCCachenode1"),
	// 		Name: to.Ptr("MccRPTest1"),
	// 		Type: to.Ptr("Microsoft.ConnectedCache/enterpriseMccCustomers/enterpriseMccCacheNodes"),
	// 		Properties: &armconnectedcache.CacheNodeProperty{
	// 			ProvisioningState: to.Ptr(armconnectedcache.ProvisioningStateSucceeded),
	// 			CacheNode: &armconnectedcache.CacheNodeEntity{
	// 				FullyQualifiedResourceID: to.Ptr("yeinlleavzbehelhsucb"),
	// 				CustomerID: to.Ptr("xjjgcknmsakfawcmwbdufdxktysp"),
	// 				CustomerName: to.Ptr("zsklcocjfjhkcpmzyefzkwamdzc"),
	// 				IPAddress: to.Ptr("gbfkdhloyphnpnhemwrcrxlk"),
	// 				CustomerIndex: to.Ptr("vafvezmelfgmjsrccjukrhppljvipg"),
	// 				CacheNodeID: to.Ptr("fmrjefyddfn"),
	// 				CacheNodeName: to.Ptr("qppvqxliajjfoalzjbgmxggr"),
	// 				CustomerAsn: to.Ptr[int32](25),
	// 				IsEnabled: to.Ptr(true),
	// 				MaxAllowableEgressInMbps: to.Ptr[int32](27),
	// 				MaxAllowableProbability: to.Ptr[float32](12),
	// 				XCid: to.Ptr("ooyrhuockukmsrqsevwaaqglhf"),
	// 				IsEnterpriseManaged: to.Ptr(true),
	// 				DeleteAsyncOperationID: to.Ptr("tvmivvnpisuboggoqgbobugld"),
	// 				ClientTenantID: to.Ptr("iwrpakllacxvtqygpaimlpxvqrjz"),
	// 				Category: to.Ptr("utfjefejkrpxcmrgygzawa"),
	// 				ReleaseVersion: to.Ptr[int32](29),
	// 				LastSyncWithAzureTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.777Z"); return t}()),
	// 				SynchWithAzureAttemptsCount: to.Ptr[int32](4),
	// 				ContainerConfigurations: to.Ptr("dyvefvbvrtsmdrdmiuphzh"),
	// 				CidrCSV: []*string{
	// 					to.Ptr("nlqlvrthafvvljuupcbcw"),
	// 				},
	// 				CidrCSVLastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.774Z"); return t}()),
	// 				BgpCidrCSVLastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.774Z"); return t}()),
	// 				BgpLastReportedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.774Z"); return t}()),
	// 				BgpReviewStateText: to.Ptr("khtriksrqrjieouoz"),
	// 				BgpReviewState: to.Ptr(armconnectedcache.BgpReviewStateEnumNotConfigured),
	// 				BgpReviewFeedback: to.Ptr("bc"),
	// 				BgpNumberOfTimesUpdated: to.Ptr[int32](23),
	// 				BgpNumberOfRecords: to.Ptr[int32](25),
	// 				BgpCidrBlocksCount: to.Ptr[int32](28),
	// 				BgpAddressSpace: to.Ptr[int32](30),
	// 				ShouldMigrate: to.Ptr(true),
	// 				BgpFileBytesTruncated: to.Ptr[int32](26),
	// 				CidrSelectionType: to.Ptr[int32](11),
	// 				IsFrozen: to.Ptr(true),
	// 				ReviewState: to.Ptr[int32](2),
	// 				ReviewStateText: to.Ptr("xubwhyrdtramrmoldbxqwauyusymqu"),
	// 				ReviewFeedback: to.Ptr("zowuztcnybt"),
	// 				ConfigurationState: to.Ptr(armconnectedcache.ConfigurationStateConfigured),
	// 				ConfigurationStateText: to.Ptr("okbofqwtzcsju"),
	// 				AddressSpace: to.Ptr[int32](1),
	// 				WorkerConnections: to.Ptr[int32](19),
	// 				WorkerConnectionsLastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.774Z"); return t}()),
	// 				ContainerResyncTrigger: to.Ptr[int32](23),
	// 				ImageURI: to.Ptr("wkowdrixfxvjmdcsy"),
	// 			},
	// 			StatusCode: to.Ptr("1"),
	// 			StatusText: to.Ptr("Success"),
	// 			StatusDetails: to.Ptr("lgljxmyyoakw"),
	// 			AdditionalCacheNodeProperties: &armconnectedcache.AdditionalCacheNodeProperties{
	// 				CacheNodePropertiesDetailsIssuesList: []*string{
	// 					to.Ptr("ennbzfpuszgalzpawmwicaofqcwcj"),
	// 				},
	// 				AggregatedStatusDetails: to.Ptr("nztcstlyjhmllbnrrluhzusmoa"),
	// 				AggregatedStatusText: to.Ptr("ofiedyvzpbijnfghnrrigqaws"),
	// 				AggregatedStatusCode: to.Ptr[int32](25),
	// 				ProductVersion: to.Ptr("hfbuukpxoerpkziym"),
	// 				IsProvisioned: to.Ptr(true),
	// 				CacheNodeStateDetailedText: to.Ptr("onckp"),
	// 				CacheNodeStateShortText: to.Ptr("yxkwlubqlzhkmemzjzlxksho"),
	// 				CacheNodeState: to.Ptr[int32](30),
	// 				DriveConfiguration: []*armconnectedcache.CacheNodeDriveConfiguration{
	// 					{
	// 						PhysicalPath: to.Ptr("pcbkezoofhamkycot"),
	// 						SizeInGb: to.Ptr[int32](14),
	// 						CacheNumber: to.Ptr[int32](11),
	// 						NginxMapping: to.Ptr("cirlpkpuxg"),
	// 					},
	// 				},
	// 				BgpConfiguration: &armconnectedcache.BgpConfiguration{
	// 					AsnToIPAddressMapping: to.Ptr("fjbggfvumrn"),
	// 				},
	// 				ProxyURLConfiguration: &armconnectedcache.ProxyURLConfiguration{
	// 					ProxyURL: to.Ptr("hplstyg"),
	// 				},
	// 				OptionalProperty1: to.Ptr("ph"),
	// 				OptionalProperty2: to.Ptr("soqqgpgcbhb"),
	// 				OptionalProperty3: to.Ptr("fpnycrbagptsujiotnjfuhlm"),
	// 				OptionalProperty4: to.Ptr("gesqugrxvhxlxxyvatgrautxwlmxsf"),
	// 				OptionalProperty5: to.Ptr("zknjgzpaqtvdqjydd"),
	// 			},
	// 			Status: to.Ptr("utgrdjsw"),
	// 			Error: &armconnectedcache.ErrorDetail{
	// 				Code: to.Ptr("dkvgvtftpsjsbhlnapvihefxneoggs"),
	// 				Message: to.Ptr("okakgyfnmyob"),
	// 				Details: []*armconnectedcache.ErrorDetail{
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key259": to.Ptr("qbkixjuyjkoj"),
	// 		},
	// 		SystemData: &armconnectedcache.SystemData{
	// 			CreatedBy: to.Ptr("gambtqj"),
	// 			CreatedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("qomgaceiessgnuogz"),
	// 			LastModifiedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
	// 		},
	// 	},
	// }
}

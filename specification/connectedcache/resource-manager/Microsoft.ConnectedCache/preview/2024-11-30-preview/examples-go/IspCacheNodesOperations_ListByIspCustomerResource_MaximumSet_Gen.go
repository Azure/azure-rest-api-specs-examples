package armconnectedcache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
)

// Generated from example definition: 2024-11-30-preview/IspCacheNodesOperations_ListByIspCustomerResource_MaximumSet_Gen.json
func ExampleIspCacheNodesOperationsClient_NewListByIspCustomerResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIspCacheNodesOperationsClient().NewListByIspCustomerResourcePager("rgConnectedCache", "MccRPTest1", nil)
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
		// page = armconnectedcache.IspCacheNodesOperationsClientListByIspCustomerResourceResponse{
		// 	IspCacheNodeResourceListResult: armconnectedcache.IspCacheNodeResourceListResult{
		// 		Value: []*armconnectedcache.IspCacheNodeResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/rgConnectedCache/providers/Microsoft.ConnectedCache/ispCustomers/MccRPTest1/ispCacheNodes/MCCCachenode1"),
		// 				Name: to.Ptr("MccRPTest1"),
		// 				Type: to.Ptr("Microsoft.ConnectedCache/ispCustomers/ispCacheNodes"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armconnectedcache.CacheNodeProperty{
		// 					ProvisioningState: to.Ptr(armconnectedcache.ProvisioningStateSucceeded),
		// 					CacheNode: &armconnectedcache.CacheNodeEntity{
		// 						FullyQualifiedResourceID: to.Ptr("hskxkpbiqbrbjiwdzrxndru"),
		// 						CustomerID: to.Ptr("ceyfqoygknpmmjojlhuklqybfl"),
		// 						CustomerName: to.Ptr("xwyqk"),
		// 						IPAddress: to.Ptr("voctagljcwqgcpnionqdcbjk"),
		// 						CustomerIndex: to.Ptr("qtoiglqaswivmkjhzogburcxtszmek"),
		// 						CacheNodeID: to.Ptr("xjzffjftwcgsehanoxsl"),
		// 						CacheNodeName: to.Ptr("mfjxb"),
		// 						CustomerAsn: to.Ptr[int32](4),
		// 						IsEnabled: to.Ptr(true),
		// 						MaxAllowableEgressInMbps: to.Ptr[int32](29),
		// 						MaxAllowableProbability: to.Ptr[float32](16),
		// 						XCid: to.Ptr("kwnwgdbeflazz"),
		// 						IsEnterpriseManaged: to.Ptr(true),
		// 						DeleteAsyncOperationID: to.Ptr("oeyevqzlpbsimi"),
		// 						ClientTenantID: to.Ptr("ds"),
		// 						Category: to.Ptr("rixlfzbl"),
		// 						ReleaseVersion: to.Ptr[int32](30),
		// 						LastSyncWithAzureTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.777Z"); return t}()),
		// 						SynchWithAzureAttemptsCount: to.Ptr[int32](11),
		// 						ContainerConfigurations: to.Ptr("waygqqgfzvnvlbufilldsqavwlshzt"),
		// 						CidrCSV: []*string{
		// 							to.Ptr("nlqlvrthafvvljuupcbcw"),
		// 						},
		// 						CidrCSVLastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.777Z"); return t}()),
		// 						BgpCidrCSVLastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.777Z"); return t}()),
		// 						BgpLastReportedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.777Z"); return t}()),
		// 						BgpReviewStateText: to.Ptr("xduneialocyri"),
		// 						BgpReviewState: to.Ptr(armconnectedcache.BgpReviewStateEnumNotConfigured),
		// 						BgpReviewFeedback: to.Ptr("us"),
		// 						BgpNumberOfTimesUpdated: to.Ptr[int32](2),
		// 						BgpNumberOfRecords: to.Ptr[int32](21),
		// 						BgpCidrBlocksCount: to.Ptr[int32](5),
		// 						BgpAddressSpace: to.Ptr[int32](6),
		// 						ShouldMigrate: to.Ptr(true),
		// 						BgpFileBytesTruncated: to.Ptr[int32](13),
		// 						CidrSelectionType: to.Ptr[int32](4),
		// 						IsFrozen: to.Ptr(true),
		// 						ReviewState: to.Ptr[int32](19),
		// 						ReviewStateText: to.Ptr("mrnragzmnscovixohmif"),
		// 						ReviewFeedback: to.Ptr("wrcfimvx"),
		// 						ConfigurationState: to.Ptr(armconnectedcache.ConfigurationStateConfigured),
		// 						ConfigurationStateText: to.Ptr("arugukqjcxmkbgfambg"),
		// 						AddressSpace: to.Ptr[int32](11),
		// 						WorkerConnections: to.Ptr[int32](18),
		// 						WorkerConnectionsLastUpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.777Z"); return t}()),
		// 						ContainerResyncTrigger: to.Ptr[int32](13),
		// 						ImageURI: to.Ptr("ezbwyfaainatxtlplyoailzxlhxy"),
		// 					},
		// 					AdditionalCacheNodeProperties: &armconnectedcache.AdditionalCacheNodeProperties{
		// 						DriveConfiguration: []*armconnectedcache.CacheNodeDriveConfiguration{
		// 							{
		// 								PhysicalPath: to.Ptr("/mcc"),
		// 								SizeInGb: to.Ptr[int32](500),
		// 								CacheNumber: to.Ptr[int32](1),
		// 								NginxMapping: to.Ptr("lijygenjq"),
		// 							},
		// 						},
		// 						BgpConfiguration: &armconnectedcache.BgpConfiguration{
		// 							AsnToIPAddressMapping: to.Ptr("pafcimhoog"),
		// 						},
		// 						CacheNodePropertiesDetailsIssuesList: []*string{
		// 							to.Ptr("ex"),
		// 						},
		// 						AggregatedStatusDetails: to.Ptr("emougql"),
		// 						AggregatedStatusText: to.Ptr("xcasvndgkob"),
		// 						AggregatedStatusCode: to.Ptr[int32](22),
		// 						ProductVersion: to.Ptr("oxhqgwlhgnuf"),
		// 						IsProvisioned: to.Ptr(true),
		// 						CacheNodeStateDetailedText: to.Ptr("ufvomikgfnmnj"),
		// 						CacheNodeStateShortText: to.Ptr("orfpuvrevhrxsaasddazigglq"),
		// 						CacheNodeState: to.Ptr[int32](9),
		// 						ProxyURLConfiguration: &armconnectedcache.ProxyURLConfiguration{
		// 							ProxyURL: to.Ptr("hplstyg"),
		// 						},
		// 						OptionalProperty1: to.Ptr("hvpmt"),
		// 						OptionalProperty2: to.Ptr("talanelmsgxvksrzoeeontqkjzbpv"),
		// 						OptionalProperty3: to.Ptr("bxkoxq"),
		// 						OptionalProperty4: to.Ptr("pqlkcekupusoc"),
		// 						OptionalProperty5: to.Ptr("nyvvmrjigqdufzjdvazdca"),
		// 					},
		// 					StatusCode: to.Ptr("1"),
		// 					StatusText: to.Ptr("Success"),
		// 					StatusDetails: to.Ptr("djruqvptzxak"),
		// 					Status: to.Ptr("tnyrntqvazk"),
		// 					Error: &armconnectedcache.ErrorDetail{
		// 						Code: to.Ptr("dkvgvtftpsjsbhlnapvihefxneoggs"),
		// 						Message: to.Ptr("okakgyfnmyob"),
		// 						Details: []*armconnectedcache.ErrorDetail{
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key4171": to.Ptr("qtjlszkawsdujzpgohsbw"),
		// 				},
		// 				SystemData: &armconnectedcache.SystemData{
		// 					CreatedBy: to.Ptr("gambtqj"),
		// 					CreatedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("qomgaceiessgnuogz"),
		// 					LastModifiedByType: to.Ptr(armconnectedcache.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.771Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

package armconnectedcache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
)

// Generated from example definition: 2023-05-01-preview/EnterpriseMccCustomers_Update_MaximumSet_Gen.json
func ExampleEnterpriseMccCustomersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedcache.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnterpriseMccCustomersClient().Update(ctx, "rgConnectedCache", "MccRPTest1", armconnectedcache.PatchResource{
		Tags: map[string]*string{
			"key1878": to.Ptr("warz"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconnectedcache.EnterpriseMccCustomersClientUpdateResponse{
	// 	EnterpriseMccCustomerResource: &armconnectedcache.EnterpriseMccCustomerResource{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/rgConnectedCache/providers/Microsoft.ConnectedCache/ispCustomers/MccRPTest2"),
	// 		Name: to.Ptr("MCCTPTest2"),
	// 		Type: to.Ptr("Microsoft.ConnectedCache/ispCustomers"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armconnectedcache.CustomerProperty{
	// 			ProvisioningState: to.Ptr(armconnectedcache.ProvisioningStateSucceeded),
	// 			Customer: &armconnectedcache.CustomerEntity{
	// 				FullyQualifiedResourceID: to.Ptr("uqsbtgae"),
	// 				CustomerID: to.Ptr("nqxuzybu"),
	// 				CustomerName: to.Ptr("mkpzynfqihnjfdbaqbqwyhd"),
	// 				ContactEmail: to.Ptr("xquos"),
	// 				ContactPhone: to.Ptr("vue"),
	// 				ContactName: to.Ptr("wxyqjoyoscmvimgwhpitxky"),
	// 				IsEntitled: to.Ptr(true),
	// 				ReleaseVersion: to.Ptr[int32](20),
	// 				CreateAsyncOperationID: to.Ptr("zjpvgirzxecwmnfyofqkikst"),
	// 				DeleteAsyncOperationID: to.Ptr("ajtdyoyecybeaxzyztjkvvtx"),
	// 				ClientTenantID: to.Ptr("fproidkpgvpdnac"),
	// 				SynchWithAzureAttemptsCount: to.Ptr[int32](17),
	// 				LastSyncWithAzureTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.773Z"); return t}()),
	// 				IsEnterpriseManaged: to.Ptr(true),
	// 				ShouldMigrate: to.Ptr(true),
	// 				ResendSignupCode: to.Ptr(true),
	// 				VerifySignupCode: to.Ptr(true),
	// 			},
	// 			AdditionalCustomerProperties: &armconnectedcache.AdditionalCustomerProperties{
	// 				PeeringDbLastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.773Z"); return t}()),
	// 				CustomerPropertiesOverviewCacheEfficiency: to.Ptr[float32](20),
	// 				CustomerPropertiesOverviewAverageEgressMbps: to.Ptr[float32](8),
	// 				CustomerPropertiesOverviewAverageMissMbps: to.Ptr[float32](19),
	// 				CustomerPropertiesOverviewEgressMbpsMax: to.Ptr[float32](15),
	// 				CustomerPropertiesOverviewEgressMbpsMaxDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.773Z"); return t}()),
	// 				CustomerPropertiesOverviewMissMbpsMax: to.Ptr[float32](28),
	// 				CustomerPropertiesOverviewMissMbpsMaxDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.773Z"); return t}()),
	// 				CustomerPropertiesOverviewCacheNodesHealthyCount: to.Ptr[int32](24),
	// 				CustomerPropertiesOverviewCacheNodesUnhealthyCount: to.Ptr[int32](15),
	// 				SignupStatus: to.Ptr(true),
	// 				SignupStatusCode: to.Ptr[int32](21),
	// 				SignupStatusText: to.Ptr("dccv"),
	// 				SignupPhaseStatusCode: to.Ptr[int32](4),
	// 				SignupPhaseStatusText: to.Ptr("q"),
	// 				PeeringDbLastUpdateDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.773Z"); return t}()),
	// 				CustomerOrgName: to.Ptr("nshrwjhhggmdaqlpc"),
	// 				CustomerEmail: to.Ptr("zdjgibsidydyzm"),
	// 				CustomerTransitAsn: to.Ptr("habgklnxqzmozqpazoyejwiphezpi"),
	// 				CustomerAsn: to.Ptr("hgrelgnrtdkleisnepfolu"),
	// 				CustomerAsnEstimatedEgressPeekGbps: to.Ptr[float32](10),
	// 				CustomerEntitlementSKUID: to.Ptr("b"),
	// 				CustomerEntitlementSKUGUID: to.Ptr("rvzmdpxyflgqetvpwupnfaxsweiiz"),
	// 				CustomerEntitlementSKUName: to.Ptr("waaqfijr"),
	// 				CustomerEntitlementExpiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T00:54:04.773Z"); return t}()),
	// 				OptionalProperty1: to.Ptr("qhmwxza"),
	// 				OptionalProperty2: to.Ptr("l"),
	// 				OptionalProperty3: to.Ptr("mblwwvbie"),
	// 				OptionalProperty4: to.Ptr("vzuek"),
	// 				OptionalProperty5: to.Ptr("fzjodscdfcdr"),
	// 			},
	// 			StatusCode: to.Ptr("jax"),
	// 			StatusText: to.Ptr("vsqydgruhuwuyipsplylgiqmkcv"),
	// 			StatusDetails: to.Ptr("wmtksbahlbxrzaksogdbozfi"),
	// 			Status: to.Ptr("rhfjbcr"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key3379": to.Ptr("dpyqeaqhcnutzezom"),
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

package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Organization_ListRegions_MaximumSet_Gen.json
func ExampleOrganizationClient_ListRegions_organizationListRegionsMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationClient().ListRegions(ctx, "rgconfluent", "bnu", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"key8083": to.Ptr("ft"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconfluent.OrganizationClientListRegionsResponse{
	// 	ListRegionsSuccessResponse: &armconfluent.ListRegionsSuccessResponse{
	// 		Data: []*armconfluent.RegionRecord{
	// 			{
	// 				Kind: to.Ptr("muabtqsllrtrsecdvjgvntpke"),
	// 				ID: to.Ptr("okvhazbpnpaxqchtznncekkwucejaa"),
	// 				Properties: &armconfluent.RegionProperties{
	// 					Metadata: &armconfluent.SCMetadataEntity{
	// 						Self: to.Ptr("bnbnbarlsvfifpzcnsnplf"),
	// 						ResourceName: to.Ptr("ciadqmxlpgllibvkz"),
	// 						CreatedTimestamp: to.Ptr("ouqjivxfggaxzrsmxm"),
	// 						UpdatedTimestamp: to.Ptr("ctrngbppcxdpzmp"),
	// 						DeletedTimestamp: to.Ptr("gn"),
	// 					},
	// 					Spec: &armconfluent.RegionSpecEntity{
	// 						Name: to.Ptr("futevtizlzyhjjrvub"),
	// 						Cloud: to.Ptr("uagbogdrkcphfjcckxfoupusycw"),
	// 						RegionName: to.Ptr("pzmdsnasogrogtatsdrwfyvetrhx"),
	// 						Packages: []*string{
	// 							to.Ptr("dnscn"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

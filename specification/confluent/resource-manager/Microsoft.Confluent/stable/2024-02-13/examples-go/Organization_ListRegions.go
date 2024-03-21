package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_ListRegions.json
func ExampleOrganizationClient_ListRegions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationClient().ListRegions(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"cloud":    to.Ptr("azure"),
			"packages": to.Ptr("ADVANCED,ESSENTIALS"),
			"region":   to.Ptr("eastus"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListRegionsSuccessResponse = armconfluent.ListRegionsSuccessResponse{
	// 	Data: []*armconfluent.RegionRecord{
	// 		{
	// 			ID: to.Ptr("sgreg-7"),
	// 			Kind: to.Ptr("Region"),
	// 			Properties: &armconfluent.RegionProperties{
	// 				Metadata: &armconfluent.SCMetadataEntity{
	// 					CreatedTimestamp: to.Ptr("2023-04-06T10:24:09.840788Z"),
	// 					ResourceName: to.Ptr("crn://confluent.cloud/schema-registry-region=sgreg-7"),
	// 					Self: to.Ptr("https://api.stag.cpdev.cloud/srcm/v2/regions/sgreg-7"),
	// 					UpdatedTimestamp: to.Ptr("2023-04-06T10:24:09.840788Z"),
	// 				},
	// 				Spec: &armconfluent.RegionSpecEntity{
	// 					Name: to.Ptr("Iowa (centralus)"),
	// 					Cloud: to.Ptr("AZURE"),
	// 					Packages: []*string{
	// 						to.Ptr("ADVANCED")},
	// 						RegionName: to.Ptr("centralus"),
	// 					},
	// 				},
	// 		}},
	// 	}
}

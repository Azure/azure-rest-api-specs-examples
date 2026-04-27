package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Organization_GetClusterAPIKey_MaximumSet_Gen.json
func ExampleOrganizationClient_GetClusterAPIKey_organizationGetClusterApiKeyMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationClient().GetClusterAPIKey(ctx, "rgconfluent", "puauqgrwsfgmolfhazfjcavnj", "xxsquwnsllkkzuyzlhdxdl", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconfluent.OrganizationClientGetClusterAPIKeyResponse{
	// 	APIKeyRecord: &armconfluent.APIKeyRecord{
	// 		Kind: to.Ptr("g"),
	// 		ID: to.Ptr("byr"),
	// 		Properties: &armconfluent.APIKeyProperties{
	// 			Metadata: &armconfluent.SCMetadataEntity{
	// 				Self: to.Ptr("bnbnbarlsvfifpzcnsnplf"),
	// 				ResourceName: to.Ptr("ciadqmxlpgllibvkz"),
	// 				CreatedTimestamp: to.Ptr("ouqjivxfggaxzrsmxm"),
	// 				UpdatedTimestamp: to.Ptr("ctrngbppcxdpzmp"),
	// 				DeletedTimestamp: to.Ptr("gn"),
	// 			},
	// 			Spec: &armconfluent.APIKeySpecEntity{
	// 				Description: to.Ptr("cgzqydt"),
	// 				Name: to.Ptr("zbsz"),
	// 				Resource: &armconfluent.APIKeyResourceEntity{
	// 					ID: to.Ptr("ncxbrkuafkzqmivoqkewsji"),
	// 					Environment: to.Ptr("spwmtilczuvrfhyykfyt"),
	// 					Related: to.Ptr("yhgzvelqihdvuqxgrnvvos"),
	// 					ResourceName: to.Ptr("cccrljvydtuuzknqrpy"),
	// 					Kind: to.Ptr("q"),
	// 				},
	// 				Owner: &armconfluent.APIKeyOwnerEntity{
	// 					ID: to.Ptr("nlbvfmehhymt"),
	// 					Related: to.Ptr("frhqdjzhi"),
	// 					ResourceName: to.Ptr("jdn"),
	// 					Kind: to.Ptr("rujp"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

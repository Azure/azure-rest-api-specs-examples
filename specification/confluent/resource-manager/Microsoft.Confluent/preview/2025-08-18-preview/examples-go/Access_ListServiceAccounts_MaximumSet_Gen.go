package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Access_ListServiceAccounts_MaximumSet_Gen.json
func ExampleAccessClient_ListServiceAccounts_accessListServiceAccountsMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListServiceAccounts(ctx, "rgconfluent", "go", armconfluent.ListAccessRequestModel{
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
	// res = armconfluent.AccessClientListServiceAccountsResponse{
	// 	AccessListServiceAccountsSuccessResponse: &armconfluent.AccessListServiceAccountsSuccessResponse{
	// 		Kind: to.Ptr("puztaopmygsvwfwufjvycnok"),
	// 		Metadata: &armconfluent.ListMetadata{
	// 			First: to.Ptr("gxyzsfkulrbudspwhbmhot"),
	// 			Last: to.Ptr("b"),
	// 			Prev: to.Ptr("ofxjgjyzkpqsqqjvd"),
	// 			Next: to.Ptr("rohrcnueuolepyisnh"),
	// 			TotalSize: to.Ptr[int32](4),
	// 		},
	// 		Data: []*armconfluent.ServiceAccountRecord{
	// 			{
	// 				Kind: to.Ptr("tozqpdeniq"),
	// 				ID: to.Ptr("uayitlece"),
	// 				Metadata: &armconfluent.MetadataEntity{
	// 					Self: to.Ptr("lnzrk"),
	// 					ResourceName: to.Ptr("vgzwrzjztaeord"),
	// 					CreatedAt: to.Ptr("mqxjit"),
	// 					UpdatedAt: to.Ptr("qmeqwokdgucbocytlfsticz"),
	// 					DeletedAt: to.Ptr("mhkhttriidtlq"),
	// 				},
	// 				DisplayName: to.Ptr("nyhanwbkwn"),
	// 				Description: to.Ptr("xcosvrzaslrokmxtdalllnc"),
	// 			},
	// 		},
	// 	},
	// }
}

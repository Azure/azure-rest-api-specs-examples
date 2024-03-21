package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_ServiceAccountsList.json
func ExampleAccessClient_ListServiceAccounts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListServiceAccounts(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"pageSize":  to.Ptr("10"),
			"pageToken": to.Ptr("asc4fts4ft"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessListServiceAccountsSuccessResponse = armconfluent.AccessListServiceAccountsSuccessResponse{
	// 	Data: []*armconfluent.ServiceAccountRecord{
	// 		{
	// 			Description: to.Ptr("Doc's repair bot for the DeLorean"),
	// 			DisplayName: to.Ptr("DeLorean_auto_repair"),
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("ServiceAccount"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/service-account=sa-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts/sa-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 	}},
	// 	Kind: to.Ptr("ServiceAccountList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts"),
	// 		Last: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}

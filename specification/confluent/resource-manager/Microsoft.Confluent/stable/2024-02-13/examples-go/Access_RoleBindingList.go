package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_RoleBindingList.json
func ExampleAccessClient_ListRoleBindings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListRoleBindings(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
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
	// res.AccessListRoleBindingsSuccessResponse = armconfluent.AccessListRoleBindingsSuccessResponse{
	// 	Data: []*armconfluent.RoleBindingRecord{
	// 		{
	// 			CrnPattern: to.Ptr("crn://confluent.cloud/organization=1111aaaa-11aa-11aa-11aa-111111aaaaaa/environment=env-aaa1111/cloud-cluster=lkc-1111aaa"),
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("RoleBinding"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/role-binding=rb-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings/rb-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 			Principal: to.Ptr("User:u-111aaa"),
	// 			RoleName: to.Ptr("CloudClusterAdmin"),
	// 	}},
	// 	Kind: to.Ptr("RoleBindingList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings"),
	// 		Last: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}

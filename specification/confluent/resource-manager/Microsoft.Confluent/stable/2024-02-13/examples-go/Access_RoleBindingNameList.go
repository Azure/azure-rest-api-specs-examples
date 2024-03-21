package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_RoleBindingNameList.json
func ExampleAccessClient_ListRoleBindingNameList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListRoleBindingNameList(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"crn_pattern": to.Ptr("crn://confluent.cloud/organization=1aa7de07-298e-479c-8f2f-16ac91fd8e76"),
			"namespace":   to.Ptr("public,dataplane,networking,identity,datagovernance,connect,streamcatalog,pipelines,ksql"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessRoleBindingNameListSuccessResponse = armconfluent.AccessRoleBindingNameListSuccessResponse{
	// 	Data: []*string{
	// 		to.Ptr("MetricsViewer"),
	// 		to.Ptr("OrganizationAdmin"),
	// 		to.Ptr("Operator"),
	// 		to.Ptr("NetworkAdmin"),
	// 		to.Ptr("ResourceKeyAdmin"),
	// 		to.Ptr("AccountAdmin")},
	// 		Kind: to.Ptr("RoleDisplayNameList"),
	// 		Metadata: &armconfluent.ListMetadata{
	// 			First: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings"),
	// 			Last: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings?page_token=bcAOehAY8F16YD84Z1wT"),
	// 			Next: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 			Prev: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings?page_token=YIXRY97wWYmwzrax4dld"),
	// 			TotalSize: to.Ptr[int32](123),
	// 		},
	// 	}
}

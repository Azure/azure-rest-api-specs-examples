package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_EnvironmentList.json
func ExampleAccessClient_ListEnvironments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListEnvironments(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
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
	// res.AccessListEnvironmentsSuccessResponse = armconfluent.AccessListEnvironmentsSuccessResponse{
	// 	Data: []*armconfluent.EnvironmentRecord{
	// 		{
	// 			DisplayName: to.Ptr("prod-finance01"),
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("Environment"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=e-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/org/v2/environments/e-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 	}},
	// 	Kind: to.Ptr("EnvironmentList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/org/v2/environments"),
	// 		Last: to.Ptr("https://api.confluent.cloud/org/v2/environments?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/org/v2/environments?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/org/v2/environments?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}

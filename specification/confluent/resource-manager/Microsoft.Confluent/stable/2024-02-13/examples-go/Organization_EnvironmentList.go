package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_EnvironmentList.json
func ExampleOrganizationClient_NewListEnvironmentsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationClient().NewListEnvironmentsPager("myResourceGroup", "myOrganization", &armconfluent.OrganizationClientListEnvironmentsOptions{PageSize: to.Ptr[int32](10),
		PageToken: nil,
	})
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
		// page.GetEnvironmentsResponse = armconfluent.GetEnvironmentsResponse{
		// 	Value: []*armconfluent.SCEnvironmentRecord{
		// 		{
		// 			Name: to.Ptr("prod-finance01"),
		// 			ID: to.Ptr("dlz-f3a90de"),
		// 			Kind: to.Ptr("Environment"),
		// 			Properties: &armconfluent.EnvironmentProperties{
		// 				Metadata: &armconfluent.SCMetadataEntity{
		// 					CreatedTimestamp: to.Ptr("2006-01-02T15:04:05-07:00"),
		// 					DeletedTimestamp: to.Ptr("2006-01-02T15:04:05-07:00"),
		// 					ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=e-12345"),
		// 					Self: to.Ptr("https://api.confluent.cloud/org/v2/environments/e-12345"),
		// 					UpdatedTimestamp: to.Ptr("2006-01-02T15:04:05-07:00"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

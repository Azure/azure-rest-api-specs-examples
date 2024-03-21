package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Validations_ValidateOrganizationsV2.json
func ExampleValidationsClient_ValidateOrganizationV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewValidationsClient().ValidateOrganizationV2(ctx, "myResourceGroup", "myOrganization", armconfluent.OrganizationResource{
		Location: to.Ptr("West US"),
		Properties: &armconfluent.OrganizationResourceProperties{
			OfferDetail: &armconfluent.OfferDetail{
				ID:             to.Ptr("string"),
				PlanID:         to.Ptr("string"),
				PlanName:       to.Ptr("string"),
				PrivateOfferID: to.Ptr("string"),
				PrivateOfferIDs: []*string{
					to.Ptr("string")},
				PublisherID: to.Ptr("string"),
				TermUnit:    to.Ptr("string"),
			},
			UserDetail: &armconfluent.UserDetail{
				AADEmail:          to.Ptr("abc@microsoft.com"),
				EmailAddress:      to.Ptr("abc@microsoft.com"),
				FirstName:         to.Ptr("string"),
				LastName:          to.Ptr("string"),
				UserPrincipalName: to.Ptr("abc@microsoft.com"),
			},
		},
		Tags: map[string]*string{
			"Environment": to.Ptr("Dev"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ValidationResponse = armconfluent.ValidationResponse{
	// 	Info: map[string]*string{
	// 		"displayMessage": to.Ptr("This is display message"),
	// 	},
	// }
}

package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Organization_Update.json
func ExampleOrganizationClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationClient().Update(ctx, "myResourceGroup", "myOrganization", &armconfluent.OrganizationClientUpdateOptions{Body: &armconfluent.OrganizationResourceUpdate{
		Tags: map[string]*string{
			"client": to.Ptr("dev-client"),
			"env":    to.Ptr("dev"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OrganizationResource = armconfluent.OrganizationResource{
	// 	Name: to.Ptr("myOrganization"),
	// 	Type: to.Ptr("Microsoft.Confluent/organizations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganization"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armconfluent.OrganizationResourceProperties{
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		OfferDetail: &armconfluent.OfferDetail{
	// 			ID: to.Ptr("string"),
	// 			PlanID: to.Ptr("string"),
	// 			PlanName: to.Ptr("string"),
	// 			PublisherID: to.Ptr("string"),
	// 			Status: to.Ptr(armconfluent.SaaSOfferStatusStarted),
	// 			TermUnit: to.Ptr("string"),
	// 		},
	// 		OrganizationID: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armconfluent.ProvisionStateSucceeded),
	// 		SsoURL: to.Ptr("string"),
	// 		UserDetail: &armconfluent.UserDetail{
	// 			EmailAddress: to.Ptr("contoso@microsoft.com"),
	// 			FirstName: to.Ptr("string"),
	// 			LastName: to.Ptr("string"),
	// 		},
	// 	},
	// 	SystemData: &armconfluent.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

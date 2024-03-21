package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_Create.json
func ExampleOrganizationClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationClient().BeginCreate(ctx, "myResourceGroup", "myOrganization", &armconfluent.OrganizationClientBeginCreateOptions{Body: &armconfluent.OrganizationResource{
		Location: to.Ptr("West US"),
		Properties: &armconfluent.OrganizationResourceProperties{
			LinkOrganization: &armconfluent.LinkOrganization{
				Token: to.Ptr("string"),
			},
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
				AADEmail:          to.Ptr("contoso@microsoft.com"),
				EmailAddress:      to.Ptr("contoso@microsoft.com"),
				FirstName:         to.Ptr("string"),
				LastName:          to.Ptr("string"),
				UserPrincipalName: to.Ptr("contoso@microsoft.com"),
			},
		},
		Tags: map[string]*string{
			"Environment": to.Ptr("Dev"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
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
	// 			PrivateOfferID: to.Ptr("string"),
	// 			PrivateOfferIDs: []*string{
	// 				to.Ptr("string")},
	// 				PublisherID: to.Ptr("string"),
	// 				Status: to.Ptr(armconfluent.SaaSOfferStatusStarted),
	// 				TermUnit: to.Ptr("string"),
	// 			},
	// 			OrganizationID: to.Ptr("string"),
	// 			ProvisioningState: to.Ptr(armconfluent.ProvisionStateSucceeded),
	// 			SsoURL: to.Ptr("string"),
	// 			UserDetail: &armconfluent.UserDetail{
	// 				AADEmail: to.Ptr("contoso@microsoft.com"),
	// 				EmailAddress: to.Ptr("contoso@microsoft.com"),
	// 				FirstName: to.Ptr("string"),
	// 				LastName: to.Ptr("string"),
	// 				UserPrincipalName: to.Ptr("contoso@microsoft.com"),
	// 			},
	// 		},
	// 		SystemData: &armconfluent.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"Environment": to.Ptr("Dev"),
	// 		},
	// 	}
}

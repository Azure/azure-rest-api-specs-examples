package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_ListBySubscription.json
func ExampleOrganizationClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationClient().NewListBySubscriptionPager(nil)
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
		// page.OrganizationResourceListResult = armconfluent.OrganizationResourceListResult{
		// 	Value: []*armconfluent.OrganizationResource{
		// 		{
		// 			Name: to.Ptr("myOrganizations"),
		// 			Type: to.Ptr("Microsoft.Confluent/organizations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganizations"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armconfluent.OrganizationResourceProperties{
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
		// 				OfferDetail: &armconfluent.OfferDetail{
		// 					ID: to.Ptr("string"),
		// 					PlanID: to.Ptr("string"),
		// 					PlanName: to.Ptr("string"),
		// 					PrivateOfferID: to.Ptr("string"),
		// 					PrivateOfferIDs: []*string{
		// 						to.Ptr("string")},
		// 						PublisherID: to.Ptr("string"),
		// 						Status: to.Ptr(armconfluent.SaaSOfferStatusStarted),
		// 						TermUnit: to.Ptr("string"),
		// 					},
		// 					OrganizationID: to.Ptr("string"),
		// 					ProvisioningState: to.Ptr(armconfluent.ProvisionStateSucceeded),
		// 					SsoURL: to.Ptr("string"),
		// 					UserDetail: &armconfluent.UserDetail{
		// 						AADEmail: to.Ptr("contoso@microsoft.com"),
		// 						EmailAddress: to.Ptr("contoso@microsoft.com"),
		// 						FirstName: to.Ptr("string"),
		// 						LastName: to.Ptr("string"),
		// 						UserPrincipalName: to.Ptr("contoso@microsoft.com"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"Environment": to.Ptr("Dev"),
		// 				},
		// 		}},
		// 	}
	}
}

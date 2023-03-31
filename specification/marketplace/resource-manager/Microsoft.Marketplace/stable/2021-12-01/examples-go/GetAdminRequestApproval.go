package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetAdminRequestApproval.json
func ExamplePrivateStoreClient_GetAdminRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().GetAdminRequestApproval(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "marketplacetestthirdparty.md-test-third-party-2", "marketplacetestthirdparty", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AdminRequestApprovalsResource = armmarketplace.AdminRequestApprovalsResource{
	// 	Name: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
	// 	Type: to.Ptr("/providers/Microsoft.Marketplace/privateStores/adminRequestApprovals"),
	// 	ID: to.Ptr("/providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/adminRequestApprovals/marketplacetestthirdparty.md-test-third-party-2"),
	// 	SystemData: &armmarketplace.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T10:23:17.6571572+02:00"); return t}()),
	// 		CreatedBy: to.Ptr("user@somedoamin.com"),
	// 		CreatedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T10:23:17.6571572+02:00"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@somedoamin.com"),
	// 		LastModifiedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 	},
	// 	Properties: &armmarketplace.AdminRequestApprovalProperties{
	// 		AdminAction: to.Ptr(armmarketplace.AdminAction("Pending")),
	// 		CollectionIDs: []*string{
	// 		},
	// 		DisplayName: to.Ptr("Offer display name"),
	// 		OfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
	// 		Plans: []*armmarketplace.PlanRequesterDetails{
	// 			{
	// 				PlanDisplayName: to.Ptr("Plan display name A"),
	// 				PlanID: to.Ptr("testA"),
	// 				Requesters: []*armmarketplace.UserRequestDetails{
	// 					{
	// 						Date: to.Ptr("2021-02-01T11:42:12.9526511+02:00"),
	// 						Justification: to.Ptr("Because I want to...."),
	// 						SubscriptionID: to.Ptr("1fs4r5t6-5a1e-4913-b849-2c68880e03c2"),
	// 						SubscriptionName: to.Ptr("Test subscription 1"),
	// 						User: to.Ptr("testUser3"),
	// 				}},
	// 			},
	// 			{
	// 				PlanDisplayName: to.Ptr("*  (this means the user requested any plan, here you will get only *)"),
	// 				PlanID: to.Ptr("*"),
	// 				Requesters: []*armmarketplace.UserRequestDetails{
	// 					{
	// 						Date: to.Ptr("2021-02-01T11:42:12.9526511+02:00"),
	// 						Justification: to.Ptr("try me :)"),
	// 						SubscriptionID: to.Ptr("4ca4753c-5a1e-4913-b849-2c68880e03c2"),
	// 						SubscriptionName: to.Ptr("Test subscription 2"),
	// 						User: to.Ptr("testUser3"),
	// 				}},
	// 		}},
	// 		PublisherID: to.Ptr("marketplacetestthirdparty"),
	// 	},
	// }
}

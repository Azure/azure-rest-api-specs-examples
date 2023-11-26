package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/UpdateAdminRequestApproval.json
func ExamplePrivateStoreClient_UpdateAdminRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().UpdateAdminRequestApproval(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "marketplacetestthirdparty.md-test-third-party-2", &armmarketplace.PrivateStoreClientUpdateAdminRequestApprovalOptions{Payload: &armmarketplace.AdminRequestApprovalsResource{
		Properties: &armmarketplace.AdminRequestApprovalProperties{
			AdminAction: to.Ptr(armmarketplace.AdminActionApproved),
			ApprovedPlans: []*string{
				to.Ptr("testPlan")},
			CollectionIDs: []*string{
				to.Ptr("f8ee227e-85d7-477d-abbf-854d6decaf70"),
				to.Ptr("39246ad6-c521-4fed-8de7-77dede2e873f")},
			Comment:     to.Ptr("I'm ok with that"),
			PublisherID: to.Ptr("marketplacetestthirdparty"),
		},
	},
	})
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
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-05T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("test@somedomain.com"),
	// 		CreatedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-05T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("testuser@somedomail.com"),
	// 		LastModifiedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 	},
	// 	Properties: &armmarketplace.AdminRequestApprovalProperties{
	// 		AdminAction: to.Ptr(armmarketplace.AdminActionApproved),
	// 		Administrator: to.Ptr("admin@someDomain.com"),
	// 		ApprovedPlans: []*string{
	// 			to.Ptr("testPlan")},
	// 			CollectionIDs: []*string{
	// 				to.Ptr("f8ee227e-85d7-477d-abbf-854d6decaf70"),
	// 				to.Ptr("39246ad6-c521-4fed-8de7-77dede2e873f")},
	// 				Comment: to.Ptr("I'm ok with that"),
	// 				DisplayName: to.Ptr("Offer display name"),
	// 				OfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
	// 				PublisherID: to.Ptr("marketplacetestthirdparty"),
	// 			},
	// 		}
}

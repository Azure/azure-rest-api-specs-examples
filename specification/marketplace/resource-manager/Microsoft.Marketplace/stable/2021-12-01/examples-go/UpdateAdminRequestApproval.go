package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/UpdateAdminRequestApproval.json
func ExamplePrivateStoreClient_UpdateAdminRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateAdminRequestApproval(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"marketplacetestthirdparty.md-test-third-party-2",
		&armmarketplace.PrivateStoreClientUpdateAdminRequestApprovalOptions{Payload: &armmarketplace.AdminRequestApprovalsResource{
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
	// TODO: use response item
	_ = res
}

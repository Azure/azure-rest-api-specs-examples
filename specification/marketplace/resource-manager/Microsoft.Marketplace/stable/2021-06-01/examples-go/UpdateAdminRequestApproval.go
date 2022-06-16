package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/UpdateAdminRequestApproval.json
func ExamplePrivateStoreClient_UpdateAdminRequestApproval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	res, err := client.UpdateAdminRequestApproval(ctx,
		"<private-store-id>",
		"<admin-request-approval-id>",
		&armmarketplace.PrivateStoreClientUpdateAdminRequestApprovalOptions{Payload: &armmarketplace.AdminRequestApprovalsResource{
			Properties: &armmarketplace.AdminRequestApprovalProperties{
				AdminAction: armmarketplace.AdminAction("Approved").ToPtr(),
				ApprovedPlans: []*string{
					to.StringPtr("testPlan")},
				CollectionIDs: []*string{
					to.StringPtr("f8ee227e-85d7-477d-abbf-854d6decaf70"),
					to.StringPtr("39246ad6-c521-4fed-8de7-77dede2e873f")},
				Comment:     to.StringPtr("<comment>"),
				PublisherID: to.StringPtr("<publisher-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateStoreClientUpdateAdminRequestApprovalResult)
}

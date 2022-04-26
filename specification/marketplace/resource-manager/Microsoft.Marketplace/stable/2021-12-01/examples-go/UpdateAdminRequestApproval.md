Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplace%2Farmmarketplace%2Fv0.4.0/sdk/resourcemanager/marketplace/armmarketplace/README.md) on how to add the SDK to your project and authenticate.

```go
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
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.UpdateAdminRequestApproval(ctx,
		"<private-store-id>",
		"<admin-request-approval-id>",
		&armmarketplace.PrivateStoreClientUpdateAdminRequestApprovalOptions{Payload: &armmarketplace.AdminRequestApprovalsResource{
			Properties: &armmarketplace.AdminRequestApprovalProperties{
				AdminAction: to.Ptr(armmarketplace.AdminActionApproved),
				ApprovedPlans: []*string{
					to.Ptr("testPlan")},
				CollectionIDs: []*string{
					to.Ptr("f8ee227e-85d7-477d-abbf-854d6decaf70"),
					to.Ptr("39246ad6-c521-4fed-8de7-77dede2e873f")},
				Comment:     to.Ptr("<comment>"),
				PublisherID: to.Ptr("<publisher-id>"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

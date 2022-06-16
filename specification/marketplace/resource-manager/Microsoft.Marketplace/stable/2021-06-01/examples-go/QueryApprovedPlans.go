package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/QueryApprovedPlans.json
func ExamplePrivateStoreClient_QueryApprovedPlans() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	res, err := client.QueryApprovedPlans(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreClientQueryApprovedPlansOptions{Payload: &armmarketplace.QueryApprovedPlansPayload{
			Properties: &armmarketplace.QueryApprovedPlans{
				OfferID: to.StringPtr("<offer-id>"),
				PlanIDs: []*string{
					to.StringPtr("testPlanA"),
					to.StringPtr("testPlanB"),
					to.StringPtr("testPlanC")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateStoreClientQueryApprovedPlansResult)
}

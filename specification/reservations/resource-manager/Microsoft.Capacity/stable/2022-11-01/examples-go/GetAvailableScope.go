package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetAvailableScope.json
func ExampleReservationClient_BeginAvailableScopes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armreservations.NewReservationClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginAvailableScopes(ctx, "276e7ae4-84d0-4da6-ab4b-d6b94f3557da", "356e7ae4-84d0-4da6-ab4b-d6b94f3557da", armreservations.AvailableScopeRequest{
		Properties: &armreservations.AvailableScopeRequestProperties{
			Scopes: []*string{
				to.Ptr("/subscriptions/efc7c997-7700-4a74-b731-55aec16c15e9")},
		},
	}, nil)
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
	// res.AvailableScopeProperties = armreservations.AvailableScopeProperties{
	// 	Properties: &armreservations.SubscriptionScopeProperties{
	// 		Scopes: []*armreservations.ScopeProperties{
	// 			{
	// 				Scope: to.Ptr("/subscriptions/efc7c997-7700-4a74-b731-55aec16c15e9"),
	// 				Valid: to.Ptr(true),
	// 		}},
	// 	},
	// }
}

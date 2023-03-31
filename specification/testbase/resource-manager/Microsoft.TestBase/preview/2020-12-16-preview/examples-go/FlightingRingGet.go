package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/FlightingRingGet.json
func ExampleFlightingRingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFlightingRingsClient().Get(ctx, "contoso-rg", "contoso-testBaseAccount", "Insider-Beta-Channel", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FlightingRingResource = armtestbase.FlightingRingResource{
	// 	Name: to.Ptr("Insider-Beta-Channel"),
	// 	Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/flightingRings"),
	// 	ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount/flightingRings/Insider-Beta-Channel"),
	// 	Properties: &armtestbase.FlightingRingProperties{
	// 		ActualFlightingRingName: to.Ptr("Insider Beta Channel"),
	// 	},
	// }
}

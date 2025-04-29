package armcarbonoptimization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/carbonoptimization/armcarbonoptimization"
)

// Generated from example definition: 2025-04-01/carbonEmissionsDataAvailableDateRange.json
func ExampleCarbonServiceClient_QueryCarbonEmissionDataAvailableDateRange() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcarbonoptimization.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCarbonServiceClient().QueryCarbonEmissionDataAvailableDateRange(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcarbonoptimization.CarbonServiceClientQueryCarbonEmissionDataAvailableDateRangeResponse{
	// 	CarbonEmissionDataAvailableDateRange: &armcarbonoptimization.CarbonEmissionDataAvailableDateRange{
	// 		StartDate: to.Ptr("2023-01-01"),
	// 		EndDate: to.Ptr("2023-05-01"),
	// 	},
	// }
}

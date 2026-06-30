package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: 2024-12-30/ServerValidateEstimateHighAvailability.json
func ExampleServersClient_ValidateEstimateHighAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().ValidateEstimateHighAvailability(ctx, "TestGroup", "testserver", armmysqlflexibleservers.HighAvailabilityValidationEstimation{
		ExpectedStandbyAvailabilityZone: to.Ptr("1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmysqlflexibleservers.ServersClientValidateEstimateHighAvailabilityResponse{
	// 	HighAvailabilityValidationEstimation: armmysqlflexibleservers.HighAvailabilityValidationEstimation{
	// 		EstimatedDowntime: to.Ptr[int32](0),
	// 		ExpectedStandbyAvailabilityZone: to.Ptr("1"),
	// 		ScheduledStandbyAvailabilityZone: to.Ptr("1"),
	// 	},
	// }
}

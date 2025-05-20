package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/TuningConfiguration_ListSessions.json
func ExampleTuningConfigurationClient_NewListSessionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTuningConfigurationClient().NewListSessionsPager("testrg", "testserver", armpostgresqlflexibleservers.TuningOptionEnumConfiguration, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SessionsListResult = armpostgresqlflexibleservers.SessionsListResult{
		// 	Value: []*armpostgresqlflexibleservers.SessionResource{
		// 		{
		// 			PostTuningAqr: to.Ptr("1.818203395454569"),
		// 			PostTuningTps: to.Ptr("2.594203762997024"),
		// 			PreTuningAqr: to.Ptr("1.9069379374999662"),
		// 			PreTuningTps: to.Ptr("2.592924534523614"),
		// 			SessionID: to.Ptr("29f70be5-74ac-498b-949c-38b0997b4290"),
		// 			SessionStartTime: to.Ptr("09/26/2024 02:13:05"),
		// 			Status: to.Ptr("completed"),
		// 	}},
		// }
	}
}

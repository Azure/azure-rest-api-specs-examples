package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/TuningConfiguration_ListSessionDetails.json
func ExampleTuningConfigurationClient_NewListSessionDetailsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTuningConfigurationClient().NewListSessionDetailsPager("testrg", "testserver", armpostgresqlflexibleservers.TuningOptionEnumConfiguration, "oooooooo-oooo-oooo-oooo-oooooooooooo", nil)
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
		// page.SessionDetailsListResult = armpostgresqlflexibleservers.SessionDetailsListResult{
		// 	Value: []*armpostgresqlflexibleservers.SessionDetailsResource{
		// 		{
		// 			AppliedConfiguration: to.Ptr("{\"work_mem\": 4096, \"max_wal_size\": 12288, \"min_wal_size\": 80, \"seq_page_cost\": 1, \"bgwriter_delay\": 20, \"random_page_cost\": 2, \"max_parallel_workers\": 8, \"bgwriter_lru_maxpages\": 100, \"effective_io_concurrency\": 1, \"max_parallel_workers_per_gather\": 2}"),
		// 			AverageQueryRuntimeMs: to.Ptr("0.38"),
		// 			IterationID: to.Ptr("1"),
		// 			IterationStartTime: to.Ptr("02/26/2025 22:57:30"),
		// 			SessionID: to.Ptr("7fb4b3c1-8d77-4fae-82d8-fcd26c892d75"),
		// 			TransactionsPerSecond: to.Ptr("2.86"),
		// 		},
		// 		{
		// 			AppliedConfiguration: to.Ptr("{\"work_mem\": 53566, \"max_wal_size\": 24576, \"min_wal_size\": 6144, \"seq_page_cost\": 1, \"bgwriter_delay\": 200, \"random_page_cost\": 3.5, \"max_parallel_workers\": 2, \"bgwriter_lru_maxpages\": 100, \"effective_io_concurrency\": 300, \"max_parallel_workers_per_gather\": 2}"),
		// 			AverageQueryRuntimeMs: to.Ptr("0.40"),
		// 			IterationID: to.Ptr("2"),
		// 			IterationStartTime: to.Ptr("02/26/2025 23:02:32"),
		// 			SessionID: to.Ptr("7fb4b3c1-8d77-4fae-82d8-fcd26c892d75"),
		// 			TransactionsPerSecond: to.Ptr("2.81"),
		// 		},
		// 		{
		// 			AppliedConfiguration: to.Ptr("{\"work_mem\": 9681, \"max_wal_size\": 4096, \"min_wal_size\": 1024, \"seq_page_cost\": 1, \"bgwriter_delay\": 150, \"random_page_cost\": 1.5, \"max_parallel_workers\": 6, \"bgwriter_lru_maxpages\": 300, \"effective_io_concurrency\": 400, \"max_parallel_workers_per_gather\": 2}"),
		// 			AverageQueryRuntimeMs: to.Ptr("0.42"),
		// 			IterationID: to.Ptr("3"),
		// 			IterationStartTime: to.Ptr("02/26/2025 23:07:33"),
		// 			SessionID: to.Ptr("7fb4b3c1-8d77-4fae-82d8-fcd26c892d75"),
		// 			TransactionsPerSecond: to.Ptr("2.67"),
		// 	}},
		// }
	}
}

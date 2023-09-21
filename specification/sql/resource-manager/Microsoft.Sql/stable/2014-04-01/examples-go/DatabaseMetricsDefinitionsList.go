package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DatabaseMetricsDefinitionsList.json
func ExampleDatabasesClient_NewListMetricDefinitionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListMetricDefinitionsPager("sqlcrudtest-6730", "sqlcrudtest-9007", "3481", nil)
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
		// page.MetricDefinitionListResult = armsql.MetricDefinitionListResult{
		// 	Value: []*armsql.MetricDefinition{
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("CPU percentage"),
		// 				Value: to.Ptr("cpu_percent"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT15S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT60S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeAverage),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypePercent),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Data IO percentage"),
		// 				Value: to.Ptr("physical_data_read_percent"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT15S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT60S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeAverage),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypePercent),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Log IO percentage"),
		// 				Value: to.Ptr("log_write_percent"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT15S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT60S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeAverage),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypePercent),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("DTU percentage"),
		// 				Value: to.Ptr("dtu_consumption_percent"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT15S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT60S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeAverage),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypePercent),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Total database size"),
		// 				Value: to.Ptr("storage"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeMaximum),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypeBytes),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("In-Memory OLTP storage percent"),
		// 				Value: to.Ptr("xtp_storage_percent"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT15S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("PT1H"),
		// 					TimeGrain: to.Ptr("PT60S"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeAverage),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypePercent),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Workers percentage"),
		// 				Value: to.Ptr("workers_percent"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeAverage),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypePercent),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Sessions percentage"),
		// 				Value: to.Ptr("sessions_percent"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeAverage),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypePercent),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("DTU limit"),
		// 				Value: to.Ptr("dtu_limit"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeAverage),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypeCount),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("DTU used"),
		// 				Value: to.Ptr("dtu_used"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeAverage),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypeCount),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Successful Connections"),
		// 				Value: to.Ptr("connection_successful"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeTotal),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypeCount),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Failed Connections"),
		// 				Value: to.Ptr("connection_failed"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeTotal),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypeCount),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Blocked by Firewall"),
		// 				Value: to.Ptr("blocked_by_firewall"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeTotal),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypeCount),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Deadlocks"),
		// 				Value: to.Ptr("deadlock"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeTotal),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypeCount),
		// 		},
		// 		{
		// 			Name: &armsql.MetricName{
		// 				LocalizedValue: to.Ptr("Database size percentage"),
		// 				Value: to.Ptr("storage_percent"),
		// 			},
		// 			MetricAvailabilities: []*armsql.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armsql.PrimaryAggregationTypeMaximum),
		// 			ResourceURI: to.Ptr("/subscriptions/b6a6e0c5-e79c-4c6d-a878-72eafbca4cf2/resourceGroups/QA/providers/Microsoft.Sql/servers/nafantest/databases/db1"),
		// 			Unit: to.Ptr(armsql.UnitDefinitionTypePercent),
		// 	}},
		// }
	}
}

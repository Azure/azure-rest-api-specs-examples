package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79d9ef3e065f2dcb6bd1db51e29c62a99dff5cb/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armbatch.OperationListResult{
		// 	Value: []*armbatch.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the available logs for the Batch service"),
		// 				Operation: to.Ptr("Read Batch service log definitions"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Account Log Definitions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 			Properties: map[string]any{
		// 				"serviceSpecification":map[string]any{
		// 					"logSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "ServiceLog",
		// 							"blobDuration": "PT1H",
		// 							"displayName": "Service Logs",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the available metrics for the Batch service"),
		// 				Operation: to.Ptr("Read Batch service metric definitions"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Account Metric Definitions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 			Properties: map[string]any{
		// 				"serviceSpecification":map[string]any{
		// 					"metricSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "CoreCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Total number of dedicated cores in the batch account",
		// 							"displayName": "Dedicated Core Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TotalNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Total number of dedicated nodes in the batch account",
		// 							"displayName": "Dedicated Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "LowPriorityCoreCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Total number of low-priority cores in the batch account",
		// 							"displayName": "LowPriority Core Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TotalLowPriorityNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Total number of low-priority nodes in the batch account",
		// 							"displayName": "Low-Priority Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "CreatingNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of nodes being created",
		// 							"displayName": "Creating Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "StartingNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of nodes starting",
		// 							"displayName": "Starting Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "WaitingForStartTaskNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of nodes waiting for the Start Task to complete",
		// 							"displayName": "Waiting For Start Task Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "StartTaskFailedNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of nodes where the Start Task has failed",
		// 							"displayName": "Start Task Failed Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "IdleNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of idle nodes",
		// 							"displayName": "Idle Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "OfflineNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of offline nodes",
		// 							"displayName": "Offline Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "RebootingNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of rebooting nodes",
		// 							"displayName": "Rebooting Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "ReimagingNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of reimaging nodes",
		// 							"displayName": "Reimaging Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "RunningNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of running nodes",
		// 							"displayName": "Running Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "LeavingPoolNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of nodes leaving the Pool",
		// 							"displayName": "Leaving Pool Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "UnusableNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of unusable nodes",
		// 							"displayName": "Unusable Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "PreemptedNodeCount",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"displayDescription": "Number of preempted nodes",
		// 							"displayName": "Preempted Node Count",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"lockAggregationType": "Total",
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TaskStartEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "poolId",
		// 									"displayName": "Pool ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of tasks that have started",
		// 							"displayName": "Task Start Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TaskCompleteEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "poolId",
		// 									"displayName": "Pool ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of tasks that have completed",
		// 							"displayName": "Task Complete Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TaskFailEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "poolId",
		// 									"displayName": "Pool ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of tasks that have completed in a failed state",
		// 							"displayName": "Task Fail Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "PoolCreateEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "poolId",
		// 									"displayName": "Pool ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of pools that have been created",
		// 							"displayName": "Pool Create Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "PoolResizeStartEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "poolId",
		// 									"displayName": "Pool ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of pool resizes that have started",
		// 							"displayName": "Pool Resize Start Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "PoolResizeCompleteEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "poolId",
		// 									"displayName": "Pool ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of pool resizes that have completed",
		// 							"displayName": "Pool Resize Complete Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "PoolDeleteStartEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "poolId",
		// 									"displayName": "Pool ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of pool deletes that have started",
		// 							"displayName": "Pool Delete Start Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "PoolDeleteCompleteEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "poolId",
		// 									"displayName": "Pool ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of pool deletes that have completed",
		// 							"displayName": "Pool Delete Complete Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "JobDeleteCompleteEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of jobs that have been successfully deleted.",
		// 							"displayName": "Job Delete Complete Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "JobDeleteStartEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of jobs that have been requested to be deleted.",
		// 							"displayName": "Job Delete Start Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "JobDisableCompleteEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of jobs that have been successfully disabled.",
		// 							"displayName": "Job Disable Complete Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "JobDisableStartEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of jobs that have been requested to be disabled.",
		// 							"displayName": "Job Disable Start Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "JobStartEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of jobs that have been successfully started.",
		// 							"displayName": "Job Start Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "JobTerminateCompleteEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of jobs that have been successfully terminated.",
		// 							"displayName": "Job Terminate Complete Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "JobTerminateStartEvent",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 								map[string]any{
		// 									"name": "jobId",
		// 									"displayName": "Job ID",
		// 									"toBeExportedForShoebox": true,
		// 								},
		// 							},
		// 							"displayDescription": "Total number of jobs that have been requested to be terminated.",
		// 							"displayName": "Job Terminate Start Events",
		// 							"enableRegionalMdmAccount": false,
		// 							"fillGapWithZero": false,
		// 							"supportsInstanceLevelAggregation": false,
		// 							"unit": "Count",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/jobs/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists jobs on a Batch account or gets the properties of a job"),
		// 				Operation: to.Ptr("List or Get Jobs"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/jobs/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Creates a new job on a Batch account or updates an existing job"),
		// 				Operation: to.Ptr("Create or Update Job"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/jobs/delete"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Deletes a job from a Batch account"),
		// 				Operation: to.Ptr("Delete Job"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/jobSchedules/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists job schedules on a Batch account or gets the properties of a job schedule"),
		// 				Operation: to.Ptr("List or Get Job Schedules"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Job Schedules"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/jobSchedules/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Creates a new job schedule on a Batch account or updates an existing job schedule"),
		// 				Operation: to.Ptr("Create or Update Job Schedule"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Job Schedules"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/jobSchedules/delete"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Deletes a job schedule from a Batch account"),
		// 				Operation: to.Ptr("Delete Job Schedule"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Job Schedules"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/operations/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists operations available on Microsoft.Batch resource provider"),
		// 				Operation: to.Ptr("List Available Batch Operations"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Available Batch Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists Batch accounts or gets the properties of a Batch account"),
		// 				Operation: to.Ptr("List or Get Batch Accounts"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Creates a new Batch account or updates an existing Batch account"),
		// 				Operation: to.Ptr("Create or Update Batch Account"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/delete"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Deletes a Batch account"),
		// 				Operation: to.Ptr("Delete Batch Account"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/listkeys/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists access keys for a Batch account"),
		// 				Operation: to.Ptr("List Batch Account Keys"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/regeneratekeys/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Regenerates access keys for a Batch account"),
		// 				Operation: to.Ptr("Regenerate Batch Account Keys"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/locations/quotas/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets Batch quotas of the specified subscription at the specified Azure region"),
		// 				Operation: to.Ptr("Get Batch Quotas"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Quotas"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/locations/checkNameAvailability/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Checks that the account name is valid and not in use."),
		// 				Operation: to.Ptr("Check Name Availability"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Name Availability"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/operationResults/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the results of a long running Batch account operation"),
		// 				Operation: to.Ptr("Get Batch account operation results"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/locations/accountOperationResults/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the results of a long running Batch account operation"),
		// 				Operation: to.Ptr("Get Batch account operation results"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/register/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the Batch Resource Provider and enables the creation of Batch accounts"),
		// 				Operation: to.Ptr("Register the Batch Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/unregister/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Unregisters the subscription for the Batch Resource Provider preventing the creation of Batch accounts"),
		// 				Operation: to.Ptr("Unregister the Batch Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/syncAutoStorageKeys/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Synchronizes access keys for the auto storage account configured for a Batch account"),
		// 				Operation: to.Ptr("Synchronize Auto Storage Account Keys"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Accounts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/applications/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists applications or gets the properties of an application"),
		// 				Operation: to.Ptr("List or Get Applications"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Applications"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/applications/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Creates a new application or updates an existing application"),
		// 				Operation: to.Ptr("Create or Update Application"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Applications"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/applications/delete"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Deletes an application"),
		// 				Operation: to.Ptr("Delete Application"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Applications"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/applications/versions/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the properties of an application package"),
		// 				Operation: to.Ptr("Get Application Package"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Application Packages"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/applications/versions/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Creates a new application package or updates an existing application package"),
		// 				Operation: to.Ptr("Create or Update Application Package"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Application Packages"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/applications/versions/delete"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Deletes an application package"),
		// 				Operation: to.Ptr("Delete Application Package"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Application Packages"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/applications/versions/activate/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Activates an application package"),
		// 				Operation: to.Ptr("Activate Application Package"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Application Packages"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/certificates/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists certificates on a Batch account or gets the properties of a certificate"),
		// 				Operation: to.Ptr("List or Get Certificates"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Certificates"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/certificates/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Creates a new certificate on a Batch account or updates an existing certificate"),
		// 				Operation: to.Ptr("Create or Update Certificate"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Certificates"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/certificates/delete"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Deletes a certificate from a Batch account"),
		// 				Operation: to.Ptr("Delete Certificate"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Certificates"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/certificates/cancelDelete/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Cancels the failed deletion of a certificate on a Batch account"),
		// 				Operation: to.Ptr("Cancel Delete Certificate"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Certificates"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/certificateOperationResults/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the results of a long running certificate operation on a Batch account"),
		// 				Operation: to.Ptr("Get Certificate Operation Results"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Certificates"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/pools/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists pools on a Batch account or gets the properties of a pool"),
		// 				Operation: to.Ptr("List or Get Pools"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Pools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/pools/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Creates a new pool on a Batch account or updates an existing pool"),
		// 				Operation: to.Ptr("Create or Update Pool"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Pools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/pools/delete"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Deletes a pool from a Batch account"),
		// 				Operation: to.Ptr("Delete Pool"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Pools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/pools/stopResize/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Stops an ongoing resize operation on a Batch account pool"),
		// 				Operation: to.Ptr("Stop Pool Resize"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Pools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/pools/disableAutoscale/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Disables automatic scaling for a Batch account pool"),
		// 				Operation: to.Ptr("Disable Pool AutoScale"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Pools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/poolOperationResults/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the results of a long running pool operation on a Batch account"),
		// 				Operation: to.Ptr("Get Pool Operation Results"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Pools"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/locations/virtualMachineSkus/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists available Batch supported Virtual Machine VM sizes at the given location"),
		// 				Operation: to.Ptr("List Supported Batch Virtual Machine VM"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Supported Skus"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/locations/cloudServiceSkus/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists available Batch supported Cloud Service VM sizes at the given location"),
		// 				Operation: to.Ptr("List Supported Batch Cloud Service VM"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Batch Supported Skus"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/privateLinkResources/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the properties of a Private link resource or Lists Private link resources on a Batch account"),
		// 				Operation: to.Ptr("Get or List Private link resources"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("PrivateLinkResources"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnections/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Update an existing Private endpoint connection on a Batch account"),
		// 				Operation: to.Ptr("Update Private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("PrivateEndpointConnections"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnections/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets Private endpoint connection or Lists Private endpoint connections on a Batch account"),
		// 				Operation: to.Ptr("Get or List Private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("PrivateEndpointConnections"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnectionResults/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the results of a long running Batch account private endpoint connection operation"),
		// 				Operation: to.Ptr("Get Batch account private endpoint connection operation results"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("PrivateEndpointConnections"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Validates a Private endpoint connection proxy on a Batch account"),
		// 				Operation: to.Ptr("Validates a Private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("PrivateEndpointConnectionProxies"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnectionProxies/write"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Create a new Private endpoint connection proxy on a Batch account"),
		// 				Operation: to.Ptr("Create or Update Private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("PrivateEndpointConnectionProxies"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnectionProxies/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets Private endpoint connection proxy on a Batch account"),
		// 				Operation: to.Ptr("Get Private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("PrivateEndpointConnectionProxies"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnectionProxies/delete"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Delete a Private endpoint connection proxy on a Batch account"),
		// 				Operation: to.Ptr("Delete Private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("PrivateEndpointConnectionProxies"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/privateEndpointConnectionProxyResults/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Gets the results of a long running Batch account private endpoint connection proxy operation"),
		// 				Operation: to.Ptr("Get Batch account private endpoint connection proxy operation results"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("PrivateEndpointConnectionProxies"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Batch/batchAccounts/outboundNetworkDependenciesEndpoints/read"),
		// 			Display: &armbatch.OperationDisplay{
		// 				Description: to.Ptr("Lists the outbound network dependency endpoints for a Batch account"),
		// 				Operation: to.Ptr("List Outbound Network Dependency Endpoints"),
		// 				Provider: to.Ptr("Microsoft Batch"),
		// 				Resource: to.Ptr("Outbound Network Dependencies Endpoints"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}

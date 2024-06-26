package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armlogic.OperationListResult{
		// 	Value: []*armlogic.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/operations/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the operation."),
		// 				Operation: to.Ptr("Get Operation"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Operation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/register/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Registers the Microsoft.Logic resource provider for a given subscription."),
		// 				Operation: to.Ptr("Register Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/locations/workflows/validate/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Validates the workflow."),
		// 				Operation: to.Ptr("Validate Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow."),
		// 				Operation: to.Ptr("Get Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the workflow."),
		// 				Operation: to.Ptr("Set Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the workflow."),
		// 				Operation: to.Ptr("Delete Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/run/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Starts a run of the workflow."),
		// 				Operation: to.Ptr("Run Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/disable/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Disables the workflow."),
		// 				Operation: to.Ptr("Disable Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/enable/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Enables the workflow."),
		// 				Operation: to.Ptr("Enable Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/suspend/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Suspends the workflow."),
		// 				Operation: to.Ptr("Suspend Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/validate/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Validates the workflow."),
		// 				Operation: to.Ptr("Validate Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/move/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Moves Workflow from its existing subscription id, resource group, and/or name to a different subscription id, resource group, and/or name."),
		// 				Operation: to.Ptr("Move Workflow"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/listSwagger/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the workflow swagger definitions."),
		// 				Operation: to.Ptr("Get workflow swagger"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/versions/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow version."),
		// 				Operation: to.Ptr("Get Workflow Version"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Version"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/versions/triggers/listCallbackUrl/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the callback URL for trigger."),
		// 				Operation: to.Ptr("List Trigger Callback URL"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Trigger"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/accessKeys/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the access key."),
		// 				Operation: to.Ptr("Get Access Key"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Access Key"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/accessKeys/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the access key."),
		// 				Operation: to.Ptr("Set Access Key"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Access Key"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/accessKeys/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the access key."),
		// 				Operation: to.Ptr("Delete Access Key"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Access Key"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/accessKeys/list/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Lists the access key secrets."),
		// 				Operation: to.Ptr("List Access Key"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Access Key"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/accessKeys/regenerate/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Regenerates the access key secrets."),
		// 				Operation: to.Ptr("Regenerate Access Key"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Access Key"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/regenerateAccessKey/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Regenerates the access key secrets."),
		// 				Operation: to.Ptr("Regenerate Access Key"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Access Key"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/listCallbackUrl/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the callback URL for workflow."),
		// 				Operation: to.Ptr("List workflow callback URL"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/triggers/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the trigger."),
		// 				Operation: to.Ptr("Get Trigger"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Trigger"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/triggers/run/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Executes the trigger."),
		// 				Operation: to.Ptr("Trigger Run"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Trigger"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/triggers/reset/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Resets the trigger."),
		// 				Operation: to.Ptr("Trigger Reset"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Trigger"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/triggers/setState/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Sets the trigger state."),
		// 				Operation: to.Ptr("Set Trigger State"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Trigger"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/triggers/histories/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the trigger histories."),
		// 				Operation: to.Ptr("Get Trigger Histories"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Trigger Histories"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/triggers/histories/resubmit/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Resubmits the workflow trigger."),
		// 				Operation: to.Ptr("Resubmit trigger"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Trigger Histories"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/triggers/listCallbackUrl/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the callback URL for trigger."),
		// 				Operation: to.Ptr("List Trigger Callback URL"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Trigger"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow run."),
		// 				Operation: to.Ptr("Get Workflow Run"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Run"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/cancel/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Cancels the run of a workflow."),
		// 				Operation: to.Ptr("Cancel Workflow Run"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Run"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/operations/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow run operation status."),
		// 				Operation: to.Ptr("Get Workflow Run Operation Status"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Run Operation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/actions/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow run action."),
		// 				Operation: to.Ptr("Get Workflow Run Action"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Run Action"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/actions/repetitions/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow run action repetition."),
		// 				Operation: to.Ptr("Get Workflow Run Action Repetition"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Run Action Repetition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/actions/scoperepetitions/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow run action scope repetition."),
		// 				Operation: to.Ptr("Get Workflow Run Action Scope Repetition"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Run Action Scope Repetition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/actions/requestHistories/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow run action request history."),
		// 				Operation: to.Ptr("Gets the workflow run action request history"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow run action request history"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/actions/repetitions/requestHistories/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow run repetition action request history."),
		// 				Operation: to.Ptr("Gets the workflow run repetition action request history"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow run repetition action request history"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow diagnostic settings."),
		// 				Operation: to.Ptr("Get Workflow Diagnostic Setting"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Diagnostic Setting"),
		// 			},
		// 			Origin: to.Ptr("System"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the workflow diagnostic setting."),
		// 				Operation: to.Ptr("Set Workflow Diagnostic Setting"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Diagnostic Setting"),
		// 			},
		// 			Origin: to.Ptr("System"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow metric definitions."),
		// 				Operation: to.Ptr("Get Workflow Metric Definition"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Metric Definition"),
		// 			},
		// 			Origin: to.Ptr("System"),
		// 			Properties: map[string]any{
		// 				"serviceSpecification":map[string]any{
		// 					"metricSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "RunsStarted",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow runs started.",
		// 							"displayName": "Runs Started",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "RunsCompleted",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow runs completed.",
		// 							"displayName": "Runs Completed",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "RunsSucceeded",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow runs succeeded.",
		// 							"displayName": "Runs Succeeded",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "RunsFailed",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow runs failed.",
		// 							"displayName": "Runs Failed",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "RunsCancelled",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow runs cancelled.",
		// 							"displayName": "Runs Cancelled",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "RunLatency",
		// 							"aggregationType": "Average",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Latency of completed workflow runs.",
		// 							"displayName": "Run Latency",
		// 							"fillGapWithZero": false,
		// 							"unit": "Seconds",
		// 						},
		// 						map[string]any{
		// 							"name": "RunSuccessLatency",
		// 							"aggregationType": "Average",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Latency of succeeded workflow runs.",
		// 							"displayName": "Run Success Latency",
		// 							"fillGapWithZero": false,
		// 							"unit": "Seconds",
		// 						},
		// 						map[string]any{
		// 							"name": "RunThrottledEvents",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow action or trigger throttled events.",
		// 							"displayName": "Run Throttled Events",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "RunFailurePercentage",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Percentage of workflow runs failed.",
		// 							"displayName": "Run Failure Percentage",
		// 							"fillGapWithZero": true,
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "ActionsStarted",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow actions started.",
		// 							"displayName": "Actions Started ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "ActionsCompleted",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow actions completed.",
		// 							"displayName": "Actions Completed ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "ActionsSucceeded",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow actions succeeded.",
		// 							"displayName": "Actions Succeeded ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "ActionsFailed",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow actions failed.",
		// 							"displayName": "Actions Failed ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "ActionsSkipped",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow actions skipped.",
		// 							"displayName": "Actions Skipped ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "ActionLatency",
		// 							"aggregationType": "Average",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Latency of completed workflow actions.",
		// 							"displayName": "Action Latency ",
		// 							"fillGapWithZero": false,
		// 							"unit": "Seconds",
		// 						},
		// 						map[string]any{
		// 							"name": "ActionSuccessLatency",
		// 							"aggregationType": "Average",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Latency of succeeded workflow actions.",
		// 							"displayName": "Action Success Latency ",
		// 							"fillGapWithZero": false,
		// 							"unit": "Seconds",
		// 						},
		// 						map[string]any{
		// 							"name": "ActionThrottledEvents",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow action throttled events..",
		// 							"displayName": "Action Throttled Events",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggersStarted",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow triggers started.",
		// 							"displayName": "Triggers Started ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggersCompleted",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow triggers completed.",
		// 							"displayName": "Triggers Completed ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggersSucceeded",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow triggers succeeded.",
		// 							"displayName": "Triggers Succeeded ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggersFailed",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow triggers failed.",
		// 							"displayName": "Triggers Failed ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggersSkipped",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow triggers skipped.",
		// 							"displayName": "Triggers Skipped",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggersFired",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow triggers fired.",
		// 							"displayName": "Triggers Fired ",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggerLatency",
		// 							"aggregationType": "Average",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Latency of completed workflow triggers.",
		// 							"displayName": "Trigger Latency ",
		// 							"fillGapWithZero": false,
		// 							"unit": "Seconds",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggerFireLatency",
		// 							"aggregationType": "Average",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Latency of fired workflow triggers.",
		// 							"displayName": "Trigger Fire Latency ",
		// 							"fillGapWithZero": false,
		// 							"unit": "Seconds",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggerSuccessLatency",
		// 							"aggregationType": "Average",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Latency of succeeded workflow triggers.",
		// 							"displayName": "Trigger Success Latency ",
		// 							"fillGapWithZero": false,
		// 							"unit": "Seconds",
		// 						},
		// 						map[string]any{
		// 							"name": "TriggerThrottledEvents",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow trigger throttled events.",
		// 							"displayName": "Trigger Throttled Events",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "BillableActionExecutions",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow action executions getting billed.",
		// 							"displayName": "Billable Action Executions",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "BillableTriggerExecutions",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow trigger executions getting billed.",
		// 							"displayName": "Billable Trigger Executions",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TotalBillableExecutions",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of workflow executions getting billed.",
		// 							"displayName": "Total Billable Executions",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "BillableNativeActionExecutions",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of native workflow action executions getting billed.",
		// 							"displayName": "Billable Native Action Executions",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "BillableNativeTriggerExecutions",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of native workflow trigger executions getting billed.",
		// 							"displayName": "Billable Native Trigger Executions",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TotalBillableNativeExecutions",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of native workflow executions getting billed.",
		// 							"displayName": "Total Native Billable Executions",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "BillableStandardActionExecutions",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of standard workflow action executions getting billed.",
		// 							"displayName": "Billable Standard Action Executions",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "BillableStandardTriggerExecutions",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of standard workflow trigger executions getting billed.",
		// 							"displayName": "Billable Standard Trigger Executions",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "TotalBillableStandardExecutions",
		// 							"aggregationType": "Total",
		// 							"availabilities":[]any{
		// 								map[string]any{
		// 									"blobDuration": "PT1H",
		// 									"timeGrain": "PT1M",
		// 								},
		// 							},
		// 							"dimensions":[]any{
		// 							},
		// 							"displayDescription": "Number of standard workflow executions getting billed.",
		// 							"displayName": "Total Standard Billable Executions",
		// 							"fillGapWithZero": true,
		// 							"unit": "Count",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the workflow log definitions."),
		// 				Operation: to.Ptr("Get Workflow Log Definition"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Log Definition"),
		// 			},
		// 			Origin: to.Ptr("System"),
		// 			Properties: map[string]any{
		// 				"serviceSpecification":map[string]any{
		// 					"logSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "WorkflowRuntime",
		// 							"description": "Diagnostic events related to workflow runtime executions.",
		// 							"blobDuration": "PT1H",
		// 							"displayName": "Workflow runtime diagnostic events",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/actions/listExpressionTraces/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the workflow run action expression traces."),
		// 				Operation: to.Ptr("List Workflow Run Action Expression Traces"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Run Action"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/workflows/runs/actions/repetitions/listExpressionTraces/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the workflow run action repetition expression traces."),
		// 				Operation: to.Ptr("List Workflow Run Action Repetition Expression Traces"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Workflow Run Action Repetition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the Integration Account log definitions."),
		// 				Operation: to.Ptr("Get Integration Account Log Definition"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Log Definition"),
		// 			},
		// 			Origin: to.Ptr("System"),
		// 			Properties: map[string]any{
		// 				"serviceSpecification":map[string]any{
		// 					"logSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "IntegrationAccountTrackingEvents",
		// 							"description": "Track events related to Integration Account.",
		// 							"blobDuration": "PT1H",
		// 							"displayName": "Integration Account track events",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the integration account."),
		// 				Operation: to.Ptr("Get Integration Account"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the integration account."),
		// 				Operation: to.Ptr("Set Integration Account"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the integration account."),
		// 				Operation: to.Ptr("Delete Integration Account"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/regenerateAccessKey/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Regenerates the access key secrets."),
		// 				Operation: to.Ptr("Regenerate Access Key"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/listCallbackUrl/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the callback URL for integration account."),
		// 				Operation: to.Ptr("List Integration Account Callback URL"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/listKeyVaultKeys/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the keys in the key vault."),
		// 				Operation: to.Ptr("List Key Vault Keys"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/logTrackingEvents/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Logs the tracking events in the integration account."),
		// 				Operation: to.Ptr("Log Integration Account Tracking Events"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/join/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Joins the Integration Account."),
		// 				Operation: to.Ptr("Join Integration Account"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/partners/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the parter in integration account."),
		// 				Operation: to.Ptr("Get Integration Account Partner"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Partner"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/partners/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the partner in integration account."),
		// 				Operation: to.Ptr("Set Integration Account Partner"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Partner"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/partners/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the partner in integration account."),
		// 				Operation: to.Ptr("Delete Integration Account Partner"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Partner"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/partners/listContentCallbackUrl/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the callback URL for partner content in integration account."),
		// 				Operation: to.Ptr("List Integration Account Partner Content Callback URL"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Partner"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/agreements/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the agreement in integration account."),
		// 				Operation: to.Ptr("Get Integration Account Agreement"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Agreement"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/agreements/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the agreement in integration account."),
		// 				Operation: to.Ptr("Set Integration Account Agreement"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Agreement"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/agreements/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the agreement in integration account."),
		// 				Operation: to.Ptr("Delete Integration Account Agreement"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Agreement"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/agreements/listContentCallbackUrl/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the callback URL for agreement content in integration account."),
		// 				Operation: to.Ptr("List Integration Account Agreement Content Callback URL"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Agreement"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/certificates/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the certificate in integration account."),
		// 				Operation: to.Ptr("Get Integration Account Certificate"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Certificate"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/certificates/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the certificate in integration account."),
		// 				Operation: to.Ptr("Set Integration Account Certificate"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Certificate"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/certificates/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the certificate in integration account."),
		// 				Operation: to.Ptr("Delete Integration Account Certificate"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Certificate"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/schemas/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the schema in integration account."),
		// 				Operation: to.Ptr("Get Integration Account Schema"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Schema"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/schemas/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the schema in integration account."),
		// 				Operation: to.Ptr("Set Integration Account Schema"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Schema"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/schemas/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the schema in integration account."),
		// 				Operation: to.Ptr("Delete Integration Account Schema"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Schema"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/schemas/listContentCallbackUrl/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the callback URL for schema content in integration account."),
		// 				Operation: to.Ptr("List Integration Account Schema Content Callback URL"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Schema"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/maps/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the map in integration account."),
		// 				Operation: to.Ptr("Get Integration Account Map"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Map"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/maps/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the map in integration account."),
		// 				Operation: to.Ptr("Set Integration Account Map"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Map"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/maps/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the map in integration account."),
		// 				Operation: to.Ptr("Delete Integration Account Map"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Map"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/maps/listContentCallbackUrl/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the callback URL for map content in integration account."),
		// 				Operation: to.Ptr("List Integration Account Map Content Callback URL"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Map"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/assemblies/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the assembly in integration account."),
		// 				Operation: to.Ptr("Get Integration Account Assembly"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Assembly"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/assemblies/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the assembly in integration account."),
		// 				Operation: to.Ptr("Set Integration Account Assembly"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Assembly"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/assemblies/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the assembly in integration account."),
		// 				Operation: to.Ptr("Delete Integration Account Assembly"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Assembly"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/assemblies/listContentCallbackUrl/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Gets the callback URL for assembly content in integration account."),
		// 				Operation: to.Ptr("List Integration Account Assembly Content Callback URL"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Assembly"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/batchConfigurations/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the batch configuration in integration account."),
		// 				Operation: to.Ptr("Get Integration Account Batch Configuration"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Batch Configuration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/batchConfigurations/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the batch configuration in integration account."),
		// 				Operation: to.Ptr("Set Integration Account Batch Configuration"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Batch Configuration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/batchConfigurations/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the batch configuration in integration account."),
		// 				Operation: to.Ptr("Delete Integration Account Batch Configuration"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Batch Configuration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/sessions/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the batch configuration in integration account."),
		// 				Operation: to.Ptr("Get Integration Account Session"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Session"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/sessions/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the session in integration account."),
		// 				Operation: to.Ptr("Set Integration Account Session"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Session"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationAccounts/sessions/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the session in integration account."),
		// 				Operation: to.Ptr("Delete Integration Account Session"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Account Session"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/read"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Reads the integration service environment."),
		// 				Operation: to.Ptr("Get Integration Service Environment"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Service Environment"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/write"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the integration service environment."),
		// 				Operation: to.Ptr("Set Integration Service Environment"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Service Environment"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/delete"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Deletes the integration service environment."),
		// 				Operation: to.Ptr("Delete Integration Service Environment"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Service Environment"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/join/action"),
		// 			Display: &armlogic.OperationDisplay{
		// 				Description: to.Ptr("Joins the Integration Service Environment."),
		// 				Operation: to.Ptr("Join Integration Service Environment"),
		// 				Provider: to.Ptr("Microsoft Logic"),
		// 				Resource: to.Ptr("Integration Service Environment"),
		// 			},
		// 	}},
		// }
	}
}

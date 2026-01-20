package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule"
)

// Generated from example definition: 2025-04-15-preview/ScheduledActions_VirtualMachinesExecuteCreate_MinimumSet_Gen.json
func ExampleScheduledActionsClient_VirtualMachinesExecuteCreate_scheduledActionsVirtualMachinesExecuteCreateMinimumSetGenGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("0505D8E4-D41A-48FB-9CA5-4AF8D93BE75F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScheduledActionsClient().VirtualMachinesExecuteCreate(ctx, "useast", armcomputeschedule.ExecuteCreateContent{
		ResourceConfigParameters: &armcomputeschedule.ResourceProvisionPayload{
			BaseProfile: map[string]any{
				"hardwareProfile": map[string]any{
					"name": "F1",
				},
				"provisioningState": 0,
				"storageProfile": map[string]any{
					"osDisk": map[string]any{
						"osType": 0,
					},
				},
				"vmExtensions": []any{
					map[string]any{
						"autoUpgradeMinorVersion": true,
						"protectedSettings":       "SomeDecryptedSecretValue",
						"provisioningState":       0,
						"enableAutomaticUpgrade":  true,
						"publisher":               "Microsoft.Azure.Monitor",
						"type":                    "AzureMonitorLinuxAgent",
						"typeHandlerVersion":      "1.0",
					},
					map[string]any{
						"name": "myExtensionName",
					},
				},
				"resourcegroupName": "RG5ABF491C-3164-42A6-8CB5-BF3CB53B018B",
				"computeApiVersion": "2024-07-01",
			},
			ResourceOverrides: []map[string]any{
				{
					"name":     "myFleet_523",
					"location": "LocalDev",
					"properties": map[string]any{
						"hardwareProfile": map[string]any{
							"vmSize": "Standard_F1s",
						},
						"provisioningState": 0,
						"osProfile": map[string]any{
							"computerName":  "myFleet000000",
							"adminUsername": "adminUser",
							"windowsConfiguration": map[string]any{
								"additionalUnattendContent": []any{
									map[string]any{
										"passName": "someValue",
										"content":  "",
									},
									map[string]any{
										"passName": "someOtherValue",
										"content":  "SomeDecryptedSecretValue",
									},
								},
							},
							"adminPassword": "SomeDecryptedSecretValue",
						},
						"priority": 0,
					},
					"zones": []any{
						"1",
					},
				},
				{
					"name":     "myFleet_524",
					"location": "LocalDev",
					"properties": map[string]any{
						"hardwareProfile": map[string]any{
							"vmSize": "Standard_G1s",
						},
						"provisioningState": 0,
						"osProfile": map[string]any{
							"computerName":  "myFleet000000",
							"adminUsername": "adminUser",
							"windowsConfiguration": map[string]any{
								"additionalUnattendContent": []any{
									map[string]any{
										"passName": "someValue",
										"content":  "",
									},
									map[string]any{
										"passName": "someOtherValue",
										"content":  "SomeDecryptedSecretValue",
									},
								},
							},
							"adminPassword": "SomeDecryptedSecretValue",
						},
						"priority": 0,
					},
					"zones": []any{
						"2",
					},
				},
			},
			ResourceCount: to.Ptr[int32](2),
		},
		ExecutionParameters: &armcomputeschedule.ExecutionParameters{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputeschedule.ScheduledActionsClientVirtualMachinesExecuteCreateResponse{
	// 	CreateResourceOperationResponse: &armcomputeschedule.CreateResourceOperationResponse{
	// 		Type: to.Ptr("virtualmachines"),
	// 		Location: to.Ptr("eastus"),
	// 		Results: []*armcomputeschedule.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/RG5ABF491C-3164-42A6-8CB5-BF3CB53B018B/providers/Microsoft.Compute/virtualMachines/TL13"),
	// 				Operation: &armcomputeschedule.ResourceOperationDetails{
	// 					OperationID: to.Ptr("5cc987d0-4a3a-4b01-b31a-c99219ece5e2"),
	// 					ResourceID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/RG5ABF491C-3164-42A6-8CB5-BF3CB53B018B/providers/Microsoft.Compute/virtualMachines/TL13"),
	// 					OpType: to.Ptr(armcomputeschedule.ResourceOperationType("Create")),
	// 					SubscriptionID: to.Ptr("1d04e8f1-ee04-4056-b0b2-718f5bb45b04"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-23T22:51:50.1355051+00:00"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armcomputeschedule.OperationStatePendingScheduling),
	// 					TimeZone: to.Ptr("UTC"),
	// 					RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](3),
	// 						RetryWindowInMinutes: to.Ptr[int32](10),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/RG5ABF491C-3164-42A6-8CB5-BF3CB53B018B/providers/Microsoft.Compute/virtualMachines/TL14"),
	// 				Operation: &armcomputeschedule.ResourceOperationDetails{
	// 					OperationID: to.Ptr("5cc987d0-4a3a-4b01-b31a-c99219ece5e2"),
	// 					ResourceID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/RG5ABF491C-3164-42A6-8CB5-BF3CB53B018B/providers/Microsoft.Compute/virtualMachines/TL14"),
	// 					OpType: to.Ptr(armcomputeschedule.ResourceOperationType("Create")),
	// 					SubscriptionID: to.Ptr("1d04e8f1-ee04-4056-b0b2-718f5bb45b04"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-23T22:51:50.1355051+00:00"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputeschedule.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armcomputeschedule.OperationStatePendingScheduling),
	// 					TimeZone: to.Ptr("UTC"),
	// 					RetryPolicy: &armcomputeschedule.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](3),
	// 						RetryWindowInMinutes: to.Ptr[int32](10),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Description: to.Ptr("Create Resource response"),
	// 	},
	// }
}

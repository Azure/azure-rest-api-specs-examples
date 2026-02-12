package armcomputebulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computebulkactions/armcomputebulkactions"
)

// Generated from example definition: 2026-02-01-preview/BulkActions_VirtualMachinesExecuteCreate_MaximumSet_Gen.json
func ExampleBulkActionsClient_VirtualMachinesExecuteCreate_bulkActionsVirtualMachinesExecuteCreateMaximumSetGenGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputebulkactions.NewClientFactory("50352BBD-59F1-4EE2-BA9C-A6E51B0B1B2B", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBulkActionsClient().VirtualMachinesExecuteCreate(ctx, "eastus2euap", armcomputebulkactions.ExecuteCreateRequest{
		ResourceConfigParameters: &armcomputebulkactions.ResourceProvisionPayload{
			ResourceCount: to.Ptr[int32](1),
			BaseProfile: map[string]any{
				"resourcegroupName": "yourresourcegroup",
				"computeApiVersion": "2023-09-01",
				"properties": map[string]any{
					"vmExtensions": []any{
						map[string]any{
							"name":     "Microsoft.Azure.Geneva.GenevaMonitoring",
							"location": "eastus2euap",
							"properties": map[string]any{
								"autoUpgradeMinorVersion": true,
								"enableAutomaticUpgrade":  true,
								"suppressFailures":        true,
								"publisher":               "Microsoft.Azure.Geneva",
								"type":                    "GenevaMonitoring",
								"typeHandlerVersion":      "2.0",
							},
						},
					},
					"hardwareProfile": map[string]any{
						"vmSize": "Standard_D2ads_v5",
					},
					"storageProfile": map[string]any{
						"imageReference": map[string]any{
							"publisher": "MicrosoftWindowsServer",
							"offer":     "WindowsServer",
							"sku":       "2022-datacenter-azure-edition",
							"version":   "latest",
						},
						"osDisk": map[string]any{
							"osType":       "Windows",
							"createOption": "FromImage",
							"caching":      "ReadWrite",
							"managedDisk": map[string]any{
								"storageAccountType": "Standard_LRS",
							},
							"deleteOption": "Detach",
							"diskSizeGB":   127,
						},
						"diskControllerType": "SCSI",
					},
					"networkProfile": map[string]any{
						"networkInterfaceConfigurations": []any{
							map[string]any{
								"name": "vmTest",
								"properties": map[string]any{
									"primary":            true,
									"enableIPForwarding": true,
									"ipConfigurations": []any{
										map[string]any{
											"name": "vmTest",
											"properties": map[string]any{
												"subnet": map[string]any{
													"id": "/subscriptions/264f0c8a-4d5f-496c-80df-b438624ce55f/resourceGroups/yourresourcegroup/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default",
												},
												"primary":                               true,
												"applicationGatewayBackendAddressPools": []any{},
												"loadBalancerBackendAddressPools":       []any{},
											},
										},
									},
								},
							},
						},
						"networkApiVersion": "2022-07-01",
					},
				},
			},
			ResourceOverrides: []map[string]any{
				{
					"name":     "testvmtestTwo",
					"location": "eastus2euap",
					"properties": map[string]any{
						"hardwareProfile": map[string]any{
							"vmSize": "Standard_D2ads_v5",
						},
						"osProfile": map[string]any{
							"computerName":  "testtestTwo",
							"adminUsername": "testUserName",
							"adminPassword": "YourStr0ngP@ssword123!",
							"windowsConfiguration": map[string]any{
								"provisionVmAgent":       true,
								"enableAutomaticUpdates": true,
								"patchSettings": map[string]any{
									"patchMode":      "AutomaticByPlatform",
									"assessmentMode": "ImageDefault",
								},
							},
						},
					},
				},
			},
		},
		ExecutionParameters: &armcomputebulkactions.ExecutionParameters{
			RetryPolicy: &armcomputebulkactions.RetryPolicy{
				RetryCount:           to.Ptr[int32](5),
				RetryWindowInMinutes: to.Ptr[int32](45),
			},
		},
		Correlationid: to.Ptr("7efcfae3-f50d-4323-9aba-1093a33368f8"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputebulkactions.BulkActionsClientVirtualMachinesExecuteCreateResponse{
	// 	CreateResourceOperationResponse: &armcomputebulkactions.CreateResourceOperationResponse{
	// 		Type: to.Ptr("VirtualMachine"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Description: to.Ptr("Create Resource Request"),
	// 		Results: []*armcomputebulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/TC_0"),
	// 				ErrorCode: to.Ptr("null"),
	// 				ErrorDetails: to.Ptr("null"),
	// 				Operation: &armcomputebulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("2a3fce8e-874c-45f4-9d27-1a804f3b7a0f"),
	// 					ResourceID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"),
	// 					OpType: to.Ptr(armcomputebulkactions.ResourceOperationTypeCreate),
	// 					SubscriptionID: to.Ptr("D8E30CC0-2763-4FCC-84A8-3C5659281032"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.667Z"); return t}()),
	// 					DeadlineType: to.Ptr(armcomputebulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armcomputebulkactions.OperationStateSucceeded),
	// 					Timezone: to.Ptr("UTC"),
	// 					ResourceOperationError: &armcomputebulkactions.ResourceOperationError{
	// 						ErrorCode: to.Ptr("null"),
	// 						ErrorDetails: to.Ptr("null"),
	// 					},
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T17:52:53.668Z"); return t}()),
	// 					RetryPolicy: &armcomputebulkactions.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](4),
	// 						RetryWindowInMinutes: to.Ptr[int32](27),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

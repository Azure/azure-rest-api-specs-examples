package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/VirtualMachineSize/list.json
func ExampleVirtualMachineSizesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineSizesClient().List(ctx, "eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineSizeListResult = armmachinelearning.VirtualMachineSizeListResult{
	// 	Value: []*armmachinelearning.VirtualMachineSize{
	// 		{
	// 			Name: to.Ptr("Standard_DS1_v2"),
	// 			EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 				BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 				UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 				Values: []*armmachinelearning.EstimatedVMPrice{
	// 					{
	// 						OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 						RetailPrice: to.Ptr[float64](0.13),
	// 						VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 					},
	// 					{
	// 						OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 						RetailPrice: to.Ptr[float64](0.01),
	// 						VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 					},
	// 					{
	// 						OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 						RetailPrice: to.Ptr[float64](0.07),
	// 						VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 					},
	// 					{
	// 						OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 						RetailPrice: to.Ptr[float64](0.05),
	// 						VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 				}},
	// 			},
	// 			Family: to.Ptr("standardDSv2Family"),
	// 			Gpus: to.Ptr[int32](0),
	// 			LowPriorityCapable: to.Ptr(true),
	// 			MaxResourceVolumeMB: to.Ptr[int32](7168),
	// 			MemoryGB: to.Ptr[float64](3.5),
	// 			OSVhdSizeMB: to.Ptr[int32](1047552),
	// 			PremiumIO: to.Ptr(true),
	// 			SupportedComputeTypes: []*string{
	// 				to.Ptr("AmlCompute"),
	// 				to.Ptr("ComputeInstance")},
	// 				VCPUs: to.Ptr[int32](1),
	// 			},
	// 			{
	// 				Name: to.Ptr("Standard_DS2_v2"),
	// 				EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 					BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 					UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 					Values: []*armmachinelearning.EstimatedVMPrice{
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](0.03),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 						},
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](0.15),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 							RetailPrice: to.Ptr[float64](0.1),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 						},
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 							RetailPrice: to.Ptr[float64](0.25),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 					}},
	// 				},
	// 				Family: to.Ptr("standardDSv2Family"),
	// 				Gpus: to.Ptr[int32](0),
	// 				LowPriorityCapable: to.Ptr(true),
	// 				MaxResourceVolumeMB: to.Ptr[int32](14336),
	// 				MemoryGB: to.Ptr[float64](7),
	// 				OSVhdSizeMB: to.Ptr[int32](1047552),
	// 				PremiumIO: to.Ptr(true),
	// 				SupportedComputeTypes: []*string{
	// 					to.Ptr("AmlCompute"),
	// 					to.Ptr("ComputeInstance"),
	// 					to.Ptr("MIR")},
	// 					VCPUs: to.Ptr[int32](2),
	// 				},
	// 				{
	// 					Name: to.Ptr("Standard_DS3_v2"),
	// 					EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 						BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 						UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 						Values: []*armmachinelearning.EstimatedVMPrice{
	// 							{
	// 								OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 								RetailPrice: to.Ptr[float64](0.2),
	// 								VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 							},
	// 							{
	// 								OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 								RetailPrice: to.Ptr[float64](0.06),
	// 								VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 							},
	// 							{
	// 								OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 								RetailPrice: to.Ptr[float64](0.5),
	// 								VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 							},
	// 							{
	// 								OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 								RetailPrice: to.Ptr[float64](0.29),
	// 								VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						}},
	// 					},
	// 					Family: to.Ptr("standardDSv2Family"),
	// 					Gpus: to.Ptr[int32](0),
	// 					LowPriorityCapable: to.Ptr(true),
	// 					MaxResourceVolumeMB: to.Ptr[int32](28672),
	// 					MemoryGB: to.Ptr[float64](14),
	// 					OSVhdSizeMB: to.Ptr[int32](1047552),
	// 					PremiumIO: to.Ptr(true),
	// 					SupportedComputeTypes: []*string{
	// 						to.Ptr("AmlCompute"),
	// 						to.Ptr("ComputeInstance"),
	// 						to.Ptr("MIR")},
	// 						VCPUs: to.Ptr[int32](4),
	// 					},
	// 					{
	// 						Name: to.Ptr("Standard_DS4_v2"),
	// 						EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 							BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 							UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 							Values: []*armmachinelearning.EstimatedVMPrice{
	// 								{
	// 									OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 									RetailPrice: to.Ptr[float64](0.12),
	// 									VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 								},
	// 								{
	// 									OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 									RetailPrice: to.Ptr[float64](0.4),
	// 									VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 								},
	// 								{
	// 									OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 									RetailPrice: to.Ptr[float64](1.01),
	// 									VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 								},
	// 								{
	// 									OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 									RetailPrice: to.Ptr[float64](0.58),
	// 									VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 							}},
	// 						},
	// 						Family: to.Ptr("standardDSv2Family"),
	// 						Gpus: to.Ptr[int32](0),
	// 						LowPriorityCapable: to.Ptr(true),
	// 						MaxResourceVolumeMB: to.Ptr[int32](57344),
	// 						MemoryGB: to.Ptr[float64](28),
	// 						OSVhdSizeMB: to.Ptr[int32](1047552),
	// 						PremiumIO: to.Ptr(true),
	// 						SupportedComputeTypes: []*string{
	// 							to.Ptr("AmlCompute"),
	// 							to.Ptr("ComputeInstance"),
	// 							to.Ptr("MIR")},
	// 							VCPUs: to.Ptr[int32](8),
	// 						},
	// 						{
	// 							Name: to.Ptr("Standard_DS5_v2"),
	// 							EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 								BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 								UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 								Values: []*armmachinelearning.EstimatedVMPrice{
	// 									{
	// 										OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 										RetailPrice: to.Ptr[float64](1.17),
	// 										VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 									},
	// 									{
	// 										OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 										RetailPrice: to.Ptr[float64](0.81),
	// 										VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 									},
	// 									{
	// 										OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 										RetailPrice: to.Ptr[float64](2.02),
	// 										VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 									},
	// 									{
	// 										OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 										RetailPrice: to.Ptr[float64](0.23),
	// 										VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 								}},
	// 							},
	// 							Family: to.Ptr("standardDSv2Family"),
	// 							Gpus: to.Ptr[int32](0),
	// 							LowPriorityCapable: to.Ptr(true),
	// 							MaxResourceVolumeMB: to.Ptr[int32](114688),
	// 							MemoryGB: to.Ptr[float64](56),
	// 							OSVhdSizeMB: to.Ptr[int32](1047552),
	// 							PremiumIO: to.Ptr(true),
	// 							SupportedComputeTypes: []*string{
	// 								to.Ptr("AmlCompute"),
	// 								to.Ptr("ComputeInstance"),
	// 								to.Ptr("MIR")},
	// 								VCPUs: to.Ptr[int32](16),
	// 							},
	// 							{
	// 								Name: to.Ptr("Standard_DS11_v2"),
	// 								EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 									BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 									UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 									Values: []*armmachinelearning.EstimatedVMPrice{
	// 										{
	// 											OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 											RetailPrice: to.Ptr[float64](0.26),
	// 											VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 										},
	// 										{
	// 											OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 											RetailPrice: to.Ptr[float64](0.18),
	// 											VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 										},
	// 										{
	// 											OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 											RetailPrice: to.Ptr[float64](0.11),
	// 											VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 										},
	// 										{
	// 											OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 											RetailPrice: to.Ptr[float64](0.04),
	// 											VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 									}},
	// 								},
	// 								Family: to.Ptr("standardDSv2Family"),
	// 								Gpus: to.Ptr[int32](0),
	// 								LowPriorityCapable: to.Ptr(true),
	// 								MaxResourceVolumeMB: to.Ptr[int32](28672),
	// 								MemoryGB: to.Ptr[float64](14),
	// 								OSVhdSizeMB: to.Ptr[int32](1047552),
	// 								PremiumIO: to.Ptr(true),
	// 								SupportedComputeTypes: []*string{
	// 									to.Ptr("AmlCompute"),
	// 									to.Ptr("ComputeInstance")},
	// 									VCPUs: to.Ptr[int32](2),
	// 								},
	// 								{
	// 									Name: to.Ptr("Standard_DS12_v2"),
	// 									EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 										BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 										UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 										Values: []*armmachinelearning.EstimatedVMPrice{
	// 											{
	// 												OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 												RetailPrice: to.Ptr[float64](0.37),
	// 												VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 											},
	// 											{
	// 												OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 												RetailPrice: to.Ptr[float64](0.53),
	// 												VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 											},
	// 											{
	// 												OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 												RetailPrice: to.Ptr[float64](0.21),
	// 												VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 											},
	// 											{
	// 												OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 												RetailPrice: to.Ptr[float64](0.07),
	// 												VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 										}},
	// 									},
	// 									Family: to.Ptr("standardDSv2Family"),
	// 									Gpus: to.Ptr[int32](0),
	// 									LowPriorityCapable: to.Ptr(true),
	// 									MaxResourceVolumeMB: to.Ptr[int32](57344),
	// 									MemoryGB: to.Ptr[float64](28),
	// 									OSVhdSizeMB: to.Ptr[int32](1047552),
	// 									PremiumIO: to.Ptr(true),
	// 									SupportedComputeTypes: []*string{
	// 										to.Ptr("AmlCompute"),
	// 										to.Ptr("ComputeInstance")},
	// 										VCPUs: to.Ptr[int32](4),
	// 									},
	// 									{
	// 										Name: to.Ptr("Standard_DS13_v2"),
	// 										EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 											BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 											UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 											Values: []*armmachinelearning.EstimatedVMPrice{
	// 												{
	// 													OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 													RetailPrice: to.Ptr[float64](0.15),
	// 													VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 												},
	// 												{
	// 													OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 													RetailPrice: to.Ptr[float64](0.42),
	// 													VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 												},
	// 												{
	// 													OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 													RetailPrice: to.Ptr[float64](0.74),
	// 													VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 												},
	// 												{
	// 													OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 													RetailPrice: to.Ptr[float64](1.06),
	// 													VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 											}},
	// 										},
	// 										Family: to.Ptr("standardDSv2Family"),
	// 										Gpus: to.Ptr[int32](0),
	// 										LowPriorityCapable: to.Ptr(true),
	// 										MaxResourceVolumeMB: to.Ptr[int32](114688),
	// 										MemoryGB: to.Ptr[float64](56),
	// 										OSVhdSizeMB: to.Ptr[int32](1047552),
	// 										PremiumIO: to.Ptr(true),
	// 										SupportedComputeTypes: []*string{
	// 											to.Ptr("AmlCompute"),
	// 											to.Ptr("ComputeInstance")},
	// 											VCPUs: to.Ptr[int32](8),
	// 										},
	// 										{
	// 											Name: to.Ptr("Standard_DS14_v2"),
	// 											EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 												BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 												UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 												Values: []*armmachinelearning.EstimatedVMPrice{
	// 													{
	// 														OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 														RetailPrice: to.Ptr[float64](0.3),
	// 														VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 													},
	// 													{
	// 														OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 														RetailPrice: to.Ptr[float64](1.48),
	// 														VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 													},
	// 													{
	// 														OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 														RetailPrice: to.Ptr[float64](0.84),
	// 														VMTier: to.Ptr(armmachinelearning.VMTierLowPriority),
	// 													},
	// 													{
	// 														OSType: to.Ptr(armmachinelearning.VMPriceOSTypeWindows),
	// 														RetailPrice: to.Ptr[float64](2.11),
	// 														VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 												}},
	// 											},
	// 											Family: to.Ptr("standardDSv2Family"),
	// 											Gpus: to.Ptr[int32](0),
	// 											LowPriorityCapable: to.Ptr(true),
	// 											MaxResourceVolumeMB: to.Ptr[int32](229376),
	// 											MemoryGB: to.Ptr[float64](112),
	// 											OSVhdSizeMB: to.Ptr[int32](1047552),
	// 											PremiumIO: to.Ptr(true),
	// 											SupportedComputeTypes: []*string{
	// 												to.Ptr("AmlCompute"),
	// 												to.Ptr("ComputeInstance")},
	// 												VCPUs: to.Ptr[int32](16),
	// 										}},
	// 									}
}

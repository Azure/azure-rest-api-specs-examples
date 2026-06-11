package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Compute/getAllowedVMSizesForResize.json
func ExampleComputeClient_GetAllowedResizeSizes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComputeClient().GetAllowedResizeSizes(ctx, "testrg123", "workspaces123", "compute123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.ComputeClientGetAllowedResizeSizesResponse{
	// 	VirtualMachineSizeListResult: armmachinelearning.VirtualMachineSizeListResult{
	// 		Value: []*armmachinelearning.VirtualMachineSize{
	// 			{
	// 				Name: to.Ptr("Standard_DS1_v2"),
	// 				EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 					BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 					UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 					Values: []*armmachinelearning.EstimatedVMPrice{
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](0.07),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 					},
	// 				},
	// 				Family: to.Ptr("standardDSv2Family"),
	// 				Gpus: to.Ptr[int32](0),
	// 				LowPriorityCapable: to.Ptr(true),
	// 				MaxResourceVolumeMB: to.Ptr[int32](7168),
	// 				MemoryGB: to.Ptr[float64](3.5),
	// 				OSVhdSizeMB: to.Ptr[int32](1047552),
	// 				PremiumIO: to.Ptr(true),
	// 				SupportedComputeTypes: []*string{
	// 					to.Ptr("AmlCompute"),
	// 					to.Ptr("ComputeInstance"),
	// 				},
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
	// 							RetailPrice: to.Ptr[float64](0.15),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 					},
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
	// 					to.Ptr("MIR"),
	// 				},
	// 				VCPUs: to.Ptr[int32](2),
	// 			},
	// 			{
	// 				Name: to.Ptr("Standard_DS3_v2"),
	// 				EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 					BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 					UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 					Values: []*armmachinelearning.EstimatedVMPrice{
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](0.29),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 					},
	// 				},
	// 				Family: to.Ptr("standardDSv2Family"),
	// 				Gpus: to.Ptr[int32](0),
	// 				LowPriorityCapable: to.Ptr(true),
	// 				MaxResourceVolumeMB: to.Ptr[int32](28672),
	// 				MemoryGB: to.Ptr[float64](14),
	// 				OSVhdSizeMB: to.Ptr[int32](1047552),
	// 				PremiumIO: to.Ptr(true),
	// 				SupportedComputeTypes: []*string{
	// 					to.Ptr("AmlCompute"),
	// 					to.Ptr("ComputeInstance"),
	// 					to.Ptr("MIR"),
	// 				},
	// 				VCPUs: to.Ptr[int32](4),
	// 			},
	// 			{
	// 				Name: to.Ptr("Standard_DS4_v2"),
	// 				EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 					BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 					UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 					Values: []*armmachinelearning.EstimatedVMPrice{
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](0.58),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 					},
	// 				},
	// 				Family: to.Ptr("standardDSv2Family"),
	// 				Gpus: to.Ptr[int32](0),
	// 				LowPriorityCapable: to.Ptr(true),
	// 				MaxResourceVolumeMB: to.Ptr[int32](57344),
	// 				MemoryGB: to.Ptr[float64](28),
	// 				OSVhdSizeMB: to.Ptr[int32](1047552),
	// 				PremiumIO: to.Ptr(true),
	// 				SupportedComputeTypes: []*string{
	// 					to.Ptr("AmlCompute"),
	// 					to.Ptr("ComputeInstance"),
	// 					to.Ptr("MIR"),
	// 				},
	// 				VCPUs: to.Ptr[int32](8),
	// 			},
	// 			{
	// 				Name: to.Ptr("Standard_DS5_v2"),
	// 				EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 					BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 					UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 					Values: []*armmachinelearning.EstimatedVMPrice{
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](1.17),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 					},
	// 				},
	// 				Family: to.Ptr("standardDSv2Family"),
	// 				Gpus: to.Ptr[int32](0),
	// 				LowPriorityCapable: to.Ptr(true),
	// 				MaxResourceVolumeMB: to.Ptr[int32](114688),
	// 				MemoryGB: to.Ptr[float64](56),
	// 				OSVhdSizeMB: to.Ptr[int32](1047552),
	// 				PremiumIO: to.Ptr(true),
	// 				SupportedComputeTypes: []*string{
	// 					to.Ptr("AmlCompute"),
	// 					to.Ptr("ComputeInstance"),
	// 					to.Ptr("MIR"),
	// 				},
	// 				VCPUs: to.Ptr[int32](16),
	// 			},
	// 			{
	// 				Name: to.Ptr("Standard_DS11_v2"),
	// 				EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 					BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 					UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 					Values: []*armmachinelearning.EstimatedVMPrice{
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](0.18),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 					},
	// 				},
	// 				Family: to.Ptr("standardDSv2Family"),
	// 				Gpus: to.Ptr[int32](0),
	// 				LowPriorityCapable: to.Ptr(true),
	// 				MaxResourceVolumeMB: to.Ptr[int32](28672),
	// 				MemoryGB: to.Ptr[float64](14),
	// 				OSVhdSizeMB: to.Ptr[int32](1047552),
	// 				PremiumIO: to.Ptr(true),
	// 				SupportedComputeTypes: []*string{
	// 					to.Ptr("AmlCompute"),
	// 					to.Ptr("ComputeInstance"),
	// 				},
	// 				VCPUs: to.Ptr[int32](2),
	// 			},
	// 			{
	// 				Name: to.Ptr("Standard_DS12_v2"),
	// 				EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 					BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 					UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 					Values: []*armmachinelearning.EstimatedVMPrice{
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](0.37),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 					},
	// 				},
	// 				Family: to.Ptr("standardDSv2Family"),
	// 				Gpus: to.Ptr[int32](0),
	// 				LowPriorityCapable: to.Ptr(true),
	// 				MaxResourceVolumeMB: to.Ptr[int32](57344),
	// 				MemoryGB: to.Ptr[float64](28),
	// 				OSVhdSizeMB: to.Ptr[int32](1047552),
	// 				PremiumIO: to.Ptr(true),
	// 				SupportedComputeTypes: []*string{
	// 					to.Ptr("AmlCompute"),
	// 					to.Ptr("ComputeInstance"),
	// 				},
	// 				VCPUs: to.Ptr[int32](4),
	// 			},
	// 			{
	// 				Name: to.Ptr("Standard_DS13_v2"),
	// 				EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 					BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 					UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 					Values: []*armmachinelearning.EstimatedVMPrice{
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](0.74),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 					},
	// 				},
	// 				Family: to.Ptr("standardDSv2Family"),
	// 				Gpus: to.Ptr[int32](0),
	// 				LowPriorityCapable: to.Ptr(true),
	// 				MaxResourceVolumeMB: to.Ptr[int32](114688),
	// 				MemoryGB: to.Ptr[float64](56),
	// 				OSVhdSizeMB: to.Ptr[int32](1047552),
	// 				PremiumIO: to.Ptr(true),
	// 				SupportedComputeTypes: []*string{
	// 					to.Ptr("AmlCompute"),
	// 					to.Ptr("ComputeInstance"),
	// 				},
	// 				VCPUs: to.Ptr[int32](8),
	// 			},
	// 			{
	// 				Name: to.Ptr("Standard_DS14_v2"),
	// 				EstimatedVMPrices: &armmachinelearning.EstimatedVMPrices{
	// 					BillingCurrency: to.Ptr(armmachinelearning.BillingCurrencyUSD),
	// 					UnitOfMeasure: to.Ptr(armmachinelearning.UnitOfMeasureOneHour),
	// 					Values: []*armmachinelearning.EstimatedVMPrice{
	// 						{
	// 							OSType: to.Ptr(armmachinelearning.VMPriceOSTypeLinux),
	// 							RetailPrice: to.Ptr[float64](1.48),
	// 							VMTier: to.Ptr(armmachinelearning.VMTierStandard),
	// 						},
	// 					},
	// 				},
	// 				Family: to.Ptr("standardDSv2Family"),
	// 				Gpus: to.Ptr[int32](0),
	// 				LowPriorityCapable: to.Ptr(true),
	// 				MaxResourceVolumeMB: to.Ptr[int32](229376),
	// 				MemoryGB: to.Ptr[float64](112),
	// 				OSVhdSizeMB: to.Ptr[int32](1047552),
	// 				PremiumIO: to.Ptr(true),
	// 				SupportedComputeTypes: []*string{
	// 					to.Ptr("AmlCompute"),
	// 					to.Ptr("ComputeInstance"),
	// 				},
	// 				VCPUs: to.Ptr[int32](16),
	// 			},
	// 		},
	// 	},
	// }
}

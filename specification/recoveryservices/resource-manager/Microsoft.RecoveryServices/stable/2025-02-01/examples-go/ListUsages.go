package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/ListUsages.json
func ExampleUsagesClient_NewListByVaultsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListByVaultsPager("Default-RecoveryServices-ResourceGroup", "swaggerExample", nil)
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
		// page.VaultUsageList = armrecoveryservices.VaultUsageList{
		// 	Value: []*armrecoveryservices.VaultUsage{
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Backup management servers"),
		// 				Value: to.Ptr("MABContainersCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](6),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Backup items/Azure virtual machine backup"),
		// 				Value: to.Ptr("ProtectedItemCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](3),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Critical"),
		// 				Value: to.Ptr("ProtectedItemCriticalCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Warning"),
		// 				Value: to.Ptr("ProtectedItemWarningCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Virtual Machines"),
		// 				Value: to.Ptr("IaaSVMProtectedItemCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Critical"),
		// 				Value: to.Ptr("IaaSVMProtectedItemCriticalCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Warning"),
		// 				Value: to.Ptr("IaaSVMProtectedItemWarningCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("File-Folders"),
		// 				Value: to.Ptr("MABProtectedItemCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("DPM Protected Items Count"),
		// 				Value: to.Ptr("DPMProtectedItemCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](1),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Azure Backup Server Protected Items Count"),
		// 				Value: to.Ptr("AzureBackupServerProtectedItemCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](2),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("In progress"),
		// 				Value: to.Ptr("InProgressJobsCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			QuotaPeriod: to.Ptr("P1D"),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Failed"),
		// 				Value: to.Ptr("FailedJobsCount"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			QuotaPeriod: to.Ptr("P1D"),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Cloud - GRS"),
		// 				Value: to.Ptr("GRSStorageUsage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](117007930),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitBytes),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Cloud - LRS"),
		// 				Value: to.Ptr("LRSStorageUsage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitBytes),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Protected Instances"),
		// 				Value: to.Ptr("ManagedInstances"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](5),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitCount),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Dedup - GRS"),
		// 				Value: to.Ptr("GRSDedupStorageUsage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitBytes),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Dedup - LRS"),
		// 				Value: to.Ptr("LRSDedupStorageUsage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitBytes),
		// 		},
		// 		{
		// 			Name: &armrecoveryservices.NameInfo{
		// 				LocalizedValue: to.Ptr("Backup Engines' Disk Used"),
		// 				Value: to.Ptr("UsedDiskSize"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](117851553792),
		// 			Limit: to.Ptr[int64](-1),
		// 			Unit: to.Ptr(armrecoveryservices.UsagesUnitBytes),
		// 	}},
		// }
	}
}

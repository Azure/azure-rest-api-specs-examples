package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79a5aa63c0551c1b5af1d2853cceb495283d334/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/CreateOrUpdateSqlVirtualMachineAutomatedBackupWeekly.json
func ExampleSQLVirtualMachinesClient_BeginCreateOrUpdate_createsOrUpdatesASqlVirtualMachineForAutomatedBackUpSettingsWithWeeklyAndDaysOfTheWeekToRunTheBackUp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLVirtualMachinesClient().BeginCreateOrUpdate(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVirtualMachine{
		Location: to.Ptr("northeurope"),
		Properties: &armsqlvirtualmachine.Properties{
			AutoBackupSettings: &armsqlvirtualmachine.AutoBackupSettings{
				BackupScheduleType: to.Ptr(armsqlvirtualmachine.BackupScheduleTypeManual),
				BackupSystemDbs:    to.Ptr(true),
				DaysOfWeek: []*armsqlvirtualmachine.AutoBackupDaysOfWeek{
					to.Ptr(armsqlvirtualmachine.AutoBackupDaysOfWeekMonday),
					to.Ptr(armsqlvirtualmachine.AutoBackupDaysOfWeekFriday)},
				Enable:                to.Ptr(true),
				EnableEncryption:      to.Ptr(true),
				FullBackupFrequency:   to.Ptr(armsqlvirtualmachine.FullBackupFrequencyTypeWeekly),
				FullBackupStartTime:   to.Ptr[int32](6),
				FullBackupWindowHours: to.Ptr[int32](11),
				LogBackupFrequency:    to.Ptr[int32](10),
				Password:              to.Ptr("<Password>"),
				RetentionPeriod:       to.Ptr[int32](17),
				StorageAccessKey:      to.Ptr("<primary storage access key>"),
				StorageAccountURL:     to.Ptr("https://teststorage.blob.core.windows.net/"),
				StorageContainerName:  to.Ptr("testcontainer"),
			},
			AutoPatchingSettings: &armsqlvirtualmachine.AutoPatchingSettings{
				DayOfWeek:                     to.Ptr(armsqlvirtualmachine.DayOfWeekSunday),
				Enable:                        to.Ptr(true),
				MaintenanceWindowDuration:     to.Ptr[int32](60),
				MaintenanceWindowStartingHour: to.Ptr[int32](2),
			},
			KeyVaultCredentialSettings: &armsqlvirtualmachine.KeyVaultCredentialSettings{
				Enable: to.Ptr(false),
			},
			ServerConfigurationsManagementSettings: &armsqlvirtualmachine.ServerConfigurationsManagementSettings{
				AdditionalFeaturesServerConfigurations: &armsqlvirtualmachine.AdditionalFeaturesServerConfigurations{
					IsRServicesEnabled: to.Ptr(false),
				},
				SQLConnectivityUpdateSettings: &armsqlvirtualmachine.SQLConnectivityUpdateSettings{
					ConnectivityType:      to.Ptr(armsqlvirtualmachine.ConnectivityTypePRIVATE),
					Port:                  to.Ptr[int32](1433),
					SQLAuthUpdatePassword: to.Ptr("<password>"),
					SQLAuthUpdateUserName: to.Ptr("sqllogin"),
				},
				SQLStorageUpdateSettings: &armsqlvirtualmachine.SQLStorageUpdateSettings{
					DiskConfigurationType: to.Ptr(armsqlvirtualmachine.DiskConfigurationTypeNEW),
					DiskCount:             to.Ptr[int32](1),
					StartingDeviceID:      to.Ptr[int32](2),
				},
				SQLWorkloadTypeUpdateSettings: &armsqlvirtualmachine.SQLWorkloadTypeUpdateSettings{
					SQLWorkloadType: to.Ptr(armsqlvirtualmachine.SQLWorkloadTypeOLTP),
				},
			},
			SQLImageSKU:              to.Ptr(armsqlvirtualmachine.SQLImageSKUEnterprise),
			SQLManagement:            to.Ptr(armsqlvirtualmachine.SQLManagementModeFull),
			SQLServerLicenseType:     to.Ptr(armsqlvirtualmachine.SQLServerLicenseTypePAYG),
			VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLVirtualMachine = armsqlvirtualmachine.SQLVirtualMachine{
	// 	Name: to.Ptr("testvm"),
	// 	Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachines"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm"),
	// 	Location: to.Ptr("northeurope"),
	// 	Properties: &armsqlvirtualmachine.Properties{
	// 		EnableAutomaticUpgrade: to.Ptr(false),
	// 		LeastPrivilegeMode: to.Ptr(armsqlvirtualmachine.LeastPrivilegeModeNotSet),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SQLImageSKU: to.Ptr(armsqlvirtualmachine.SQLImageSKUEnterprise),
	// 		SQLManagement: to.Ptr(armsqlvirtualmachine.SQLManagementModeFull),
	// 		SQLServerLicenseType: to.Ptr(armsqlvirtualmachine.SQLServerLicenseTypePAYG),
	// 		VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
	// 	},
	// }
}

package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.AvailableProviderOperationList = armstorsimple8000series.AvailableProviderOperationList{
		// 	Value: []*armstorsimple8000series.AvailableProviderOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/Managers/certificates/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("The Update Resource Certificate operation updates the resource/vault credential certificate."),
		// 				Operation: to.Ptr("Update Resource Certificate"),
		// 				Provider: to.Ptr("Microsoft.StorSimple"),
		// 				Resource: to.Ptr("Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/updateSummary/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Update Summary"),
		// 				Operation: to.Ptr("List Update Summary"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Update Summary"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/scanForUpdates/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Scan for updates in a device."),
		// 				Operation: to.Ptr("Scan Updates"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/download/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Dowload updates for a device."),
		// 				Operation: to.Ptr("Download Device Updates"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/install/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Install updates on a device."),
		// 				Operation: to.Ptr("Install Updates (1200 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/timeSettings/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Time Settings"),
		// 				Operation: to.Ptr("List Time Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Time Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/alertSettings/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Alert Settings"),
		// 				Operation: to.Ptr("List Alert Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Alert Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/networkSettings/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Network Settings"),
		// 				Operation: to.Ptr("List Network Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Network Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/alertSettings/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Alert Settings"),
		// 				Operation: to.Ptr("Create or Update Alert Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Alert Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/securitySettings/update/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Update the security settings."),
		// 				Operation: to.Ptr("Update (1200 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Security Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Devices"),
		// 				Operation: to.Ptr("List Devices"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Devices"),
		// 				Operation: to.Ptr("Create or Update Devices"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Devices"),
		// 				Operation: to.Ptr("Delete Devices"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/deactivate/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deactivates a device."),
		// 				Operation: to.Ptr("Deactivate Device"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/publishSupportPackage/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Publish support package of a device for Microsoft Support troubleshooting."),
		// 				Operation: to.Ptr("Publish Support Package"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/metricsDefinitions/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics Definitions"),
		// 				Operation: to.Ptr("List Metrics Definitions"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics Definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/metrics/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics"),
		// 				Operation: to.Ptr("List Metrics"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/failover/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Failover of the device."),
		// 				Operation: to.Ptr("Failover Device"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/shares/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Shares"),
		// 				Operation: to.Ptr("Create or Update Shares"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Shares (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/shares/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Shares"),
		// 				Operation: to.Ptr("List Shares"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Shares (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/shares/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Shares"),
		// 				Operation: to.Ptr("Delete Shares"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Shares (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/shares/metricsDefinitions/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics Definitions"),
		// 				Operation: to.Ptr("List Metrics Definitions"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics Definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/shares/metrics/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics"),
		// 				Operation: to.Ptr("List Metrics"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/disks/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Disks"),
		// 				Operation: to.Ptr("List Disks"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Disks (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/disks/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Disks"),
		// 				Operation: to.Ptr("Create or Update Disks"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Disks (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/disks/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Disks"),
		// 				Operation: to.Ptr("Delete Disks"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Disks (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/disks/metricsDefinitions/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics Definitions"),
		// 				Operation: to.Ptr("List Metrics Definitions"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics Definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/disks/metrics/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics"),
		// 				Operation: to.Ptr("List Metrics"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the iSCSI Servers"),
		// 				Operation: to.Ptr("List iSCSI Servers"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("iSCSI Servers (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the iSCSI Servers"),
		// 				Operation: to.Ptr("Create or Update iSCSI Servers"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("iSCSI Servers (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the iSCSI Servers"),
		// 				Operation: to.Ptr("Delete iSCSI Servers"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("iSCSI Servers (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/backup/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Take backup of an iSCSI server."),
		// 				Operation: to.Ptr("Take Backup"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("iSCSI Servers (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/metricsDefinitions/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics Definitions"),
		// 				Operation: to.Ptr("List Metrics Definitions"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics Definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiservers/metrics/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics"),
		// 				Operation: to.Ptr("List Metrics"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupScheduleGroups/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Backup Schedule Groups"),
		// 				Operation: to.Ptr("List Backup Schedule Groups"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Schedule Groups (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupScheduleGroups/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Backup Schedule Groups"),
		// 				Operation: to.Ptr("Create or Update Backup Schedule Groups"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Schedule Groups (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupScheduleGroups/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Backup Schedule Groups"),
		// 				Operation: to.Ptr("Delete Backup Schedule Groups"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Schedule Groups (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/alerts/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Alerts"),
		// 				Operation: to.Ptr("List Alerts"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Alerts"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/clearAlerts/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Clear all the alerts associated with the device manager."),
		// 				Operation: to.Ptr("Clears All Alerts"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/sendTestAlertEmail/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Send test alert email to configured email recipients."),
		// 				Operation: to.Ptr("Send Test Alert Email"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/storageAccountCredentials/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Storage Account Credentials"),
		// 				Operation: to.Ptr("Create or Update Storage Account Credentials"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Storage Account Credentials"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/storageAccountCredentials/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Storage Account Credentials"),
		// 				Operation: to.Ptr("List Storage Account Credentials"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Storage Account Credentials"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/storageAccountCredentials/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Storage Account Credentials"),
		// 				Operation: to.Ptr("Delete Storage Account Credentials"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Storage Account Credentials"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/getActivationKey/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Get activation key for the device manager."),
		// 				Operation: to.Ptr("Get Activation Key"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/regenerateActivationKey/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Regenerate activation key for the device manager."),
		// 				Operation: to.Ptr("Regenerate Activation Key"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/regenarateRegistationCertificate/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Regenerate registration certificate for the device managers."),
		// 				Operation: to.Ptr("Regenerate Registration Certificate"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/encryptionSettings/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Encryption Settings"),
		// 				Operation: to.Ptr("List Encryption Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Encryption Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/getEncryptionKey/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Get encryption key for the device manager."),
		// 				Operation: to.Ptr("Get Encryption Key"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Device Managers"),
		// 				Operation: to.Ptr("List Device Managers"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Device Managers"),
		// 				Operation: to.Ptr("Delete Device Managers"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Device Managers"),
		// 				Operation: to.Ptr("Create or Update Device Managers"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/metricsDefinitions/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics Definitions"),
		// 				Operation: to.Ptr("List Metrics Definitions"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics Definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/metrics/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics"),
		// 				Operation: to.Ptr("List Metrics"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/chapSettings/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Chap Settings"),
		// 				Operation: to.Ptr("Create or Update Chap Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Chap Settings (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/chapSettings/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Chap Settings"),
		// 				Operation: to.Ptr("List Chap Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Chap Settings (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/chapSettings/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Chap Settings"),
		// 				Operation: to.Ptr("Delete Chap Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Chap Settings (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backups/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Backup Set"),
		// 				Operation: to.Ptr("List Backup Set"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Set"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backups/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Backup Set"),
		// 				Operation: to.Ptr("Delete Backup Set"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Set"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backups/elements/clone/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Clone a share or volume using a backup element."),
		// 				Operation: to.Ptr("Clone"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Elements (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/storageDomains/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Storage Domains"),
		// 				Operation: to.Ptr("List Storage Domains"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Storage Domains (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/storageDomains/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Storage Domains"),
		// 				Operation: to.Ptr("Create or Update Storage Domains"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Storage Domains (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/storageDomains/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Storage Domains"),
		// 				Operation: to.Ptr("Delete Storage Domains"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Storage Domains (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/jobs/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Jobs"),
		// 				Operation: to.Ptr("List Jobs"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/accessControlRecords/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Access Control Records"),
		// 				Operation: to.Ptr("List Access Control Records"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Access Control Records"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/accessControlRecords/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the Access Control Records"),
		// 				Operation: to.Ptr("Create or Update Access Control Records"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Access Control Records"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/accessControlRecords/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the Access Control Records"),
		// 				Operation: to.Ptr("Delete Access Control Records"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Access Control Records"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the File Servers"),
		// 				Operation: to.Ptr("List File Servers"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("File Servers (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create or update the File Servers"),
		// 				Operation: to.Ptr("Create or Update File Servers"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("File Servers (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes the File Servers"),
		// 				Operation: to.Ptr("Delete File Servers"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("File Servers (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/backup/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Take backup of an File Server."),
		// 				Operation: to.Ptr("Take Backup"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("File Servers (1200 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/metricsDefinitions/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics Definitions"),
		// 				Operation: to.Ptr("List Metrics Definitions"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics Definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/fileservers/metrics/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Metrics"),
		// 				Operation: to.Ptr("List Metrics"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/timeSettings/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Creates a new or updates Time Settings"),
		// 				Operation: to.Ptr("Creates or Updates Time Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Time Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/networkSettings/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Creates a new or updates Network Settings"),
		// 				Operation: to.Ptr("Creates or Updates Network Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Network Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/securitySettings/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Security Settings"),
		// 				Operation: to.Ptr("List Security Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Security Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/securitySettings/syncRemoteManagementCertificate/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Synchronize the remote management certificate for a device."),
		// 				Operation: to.Ptr("Synchronize Remote Management Certificate (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Security Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/securitySettings/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Creates a new or updates Security Settings"),
		// 				Operation: to.Ptr("Creates or Updates Security Settings"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Security Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/configureDevice/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Configures a device"),
		// 				Operation: to.Ptr("Configure Device"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/installUpdates/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Installs updates on the devices"),
		// 				Operation: to.Ptr("Install Updates (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/listFailoverSets/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the failover sets for an existing device."),
		// 				Operation: to.Ptr("List Failover Sets"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/listFailoverTargets/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List failover targets of the devices"),
		// 				Operation: to.Ptr("List Failover Targets"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/listActivationKey/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Gets the activation key of the StorSimple Device Manager."),
		// 				Operation: to.Ptr("Gets Activation Key"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/listPublicEncryptionKey/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List public encryption keys of a StorSimple Device Manager."),
		// 				Operation: to.Ptr("List Public Encryption Key"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/publicEncryptionKey/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List public encryption key of the device manager"),
		// 				Operation: to.Ptr("List Public Encryption Key"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/listPrivateEncryptionKey/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Gets private encryption key for a StorSimple Device Manager."),
		// 				Operation: to.Ptr("List Private Encryption Key"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/storageAccountCredentials/listAccessKey/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List access keys of Storage Account Credentials"),
		// 				Operation: to.Ptr("List Access Keys"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Storage Account Credentials"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/jobs/cancel/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Cancel a running job"),
		// 				Operation: to.Ptr("Cancel Job"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Creates a new or updates Volume Containers (8000 Series Only)"),
		// 				Operation: to.Ptr("Creates or Updates Volume Containers (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Volume Containers (8000 Series Only) (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Volume Containers (8000 Series Only)"),
		// 				Operation: to.Ptr("List Volume Containers (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Volume Containers (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes an existing Volume Containers (8000 Series Only)"),
		// 				Operation: to.Ptr("Deletes Volume Containers (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Volume Containers (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/listEncryptionKeys/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List encryption keys of Volume Containers"),
		// 				Operation: to.Ptr("List Encryption Keys"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Volume Containers (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/rolloverEncryptionKey/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Rollover encryption keys of Volume Containers"),
		// 				Operation: to.Ptr("Rollover Encryption Keys"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Volume Containers (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/metricsDefinitions/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Metrics Definitions"),
		// 				Operation: to.Ptr("List Metrics Definitions"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics Definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/metrics/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Metrics"),
		// 				Operation: to.Ptr("List Metrics"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/volumes/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Volumes"),
		// 				Operation: to.Ptr("List Volumes"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Volumes (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/volumes/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Creates a new or updates Volumes"),
		// 				Operation: to.Ptr("Creates or Updates Volumes"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Volumes"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/volumes/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes an existing Volumes"),
		// 				Operation: to.Ptr("Deletes Volumes"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Volumes"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/volumes/metricsDefinitions/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Metrics Definitions"),
		// 				Operation: to.Ptr("List Metrics Definitions"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics Definitions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/volumeContainers/volumes/metrics/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Metrics"),
		// 				Operation: to.Ptr("List Metrics"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Metrics"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/hardwareComponentGroups/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Hardware Component Groups"),
		// 				Operation: to.Ptr("List Hardware Component Groups"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Hardware Component Groups (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/hardwareComponentGroups/changeControllerPowerState/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Change controller power state of hardware component groups"),
		// 				Operation: to.Ptr("Change controller power state of hardware component groups"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Hardware Component Groups"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/bandwidthSettings/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Bandwidth Settings (8000 Series Only)"),
		// 				Operation: to.Ptr("List Bandwidth Settings (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Bandwidth Settings (8000 Series Only) (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/bandwidthSettings/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Creates a new or updates Bandwidth Settings (8000 Series Only)"),
		// 				Operation: to.Ptr("Creates or Updates Bandwidth Settings (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Bandwidth Settings (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/bandwidthSettings/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes an existing Bandwidth Settings (8000 Series Only)"),
		// 				Operation: to.Ptr("Deletes Bandwidth Settings (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Bandwidth Settings (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backups/restore/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Restore all the volumes from a backup set."),
		// 				Operation: to.Ptr("Restore from Backup Set"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Set"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/schedules/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Creates a new or updates Schedules"),
		// 				Operation: to.Ptr("Creates or Updates Schedules"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Schedules (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/schedules/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Schedules"),
		// 				Operation: to.Ptr("List Schedules"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Schedules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/schedules/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes an existing Schedules"),
		// 				Operation: to.Ptr("Deletes Schedules"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Schedules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Creates a new or updates Backup Polices (8000 Series Only)"),
		// 				Operation: to.Ptr("Creates or Updates Backup Polices (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Polices (8000 Series Only) (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Backup Polices (8000 Series Only)"),
		// 				Operation: to.Ptr("List Backup Polices (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Polices (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes an existing Backup Polices (8000 Series Only)"),
		// 				Operation: to.Ptr("Deletes Backup Polices (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Polices (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies/backup/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Take a manual backup to create an on-demand backup of all the volumes protected by the policy."),
		// 				Operation: to.Ptr("Take Backup"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Backup Polices (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/migrationSourceConfigurations/import/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Import source configurations for migration"),
		// 				Operation: to.Ptr("Import Source Configurations"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Migration Source Configurations (8000 Series Only) (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/migrationSourceConfigurations/startMigrationEstimate/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Start a job to estimate the duration of the migration process."),
		// 				Operation: to.Ptr("Start Migration Estimate"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Migration Source Configurations (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/migrationSourceConfigurations/startMigration/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Start migration using source configurations"),
		// 				Operation: to.Ptr("Start Migration"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Migration Source Configurations (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/migrationSourceConfigurations/confirmMigration/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Confirms a successful migration and commit it."),
		// 				Operation: to.Ptr("Confirm Migration"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Migration Source Configurations (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/migrationSourceConfigurations/fetchMigrationEstimate/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Fetch the status for the migration estimation job."),
		// 				Operation: to.Ptr("Fetch Migration Estimate"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Migration Source Configurations (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/migrationSourceConfigurations/fetchMigrationStatus/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Fetch the status for the migration."),
		// 				Operation: to.Ptr("Fetch Migration Status"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Migration Source Configurations (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/devices/migrationSourceConfigurations/fetchConfirmMigrationStatus/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Fetch the confirm status of migration."),
		// 				Operation: to.Ptr("Fetch Migration Confirm Status"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Migration Source Configurations (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/provisionCloudAppliance/action"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create a new cloud appliance."),
		// 				Operation: to.Ptr("Create Cloud Appliance (8000 Series Only)"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Device Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/managers/cloudApplianceConfigurations/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("List the Cloud Appliance Supported Configurations"),
		// 				Operation: to.Ptr("List Cloud Appliance Supported Configurations"),
		// 				Provider: to.Ptr("Microsoft StorSimple Device Manager"),
		// 				Resource: to.Ptr("Cloud Appliance Supported Configurations (8000 Series Only)"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/Managers/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("Create Vault operation creates an Azure resource of type 'vault'"),
		// 				Operation: to.Ptr("Create Vault"),
		// 				Provider: to.Ptr("Microsoft.StorSimple"),
		// 				Resource: to.Ptr("Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/Managers/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("The Get Vault operation gets an object representing the Azure resource of type 'vault'"),
		// 				Operation: to.Ptr("Get Vault"),
		// 				Provider: to.Ptr("Microsoft.StorSimple"),
		// 				Resource: to.Ptr("Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/Managers/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("The Delete Vault operation deletes the specified Azure resource of type 'vault'"),
		// 				Operation: to.Ptr("Delete Vault"),
		// 				Provider: to.Ptr("Microsoft.StorSimple"),
		// 				Resource: to.Ptr("Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/Managers/extendedInformation/read"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("The Get Extended Info operation gets an object's Extended Info representing the Azure resource of type ?vault?"),
		// 				Operation: to.Ptr("Get Extended Info"),
		// 				Provider: to.Ptr("Microsoft.StorSimple"),
		// 				Resource: to.Ptr("Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/Managers/extendedInformation/write"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("The Get Extended Info operation gets an object's Extended Info representing the Azure resource of type ?vault?"),
		// 				Operation: to.Ptr("Get Extended Info"),
		// 				Provider: to.Ptr("Microsoft.StorSimple"),
		// 				Resource: to.Ptr("Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorSimple/Managers/extendedInformation/delete"),
		// 			Display: &armstorsimple8000series.AvailableProviderOperationDisplay{
		// 				Description: to.Ptr("The Get Extended Info operation gets an object's Extended Info representing the Azure resource of type ?vault?"),
		// 				Operation: to.Ptr("Get Extended Info"),
		// 				Provider: to.Ptr("Microsoft.StorSimple"),
		// 				Resource: to.Ptr("Managers"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 	}},
		// }
	}
}

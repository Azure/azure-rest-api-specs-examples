const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a SQL virtual machine.
 *
 * @summary creates or updates a SQL virtual machine.
 * x-ms-original-file: 2023-10-01/CreateOrUpdateSqlVirtualMachineAutomatedBackupWeekly.json
 */
async function createsOrUpdatesASQLVirtualMachineForAutomatedBackUpSettingsWithWeeklyAndDaysOfTheWeekToRunTheBackUp() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.sqlVirtualMachines.createOrUpdate("testrg", "testvm", {
    location: "northeurope",
    autoBackupSettings: {
      backupScheduleType: "Manual",
      backupSystemDbs: true,
      daysOfWeek: ["Monday", "Friday"],
      enable: true,
      enableEncryption: true,
      fullBackupFrequency: "Weekly",
      fullBackupStartTime: 6,
      fullBackupWindowHours: 11,
      logBackupFrequency: 10,
      password: "<Password>",
      retentionPeriod: 17,
      storageAccessKey: "<primary storage access key>",
      storageAccountUrl: "https://teststorage.blob.core.windows.net/",
      storageContainerName: "testcontainer",
    },
    autoPatchingSettings: {
      dayOfWeek: "Sunday",
      enable: true,
      maintenanceWindowDuration: 60,
      maintenanceWindowStartingHour: 2,
    },
    keyVaultCredentialSettings: { enable: false },
    serverConfigurationsManagementSettings: {
      additionalFeaturesServerConfigurations: { isRServicesEnabled: false },
      sqlConnectivityUpdateSettings: {
        connectivityType: "PRIVATE",
        port: 1433,
        sqlAuthUpdatePassword: "<password>",
        sqlAuthUpdateUserName: "sqllogin",
      },
      sqlStorageUpdateSettings: { diskConfigurationType: "NEW", diskCount: 1, startingDeviceId: 2 },
      sqlWorkloadTypeUpdateSettings: { sqlWorkloadType: "OLTP" },
    },
    sqlImageSku: "Enterprise",
    sqlServerLicenseType: "PAYG",
    virtualMachineResourceId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm",
  });
  console.log(result);
}

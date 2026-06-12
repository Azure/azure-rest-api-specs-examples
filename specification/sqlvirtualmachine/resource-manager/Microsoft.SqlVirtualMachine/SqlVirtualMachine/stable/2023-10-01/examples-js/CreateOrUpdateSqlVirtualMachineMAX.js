const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a SQL virtual machine.
 *
 * @summary creates or updates a SQL virtual machine.
 * x-ms-original-file: 2023-10-01/CreateOrUpdateSqlVirtualMachineMAX.json
 */
async function createsOrUpdatesASQLVirtualMachineWithMaxParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.sqlVirtualMachines.createOrUpdate("testrg", "testvm", {
    location: "northeurope",
    assessmentSettings: {
      enable: true,
      runImmediately: true,
      schedule: { dayOfWeek: "Sunday", enable: true, startTime: "23:17", weeklyInterval: 1 },
    },
    autoBackupSettings: {
      backupScheduleType: "Manual",
      backupSystemDbs: true,
      enable: true,
      enableEncryption: true,
      fullBackupFrequency: "Daily",
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
    enableAutomaticUpgrade: true,
    keyVaultCredentialSettings: { enable: false },
    leastPrivilegeMode: "Enabled",
    serverConfigurationsManagementSettings: {
      additionalFeaturesServerConfigurations: { isRServicesEnabled: false },
      azureAdAuthenticationSettings: { clientId: "11111111-2222-3333-4444-555555555555" },
      sqlConnectivityUpdateSettings: {
        connectivityType: "PRIVATE",
        port: 1433,
        sqlAuthUpdatePassword: "<password>",
        sqlAuthUpdateUserName: "sqllogin",
      },
      sqlInstanceSettings: {
        collation: "SQL_Latin1_General_CP1_CI_AS",
        isIfiEnabled: true,
        isLpimEnabled: true,
        isOptimizeForAdHocWorkloadsEnabled: true,
        maxDop: 8,
        maxServerMemoryMB: 128,
        minServerMemoryMB: 0,
      },
      sqlStorageUpdateSettings: { diskConfigurationType: "NEW", diskCount: 1, startingDeviceId: 2 },
      sqlWorkloadTypeUpdateSettings: { sqlWorkloadType: "OLTP" },
    },
    sqlImageSku: "Enterprise",
    sqlServerLicenseType: "PAYG",
    storageConfigurationSettings: {
      diskConfigurationType: "NEW",
      enableStorageConfigBlade: true,
      sqlDataSettings: { defaultFilePath: "F:\\folderpath\\", luns: [0], useStoragePool: false },
      sqlLogSettings: { defaultFilePath: "G:\\folderpath\\", luns: [1], useStoragePool: false },
      sqlSystemDbOnDataDisk: true,
      sqlTempDbSettings: {
        dataFileCount: 8,
        dataFileSize: 256,
        dataGrowth: 512,
        defaultFilePath: "D:\\TEMP",
        logFileSize: 256,
        logGrowth: 512,
        luns: [2],
        useStoragePool: false,
      },
      storageWorkloadType: "OLTP",
    },
    virtualMachineResourceId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm",
  });
  console.log(result);
}

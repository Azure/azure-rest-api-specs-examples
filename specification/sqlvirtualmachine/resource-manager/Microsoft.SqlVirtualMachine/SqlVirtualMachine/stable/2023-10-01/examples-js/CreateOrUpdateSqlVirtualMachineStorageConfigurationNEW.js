const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a SQL virtual machine.
 *
 * @summary creates or updates a SQL virtual machine.
 * x-ms-original-file: 2023-10-01/CreateOrUpdateSqlVirtualMachineStorageConfigurationNEW.json
 */
async function createsOrUpdatesASQLVirtualMachineForStorageConfigurationSettingsToNEWDataLogAndTempDBStoragePool() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.sqlVirtualMachines.createOrUpdate("testrg", "testvm", {
    location: "northeurope",
    storageConfigurationSettings: {
      diskConfigurationType: "NEW",
      sqlDataSettings: { defaultFilePath: "F:\\folderpath\\", luns: [0] },
      sqlLogSettings: { defaultFilePath: "G:\\folderpath\\", luns: [1] },
      sqlSystemDbOnDataDisk: true,
      sqlTempDbSettings: {
        dataFileCount: 8,
        dataFileSize: 256,
        dataGrowth: 512,
        defaultFilePath: "D:\\TEMP",
        logFileSize: 256,
        logGrowth: 512,
      },
      storageWorkloadType: "OLTP",
    },
    virtualMachineResourceId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm",
  });
  console.log(result);
}

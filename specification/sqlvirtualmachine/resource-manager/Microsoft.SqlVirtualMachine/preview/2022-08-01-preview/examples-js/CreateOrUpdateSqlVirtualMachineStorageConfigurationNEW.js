const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a SQL virtual machine.
 *
 * @summary Creates or updates a SQL virtual machine.
 * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/CreateOrUpdateSqlVirtualMachineStorageConfigurationNEW.json
 */
async function createsOrUpdatesASqlVirtualMachineForStorageConfigurationSettingsToNewDataLogAndTempDbStoragePool() {
  const subscriptionId =
    process.env["SQLVIRTUALMACHINE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQLVIRTUALMACHINE_RESOURCE_GROUP"] || "testrg";
  const sqlVirtualMachineName = "testvm";
  const parameters = {
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
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.sqlVirtualMachines.beginCreateOrUpdateAndWait(
    resourceGroupName,
    sqlVirtualMachineName,
    parameters
  );
  console.log(result);
}

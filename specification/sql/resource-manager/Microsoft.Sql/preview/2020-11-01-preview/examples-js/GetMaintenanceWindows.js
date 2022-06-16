const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets maintenance windows settings for a database.
 *
 * @summary Gets maintenance windows settings for a database.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetMaintenanceWindows.json
 */
async function getsMaintenanceWindowSettingsForASelectedDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const maintenanceWindowName = "current";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.maintenanceWindowsOperations.get(
    resourceGroupName,
    serverName,
    databaseName,
    maintenanceWindowName
  );
  console.log(result);
}

getsMaintenanceWindowSettingsForASelectedDatabase().catch(console.error);

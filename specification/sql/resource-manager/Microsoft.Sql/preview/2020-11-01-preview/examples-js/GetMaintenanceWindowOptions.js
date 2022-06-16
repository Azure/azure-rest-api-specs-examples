const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of available maintenance windows.
 *
 * @summary Gets a list of available maintenance windows.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetMaintenanceWindowOptions.json
 */
async function getsAListOfAvailableMaintenanceWindowsForASelectedDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const maintenanceWindowOptionsName = "current";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.maintenanceWindowOptionsOperations.get(
    resourceGroupName,
    serverName,
    databaseName,
    maintenanceWindowOptionsName
  );
  console.log(result);
}

getsAListOfAvailableMaintenanceWindowsForASelectedDatabase().catch(console.error);

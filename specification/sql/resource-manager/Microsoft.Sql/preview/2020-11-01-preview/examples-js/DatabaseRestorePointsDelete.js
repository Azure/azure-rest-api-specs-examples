const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a restore point.
 *
 * @summary Deletes a restore point.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRestorePointsDelete.json
 */
async function deletesARestorePoint() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const serverName = "testserver";
  const databaseName = "testDatabase";
  const restorePointName = "131546477590000000";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.restorePoints.delete(
    resourceGroupName,
    serverName,
    databaseName,
    restorePointName
  );
  console.log(result);
}

deletesARestorePoint().catch(console.error);

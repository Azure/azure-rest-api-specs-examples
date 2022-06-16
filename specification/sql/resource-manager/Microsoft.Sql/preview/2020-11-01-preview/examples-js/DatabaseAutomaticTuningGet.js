const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a database's automatic tuning.
 *
 * @summary Gets a database's automatic tuning.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseAutomaticTuningGet.json
 */
async function getADatabaseAutomaticTuningSettings() {
  const subscriptionId = "c3aa9078-0000-0000-0000-e36f151182d7";
  const resourceGroupName = "default-sql-onebox";
  const serverName = "testsvr11";
  const databaseName = "db1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseAutomaticTuningOperations.get(
    resourceGroupName,
    serverName,
    databaseName
  );
  console.log(result);
}

getADatabaseAutomaticTuningSettings().catch(console.error);

const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves server automatic tuning options.
 *
 * @summary Retrieves server automatic tuning options.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerAutomaticTuningGet.json
 */
async function getAServerAutomaticTuningSettings() {
  const subscriptionId = "c3aa9078-0000-0000-0000-e36f151182d7";
  const resourceGroupName = "default-sql-onebox";
  const serverName = "testsvr11";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverAutomaticTuningOperations.get(resourceGroupName, serverName);
  console.log(result);
}

getAServerAutomaticTuningSettings().catch(console.error);

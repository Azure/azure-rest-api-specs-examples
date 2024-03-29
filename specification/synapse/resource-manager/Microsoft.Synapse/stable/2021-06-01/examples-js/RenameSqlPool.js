const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Rename a SQL pool.
 *
 * @summary Rename a SQL pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RenameSqlPool.json
 */
async function renameASqlAnalyticsPool() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const workspaceName = "testsvr";
  const sqlPoolName = "testdb";
  const parameters = {
    id: "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Synapse/workspaces/testsvr/sqlPools/newtestdb",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPools.rename(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    parameters
  );
  console.log(result);
}

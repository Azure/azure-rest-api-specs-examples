const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the user activities of a SQL pool which includes running and suspended queries
 *
 * @summary Gets the user activities of a SQL pool which includes running and suspended queries
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolUserActivity.json
 */
async function getASqlAnalyticsPoolUserActivity() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const workspaceName = "testsvr";
  const sqlPoolName = "testdb";
  const dataWarehouseUserActivityName = "current";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolDataWarehouseUserActivities.get(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    dataWarehouseUserActivityName
  );
  console.log(result);
}

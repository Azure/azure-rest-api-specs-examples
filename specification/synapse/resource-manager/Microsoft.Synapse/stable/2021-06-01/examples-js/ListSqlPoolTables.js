const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets tables of a given schema in a SQL pool.
 *
 * @summary Gets tables of a given schema in a SQL pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolTables.json
 */
async function listTheTablesOfAGivenSchemaInASqlAnalyticsPool() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "myRG";
  const workspaceName = "serverName";
  const sqlPoolName = "myDatabase";
  const schemaName = "dbo";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlPoolTables.listBySchema(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    schemaName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

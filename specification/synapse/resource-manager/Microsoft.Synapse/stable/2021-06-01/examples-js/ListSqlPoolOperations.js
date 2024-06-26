const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of operations performed on the SQL pool.
 *
 * @summary Gets a list of operations performed on the SQL pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolOperations.json
 */
async function listTheSqlAnalyticsPoolManagementOperations() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "sqlcrudtest-7398";
  const workspaceName = "sqlcrudtest-4645";
  const sqlPoolName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlPoolOperations.list(
    resourceGroupName,
    workspaceName,
    sqlPoolName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

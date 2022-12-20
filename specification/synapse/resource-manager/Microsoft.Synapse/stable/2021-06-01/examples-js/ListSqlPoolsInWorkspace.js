const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all SQL pools
 *
 * @summary List all SQL pools
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolsInWorkspace.json
 */
async function listSqlAnalyticsPoolsInAWorkspace() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "sqlcrudtest-6845";
  const workspaceName = "sqlcrudtest-7177";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlPools.listByWorkspace(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

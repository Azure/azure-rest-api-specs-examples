const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of data connections of the given Kusto pool database.
 *
 * @summary Returns the list of data connections of the given Kusto pool database.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsListByDatabase.json
 */
async function kustoPoolDataConnectionsListByDatabase() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "kustorptest";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "KustoDatabase8";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.kustoPoolDataConnections.listByDatabase(
    resourceGroupName,
    workspaceName,
    kustoPoolName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

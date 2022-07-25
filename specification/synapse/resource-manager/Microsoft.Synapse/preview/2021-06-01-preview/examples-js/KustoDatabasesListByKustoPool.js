const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of databases of the given Kusto pool.
 *
 * @summary Returns the list of databases of the given Kusto pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoDatabasesListByKustoPool.json
 */
async function kustoDatabasesListByKustoPool() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const workspaceName = "synapseWorkspaceName";
  const kustoPoolName = "kustoclusterrptest4";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.kustoPoolDatabases.listByKustoPool(
    resourceGroupName,
    workspaceName,
    kustoPoolName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

kustoDatabasesListByKustoPool().catch(console.error);

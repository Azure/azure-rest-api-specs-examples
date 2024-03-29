const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of attached database configurations of the given Kusto Pool.
 *
 * @summary Returns the list of attached database configurations of the given Kusto Pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsListByKustoPool.json
 */
async function kustoPoolAttachedDatabaseConfigurationsListByKustoPool() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const workspaceName = "kustorptest";
  const kustoPoolName = "kustoclusterrptest4";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "kustorptest";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.kustoPoolAttachedDatabaseConfigurations.listByKustoPool(
    workspaceName,
    kustoPoolName,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the hybrid connection within the namespace.
 *
 * @summary Lists the hybrid connection within the namespace.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionListAll.json
 */
async function relayHybridConnectionListAll() {
  const subscriptionId =
    process.env["RELAY_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["RELAY_RESOURCE_GROUP"] || "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hybridConnections.listByNamespace(
    resourceGroupName,
    namespaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

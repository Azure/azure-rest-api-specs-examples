const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all private endpoints in a subscription.
 *
 * @summary Gets all private endpoints in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PrivateEndpointListAll.json
 */
async function listAllPrivateEndpoints() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpoints.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllPrivateEndpoints().catch(console.error);

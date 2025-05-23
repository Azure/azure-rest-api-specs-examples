const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets all private endpoints in a subscription.
 *
 * @summary Gets all private endpoints in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/PrivateEndpointListAll.json
 */
async function listAllPrivateEndpoints() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subId";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.privateEndpoints.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

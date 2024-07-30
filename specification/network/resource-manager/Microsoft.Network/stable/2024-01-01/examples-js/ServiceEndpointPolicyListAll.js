const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the service endpoint policies in a subscription.
 *
 * @summary Gets all the service endpoint policies in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ServiceEndpointPolicyListAll.json
 */
async function listAllServiceEndpointPolicy() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serviceEndpointPolicies.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

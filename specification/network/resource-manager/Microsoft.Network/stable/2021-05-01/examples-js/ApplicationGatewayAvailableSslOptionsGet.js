const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists available Ssl options for configuring Ssl policy.
 *
 * @summary Lists available Ssl options for configuring Ssl policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayAvailableSslOptionsGet.json
 */
async function getAvailableSslOptions() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationGateways.listAvailableSslOptions();
  console.log(result);
}

getAvailableSslOptions().catch(console.error);

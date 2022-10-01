const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all available response headers.
 *
 * @summary Lists all available response headers.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ApplicationGatewayAvailableResponseHeadersGet.json
 */
async function getAvailableResponseHeaders() {
  const subscriptionId = "72f988bf-86f1-41af-91ab-2d7cd0dddd4";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationGateways.listAvailableResponseHeaders();
  console.log(result);
}

getAvailableResponseHeaders().catch(console.error);

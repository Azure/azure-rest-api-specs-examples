const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List what values of endpoint services are available for use.
 *
 * @summary List what values of endpoint services are available for use.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/EndpointServicesList.json
 */
async function endpointServicesList() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availableEndpointServices.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

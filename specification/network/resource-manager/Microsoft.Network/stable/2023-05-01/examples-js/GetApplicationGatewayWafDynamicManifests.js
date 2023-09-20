const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the regional application gateway waf manifest.
 *
 * @summary Gets the regional application gateway waf manifest.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/GetApplicationGatewayWafDynamicManifests.json
 */
async function getsWafManifests() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationGatewayWafDynamicManifests.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

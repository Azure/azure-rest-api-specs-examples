const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the regional application gateway waf manifest.
 *
 * @summary Gets the regional application gateway waf manifest.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/GetApplicationGatewayWafDynamicManifestsDefault.json
 */
async function getsWafDefaultManifest() {
  const subscriptionId = "subid";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationGatewayWafDynamicManifestsDefault.get(location);
  console.log(result);
}

getsWafDefaultManifest().catch(console.error);

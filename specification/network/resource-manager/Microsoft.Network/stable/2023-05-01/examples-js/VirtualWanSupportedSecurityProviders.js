const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gives the supported security providers for the virtual wan.
 *
 * @summary Gives the supported security providers for the virtual wan.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/VirtualWanSupportedSecurityProviders.json
 */
async function supportedSecurityProviders() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualWANName = "wan1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.supportedSecurityProviders(resourceGroupName, virtualWANName);
  console.log(result);
}

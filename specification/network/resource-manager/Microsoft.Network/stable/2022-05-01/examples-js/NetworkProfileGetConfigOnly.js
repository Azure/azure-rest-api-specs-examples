const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified network profile in a specified resource group.
 *
 * @summary Gets the specified network profile in a specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkProfileGetConfigOnly.json
 */
async function getNetworkProfile() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkProfileName = "networkProfile1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkProfiles.get(resourceGroupName, networkProfileName);
  console.log(result);
}

getNetworkProfile().catch(console.error);

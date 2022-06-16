const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified public IP prefix in a specified resource group.
 *
 * @summary Gets the specified public IP prefix in a specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PublicIpPrefixGet.json
 */
async function getPublicIPPrefix() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const publicIpPrefixName = "test-ipprefix";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPPrefixes.get(resourceGroupName, publicIpPrefixName);
  console.log(result);
}

getPublicIPPrefix().catch(console.error);

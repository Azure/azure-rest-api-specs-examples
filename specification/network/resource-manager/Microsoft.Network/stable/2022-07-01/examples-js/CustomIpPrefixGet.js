const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified custom IP prefix in a specified resource group.
 *
 * @summary Gets the specified custom IP prefix in a specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/CustomIpPrefixGet.json
 */
async function getCustomIPPrefix() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const customIpPrefixName = "test-customipprefix";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.customIPPrefixes.get(resourceGroupName, customIpPrefixName);
  console.log(result);
}

getCustomIPPrefix().catch(console.error);

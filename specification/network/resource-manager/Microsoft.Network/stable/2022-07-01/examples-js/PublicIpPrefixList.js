const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all public IP prefixes in a resource group.
 *
 * @summary Gets all public IP prefixes in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/PublicIpPrefixList.json
 */
async function listResourceGroupPublicIPPrefixes() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.publicIPPrefixes.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listResourceGroupPublicIPPrefixes().catch(console.error);

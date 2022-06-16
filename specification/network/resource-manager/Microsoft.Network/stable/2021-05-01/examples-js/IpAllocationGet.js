const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified IpAllocation by resource group.
 *
 * @summary Gets the specified IpAllocation by resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/IpAllocationGet.json
 */
async function getIPAllocation() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ipAllocationName = "test-ipallocation";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ipAllocations.get(resourceGroupName, ipAllocationName);
  console.log(result);
}

getIPAllocation().catch(console.error);

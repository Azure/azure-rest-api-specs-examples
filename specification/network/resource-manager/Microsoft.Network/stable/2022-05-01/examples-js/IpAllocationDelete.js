const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified IpAllocation.
 *
 * @summary Deletes the specified IpAllocation.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/IpAllocationDelete.json
 */
async function deleteIPAllocation() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const ipAllocationName = "test-ipallocation";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.ipAllocations.beginDeleteAndWait(resourceGroupName, ipAllocationName);
  console.log(result);
}

deleteIPAllocation().catch(console.error);

const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a virtual network in the specified resource group.
 *
 * @summary creates or updates a virtual network in the specified resource group.
 * x-ms-original-file: 2025-07-01/VirtualNetworkCreateWithIpamPool.json
 */
async function createVirtualNetworkWithIpamPool() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworks.createOrUpdate("rg1", "test-vnet", {
    location: "eastus",
    addressSpace: {
      ipamPoolPrefixAllocations: [
        {
          numberOfIpAddresses: "65536",
          id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/nm1/ipamPools/testIpamPool",
        },
      ],
    },
    subnets: [
      {
        ipamPoolPrefixAllocations: [
          {
            numberOfIpAddresses: "80",
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/nm1/ipamPools/testIpamPool",
          },
        ],
      },
    ],
  });
  console.log(result);
}

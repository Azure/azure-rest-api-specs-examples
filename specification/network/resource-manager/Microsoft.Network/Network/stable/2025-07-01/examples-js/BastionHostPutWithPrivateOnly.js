const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified Bastion Host.
 *
 * @summary creates or updates the specified Bastion Host.
 * x-ms-original-file: 2025-07-01/BastionHostPutWithPrivateOnly.json
 */
async function createBastionHostWithPrivateOnly() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.bastionHosts.createOrUpdate("rg1", "bastionhosttenant", {
    enablePrivateOnlyBastion: true,
    ipConfigurations: [
      {
        name: "bastionHostIpConfiguration",
        subnet: {
          id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet",
        },
      },
    ],
  });
  console.log(result);
}

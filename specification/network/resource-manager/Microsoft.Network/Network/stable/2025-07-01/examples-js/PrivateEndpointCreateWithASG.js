const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an private endpoint in the specified resource group.
 *
 * @summary creates or updates an private endpoint in the specified resource group.
 * x-ms-original-file: 2025-07-01/PrivateEndpointCreateWithASG.json
 */
async function createPrivateEndpointWithApplicationSecurityGroups() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateEndpoints.createOrUpdate("rg1", "testPe", {
    location: "eastus2euap",
    applicationSecurityGroups: [
      {
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/provders/Microsoft.Network/applicationSecurityGroup/asg1",
      },
    ],
    privateLinkServiceConnections: [
      {
        groupIds: ["groupIdFromResource"],
        privateLinkServiceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls",
        requestMessage: "Please approve my connection.",
      },
    ],
    subnet: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet",
    },
  });
  console.log(result);
}

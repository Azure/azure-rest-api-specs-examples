const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an private endpoint in the specified resource group.
 *
 * @summary Creates or updates an private endpoint in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/PrivateEndpointCreateForManualApproval.json
 */
async function createPrivateEndpointWithManualApprovalConnection() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const privateEndpointName = "testPe";
  const parameters = {
    customNetworkInterfaceName: "testPeNic",
    ipConfigurations: [
      {
        name: "pestaticconfig",
        groupId: "file",
        memberName: "file",
        privateIPAddress: "192.168.0.5",
      },
    ],
    location: "eastus",
    manualPrivateLinkServiceConnections: [
      {
        groupIds: ["groupIdFromResource"],
        privateLinkServiceId:
          "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls",
        requestMessage: "Please manually approve my connection.",
      },
    ],
    subnet: {
      id: "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateEndpoints.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateEndpointName,
    parameters
  );
  console.log(result);
}

createPrivateEndpointWithManualApprovalConnection().catch(console.error);

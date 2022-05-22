Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an private endpoint in the specified resource group.
 *
 * @summary Creates or updates an private endpoint in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PrivateEndpointCreate.json
 */
async function createPrivateEndpoint() {
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
        privateIPAddress: "192.168.0.6",
      },
    ],
    location: "eastus2euap",
    privateLinkServiceConnections: [
      {
        groupIds: ["groupIdFromResource"],
        privateLinkServiceId:
          "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls",
        requestMessage: "Please approve my connection.",
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

createPrivateEndpoint().catch(console.error);
```

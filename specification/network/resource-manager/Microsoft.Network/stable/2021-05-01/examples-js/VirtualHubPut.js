const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a VirtualHub resource if it doesn't exist else updates the existing VirtualHub.
 *
 * @summary Creates a VirtualHub resource if it doesn't exist else updates the existing VirtualHub.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualHubPut.json
 */
async function virtualHubPut() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub2";
  const virtualHubParameters = {
    addressPrefix: "10.168.0.0/24",
    location: "West US",
    sku: "Basic",
    tags: { key1: "value1" },
    virtualWan: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubs.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualHubName,
    virtualHubParameters
  );
  console.log(result);
}

virtualHubPut().catch(console.error);

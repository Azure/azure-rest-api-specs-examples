const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a VirtualWAN.
 *
 * @summary Deletes a VirtualWAN.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualWANDelete.json
 */
async function virtualWanDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualWANName = "virtualWan1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualWans.beginDeleteAndWait(resourceGroupName, virtualWANName);
  console.log(result);
}

virtualWanDelete().catch(console.error);

const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete virtual network. This operation can take a while to complete.
 *
 * @summary Delete virtual network. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_Delete.json
 */
async function virtualNetworksDelete() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{virtualNetworkName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualNetworks.beginDeleteAndWait(resourceGroupName, labName, name);
  console.log(result);
}

virtualNetworksDelete().catch(console.error);

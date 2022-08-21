const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to update certain properties of the virtual network resource.
 *
 * @summary API to update certain properties of the virtual network resource.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/UpdateVirtualNetwork.json
 */
async function updateVirtualNetwork() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "testrg";
  const virtualNetworkName = "ProdNetwork";
  const body = { tags: { tag1: "value1", tag2: "value2" } };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.virtualNetworks.update(
    resourceGroupName,
    virtualNetworkName,
    options
  );
  console.log(result);
}

updateVirtualNetwork().catch(console.error);

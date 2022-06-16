const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Network Virtual Appliance.
 *
 * @summary Updates a Network Virtual Appliance.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkVirtualApplianceUpdateTags.json
 */
async function updateNetworkVirtualAppliance() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkVirtualApplianceName = "nva";
  const parameters = { tags: { key1: "value1", key2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkVirtualAppliances.updateTags(
    resourceGroupName,
    networkVirtualApplianceName,
    parameters
  );
  console.log(result);
}

updateNetworkVirtualAppliance().catch(console.error);

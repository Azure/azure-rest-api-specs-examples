const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Network Virtual Appliance.
 *
 * @summary Gets the specified Network Virtual Appliance.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkVirtualApplianceGet.json
 */
async function getNetworkVirtualAppliance() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkVirtualApplianceName = "nva";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkVirtualAppliances.get(
    resourceGroupName,
    networkVirtualApplianceName
  );
  console.log(result);
}

getNetworkVirtualAppliance().catch(console.error);

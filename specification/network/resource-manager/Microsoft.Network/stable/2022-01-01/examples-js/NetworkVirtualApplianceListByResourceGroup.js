const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Network Virtual Appliances in a resource group.
 *
 * @summary Lists all Network Virtual Appliances in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkVirtualApplianceListByResourceGroup.json
 */
async function listAllNetworkVirtualApplianceForAGivenResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkVirtualAppliances.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllNetworkVirtualApplianceForAGivenResourceGroup().catch(console.error);

const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all network security groups in a resource group.
 *
 * @summary Gets all network security groups in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkSecurityGroupList.json
 */
async function listNetworkSecurityGroupsInResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkSecurityGroups.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listNetworkSecurityGroupsInResourceGroup().catch(console.error);

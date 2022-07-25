const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of privateLinkHubs in a resource group
 *
 * @summary Returns a list of privateLinkHubs in a resource group
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListPrivateLinkHubsInResourceGroup.json
 */
async function listPrivateLinkHubsInResourceGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup1";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkHubs.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPrivateLinkHubsInResourceGroup().catch(console.error);

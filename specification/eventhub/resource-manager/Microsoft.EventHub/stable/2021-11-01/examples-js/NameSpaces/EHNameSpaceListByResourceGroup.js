const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the available Namespaces within a resource group.
 *
 * @summary Lists the available Namespaces within a resource group.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceListByResourceGroup.json
 */
async function namespaceListByResourceGroup() {
  const subscriptionId = "SampleSubscription";
  const resourceGroupName = "ResurceGroupSample";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaces.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

namespaceListByResourceGroup().catch(console.error);

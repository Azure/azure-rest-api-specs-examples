const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the namespaces under a resource group.
 *
 * @summary List all the namespaces under a resource group.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Namespaces_ListByResourceGroup.json
 */
async function namespacesListByResourceGroup() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaces.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

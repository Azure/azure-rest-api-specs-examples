const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List event types for a topic.
 *
 * @summary List event types for a topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/Topics_ListEventTypes.json
 */
async function topicsListEventTypes() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const providerNamespace = "Microsoft.Storage";
  const resourceTypeName = "storageAccounts";
  const resourceName = "ExampleStorageAccount";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.topics.listEventTypes(
    resourceGroupName,
    providerNamespace,
    resourceTypeName,
    resourceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List event types for a topic type.
 *
 * @summary List event types for a topic type.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/TopicTypes_ListEventTypes.json
 */
async function topicTypesListEventTypes() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const topicTypeName = "Microsoft.Storage.StorageAccounts";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.topicTypes.listEventTypes(topicTypeName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

topicTypesListEventTypes().catch(console.error);

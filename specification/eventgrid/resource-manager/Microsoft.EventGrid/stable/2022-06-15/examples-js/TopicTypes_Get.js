const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get information about a topic type.
 *
 * @summary Get information about a topic type.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/TopicTypes_Get.json
 */
async function topicTypesGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const topicTypeName = "Microsoft.Storage.StorageAccounts";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.topicTypes.get(topicTypeName);
  console.log(result);
}

topicTypesGet().catch(console.error);

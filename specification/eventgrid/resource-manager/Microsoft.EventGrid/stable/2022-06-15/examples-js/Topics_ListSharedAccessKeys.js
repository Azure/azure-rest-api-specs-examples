const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the two keys used to publish to a topic.
 *
 * @summary List the two keys used to publish to a topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Topics_ListSharedAccessKeys.json
 */
async function topicsListSharedAccessKeys() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const topicName = "exampletopic2";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.topics.listSharedAccessKeys(resourceGroupName, topicName);
  console.log(result);
}

topicsListSharedAccessKeys().catch(console.error);

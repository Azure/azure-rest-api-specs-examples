const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete existing system topic.
 *
 * @summary Delete existing system topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/SystemTopics_Delete.json
 */
async function systemTopicsDelete() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const systemTopicName = "exampleSystemTopic1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.systemTopics.beginDeleteAndWait(resourceGroupName, systemTopicName);
  console.log(result);
}

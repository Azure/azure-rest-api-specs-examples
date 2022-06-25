const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously updates a topic with the specified parameters.
 *
 * @summary Asynchronously updates a topic with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Topics_Update.json
 */
async function topicsUpdate() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const topicName = "exampletopic1";
  const topicUpdateParameters = {
    inboundIpRules: [
      { action: "Allow", ipMask: "12.18.30.15" },
      { action: "Allow", ipMask: "12.18.176.1" },
    ],
    publicNetworkAccess: "Enabled",
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.topics.beginUpdateAndWait(
    resourceGroupName,
    topicName,
    topicUpdateParameters
  );
  console.log(result);
}

topicsUpdate().catch(console.error);

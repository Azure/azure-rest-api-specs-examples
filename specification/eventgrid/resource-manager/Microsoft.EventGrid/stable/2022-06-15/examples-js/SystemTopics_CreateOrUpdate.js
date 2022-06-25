const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new system topic with the specified parameters.
 *
 * @summary Asynchronously creates a new system topic with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopics_CreateOrUpdate.json
 */
async function systemTopicsCreateOrUpdate() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const systemTopicName = "exampleSystemTopic1";
  const systemTopicInfo = {
    location: "westus2",
    source:
      "/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/microsoft.storage/storageaccounts/pubstgrunnerb71cd29e",
    tags: { tag1: "value1", tag2: "value2" },
    topicType: "microsoft.storage.storageaccounts",
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.systemTopics.beginCreateOrUpdateAndWait(
    resourceGroupName,
    systemTopicName,
    systemTopicInfo
  );
  console.log(result);
}

systemTopicsCreateOrUpdate().catch(console.error);

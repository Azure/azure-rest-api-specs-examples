const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new topic with the specified parameters.
 *
 * @summary Asynchronously creates a new topic with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Topics_CreateOrUpdateForAzureArc.json
 */
async function topicsCreateOrUpdateForAzureArc() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const topicName = "exampletopic1";
  const topicInfo = {
    extendedLocation: {
      name: "/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourcegroups/examplerg/providers/Microsoft.ExtendedLocation/CustomLocations/exampleCustomLocation",
      type: "CustomLocation",
    },
    inputSchema: "CloudEventSchemaV1_0",
    kind: "AzureArc",
    location: "westus2",
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.topics.beginCreateOrUpdateAndWait(
    resourceGroupName,
    topicName,
    topicInfo
  );
  console.log(result);
}

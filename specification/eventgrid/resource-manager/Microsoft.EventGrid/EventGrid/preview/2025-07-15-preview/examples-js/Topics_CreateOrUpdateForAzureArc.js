const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to asynchronously creates a new topic with the specified parameters.
 *
 * @summary asynchronously creates a new topic with the specified parameters.
 * x-ms-original-file: 2025-07-15-preview/Topics_CreateOrUpdateForAzureArc.json
 */
async function topicsCreateOrUpdateForAzureArc() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.topics.createOrUpdate("examplerg", "exampletopic1", {
    extendedLocation: {
      name: "/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourcegroups/examplerg/providers/Microsoft.ExtendedLocation/CustomLocations/exampleCustomLocation",
      type: "CustomLocation",
    },
    kind: "AzureArc",
    location: "westus2",
    inputSchema: "CloudEventSchemaV1_0",
    tags: { tag1: "value1", tag2: "value2" },
  });
  console.log(result);
}

const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 *
 * @summary asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 * x-ms-original-file: 2025-07-15-preview/EventSubscriptions_CreateOrUpdateForCustomTopic_StorageQueueDestination.json
 */
async function eventSubscriptionsCreateOrUpdateForCustomTopicStorageQueueDestination() {
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.createOrUpdate(
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
    "examplesubscription1",
    {
      deadLetterDestination: {
        endpointType: "StorageBlob",
        blobContainerName: "contosocontainer",
        resourceId:
          "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg",
      },
      destination: {
        endpointType: "StorageQueue",
        queueMessageTimeToLiveInSeconds: 300,
        queueName: "queue1",
        resourceId:
          "/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg",
      },
      filter: {
        isSubjectCaseSensitive: false,
        subjectBeginsWith: "ExamplePrefix",
        subjectEndsWith: "ExampleSuffix",
      },
    },
  );
  console.log(result);
}

const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 *
 * @summary Asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_CreateOrUpdateForCustomTopic_HybridConnectionDestination.json
 */
async function eventSubscriptionsCreateOrUpdateForCustomTopicHybridConnectionDestination() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope =
    "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1";
  const eventSubscriptionName = "examplesubscription1";
  const eventSubscriptionInfo = {
    deadLetterDestination: {
      blobContainerName: "contosocontainer",
      endpointType: "StorageBlob",
      resourceId:
        "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg",
    },
    destination: {
      endpointType: "HybridConnection",
      resourceId:
        "/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Relay/namespaces/ContosoNamespace/hybridConnections/HC1",
    },
    filter: {
      isSubjectCaseSensitive: false,
      subjectBeginsWith: "ExamplePrefix",
      subjectEndsWith: "ExampleSuffix",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.eventSubscriptions.beginCreateOrUpdateAndWait(
    scope,
    eventSubscriptionName,
    eventSubscriptionInfo
  );
  console.log(result);
}

eventSubscriptionsCreateOrUpdateForCustomTopicHybridConnectionDestination().catch(console.error);

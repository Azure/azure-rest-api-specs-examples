const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Asynchronously updates an existing event subscription.
 *
 * @summary Asynchronously updates an existing event subscription.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/EventSubscriptions_UpdateForCustomTopic_AzureFunctionDestination.json
 */
async function eventSubscriptionsUpdateForCustomTopicAzureFunctionDestination() {
  const scope =
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1";
  const eventSubscriptionName = "examplesubscription1";
  const eventSubscriptionUpdateParameters = {
    deadLetterDestination: {
      blobContainerName: "contosocontainer",
      endpointType: "StorageBlob",
      resourceId:
        "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg",
    },
    destination: {
      endpointType: "AzureFunction",
      resourceId:
        "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Web/sites/ContosoSite/funtions/ContosoFunc",
    },
    filter: {
      isSubjectCaseSensitive: false,
      subjectBeginsWith: "ExamplePrefix",
      subjectEndsWith: "ExampleSuffix",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.beginUpdateAndWait(
    scope,
    eventSubscriptionName,
    eventSubscriptionUpdateParameters,
  );
  console.log(result);
}

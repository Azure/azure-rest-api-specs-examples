const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 *
 * @summary asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 * x-ms-original-file: 2025-07-15-preview/EventSubscriptions_CreateOrUpdateForResourceGroup.json
 */
async function eventSubscriptionsCreateOrUpdateForResourceGroup() {
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.createOrUpdate(
    "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg",
    "examplesubscription2",
    {
      destination: { endpointType: "WebHook", endpointUrl: "https://requestb.in/15ksip71" },
      filter: {
        isSubjectCaseSensitive: false,
        subjectBeginsWith: "ExamplePrefix",
        subjectEndsWith: "ExampleSuffix",
      },
    },
  );
  console.log(result);
}

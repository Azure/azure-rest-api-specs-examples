const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 *
 * @summary Asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/EventSubscriptions_CreateOrUpdateForSubscription.json
 */
async function eventSubscriptionsCreateOrUpdateForSubscription() {
  const scope = "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const eventSubscriptionName = "examplesubscription3";
  const eventSubscriptionInfo = {
    destination: {
      endpointType: "WebHook",
      endpointUrl: "https://requestb.in/15ksip71",
    },
    filter: { isSubjectCaseSensitive: false },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.eventSubscriptions.beginCreateOrUpdateAndWait(
    scope,
    eventSubscriptionName,
    eventSubscriptionInfo,
  );
  console.log(result);
}

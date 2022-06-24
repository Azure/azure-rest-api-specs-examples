const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 *
 * @summary Asynchronously creates a new event subscription or updates an existing event subscription based on the specified scope.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_CreateOrUpdateForSubscription.json
 */
async function eventSubscriptionsCreateOrUpdateForSubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const eventSubscriptionName = "examplesubscription3";
  const eventSubscriptionInfo = {
    destination: {
      endpointType: "WebHook",
      endpointUrl: "https://requestb.in/15ksip71",
    },
    filter: { isSubjectCaseSensitive: false },
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

eventSubscriptionsCreateOrUpdateForSubscription().catch(console.error);

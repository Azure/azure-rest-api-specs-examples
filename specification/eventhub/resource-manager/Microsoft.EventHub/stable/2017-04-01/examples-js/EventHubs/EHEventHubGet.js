const { EventHubManagementClient } = require("@azure/arm-eventhub-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an Event Hubs description for the specified Event Hub.
 *
 * @summary Gets an Event Hubs description for the specified Event Hub.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2017-04-01/examples/EventHubs/EHEventHubGet.json
 */
async function eventHubGet() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "e2f361f0-3b27-4503-a9cc-21cfba380093";
  const resourceGroupName =
    process.env["EVENTHUB_RESOURCE_GROUP"] || "Default-NotificationHubs-AustraliaEast";
  const namespaceName = "sdk-Namespace-716";
  const eventHubName = "sdk-EventHub-10";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.eventHubs.get(resourceGroupName, namespaceName, eventHubName);
  console.log(result);
}

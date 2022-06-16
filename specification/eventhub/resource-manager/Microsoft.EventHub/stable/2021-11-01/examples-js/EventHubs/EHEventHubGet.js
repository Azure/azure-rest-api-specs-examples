const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an Event Hubs description for the specified Event Hub.
 *
 * @summary Gets an Event Hubs description for the specified Event Hub.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubGet.json
 */
async function eventHubGet() {
  const subscriptionId = "e2f361f0-3b27-4503-a9cc-21cfba380093";
  const resourceGroupName = "Default-NotificationHubs-AustraliaEast";
  const namespaceName = "sdk-Namespace-716";
  const eventHubName = "sdk-EventHub-10";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.eventHubs.get(resourceGroupName, namespaceName, eventHubName);
  console.log(result);
}

eventHubGet().catch(console.error);

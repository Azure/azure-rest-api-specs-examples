const { NotificationHubsManagementClient } = require("@azure/arm-notificationhubs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch a NotificationHub in a namespace.
 *
 * @summary Patch a NotificationHub in a namespace.
 * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubPatch.json
 */
async function notificationHubPatch() {
  const subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
  const resourceGroupName = "sdkresourceGroup";
  const namespaceName = "nh-sdk-ns";
  const notificationHubName = "sdk-notificationHubs-8708";
  const parameters = {};
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new NotificationHubsManagementClient(credential, subscriptionId);
  const result = await client.notificationHubs.patch(
    resourceGroupName,
    namespaceName,
    notificationHubName,
    options
  );
  console.log(result);
}

notificationHubPatch().catch(console.error);

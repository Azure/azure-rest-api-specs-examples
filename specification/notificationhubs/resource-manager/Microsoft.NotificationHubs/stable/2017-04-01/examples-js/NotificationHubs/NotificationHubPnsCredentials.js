const { NotificationHubsManagementClient } = require("@azure/arm-notificationhubs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the PNS Credentials associated with a notification hub .
 *
 * @summary Lists the PNS Credentials associated with a notification hub .
 * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubPnsCredentials.json
 */
async function notificationHubPnsCredentials() {
  const subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
  const resourceGroupName = "5ktrial";
  const namespaceName = "nh-sdk-ns";
  const notificationHubName = "nh-sdk-hub";
  const credential = new DefaultAzureCredential();
  const client = new NotificationHubsManagementClient(credential, subscriptionId);
  const result = await client.notificationHubs.getPnsCredentials(
    resourceGroupName,
    namespaceName,
    notificationHubName
  );
  console.log(result);
}

notificationHubPnsCredentials().catch(console.error);

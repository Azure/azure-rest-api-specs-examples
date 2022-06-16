const { NotificationHubsManagementClient } = require("@azure/arm-notificationhubs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the notification hubs associated with a namespace.
 *
 * @summary Lists the notification hubs associated with a namespace.
 * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubListByNameSpace.json
 */
async function notificationHubListByNameSpace() {
  const subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
  const resourceGroupName = "5ktrial";
  const namespaceName = "nh-sdk-ns";
  const credential = new DefaultAzureCredential();
  const client = new NotificationHubsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.notificationHubs.list(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

notificationHubListByNameSpace().catch(console.error);

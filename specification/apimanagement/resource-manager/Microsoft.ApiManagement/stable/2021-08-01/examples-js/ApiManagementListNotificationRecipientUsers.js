const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of the Notification Recipient User subscribed to the notification.
 *
 * @summary Gets the list of the Notification Recipient User subscribed to the notification.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListNotificationRecipientUsers.json
 */
async function apiManagementListNotificationRecipientUsers() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const notificationName = "RequestPublisherNotificationMessage";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.notificationRecipientUser.listByNotification(
    resourceGroupName,
    serviceName,
    notificationName
  );
  console.log(result);
}

apiManagementListNotificationRecipientUsers().catch(console.error);

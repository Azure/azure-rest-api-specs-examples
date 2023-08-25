const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of the Notification Recipient Emails subscribed to a notification.
 *
 * @summary Gets the list of the Notification Recipient Emails subscribed to a notification.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListNotificationRecipientEmails.json
 */
async function apiManagementListNotificationRecipientEmails() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const notificationName = "RequestPublisherNotificationMessage";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.notificationRecipientEmail.listByNotification(
    resourceGroupName,
    serviceName,
    notificationName
  );
  console.log(result);
}

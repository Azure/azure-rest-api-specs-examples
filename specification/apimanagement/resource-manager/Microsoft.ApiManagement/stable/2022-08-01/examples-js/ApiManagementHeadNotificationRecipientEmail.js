const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Determine if Notification Recipient Email subscribed to the notification.
 *
 * @summary Determine if Notification Recipient Email subscribed to the notification.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadNotificationRecipientEmail.json
 */
async function apiManagementHeadNotificationRecipientEmail() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const notificationName = "RequestPublisherNotificationMessage";
  const email = "contoso@live.com";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.notificationRecipientEmail.checkEntityExists(
    resourceGroupName,
    serviceName,
    notificationName,
    email
  );
  console.log(result);
}

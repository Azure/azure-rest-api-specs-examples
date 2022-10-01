const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Removes the email from the list of Notification.
 *
 * @summary Removes the email from the list of Notification.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteNotificationRecipientEmail.json
 */
async function apiManagementDeleteNotificationRecipientEmail() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const notificationName = "RequestPublisherNotificationMessage";
  const email = "contoso@live.com";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.notificationRecipientEmail.delete(
    resourceGroupName,
    serviceName,
    notificationName,
    email
  );
  console.log(result);
}

apiManagementDeleteNotificationRecipientEmail().catch(console.error);

const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Adds the API Management User to the list of Recipients for the Notification.
 *
 * @summary Adds the API Management User to the list of Recipients for the Notification.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateNotificationRecipientUser.json
 */
async function apiManagementCreateNotificationRecipientUser() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const notificationName = "RequestPublisherNotificationMessage";
  const userId = "576823d0a40f7e74ec07d642";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.notificationRecipientUser.createOrUpdate(
    resourceGroupName,
    serviceName,
    notificationName,
    userId
  );
  console.log(result);
}

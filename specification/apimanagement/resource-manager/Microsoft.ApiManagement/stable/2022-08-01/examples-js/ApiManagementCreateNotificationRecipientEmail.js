const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Adds the Email address to the list of Recipients for the Notification.
 *
 * @summary Adds the Email address to the list of Recipients for the Notification.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateNotificationRecipientEmail.json
 */
async function apiManagementCreateNotificationRecipientEmail() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const notificationName = "RequestPublisherNotificationMessage";
  const email = "foobar@live.com";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.notificationRecipientEmail.createOrUpdate(
    resourceGroupName,
    serviceName,
    notificationName,
    email
  );
  console.log(result);
}

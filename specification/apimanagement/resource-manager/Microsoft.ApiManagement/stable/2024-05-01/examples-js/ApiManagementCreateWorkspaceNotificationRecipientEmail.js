const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Adds the Email address to the list of Recipients for the Notification.
 *
 * @summary Adds the Email address to the list of Recipients for the Notification.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceNotificationRecipientEmail.json
 */
async function apiManagementCreateWorkspaceNotificationRecipientEmail() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const workspaceId = "wks1";
  const notificationName = "RequestPublisherNotificationMessage";
  const email = "foobar@live.com";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceNotificationRecipientEmail.createOrUpdate(
    resourceGroupName,
    serviceName,
    workspaceId,
    notificationName,
    email,
  );
  console.log(result);
}

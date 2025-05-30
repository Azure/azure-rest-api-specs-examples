const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the specified Subscription keys.
 *
 * @summary Gets the specified Subscription keys.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementWorkspaceSubscriptionListSecrets.json
 */
async function apiManagementWorkspaceSubscriptionListSecrets() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const workspaceId = "wks1";
  const sid = "5931a769d8d14f0ad8ce13b8";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceSubscription.listSecrets(
    resourceGroupName,
    serviceName,
    workspaceId,
    sid,
  );
  console.log(result);
}

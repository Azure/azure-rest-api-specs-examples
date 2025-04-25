const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Regenerates secondary key of existing subscription of the API Management service instance.
 *
 * @summary Regenerates secondary key of existing subscription of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementSubscriptionRegenerateSecondaryKey.json
 */
async function apiManagementSubscriptionRegenerateSecondaryKey() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const sid = "testsub";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.subscription.regenerateSecondaryKey(
    resourceGroupName,
    serviceName,
    sid,
  );
  console.log(result);
}

const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves a redirection URL containing an authentication token for signing a given user into the developer portal.
 *
 * @summary Retrieves a redirection URL containing an authentication token for signing a given user into the developer portal.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUserGenerateSsoUrl.json
 */
async function apiManagementUserGenerateSsoUrl() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const userId = "57127d485157a511ace86ae7";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.user.generateSsoUrl(resourceGroupName, serviceName, userId);
  console.log(result);
}

const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Single-Sign-On token for the API Management Service which is valid for 5 Minutes.
 *
 * @summary Gets the Single-Sign-On token for the API Management Service which is valid for 5 Minutes.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetSsoToken.json
 */
async function apiManagementServiceGetSsoToken() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.getSsoToken(resourceGroupName, serviceName);
  console.log(result);
}

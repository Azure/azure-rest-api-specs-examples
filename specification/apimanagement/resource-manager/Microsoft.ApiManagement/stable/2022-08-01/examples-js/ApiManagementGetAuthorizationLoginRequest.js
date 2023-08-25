const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets authorization login links.
 *
 * @summary Gets authorization login links.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetAuthorizationLoginRequest.json
 */
async function apiManagementGetAuthorizationLoginRequest() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const authorizationProviderId = "aadwithauthcode";
  const authorizationId = "authz1";
  const parameters = {
    postLoginRedirectUrl: "https://www.bing.com/",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorizationLoginLinks.post(
    resourceGroupName,
    serviceName,
    authorizationProviderId,
    authorizationId,
    parameters
  );
  console.log(result);
}

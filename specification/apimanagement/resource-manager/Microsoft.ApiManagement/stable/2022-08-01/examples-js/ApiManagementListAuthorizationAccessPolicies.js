const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a collection of authorization access policy defined within a authorization.
 *
 * @summary Lists a collection of authorization access policy defined within a authorization.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListAuthorizationAccessPolicies.json
 */
async function apiManagementListAuthorizationAccessPolicies() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const authorizationProviderId = "aadwithauthcode";
  const authorizationId = "authz1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.authorizationAccessPolicy.listByAuthorization(
    resourceGroupName,
    serviceName,
    authorizationProviderId,
    authorizationId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

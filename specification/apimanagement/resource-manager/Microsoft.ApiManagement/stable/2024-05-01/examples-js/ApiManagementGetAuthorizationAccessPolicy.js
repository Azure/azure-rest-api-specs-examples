const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the details of the authorization access policy specified by its identifier.
 *
 * @summary Gets the details of the authorization access policy specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetAuthorizationAccessPolicy.json
 */
async function apiManagementGetAuthorizationAccessPolicy() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const authorizationProviderId = "aadwithauthcode";
  const authorizationId = "authz1";
  const authorizationAccessPolicyId = "fe0bed83-631f-4149-bd0b-0464b1bc7cab";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorizationAccessPolicy.get(
    resourceGroupName,
    serviceName,
    authorizationProviderId,
    authorizationId,
    authorizationAccessPolicyId,
  );
  console.log(result);
}

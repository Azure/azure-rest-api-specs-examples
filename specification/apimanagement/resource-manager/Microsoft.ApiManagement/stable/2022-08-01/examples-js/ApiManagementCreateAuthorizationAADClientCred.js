const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates authorization.
 *
 * @summary Creates or updates authorization.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationAADClientCred.json
 */
async function apiManagementCreateAuthorizationAadClientCred() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const authorizationProviderId = "aadwithclientcred";
  const authorizationId = "authz1";
  const parameters = {
    authorizationType: "OAuth2",
    oAuth2GrantType: "AuthorizationCode",
    parameters: {
      clientId: "",
      clientSecret: "",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorization.createOrUpdate(
    resourceGroupName,
    serviceName,
    authorizationProviderId,
    authorizationId,
    parameters
  );
  console.log(result);
}

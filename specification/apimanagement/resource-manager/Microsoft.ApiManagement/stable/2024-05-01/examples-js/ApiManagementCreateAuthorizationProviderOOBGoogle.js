const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates authorization provider.
 *
 * @summary Creates or updates authorization provider.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationProviderOOBGoogle.json
 */
async function apiManagementCreateAuthorizationProviderOobGoogle() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const authorizationProviderId = "google";
  const parameters = {
    displayName: "google",
    identityProvider: "google",
    oauth2: {
      grantTypes: {
        authorizationCode: {
          clientId: "99999999-xxxxxxxxxxxxxxxxxxxxxxxx.apps.googleusercontent.com",
          clientSecret: "clientsecretvalue",
          scopes:
            "openid https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email",
        },
      },
      redirectUrl:
        "https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorizationProvider.createOrUpdate(
    resourceGroupName,
    serviceName,
    authorizationProviderId,
    parameters,
  );
  console.log(result);
}

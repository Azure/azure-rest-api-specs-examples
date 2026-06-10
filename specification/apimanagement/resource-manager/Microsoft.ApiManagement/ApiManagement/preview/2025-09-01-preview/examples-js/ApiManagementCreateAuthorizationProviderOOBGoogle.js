const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates authorization provider.
 *
 * @summary creates or updates authorization provider.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateAuthorizationProviderOOBGoogle.json
 */
async function apiManagementCreateAuthorizationProviderOOBGoogle() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorizationProvider.createOrUpdate(
    "rg1",
    "apimService1",
    "google",
    {
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
    },
  );
  console.log(result);
}

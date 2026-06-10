const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates authorization provider.
 *
 * @summary creates or updates authorization provider.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateAuthorizationProviderAADClientCred.json
 */
async function apiManagementCreateAuthorizationProviderAADClientCred() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorizationProvider.createOrUpdate(
    "rg1",
    "apimService1",
    "aadwithclientcred",
    {
      displayName: "aadwithclientcred",
      identityProvider: "aad",
      oauth2: {
        grantTypes: {
          authorizationCode: {
            resourceUri: "https://graph.microsoft.com",
            scopes: "User.Read.All Group.Read.All",
          },
        },
        redirectUrl:
          "https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1",
      },
    },
  );
  console.log(result);
}

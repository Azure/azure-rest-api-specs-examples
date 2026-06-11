const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates authorization provider.
 *
 * @summary creates or updates authorization provider.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateAuthorizationProviderGenericOAuth2.json
 */
async function apiManagementCreateAuthorizationProviderGenericOAuth2() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorizationProvider.createOrUpdate(
    "rg1",
    "apimService1",
    "eventbrite",
    {
      displayName: "eventbrite",
      identityProvider: "oauth2",
      oauth2: {
        grantTypes: {
          authorizationCode: {
            authorizationUrl: "https://www.eventbrite.com/oauth/authorize",
            clientId: "clientid",
            clientSecret: "clientsecretvalue",
            refreshUrl: "https://www.eventbrite.com/oauth/token",
            tokenUrl: "https://www.eventbrite.com/oauth/token",
          },
        },
        redirectUrl:
          "https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1",
      },
    },
  );
  console.log(result);
}

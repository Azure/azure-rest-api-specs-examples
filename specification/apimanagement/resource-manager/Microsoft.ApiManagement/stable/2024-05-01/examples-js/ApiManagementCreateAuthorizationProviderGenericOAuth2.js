const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates authorization provider.
 *
 * @summary Creates or updates authorization provider.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationProviderGenericOAuth2.json
 */
async function apiManagementCreateAuthorizationProviderGenericOAuth2() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const authorizationProviderId = "eventbrite";
  const parameters = {
    displayName: "eventbrite",
    identityProvider: "oauth2",
    oauth2: {
      grantTypes: {
        authorizationCode: {
          authorizationUrl: "https://www.eventbrite.com/oauth/authorize",
          clientId: "clientid",
          clientSecret: "clientsecretvalue",
          refreshUrl: "https://www.eventbrite.com/oauth/token",
          scopes: undefined,
          tokenUrl: "https://www.eventbrite.com/oauth/token",
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

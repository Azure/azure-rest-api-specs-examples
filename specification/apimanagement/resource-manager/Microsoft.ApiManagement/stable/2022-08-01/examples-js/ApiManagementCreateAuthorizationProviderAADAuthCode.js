const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates authorization provider.
 *
 * @summary Creates or updates authorization provider.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationProviderAADAuthCode.json
 */
async function apiManagementCreateAuthorizationProviderAadAuthCode() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const authorizationProviderId = "aadwithauthcode";
  const parameters = {
    displayName: "aadwithauthcode",
    identityProvider: "aad",
    oauth2: {
      grantTypes: {
        authorizationCode: {
          clientId: "",
          clientSecret: "",
          resourceUri: "https://graph.microsoft.com",
          scopes: "User.Read.All Group.Read.All",
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
    parameters
  );
  console.log(result);
}

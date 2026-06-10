const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates authorization provider.
 *
 * @summary creates or updates authorization provider.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateAuthorizationProviderAADAuthCodeWithKeyVault.json
 */
async function apiManagementCreateAuthorizationProviderAADAuthCodeWithKeyVault() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.authorizationProvider.createOrUpdate(
    "rg1",
    "apimService1",
    "aadwithkeyvault",
    {
      displayName: "Azure AD with Key Vault",
      identityProvider: "aad",
      oauth2: {
        grantTypes: {
          authorizationCode: {
            clientId: "53790825-fdd3-4b80-bc7a-4c3aaf25801d",
            resourceUri: "https://graph.microsoft.com",
            scopes: "User.Read.All Group.Read.All",
          },
        },
        keyVault: { secretIdentifier: "https://my.vault.azure.net/secrets/clientSecret" },
        redirectUrl:
          "https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1",
      },
    },
  );
  console.log(result);
}

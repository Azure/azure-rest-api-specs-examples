const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update the AuthConfig for a Container App.
 *
 * @summary create or update the AuthConfig for a Container App.
 * x-ms-original-file: 2025-10-02-preview/AuthConfigs_CreateOrUpdate.json
 */
async function createOrUpdateContainerAppAuthConfig() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsAuthConfigs.createOrUpdate(
    "workerapps-rg-xj",
    "testcanadacentral",
    "current",
    {
      encryptionSettings: {
        containerAppAuthEncryptionSecretName: "testEncryptionSecretName",
        containerAppAuthSigningSecretName: "testSigningSecretName",
      },
      globalValidation: { unauthenticatedClientAction: "AllowAnonymous" },
      identityProviders: {
        facebook: { registration: { appId: "123", appSecretSettingName: "facebook-secret" } },
      },
      platform: { enabled: true },
    },
  );
  console.log(result);
}

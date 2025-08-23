const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update the AuthConfig for a Container App.
 *
 * @summary Create or update the AuthConfig for a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2025-02-02-preview/examples/AuthConfigs_BlobStorageTokenStore_ClientId_CreateOrUpdate.json
 */
async function createOrUpdateContainerAppAuthConfigWithMsiClientIdBlobStorageTokenStore() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg1";
  const containerAppName = "myapp";
  const authConfigName = "current";
  const authConfigEnvelope = {
    encryptionSettings: {
      containerAppAuthEncryptionSecretName: "testEncryptionSecretName",
      containerAppAuthSigningSecretName: "testSigningSecretName",
    },
    globalValidation: { unauthenticatedClientAction: "AllowAnonymous" },
    identityProviders: {
      facebook: {
        registration: { appId: "123", appSecretSettingName: "facebook-secret" },
      },
    },
    login: {
      tokenStore: {
        azureBlobStorage: {
          blobContainerUri: "https://test.blob.core.windows.net/container1",
          clientId: "00000000-0000-0000-0000-000000000000",
        },
      },
    },
    platform: { enabled: true },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsAuthConfigs.createOrUpdate(
    resourceGroupName,
    containerAppName,
    authConfigName,
    authConfigEnvelope,
  );
  console.log(result);
}

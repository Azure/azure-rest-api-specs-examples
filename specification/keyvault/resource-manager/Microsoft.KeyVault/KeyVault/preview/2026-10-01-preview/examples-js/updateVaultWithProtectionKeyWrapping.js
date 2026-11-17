const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a key vault in the specified subscription.
 *
 * @summary update a key vault in the specified subscription.
 * x-ms-original-file: 2026-10-01-preview/updateVaultWithProtectionKeyWrapping.json
 */
async function configureProtectionKeyWrappingOnAnExistingVault() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.vaults.update("sample-resource-group", "sample-vault", {
    properties: {
      protectionKeysWrappingKey: {
        keyValue: { url: "https://sample-hsm.managedhsm.azure.net/keys/wrapping-key", kty: "oct" },
        keyAuthenticationData: { authType: "ServiceIdentity" },
      },
    },
  });
  console.log(result);
}

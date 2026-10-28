const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new Secret within the specified profile.
 *
 * @summary creates a new Secret within the specified profile.
 * x-ms-original-file: 2026-10-01/Secrets_CreateUrlSigningKeyType.json
 */
async function secretsCreateUrlSigningKeyType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.secrets.create("RG", "profile1", "secret1", {
    parameters: {
      type: "UrlSigningKey",
      keyId: "customKeyId",
      secretSource: {
        id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/urlsigningkeyname",
      },
      secretVersion: "abcdef1234578900abcdef1234567890",
    },
  });
  console.log(result);
}

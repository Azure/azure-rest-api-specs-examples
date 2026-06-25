const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes an existing Secret within profile.
 *
 * @summary deletes an existing Secret within profile.
 * x-ms-original-file: 2025-12-01/Secrets_Delete.json
 */
async function secretsDelete() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CdnManagementClient(credential, subscriptionId);
  await client.secrets.delete("RG", "profile1", "secret1");
}

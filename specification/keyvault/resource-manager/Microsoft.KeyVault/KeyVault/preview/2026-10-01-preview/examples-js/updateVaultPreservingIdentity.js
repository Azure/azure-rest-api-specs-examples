const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a key vault in the specified subscription.
 *
 * @summary update a key vault in the specified subscription.
 * x-ms-original-file: 2026-10-01-preview/updateVaultPreservingIdentity.json
 */
async function preserveExistingUserAssignedIdentitiesByOmittingIdentityFromPatch() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.vaults.update("sample-resource-group", "sample-vault", {
    tags: { environment: "production" },
  });
  console.log(result);
}

const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a key vault in the specified subscription.
 *
 * @summary create or update a key vault in the specified subscription.
 * x-ms-original-file: 2026-10-01-preview/createVaultWithIdentity.json
 */
async function createOrUpdateAVaultByFullyReplacingTheUserAssignedIdentitySet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.vaults.createOrUpdate("sample-resource-group", "sample-vault", {
    location: "westus",
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sample-identity":
          {},
      },
    },
    properties: {
      accessPolicies: [],
      enableRbacAuthorization: true,
      sku: { name: "standard", family: "A" },
      tenantId: "00000000-0000-0000-0000-000000000000",
    },
  });
  console.log(result);
}

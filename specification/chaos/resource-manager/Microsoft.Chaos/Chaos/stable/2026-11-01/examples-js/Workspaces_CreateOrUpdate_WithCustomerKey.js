const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Workspace resource.
 *
 * @summary create or update a Workspace resource.
 * x-ms-original-file: 2026-11-01/Workspaces_CreateOrUpdate_WithCustomerKey.json
 */
async function createOrUpdateAWorkspaceWithACustomerManagedKey() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.workspaces.createOrUpdate("exampleRG", "exampleWorkspace", {
    location: "westus2",
    properties: {
      scopes: ["/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG"],
      encryption: {
        customerManagedKeyEncryption: {
          keyEncryptionKeyUrl: "https://contoso-vault.vault.azure.net/keys/workspace-key",
          keyVaultResourceId:
            "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/customer-keys/providers/Microsoft.KeyVault/vaults/contoso-vault",
        },
      },
    },
  });
  console.log(result);
}

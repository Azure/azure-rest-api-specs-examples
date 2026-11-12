const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Workspace.
 *
 * @summary update a Workspace.
 * x-ms-original-file: 2026-11-01/Workspaces_Update_ChangeKeyVault.json
 */
async function updateAWorkspaceWithAValidatedKeyAndResourceIDInAnotherVault() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.workspaces.update("exampleRG", "exampleWorkspace", {
    properties: {
      encryption: {
        customerManagedKeyEncryption: {
          keyEncryptionKeyUrl: "https://contoso-vault-b.vault.azure.net/keys/Workspace-Key-Next",
          keyVaultResourceId:
            "/subscriptions/22222222-2222-2222-2222-222222222222/resourceGroups/customer-keys/providers/Microsoft.KeyVault/vaults/contoso-vault-b",
        },
      },
    },
  });
  console.log(result);
}

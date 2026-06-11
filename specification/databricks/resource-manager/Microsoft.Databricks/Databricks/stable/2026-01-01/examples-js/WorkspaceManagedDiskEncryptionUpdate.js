const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new workspace.
 *
 * @summary creates a new workspace.
 * x-ms-original-file: 2026-01-01/WorkspaceManagedDiskEncryptionUpdate.json
 */
async function updateAWorkspaceWithCustomerManagedKeyCMKEncryptionForManagedDisks() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.workspaces.createOrUpdate("rg", "myWorkspace", {
    location: "westus",
    computeMode: "Hybrid",
    encryption: {
      entities: {
        managedDisk: {
          keySource: "Microsoft.Keyvault",
          keyVaultProperties: {
            keyName: "test-cmk-key",
            keyVaultUri: "https://test-vault-name.vault.azure.net/",
            keyVersion: "00000000000000000000000000000000",
          },
          rotationToLatestKeyVersionEnabled: true,
        },
      },
    },
    managedResourceGroupId:
      "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myManagedRG",
    sku: { name: "premium" },
    tags: { mytag1: "myvalue1" },
  });
  console.log(result);
}

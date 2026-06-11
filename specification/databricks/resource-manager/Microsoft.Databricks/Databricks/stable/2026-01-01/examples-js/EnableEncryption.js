const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new workspace.
 *
 * @summary creates a new workspace.
 * x-ms-original-file: 2026-01-01/EnableEncryption.json
 */
async function enableCustomerManagedKeyCMKEncryptionOnAWorkspaceWhichIsPreparedForEncryption() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.workspaces.createOrUpdate("rg", "myWorkspace", {
    location: "westus",
    computeMode: "Hybrid",
    managedResourceGroupId:
      "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myManagedRG",
    parameters: {
      encryption: {
        value: {
          keyName: "myKeyName",
          keySource: "Microsoft.Keyvault",
          keyVaultUri: "https://myKeyVault.vault.azure.net/",
          keyVersion: "00000000000000000000000000000000",
        },
      },
      prepareEncryption: { value: true },
    },
    sku: { name: "premium" },
  });
  console.log(result);
}

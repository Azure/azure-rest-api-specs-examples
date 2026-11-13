const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Workspace.
 *
 * @summary update a Workspace.
 * x-ms-original-file: 2026-11-01/Workspaces_Update_ReplaceCustomerKey.json
 */
async function updateOnlyTheCustomerKeyURIOfAnEncryptedWorkspace() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.workspaces.update("exampleRG", "exampleWorkspace", {
    properties: {
      encryption: {
        customerManagedKeyEncryption: {
          keyEncryptionKeyUrl: "https://contoso-vault.vault.azure.net/keys/replacement-key",
        },
      },
    },
  });
  console.log(result);
}

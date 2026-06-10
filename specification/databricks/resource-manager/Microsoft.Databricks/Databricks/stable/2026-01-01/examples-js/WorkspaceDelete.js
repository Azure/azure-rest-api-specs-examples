const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes the workspace.
 *
 * @summary deletes the workspace.
 * x-ms-original-file: 2026-01-01/WorkspaceDelete.json
 */
async function deleteAWorkspace() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  await client.workspaces.delete("rg", "myWorkspace");
}

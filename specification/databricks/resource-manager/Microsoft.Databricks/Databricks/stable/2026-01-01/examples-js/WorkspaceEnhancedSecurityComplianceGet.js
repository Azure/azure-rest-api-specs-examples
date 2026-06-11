const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the workspace.
 *
 * @summary gets the workspace.
 * x-ms-original-file: 2026-01-01/WorkspaceEnhancedSecurityComplianceGet.json
 */
async function getAWorkspaceWithEnhancedSecurityComplianceAddOn() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.workspaces.get("rg", "myWorkspace");
  console.log(result);
}

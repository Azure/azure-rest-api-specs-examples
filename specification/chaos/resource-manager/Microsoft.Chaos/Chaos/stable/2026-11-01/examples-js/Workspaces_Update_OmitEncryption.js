const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Workspace.
 *
 * @summary update a Workspace.
 * x-ms-original-file: 2026-11-01/Workspaces_Update_OmitEncryption.json
 */
async function updateWorkspaceTagsWithoutChangingCustomerManagedEncryption() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.workspaces.update("exampleRG", "exampleWorkspace", {
    tags: { environment: "production" },
  });
  console.log(result);
}

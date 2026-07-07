const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes Verifier Workspace.
 *
 * @summary deletes Verifier Workspace.
 * x-ms-original-file: 2025-07-01/VerifierWorkspaceDelete.json
 */
async function verifierWorkspaceDelete() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  await client.verifierWorkspaces.delete("rg1", "testNetworkManager", "testWorkspace");
}

const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a retained discovery result by its ID, or uses latest to select the most recently durably admitted discovery. Returns 200 with the current representation while processing or terminal. A Failed result can contain errors without a snapshot. Returns 404 if no qualifying retained result exists. This is not an LRO polling endpoint.
 *
 * @summary gets a retained discovery result by its ID, or uses latest to select the most recently durably admitted discovery. Returns 200 with the current representation while processing or terminal. A Failed result can contain errors without a snapshot. Returns 404 if no qualifying retained result exists. This is not an LRO polling endpoint.
 * x-ms-original-file: 2026-11-01/WorkspaceDiscoveries_Get_InProgress.json
 */
async function readARunningDiscoveryWithHttp200WithoutInventingAcceptedCoverage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.workspaceDiscoveries.get(
    "exampleRG",
    "exampleWorkspace",
    "00000000-0000-0000-0000-000000000009",
  );
  console.log(result);
}

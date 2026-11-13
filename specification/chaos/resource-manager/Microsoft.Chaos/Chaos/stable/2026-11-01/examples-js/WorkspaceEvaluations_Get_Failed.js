const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a retained template evaluation by its ID, or uses latest to select the most recently durably admitted GA template evaluation. Returns 200 with the current representation while processing or terminal. Returns 404 if no qualifying retained result exists; scenario-backed legacy operations are not substituted. This is not an LRO polling endpoint.
 *
 * @summary gets a retained template evaluation by its ID, or uses latest to select the most recently durably admitted GA template evaluation. Returns 200 with the current representation while processing or terminal. Returns 404 if no qualifying retained result exists; scenario-backed legacy operations are not substituted. This is not an LRO polling endpoint.
 * x-ms-original-file: 2026-11-01/WorkspaceEvaluations_Get_Failed.json
 */
async function readAFailedTemplateEvaluationIndependentlyOfWorkspaceScenarios() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.workspaceEvaluations.get(
    "exampleRG",
    "exampleWorkspace",
    "00000000-0000-0000-0000-000000000003",
  );
  console.log(result);
}

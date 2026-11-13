const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the state of specified Content Filters associated with the Azure OpenAI account.
 *
 * @summary update the state of specified Content Filters associated with the Azure OpenAI account.
 * x-ms-original-file: 2026-11-01/PutRaiPolicyAcs.json
 */
async function createAnACSPolicyWithOptionalPolicyDependencies() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.raiPolicies.createOrUpdate(
    "resource-group",
    "safety-account",
    "agent-guard",
    {
      properties: {
        customBlocklists: [{ blocklistName: "blocked-terms", source: "Prompt", blocking: true }],
      },
    },
  );
  console.log(result);
}

const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the state of specified Content Filters associated with the Azure OpenAI account.
 *
 * @summary update the state of specified Content Filters associated with the Azure OpenAI account.
 * x-ms-original-file: 2026-09-15-preview/PutRaiPolicyAcsWithoutTools.json
 */
async function createAnACSPolicyWithoutAToolCatalog() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.raiPolicies.createOrUpdate(
    "resource-group",
    "safety-account",
    "input-guard",
    {
      properties: {
        format: "ACS",
        acs: {
          agentControlSpecificationVersion: "0.4.0-alpha.1",
          policies: { "input-guard": { type: "rego", query: "data.input_guard.verdict" } },
          interventionPoints: {
            input: {
              policyTarget: "$snap.input",
              policyTargetKind: "user_input",
              policy: { id: "input-guard" },
            },
          },
        },
        acsRegos: [{ regoName: "input-guard" }],
      },
    },
    { ifNoneMatch: "*" },
  );
  console.log(result);
}

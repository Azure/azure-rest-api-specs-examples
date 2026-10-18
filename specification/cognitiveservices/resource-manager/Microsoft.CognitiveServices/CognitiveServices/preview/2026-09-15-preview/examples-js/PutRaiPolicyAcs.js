const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the state of specified Content Filters associated with the Azure OpenAI account.
 *
 * @summary update the state of specified Content Filters associated with the Azure OpenAI account.
 * x-ms-original-file: 2026-09-15-preview/PutRaiPolicyAcs.json
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
        format: "ACS",
        acs: {
          agentControlSpecificationVersion: "0.4.0-alpha.1",
          metadata: { name: "agent-guard" },
          policies: {
            "input-guard": { type: "rego", query: "data.input_guard.verdict" },
            "tool-guard": { type: "rego", query: "data.tool_guard.verdict" },
          },
          interventionPoints: {
            input: {
              policyTarget: "$snap.input",
              policyTargetKind: "user_input",
              policy: {
                id: "input-guard",
                aacsModeration: {
                  subjectFormat: "text",
                  harmConfigs: [{ category: "PromptInjection" }],
                },
              },
            },
            preToolCall: {
              policyTarget: "$snap.tool_call.args",
              policyTargetKind: "tool_args",
              toolNameFrom: "$snap.tool_call.name",
              policy: { id: "tool-guard" },
            },
            postToolCall: {
              policyTarget: "$snap.tool_result.value",
              policyTargetKind: "tool_result",
              toolNameFrom: "$snap.tool_call.name",
              policy: { id: "tool-guard" },
            },
          },
          tools: {
            web_search: {
              id: "web_search",
              type: "retrieval",
              description: "Search approved public documentation",
              securityLabels: ["network_egress", "untrusted_content"],
              clearance: "public",
              additionalProperties: {
                allowed_domains: ["learn.microsoft.com"],
              },
            },
            wire_transfer: {
              id: "wire_transfer",
              type: "financial_action",
              securityLabels: ["financial_write"],
              clearance: "confidential",
            },
          },
          annotators: {},
        },
        acsRegos: [{ regoName: "input-guard" }, { regoName: "tool-guard" }],
        customBlocklists: [{ blocklistName: "blocked-terms", source: "Prompt", blocking: true }],
        customExternalSafetyProviders: [
          {
            externalSafetyProviderName: "contoso-safety-provider",
            managedIdentityResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/safety-provider-identity",
            source: "Prompt",
            blocking: true,
          },
        ],
      },
    },
    { ifNoneMatch: "*" },
  );
  console.log(result);
}

const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the state of specified Content Filters associated with the Azure OpenAI account.
 *
 * @summary update the state of specified Content Filters associated with the Azure OpenAI account.
 * x-ms-original-file: 2026-11-15-preview/PutRaiPolicyAcsWithAnnotationsAndEgress.json
 */
async function createAnACSPolicyWithOptionalAnnotationsAndAnEgressCompanion() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.raiPolicies.createOrUpdate(
    "resource-group",
    "safety-account",
    "annotated-egress-guard",
    {
      properties: {
        format: "ACS",
        mode: "Asynchronous_filter",
        acs: {
          agentControlSpecificationVersion: "0.4.0-alpha.1",
          metadata: {
            description:
              "Illustrative annotation authoring; supported annotators and configuration depend on the evaluation host.",
          },
          policies: { "agent-guard": { type: "rego", query: "data.agent_guard.verdict" } },
          annotators: {
            risk: {
              type: "llm",
              additionalProperties: {
                provider: "azure_openai",
                endpoint: "https://synthetic-judge.openai.azure.com/",
                deployment: "judge",
                api_version: "2024-10-21",
                result_schema: "acs.judge.observation/v2",
                system_prompt:
                  "Classify the complete subject for harmful intent. Score from 0 to 10: 0 is benign, 5 is ambiguous or potentially harmful, and 10 is clear harmful intent. Use label clear for scores below 4, review for scores from 4 to below 7, and high_risk for scores of 7 or higher. Report observations only; Rego decides the action.",
                timeout_ms: 10000,
              },
            },
          },
          interventionPoints: {
            agentStartup: {
              policyTarget: "$snap.agent_init",
              policyTargetKind: "agent_init",
              policy: { id: "agent-guard" },
            },
            input: {
              policyTarget: "$snap.input",
              policyTargetKind: "user_input",
              policy: { id: "agent-guard" },
              annotations: { risk: { from: "$target.content" } },
            },
            preModelCall: {
              policyTarget: "$snap.messages",
              policyTargetKind: "messages",
              policy: { id: "agent-guard" },
            },
            postModelCall: {
              policyTarget: "$snap.response",
              policyTargetKind: "model_response",
              policy: { id: "agent-guard" },
            },
            preToolCall: {
              policyTarget: "$snap.tool_call.args",
              policyTargetKind: "tool_args",
              policy: { id: "agent-guard" },
            },
            postToolCall: {
              policyTarget: "$snap.tool_result.value",
              policyTargetKind: "tool_result",
              policy: { id: "agent-guard" },
            },
            output: {
              policyTarget: "$snap.output",
              policyTargetKind: "assistant_output",
              policy: { id: "agent-guard" },
              annotations: { risk: { from: "$target.content" } },
            },
            agentShutdown: {
              policyTarget: "$snap.summary",
              policyTargetKind: "agent_summary",
              policy: { id: "agent-guard" },
            },
          },
        },
        acsRegos: [{ regoName: "agent-guard" }],
        egressPolicy: {
          mode: "Audit",
          defaultAction: "Deny",
          rules: [
            {
              name: "allow-docs",
              ruleType: "Fqdn",
              match: { host: "learn.microsoft.com", path: "/azure/*" },
              action: { actionType: "Allow" },
            },
            {
              name: "label-api",
              ruleType: "Fqdn",
              match: { host: "api.contoso.com", path: "/v1/*" },
              action: {
                actionType: "Transform",
                headers: [{ operation: "Set", name: "X-Workload-Class", value: "public-demo" }],
              },
            },
          ],
        },
      },
    },
    { ifNoneMatch: "*" },
  );
  console.log(result);
}

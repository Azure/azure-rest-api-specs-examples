const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or replaces one reusable Rego artifact.
 *
 * @summary creates or replaces one reusable Rego artifact.
 * x-ms-original-file: 2026-09-15-preview/PutRaiRego.json
 */
async function createAReusableRegoArtifact() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.raiRegos.createOrUpdate(
    "resource-group",
    "safety-account",
    "input-guard",
    {
      properties: {
        encoding: "None",
        rego: 'package input_guard\n\nimport rego.v1\n\ndefault verdict := {"decision": "allow"}\n\nverdict := {"decision": "deny", "reason": "prompt_injection_detected"} if {\n  input.snapshot.moderation.harm.PromptInjection.detected\n}',
      },
      tags: { environment: "production" },
    },
    { ifNoneMatch: "*" },
  );
  console.log(result);
}

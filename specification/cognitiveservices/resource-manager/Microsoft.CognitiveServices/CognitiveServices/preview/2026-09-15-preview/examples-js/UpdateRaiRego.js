const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or replaces one reusable Rego artifact.
 *
 * @summary creates or replaces one reusable Rego artifact.
 * x-ms-original-file: 2026-09-15-preview/UpdateRaiRego.json
 */
async function replaceAReusableRegoArtifactConditionally() {
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
        rego: 'package input_guard\n\nimport rego.v1\n\ndefault verdict := {"decision": "allow"}\n\nblocking_signal if {\n  input.snapshot.moderation.harm.PromptInjection.detected\n}\n\nblocking_signal if {\n  input.snapshot.moderation.harm.Hate.severity >= 4\n}\n\nverdict := {"decision": "deny", "reason": "unsafe_input"} if {\n  blocking_signal\n}',
      },
    },
    { ifMatch: '"00000000-0000-0000-0000-000000000001"' },
  );
  console.log(result);
}

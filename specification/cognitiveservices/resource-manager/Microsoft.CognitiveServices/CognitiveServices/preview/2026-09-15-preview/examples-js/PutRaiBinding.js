const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or replaces one RAI binding.
 *
 * @summary creates or replaces one RAI binding.
 * x-ms-original-file: 2026-09-15-preview/PutRaiBinding.json
 */
async function bindAnAzureOpenAIDeploymentToAnACSPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.raiBindings.createOrUpdate(
    "resource-group",
    "safety-account",
    "chat-guard",
    {
      properties: {
        boundResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/aoai-account/deployments/chat-prod",
        targetPolicyName: "agent-guard",
      },
    },
    { ifNoneMatch: "*" },
  );
  console.log(result);
}

const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Arc deployment associated with the Cognitive Services account.
 *
 * @summary creates or updates an Arc deployment associated with the Cognitive Services account.
 * x-ms-original-file: 2026-11-15-preview/CreateOrUpdateArcDeploymentWithVllm.json
 */
async function createOrUpdateArcDeploymentWithVllm() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.arcDeployments.createOrUpdate(
    "resourceGroupName",
    "accountName",
    "gpt-oss-vllm-arc",
    {
      properties: {
        model: {
          assetId:
            "azureml://registries/azure-foundrylocalonazurelocal/models/openai-gpt-oss-20b/versions/1",
        },
        extensionId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/edge-cluster/providers/Microsoft.KubernetesConfiguration/extensions/inference-operator",
        replicas: 2,
        resources: { limits: { gpu: 5 } },
        nodeSelector: { agentpool: "a100" },
      },
      sku: { name: "Arc" },
    },
  );
  console.log(result);
}

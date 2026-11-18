const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Arc deployment associated with the Cognitive Services account.
 *
 * @summary creates or updates an Arc deployment associated with the Cognitive Services account.
 * x-ms-original-file: 2026-11-15-preview/CreateOrUpdateArcDeployment.json
 */
async function createOrUpdateArcDeployment() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.arcDeployments.createOrUpdate(
    "resourceGroupName",
    "accountName",
    "phi-3-arc",
    {
      properties: {
        model: {
          assetId:
            "azureml://registries/azureml/models/Phi-3-mini-128k-instruct-generic-cpu/versions/3",
        },
        extensionId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/edge-cluster/providers/Microsoft.KubernetesConfiguration/extensions/inference-operator",
        replicas: 2,
        resources: { requests: { cpu: "8", memory: "16Gi" }, limits: { cpu: "8", memory: "16Gi" } },
        nodeSelector: { agentpool: "cpu" },
      },
      sku: { name: "Arc" },
    },
  );
  console.log(result);
}

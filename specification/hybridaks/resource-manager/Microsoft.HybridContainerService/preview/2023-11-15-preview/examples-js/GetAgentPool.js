const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the agent pool in the Hybrid AKS provisioned cluster instance
 *
 * @summary Gets the agent pool in the Hybrid AKS provisioned cluster instance
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/GetAgentPool.json
 */
async function getAgentPool() {
  const connectedClusterResourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster";
  const agentPoolName = "test-hybridaksnodepool";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential);
  const result = await client.agentPoolOperations.get(connectedClusterResourceUri, agentPoolName);
  console.log(result);
}

const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an agent pool in the specified managed cluster. Visit https://aka.ms/aks/concurrent-node-operations for more information.
 *
 * @summary updates an agent pool in the specified managed cluster. Visit https://aka.ms/aks/concurrent-node-operations for more information.
 * x-ms-original-file: 2026-06-02-preview/AgentPoolsUpdate_Scale.json
 */
async function updateAgentPoolScaleVmss() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.update("rg1", "clustername1", "agentpool1", {
    properties: { count: 5 },
  });
  console.log(result);
}

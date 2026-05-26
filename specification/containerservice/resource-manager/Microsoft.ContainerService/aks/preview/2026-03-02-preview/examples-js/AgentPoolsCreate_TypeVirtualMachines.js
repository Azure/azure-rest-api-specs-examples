const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an agent pool in the specified managed cluster.
 *
 * @summary creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: 2026-03-02-preview/AgentPoolsCreate_TypeVirtualMachines.json
 */
async function createAgentPoolWithVirtualMachinesPoolType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.createOrUpdate("rg1", "clustername1", "agentpool1", {
    typePropertiesType: "VirtualMachines",
    nodeLabels: { key1: "val1" },
    nodeTaints: ["Key1=Value1:NoSchedule"],
    orchestratorVersion: "1.9.6",
    osType: "Linux",
    tags: { name1: "val1" },
    virtualMachinesProfile: {
      scale: {
        manual: [
          { count: 3, size: "Standard_D2_v2" },
          { count: 2, size: "Standard_D2_v3" },
        ],
      },
    },
  });
  console.log(result);
}

const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an agent pool in the specified managed cluster.
 *
 * @summary Creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/examples/AgentPoolsCreate_TypeVirtualMachines.json
 */
async function createAgentPoolWithVirtualMachinesPoolType() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "agentpool1";
  const parameters = {
    typePropertiesType: "VirtualMachines",
    nodeLabels: { key1: "val1" },
    nodeTaints: ["Key1=Value1:NoSchedule"],
    orchestratorVersion: "1.9.6",
    osType: "Linux",
    tags: { name1: "val1" },
    virtualMachinesProfile: {
      scale: {
        manual: [{ count: 5, sizes: ["Standard_D2_v2", "Standard_D2_v3"] }],
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    agentPoolName,
    parameters,
  );
  console.log(result);
}

const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an agent pool in the specified managed cluster.
 *
 * @summary Creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/AgentPoolsCreate_Spot.json
 */
async function createSpotAgentPool() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "agentpool1";
  const parameters = {
    count: 3,
    nodeLabels: { key1: "val1" },
    nodeTaints: ["Key1=Value1:NoSchedule"],
    orchestratorVersion: "",
    osType: "Linux",
    scaleSetEvictionPolicy: "Delete",
    scaleSetPriority: "Spot",
    tags: { name1: "val1" },
    vmSize: "Standard_DS1_v2",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    agentPoolName,
    parameters
  );
  console.log(result);
}

createSpotAgentPool().catch(console.error);

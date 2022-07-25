const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an agent pool for a container registry with the specified parameters.
 *
 * @summary Creates an agent pool for a container registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsCreate.json
 */
async function agentPoolsCreate() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const agentPoolName = "myAgentPool";
  const agentPool = {
    count: 1,
    location: "WESTUS",
    os: "Linux",
    tags: { key: "value" },
    tier: "S1",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.agentPools.beginCreateAndWait(
    resourceGroupName,
    registryName,
    agentPoolName,
    agentPool
  );
  console.log(result);
}

agentPoolsCreate().catch(console.error);

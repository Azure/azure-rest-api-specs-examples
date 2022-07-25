const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an agent pool with the specified parameters.
 *
 * @summary Updates an agent pool with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsUpdate.json
 */
async function agentPoolsUpdate() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const agentPoolName = "myAgentPool";
  const updateParameters = { count: 1 };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.agentPools.beginUpdateAndWait(
    resourceGroupName,
    registryName,
    agentPoolName,
    updateParameters
  );
  console.log(result);
}

agentPoolsUpdate().catch(console.error);

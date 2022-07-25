const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified managed cluster agent pool.
 *
 * @summary Gets the specified managed cluster agent pool.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-06-01/examples/AgentPoolsGet.json
 */
async function getAgentPool() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "agentpool1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.get(resourceGroupName, resourceName, agentPoolName);
  console.log(result);
}

getAgentPool().catch(console.error);

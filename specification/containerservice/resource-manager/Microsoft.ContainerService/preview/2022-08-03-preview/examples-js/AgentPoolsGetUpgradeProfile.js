const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the upgrade profile for an agent pool.
 *
 * @summary Gets the upgrade profile for an agent pool.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-08-03-preview/examples/AgentPoolsGetUpgradeProfile.json
 */
async function getUpgradeProfileForAgentPool() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "agentpool1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.getUpgradeProfile(
    resourceGroupName,
    resourceName,
    agentPoolName
  );
  console.log(result);
}

getUpgradeProfileForAgentPool().catch(console.error);

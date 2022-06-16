const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function upgradeAgentPoolNodeImageVersion() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "agentpool1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.beginUpgradeNodeImageVersionAndWait(
    resourceGroupName,
    resourceName,
    agentPoolName
  );
  console.log(result);
}

upgradeAgentPoolNodeImageVersion().catch(console.error);

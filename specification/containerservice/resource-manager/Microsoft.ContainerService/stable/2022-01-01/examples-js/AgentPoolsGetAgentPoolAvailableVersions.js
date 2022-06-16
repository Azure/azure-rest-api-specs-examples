const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function getAvailableVersionsForAgentPool() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.getAvailableAgentPoolVersions(
    resourceGroupName,
    resourceName
  );
  console.log(result);
}

getAvailableVersionsForAgentPool().catch(console.error);

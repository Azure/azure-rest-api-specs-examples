const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAgentPoolWithPpg() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "agentpool1";
  const parameters = {
    count: 3,
    orchestratorVersion: "",
    osType: "Linux",
    proximityPlacementGroupID:
      "/subscriptions/subid1/resourcegroups/rg1/providers//Microsoft.Compute/proximityPlacementGroups/ppg1",
    vmSize: "Standard_DS2_v2",
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

createAgentPoolWithPpg().catch(console.error);

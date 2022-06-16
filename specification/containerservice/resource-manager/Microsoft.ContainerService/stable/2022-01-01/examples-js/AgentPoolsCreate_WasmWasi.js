const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAgentPoolWithKrustletAndTheWasiRuntime() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "agentpool1";
  const parameters = {
    count: 3,
    mode: "User",
    orchestratorVersion: "",
    osDiskSizeGB: 64,
    osType: "Linux",
    vmSize: "Standard_DS2_v2",
    workloadRuntime: "WasmWasi",
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

createAgentPoolWithKrustletAndTheWasiRuntime().catch(console.error);

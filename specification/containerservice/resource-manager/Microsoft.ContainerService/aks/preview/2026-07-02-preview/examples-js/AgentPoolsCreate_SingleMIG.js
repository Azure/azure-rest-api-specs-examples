const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an agent pool in the specified managed cluster.
 *
 * @summary creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: 2026-07-02-preview/AgentPoolsCreate_SingleMIG.json
 */
async function createAgentPoolWithSingleMIGProfile() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.createOrUpdate("rg1", "clustername1", "agentpool1", {
    count: 3,
    gpuProfile: {
      driver: "Install",
      nvidia: { managementMode: "Managed", migStrategy: "Single", migProfiles: ["MIG3g"] },
    },
    mode: "User",
    osType: "Linux",
    vmSize: "Standard_ND96asr_v4",
  });
  console.log(result);
}

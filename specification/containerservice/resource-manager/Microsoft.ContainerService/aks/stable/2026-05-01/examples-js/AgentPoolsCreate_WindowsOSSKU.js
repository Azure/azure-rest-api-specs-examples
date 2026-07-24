const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an agent pool in the specified managed cluster.
 *
 * @summary creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: 2026-05-01/AgentPoolsCreate_WindowsOSSKU.json
 */
async function createAgentPoolWithWindowsOssku() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.createOrUpdate("rg1", "clustername1", "wnp2", {
    count: 3,
    orchestratorVersion: "1.23.3",
    osSKU: "Windows2022",
    osType: "Windows",
    vmSize: "Standard_D4s_v3",
  });
  console.log(result);
}

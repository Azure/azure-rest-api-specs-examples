const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an agent pool in the specified managed cluster.
 *
 * @summary Creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-09-02-preview/examples/AgentPoolsCreate_WindowsDisableOutboundNAT.json
 */
async function createWindowsAgentPoolWithDisablingOutboundNat() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "wnp2";
  const parameters = {
    count: 3,
    orchestratorVersion: "1.23.8",
    osSKU: "Windows2022",
    osType: "Windows",
    vmSize: "Standard_D4s_v3",
    windowsProfile: { disableOutboundNat: true },
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

createWindowsAgentPoolWithDisablingOutboundNat().catch(console.error);

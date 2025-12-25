const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates an agent pool in the specified managed cluster.
 *
 * @summary Creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/AgentPoolsCreate_WindowsDisableOutboundNAT.json
 */
async function createWindowsAgentPoolWithDisablingOutboundNat() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
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
    parameters,
  );
  console.log(result);
}

const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates an agent pool in the specified managed cluster.
 *
 * @summary Creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-04-02-preview/examples/AgentPoolsCreate_Snapshot.json
 */
async function createAgentPoolUsingAnAgentPoolSnapshot() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const agentPoolName = "agentpool1";
  const parameters = {
    count: 3,
    creationData: {
      sourceResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/snapshots/snapshot1",
    },
    enableFips: true,
    orchestratorVersion: "",
    osType: "Linux",
    vmSize: "Standard_DS2_v2",
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

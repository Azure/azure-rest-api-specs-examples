const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch the properties of the provided Kubernetes cluster agent pool, or update the tags associated with the Kubernetes cluster agent pool. Properties and tag updates can be done independently.
 *
 * @summary Patch the properties of the provided Kubernetes cluster agent pool, or update the tags associated with the Kubernetes cluster agent pool. Properties and tag updates can be done independently.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/AgentPools_Patch.json
 */
async function patchKubernetesClusterAgentPool() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const kubernetesClusterName = "kubernetesClusterName";
  const agentPoolName = "agentPoolName";
  const agentPoolUpdateParameters = {
    count: 3,
    tags: { key1: "myvalue1", key2: "myvalue2" },
    upgradeSettings: { maxSurge: "1" },
  };
  const options = { agentPoolUpdateParameters };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.agentPools.beginUpdateAndWait(
    resourceGroupName,
    kubernetesClusterName,
    agentPoolName,
    options
  );
  console.log(result);
}

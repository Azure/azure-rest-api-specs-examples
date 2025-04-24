const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete the provided Kubernetes cluster agent pool.
 *
 * @summary Delete the provided Kubernetes cluster agent pool.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/AgentPools_Delete.json
 */
async function deleteKubernetesClusterAgentPool() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const kubernetesClusterName = "kubernetesClusterName";
  const agentPoolName = "agentPoolName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.agentPools.beginDeleteAndWait(
    resourceGroupName,
    kubernetesClusterName,
    agentPoolName,
  );
  console.log(result);
}

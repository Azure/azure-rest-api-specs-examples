const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get a list of agent pools for the provided Kubernetes cluster.
 *
 * @summary Get a list of agent pools for the provided Kubernetes cluster.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/AgentPools_ListByKubernetesCluster.json
 */
async function listAgentPoolsOfTheKubernetesCluster() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const kubernetesClusterName = "kubernetesClusterName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.agentPools.listByKubernetesCluster(
    resourceGroupName,
    kubernetesClusterName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

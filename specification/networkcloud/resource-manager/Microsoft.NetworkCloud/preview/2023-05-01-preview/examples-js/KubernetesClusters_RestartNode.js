const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restart a targeted node of a Kubernetes cluster.
 *
 * @summary Restart a targeted node of a Kubernetes cluster.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/KubernetesClusters_RestartNode.json
 */
async function restartAKubernetesClusterNode() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const kubernetesClusterName = "kubernetesClusterName";
  const kubernetesClusterRestartNodeParameters = {
    nodeName: "nodeName",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.kubernetesClusters.beginRestartNodeAndWait(
    resourceGroupName,
    kubernetesClusterName,
    kubernetesClusterRestartNodeParameters
  );
  console.log(result);
}

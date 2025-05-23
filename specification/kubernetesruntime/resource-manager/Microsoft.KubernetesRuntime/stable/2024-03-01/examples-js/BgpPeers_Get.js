const { KubernetesRuntimeClient } = require("@azure/arm-containerorchestratorruntime");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a BgpPeer
 *
 * @summary get a BgpPeer
 * x-ms-original-file: 2024-03-01/BgpPeers_Get.json
 */
async function bgpPeersGet() {
  const credential = new DefaultAzureCredential();
  const client = new KubernetesRuntimeClient(credential);
  const result = await client.bgpPeers.get(
    "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/example/providers/Microsoft.Kubernetes/connectedClusters/cluster1",
    "testpeer",
  );
  console.log(result);
}

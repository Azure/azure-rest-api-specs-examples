const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to patch the properties of the provided cluster, or update the tags associated with the cluster. Properties and tag updates can be done independently.
 *
 * @summary patch the properties of the provided cluster, or update the tags associated with the cluster. Properties and tag updates can be done independently.
 * x-ms-original-file: 2026-05-01-preview/Clusters_Patch_UpdateStrategy.json
 */
async function patchUpdateStrategy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.clusters.update("resourceGroupName", "clusterName", {
    clusterUpdateParameters: {
      updateStrategy: {
        maxUnavailable: 4,
        strategyType: "Rack",
        thresholdType: "CountSuccess",
        thresholdValue: 4,
        waitTimeMinutes: 10,
      },
      tags: { key1: "myvalue1", key2: "myvalue2" },
    },
  });
  console.log(result);
}

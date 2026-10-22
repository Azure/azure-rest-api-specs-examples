const { HorizonDbClient } = require("@azure/arm-horizondb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new HorizonDB cluster or replaces an existing cluster. PointInTimeRestore creates the new Cluster resource and hydrates its service-owned database storage from an available source backup. It does not clone source pools, Pool-scoped endpoints, firewall rules, parameter-group associations, or other child resources. The new Cluster receives its service-generated Cluster-scoped Endpoint named cluster. For restore, postgreSqlVersion must be omitted because the service obtains it from backup metadata; all other new-cluster configuration comes from the request or service defaults. Same-location restore requires and honors pointInTimeUtc and allows creationSourceClusterLocation to be omitted. Cross-location restore requires creationSourceClusterLocation to identify the source cluster's Azure location and requires pointInTimeUtc, but ignores the supplied timestamp and restores the latest available backup. Cross-location responses omit pointInTimeUtc.
 *
 * @summary creates a new HorizonDB cluster or replaces an existing cluster. PointInTimeRestore creates the new Cluster resource and hydrates its service-owned database storage from an available source backup. It does not clone source pools, Pool-scoped endpoints, firewall rules, parameter-group associations, or other child resources. The new Cluster receives its service-generated Cluster-scoped Endpoint named cluster. For restore, postgreSqlVersion must be omitted because the service obtains it from backup metadata; all other new-cluster configuration comes from the request or service defaults. Same-location restore requires and honors pointInTimeUtc and allows creationSourceClusterLocation to be omitted. Cross-location restore requires creationSourceClusterLocation to identify the source cluster's Azure location and requires pointInTimeUtc, but ignores the supplied timestamp and restores the latest available backup. Cross-location responses omit pointInTimeUtc.
 * x-ms-original-file: 2026-10-01-preview/Clusters_CreateOrReplacePointInTimeRestore.json
 */
async function restoreAHorizonDBClusterToAPointInTime() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new HorizonDbClient(credential, subscriptionId);
  const result = await client.clusters.createOrReplace(
    "example-resource-group",
    "restored-cluster",
    {
      location: "eastus2",
      properties: {
        creationMode: "PointInTimeRestore",
        pointInTimeUtc: new Date("2026-07-27T12:00:00Z"),
        creationSourceClusterResourceId:
          "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/source-cluster",
        network: { publicAccess: "Disabled" },
        backup: { geoRedundancy: "Enabled", retentionDays: 14 },
      },
    },
  );
  console.log(result);
}

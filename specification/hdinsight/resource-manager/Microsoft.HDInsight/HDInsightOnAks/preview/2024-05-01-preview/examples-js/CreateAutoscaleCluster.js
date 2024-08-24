const { HDInsightContainersManagementClient } = require("@azure/arm-hdinsightcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a cluster.
 *
 * @summary Creates a cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/CreateAutoscaleCluster.json
 */
async function hdInsightClusterPut() {
  const subscriptionId =
    process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "10e32bab-26da-4cc4-a441-52b318f824e6";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "hiloResourcegroup";
  const clusterPoolName = "clusterpool1";
  const clusterName = "cluster1";
  const hDInsightCluster = {
    location: "West US 2",
    properties: {
      clusterProfile: {
        authorizationProfile: { userIds: ["testuser1", "testuser2"] },
        autoscaleProfile: {
          autoscaleType: "ScheduleBased",
          enabled: true,
          gracefulDecommissionTimeout: 3600,
          loadBasedConfig: {
            cooldownPeriod: 300,
            maxNodes: 20,
            minNodes: 10,
            pollInterval: 60,
            scalingRules: [
              {
                actionType: "scaleup",
                comparisonRule: { operator: "greaterThan", threshold: 90 },
                evaluationCount: 3,
                scalingMetric: "cpu",
              },
              {
                actionType: "scaledown",
                comparisonRule: { operator: "lessThan", threshold: 20 },
                evaluationCount: 3,
                scalingMetric: "cpu",
              },
            ],
          },
          scheduleBasedConfig: {
            defaultCount: 10,
            schedules: [
              {
                count: 20,
                days: ["Monday"],
                endTime: "12:00",
                startTime: "00:00",
              },
              {
                count: 25,
                days: ["Sunday"],
                endTime: "12:00",
                startTime: "00:00",
              },
            ],
            timeZone: "Cen. Australia Standard Time",
          },
        },
        clusterVersion: "1.0.6",
        managedIdentityProfile: {
          identityList: [
            {
              type: "cluster",
              clientId: "de91f1d8-767f-460a-ac11-3cf103f74b34",
              objectId: "40491351-c240-4042-91e0-f644a1d2b441",
              resourceId:
                "/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-msi",
            },
          ],
        },
        ossVersion: "0.410.0",
        sshProfile: { count: 2, vmSize: "Standard_E8as_v5" },
        trinoProfile: {},
      },
      clusterType: "Trino",
      computeProfile: {
        availabilityZones: ["1", "2", "3"],
        nodes: [
          { type: "Head", count: 2, vmSize: "Standard_E8as_v5" },
          { type: "Worker", count: 3, vmSize: "Standard_E8as_v5" },
        ],
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightContainersManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginCreateAndWait(
    resourceGroupName,
    clusterPoolName,
    clusterName,
    hDInsightCluster,
  );
  console.log(result);
}

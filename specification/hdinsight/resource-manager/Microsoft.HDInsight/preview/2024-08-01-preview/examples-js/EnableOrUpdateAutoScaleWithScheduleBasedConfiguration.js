const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the Autoscale Configuration for HDInsight cluster.
 *
 * @summary Updates the Autoscale Configuration for HDInsight cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/EnableOrUpdateAutoScaleWithScheduleBasedConfiguration.json
 */
async function enableOrUpdateAutoscaleWithTheScheduleBasedConfigurationForHdInsightCluster() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const roleName = "workernode";
  const parameters = {
    autoscale: {
      recurrence: {
        schedule: [
          {
            days: ["Thursday"],
            timeAndCapacity: {
              maxInstanceCount: 4,
              minInstanceCount: 4,
              time: "16:00",
            },
          },
        ],
        timeZone: "China Standard Time",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginUpdateAutoScaleConfigurationAndWait(
    resourceGroupName,
    clusterName,
    roleName,
    parameters,
  );
  console.log(result);
}

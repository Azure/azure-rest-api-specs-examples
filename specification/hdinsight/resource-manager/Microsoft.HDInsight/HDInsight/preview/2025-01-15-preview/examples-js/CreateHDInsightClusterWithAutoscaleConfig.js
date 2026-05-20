const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new HDInsight cluster with the specified parameters.
 *
 * @summary creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: 2025-01-15-preview/CreateHDInsightClusterWithAutoscaleConfig.json
 */
async function createHDInsightClusterWithAutoscaleConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.create("rg1", "cluster1", {
    properties: {
      clusterDefinition: {
        componentVersion: { Hadoop: "2.7" },
        configurations: {
          gateway: {
            "restAuthCredential.isEnabled": true,
            "restAuthCredential.password": "**********",
            "restAuthCredential.username": "admin",
          },
        },
        kind: "hadoop",
      },
      clusterVersion: "3.6",
      computeProfile: {
        roles: [
          {
            name: "workernode",
            autoscaleConfiguration: {
              recurrence: {
                schedule: [
                  {
                    days: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"],
                    timeAndCapacity: { maxInstanceCount: 3, minInstanceCount: 3, time: "09:00" },
                  },
                  {
                    days: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"],
                    timeAndCapacity: { maxInstanceCount: 6, minInstanceCount: 6, time: "18:00" },
                  },
                  {
                    days: ["Saturday", "Sunday"],
                    timeAndCapacity: { maxInstanceCount: 2, minInstanceCount: 2, time: "09:00" },
                  },
                  {
                    days: ["Saturday", "Sunday"],
                    timeAndCapacity: { maxInstanceCount: 4, minInstanceCount: 4, time: "18:00" },
                  },
                ],
                timeZone: "China Standard Time",
              },
            },
            hardwareProfile: { vmSize: "Standard_D4_V2" },
            osProfile: {
              linuxOperatingSystemProfile: { password: "**********", username: "sshuser" },
            },
            scriptActions: [],
            targetInstanceCount: 4,
          },
        ],
      },
      osType: "Linux",
      storageProfile: {
        storageaccounts: [
          {
            name: "mystorage.blob.core.windows.net",
            container: "hdinsight-autoscale-tes-2019-06-18t05-49-16-591z",
            enableSecureChannel: true,
            isDefault: true,
            key: "storagekey",
          },
        ],
      },
      tier: "Standard",
    },
  });
  console.log(result);
}

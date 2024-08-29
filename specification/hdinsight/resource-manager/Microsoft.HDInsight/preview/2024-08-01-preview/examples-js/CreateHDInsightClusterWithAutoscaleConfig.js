const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new HDInsight cluster with the specified parameters.
 *
 * @summary Creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/CreateHDInsightClusterWithAutoscaleConfig.json
 */
async function createHdInsightClusterWithAutoscaleConfiguration() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const parameters = {
    properties: {
      clusterDefinition: {
        componentVersion: { hadoop: "2.7" },
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
              capacity: {},
              recurrence: {
                schedule: [
                  {
                    days: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"],
                    timeAndCapacity: {
                      maxInstanceCount: 3,
                      minInstanceCount: 3,
                      time: "09:00",
                    },
                  },
                  {
                    days: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"],
                    timeAndCapacity: {
                      maxInstanceCount: 6,
                      minInstanceCount: 6,
                      time: "18:00",
                    },
                  },
                  {
                    days: ["Saturday", "Sunday"],
                    timeAndCapacity: {
                      maxInstanceCount: 2,
                      minInstanceCount: 2,
                      time: "09:00",
                    },
                  },
                  {
                    days: ["Saturday", "Sunday"],
                    timeAndCapacity: {
                      maxInstanceCount: 4,
                      minInstanceCount: 4,
                      time: "18:00",
                    },
                  },
                ],
                timeZone: "China Standard Time",
              },
            },
            dataDisksGroups: [],
            hardwareProfile: { vmSize: "Standard_D4_V2" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                username: "sshuser",
              },
            },
            scriptActions: [],
            targetInstanceCount: 4,
            virtualNetworkProfile: {},
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
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    parameters,
  );
  console.log(result);
}

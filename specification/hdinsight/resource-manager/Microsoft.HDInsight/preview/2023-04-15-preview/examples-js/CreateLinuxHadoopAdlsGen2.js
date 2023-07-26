const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new HDInsight cluster with the specified parameters.
 *
 * @summary Creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/CreateLinuxHadoopAdlsGen2.json
 */
async function createHadoopClusterWithAzureDataLakeStorageGen2() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const parameters = {
    properties: {
      clusterDefinition: {
        configurations: {
          gateway: {
            "restAuthCredential.isEnabled": "true",
            "restAuthCredential.password": "**********",
            "restAuthCredential.username": "admin",
          },
        },
        kind: "Hadoop",
      },
      clusterVersion: "3.6",
      computeProfile: {
        roles: [
          {
            name: "headnode",
            hardwareProfile: { vmSize: "Standard_D3_V2" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                username: "sshuser",
              },
            },
            targetInstanceCount: 2,
          },
          {
            name: "workernode",
            hardwareProfile: { vmSize: "Standard_D3_V2" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                username: "sshuser",
              },
            },
            targetInstanceCount: 4,
          },
          {
            name: "zookeepernode",
            hardwareProfile: { vmSize: "Small" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                username: "sshuser",
              },
            },
            targetInstanceCount: 3,
          },
        ],
      },
      osType: "Linux",
      storageProfile: {
        storageaccounts: [
          {
            name: "mystorage.dfs.core.windows.net",
            enableSecureChannel: true,
            fileSystem: "default",
            isDefault: true,
            key: "storagekey",
          },
        ],
      },
      tier: "Standard",
    },
    tags: { key1: "val1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}

const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new HDInsight cluster with the specified parameters.
 *
 * @summary Creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/CreateKafkaClusterWithKafkaRestProxy.json
 */
async function createKafkaClusterWithKafkaRestProxy() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const parameters = {
    properties: {
      clusterDefinition: {
        componentVersion: { kafka: "2.1" },
        configurations: {
          gateway: {
            "restAuthCredential.isEnabled": true,
            "restAuthCredential.password": "**********",
            "restAuthCredential.username": "admin",
          },
        },
        kind: "kafka",
      },
      clusterVersion: "4.0",
      computeProfile: {
        roles: [
          {
            name: "headnode",
            hardwareProfile: { vmSize: "Large" },
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
            dataDisksGroups: [{ disksPerNode: 8 }],
            hardwareProfile: { vmSize: "Large" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                username: "sshuser",
              },
            },
            targetInstanceCount: 3,
          },
          {
            name: "zookeepernode",
            hardwareProfile: { vmSize: "Small" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                username: "sshuser",
              },
            },
            targetInstanceCount: 3,
          },
          {
            name: "kafkamanagementnode",
            hardwareProfile: { vmSize: "Standard_D4_v2" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                username: "kafkauser",
              },
            },
            targetInstanceCount: 2,
          },
        ],
      },
      kafkaRestProperties: {
        clientGroupInfo: {
          groupId: "00000000-0000-0000-0000-111111111111",
          groupName: "Kafka security group name",
        },
      },
      osType: "Linux",
      storageProfile: {
        storageaccounts: [
          {
            name: "mystorage.blob.core.windows.net",
            container: "containername",
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

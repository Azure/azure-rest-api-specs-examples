const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new HDInsight cluster with the specified parameters.
 *
 * @summary Creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/CreateHDInsightClusterWithComputeIsolationProperties.json
 */
async function createClusterWithComputeIsolationProperties() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const parameters = {
    properties: {
      clusterDefinition: {
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
      computeIsolationProperties: {
        enableComputeIsolation: true,
        hostSku: undefined,
      },
      computeProfile: {
        roles: [
          {
            name: "headnode",
            hardwareProfile: { vmSize: "standard_d3" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                sshProfile: { publicKeys: [{ certificateData: "**********" }] },
                username: "sshuser",
              },
            },
            targetInstanceCount: 2,
          },
          {
            name: "workernode",
            hardwareProfile: { vmSize: "standard_d3" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                sshProfile: { publicKeys: [{ certificateData: "**********" }] },
                username: "sshuser",
              },
            },
            targetInstanceCount: 2,
          },
        ],
      },
      osType: "Linux",
      storageProfile: {
        storageaccounts: [
          {
            name: "mystorage",
            container: "containername",
            enableSecureChannel: true,
            isDefault: true,
            key: "storage account key",
          },
        ],
      },
    },
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

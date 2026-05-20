const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new HDInsight cluster with the specified parameters.
 *
 * @summary creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: 2025-01-15-preview/CreateHDInsightClusterWithComputeIsolationProperties.json
 */
async function createClusterWithComputeIsolationProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId";
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.create("rg1", "cluster1", {
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
      computeIsolationProperties: { enableComputeIsolation: true },
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
  });
  console.log(result);
}

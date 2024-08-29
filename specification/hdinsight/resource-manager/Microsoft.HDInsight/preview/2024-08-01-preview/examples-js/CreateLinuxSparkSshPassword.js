const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new HDInsight cluster with the specified parameters.
 *
 * @summary Creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/CreateLinuxSparkSshPassword.json
 */
async function createSparkOnLinuxClusterWithSshPassword() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const parameters = {
    properties: {
      clusterDefinition: {
        componentVersion: { spark: "2.0" },
        configurations: {
          gateway: {
            "restAuthCredential.isEnabled": true,
            "restAuthCredential.password": "**********",
            "restAuthCredential.username": "admin",
          },
        },
        kind: "Spark",
      },
      clusterVersion: "3.5",
      computeProfile: {
        roles: [
          {
            name: "headnode",
            hardwareProfile: { vmSize: "Standard_D12_V2" },
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
            hardwareProfile: { vmSize: "Standard_D4_V2" },
            minInstanceCount: 1,
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                username: "sshuser",
              },
            },
            targetInstanceCount: 4,
          },
        ],
      },
      osType: "Linux",
      storageProfile: {
        storageaccounts: [
          {
            name: "mystorage.blob.core.windows.net",
            container: "containername",
            enableSecureChannel: true,
            isDefault: true,
            key: "storageapikey*",
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
    parameters,
  );
  console.log(result);
}

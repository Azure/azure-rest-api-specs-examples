const { HDInsightContainersManagementClient } = require("@azure/arm-hdinsightcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a cluster.
 *
 * @summary Creates a cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-11-01-preview/examples/CreateSparkCluster.json
 */
async function hdInsightSparkClusterPut() {
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
        clusterVersion: "0.0.1",
        identityProfile: {
          msiClientId: "de91f1d8-767f-460a-ac11-3cf103f74b34",
          msiObjectId: "40491351-c240-4042-91e0-f644a1d2b441",
          msiResourceId:
            "/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-msi",
        },
        ossVersion: "2.2.3",
        serviceConfigsProfiles: [
          {
            configs: [
              {
                component: "spark-config",
                files: [
                  {
                    fileName: "spark-defaults.conf",
                    values: { sparkEventLogEnabled: "true" },
                  },
                ],
              },
            ],
            serviceName: "spark-service",
          },
          {
            configs: [
              {
                component: "yarn-config",
                files: [
                  {
                    fileName: "core-site.xml",
                    values: {
                      fsDefaultFS: "wasb://testcontainer@teststorage.dfs.core.windows.net/",
                      storageContainer: "testcontainer",
                      storageKey: "test key",
                      storageName: "teststorage",
                      storageProtocol: "wasb",
                    },
                  },
                  {
                    fileName: "yarn-site.xml",
                    values: { yarnWebappUi2Enable: "false" },
                  },
                ],
              },
            ],
            serviceName: "yarn-service",
          },
        ],
        sparkProfile: {},
        sshProfile: { count: 2 },
      },
      clusterType: "spark",
      computeProfile: {
        nodes: [{ type: "worker", count: 4, vmSize: "Standard_D3_v2" }],
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

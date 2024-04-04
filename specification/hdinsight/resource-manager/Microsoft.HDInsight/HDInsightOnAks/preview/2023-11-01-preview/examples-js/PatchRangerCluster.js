const { HDInsightContainersManagementClient } = require("@azure/arm-hdinsightcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing Cluster.
 *
 * @summary Updates an existing Cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-11-01-preview/examples/PatchRangerCluster.json
 */
async function hdInsightRangerClusterPatchTags() {
  const subscriptionId =
    process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "10e32bab-26da-4cc4-a441-52b318f824e6";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "hiloResourcegroup";
  const clusterPoolName = "clusterpool1";
  const clusterName = "cluster1";
  const clusterPatchRequest = {
    properties: {
      clusterProfile: {
        rangerProfile: {
          rangerAdmin: {
            admins: ["testuser1@contoso.com", "testuser2@contoso.com"],
            database: {
              name: "testdb",
              host: "testsqlserver.database.windows.net",
              passwordSecretRef:
                "https://testkv.vault.azure.net/secrets/mysecret/5df6584d9c25418c8d900240aa6c3452",
              username: "admin",
            },
          },
          rangerAudit: {
            storageAccount: "https://teststorage.blob.core.windows.net/testblob",
          },
          rangerUsersync: {
            enabled: true,
            groups: [
              "0a53828f-36c9-44c3-be3d-99a7fce977ad",
              "13be6971-79db-4f33-9d41-b25589ca25ac",
            ],
            mode: "automatic",
            users: ["testuser1@contoso.com", "testuser2@contoso.com"],
          },
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightContainersManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginUpdateAndWait(
    resourceGroupName,
    clusterPoolName,
    clusterName,
    clusterPatchRequest,
  );
  console.log(result);
}

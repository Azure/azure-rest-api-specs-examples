const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the cluster identity certificate.
 *
 * @summary Updates the cluster identity certificate.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/HDI_Clusters_UpdateClusterIdentityCertificate.json
 */
async function updateClusterIdentityCertificate() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const parameters = {
    applicationId: "applicationId",
    certificate: "base64encodedcertificate",
    certificatePassword: "**********",
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginUpdateIdentityCertificateAndWait(
    resourceGroupName,
    clusterName,
    parameters,
  );
  console.log(result);
}

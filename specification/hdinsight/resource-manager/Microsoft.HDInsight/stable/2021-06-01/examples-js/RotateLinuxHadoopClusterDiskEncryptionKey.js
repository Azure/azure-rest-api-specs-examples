const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Rotate disk encryption key of the specified HDInsight cluster.
 *
 * @summary Rotate disk encryption key of the specified HDInsight cluster.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/RotateLinuxHadoopClusterDiskEncryptionKey.json
 */
async function rotateDiskEncryptionKeyOfTheSpecifiedHdInsightCluster() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cluster1";
  const parameters = {
    keyName: "newkeyname",
    keyVersion: "newkeyversion",
    vaultUri: "https://newkeyvault.vault.azure.net/",
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginRotateDiskEncryptionKeyAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}

rotateDiskEncryptionKeyOfTheSpecifiedHdInsightCluster().catch(console.error);

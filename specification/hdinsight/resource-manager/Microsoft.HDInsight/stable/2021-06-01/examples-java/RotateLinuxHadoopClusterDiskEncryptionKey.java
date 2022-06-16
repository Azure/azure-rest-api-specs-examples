import com.azure.core.util.Context;
import com.azure.resourcemanager.hdinsight.models.ClusterDiskEncryptionParameters;

/** Samples for Clusters RotateDiskEncryptionKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/RotateLinuxHadoopClusterDiskEncryptionKey.json
     */
    /**
     * Sample code: Rotate disk encryption key of the specified HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void rotateDiskEncryptionKeyOfTheSpecifiedHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .clusters()
            .rotateDiskEncryptionKey(
                "rg1",
                "cluster1",
                new ClusterDiskEncryptionParameters()
                    .withVaultUri("https://newkeyvault.vault.azure.net/")
                    .withKeyName("newkeyname")
                    .withKeyVersion("newkeyversion"),
                Context.NONE);
    }
}

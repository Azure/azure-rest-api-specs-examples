import com.azure.core.util.Context;

/** Samples for Clusters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetLinuxHadoopAllClustersInResourceGroup.json
     */
    /**
     * Sample code: Get All Hadoop on Linux clusters in a resource group.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAllHadoopOnLinuxClustersInAResourceGroup(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.clusters().listByResourceGroup("rg1", Context.NONE);
    }
}

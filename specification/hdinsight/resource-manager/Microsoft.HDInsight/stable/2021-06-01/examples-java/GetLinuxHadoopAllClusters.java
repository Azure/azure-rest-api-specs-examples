import com.azure.core.util.Context;

/** Samples for Clusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetLinuxHadoopAllClusters.json
     */
    /**
     * Sample code: Get All Hadoop on Linux clusters.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAllHadoopOnLinuxClusters(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.clusters().list(Context.NONE);
    }
}

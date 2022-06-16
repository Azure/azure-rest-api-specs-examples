import com.azure.core.util.Context;

/** Samples for Applications ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetAllApplications.json
     */
    /**
     * Sample code: Get All Applications for an HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAllApplicationsForAnHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.applications().listByCluster("rg1", "cluster1", Context.NONE);
    }
}

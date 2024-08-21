
/**
 * Samples for Extensions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetExtension.
     * json
     */
    /**
     * Sample code: Get an extension.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAnExtension(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.extensions().getWithResponse("rg1", "cluster1", "clustermonitoring", com.azure.core.util.Context.NONE);
    }
}

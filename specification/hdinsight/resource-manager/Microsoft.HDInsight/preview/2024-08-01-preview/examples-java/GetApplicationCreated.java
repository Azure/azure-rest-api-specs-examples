
/**
 * Samples for Applications Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetApplicationCreated.json
     */
    /**
     * Sample code: Get application on HDInsight cluster successfully created.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getApplicationOnHDInsightClusterSuccessfullyCreated(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.applications().getWithResponse("rg1", "cluster1", "app", com.azure.core.util.Context.NONE);
    }
}

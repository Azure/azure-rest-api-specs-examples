
/**
 * Samples for Configurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * HDI_Configurations_Get.json
     */
    /**
     * Sample code: Get Core site settings.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getCoreSiteSettings(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.configurations().getWithResponse("rg1", "cluster1", "core-site", com.azure.core.util.Context.NONE);
    }
}

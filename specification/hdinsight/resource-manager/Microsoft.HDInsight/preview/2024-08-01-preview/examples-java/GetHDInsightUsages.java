
/**
 * Samples for Locations ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetHDInsightUsages.json
     */
    /**
     * Sample code: Get the subscription usages for specific location.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void
        getTheSubscriptionUsagesForSpecificLocation(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.locations().listUsagesWithResponse("West US", com.azure.core.util.Context.NONE);
    }
}

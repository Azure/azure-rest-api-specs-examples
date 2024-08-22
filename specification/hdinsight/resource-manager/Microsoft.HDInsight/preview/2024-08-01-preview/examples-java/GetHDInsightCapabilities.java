
/**
 * Samples for Locations GetCapabilities.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetHDInsightCapabilities.json
     */
    /**
     * Sample code: Get the subscription capabilities for specific location.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getTheSubscriptionCapabilitiesForSpecificLocation(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.locations().getCapabilitiesWithResponse("West US", com.azure.core.util.Context.NONE);
    }
}

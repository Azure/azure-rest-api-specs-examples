import com.azure.core.util.Context;

/** Samples for Locations GetCapabilities. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetHDInsightCapabilities.json
     */
    /**
     * Sample code: Get the subscription capabilities for specific location.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getTheSubscriptionCapabilitiesForSpecificLocation(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.locations().getCapabilitiesWithResponse("West US", Context.NONE);
    }
}

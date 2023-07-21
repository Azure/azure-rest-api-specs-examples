import com.azure.resourcemanager.hdinsight.models.NameAvailabilityCheckRequestParameters;

/** Samples for Locations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/HDI_Locations_CheckClusterNameAvailability.json
     */
    /**
     * Sample code: Get the subscription usages for specific location.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getTheSubscriptionUsagesForSpecificLocation(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .locations()
            .checkNameAvailabilityWithResponse(
                "westus",
                new NameAvailabilityCheckRequestParameters().withName("test123").withType("clusters"),
                com.azure.core.util.Context.NONE);
    }
}

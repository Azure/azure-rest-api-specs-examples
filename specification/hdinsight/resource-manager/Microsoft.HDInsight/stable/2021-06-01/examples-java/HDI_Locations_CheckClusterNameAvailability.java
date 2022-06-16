import com.azure.core.util.Context;
import com.azure.resourcemanager.hdinsight.models.NameAvailabilityCheckRequestParameters;

/** Samples for Locations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Locations_CheckClusterNameAvailability.json
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
                Context.NONE);
    }
}

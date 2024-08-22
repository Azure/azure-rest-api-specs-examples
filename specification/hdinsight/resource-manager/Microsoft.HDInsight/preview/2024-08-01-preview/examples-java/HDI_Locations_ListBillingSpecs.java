
/**
 * Samples for Locations ListBillingSpecs.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * HDI_Locations_ListBillingSpecs.json
     */
    /**
     * Sample code: Get the subscription billingSpecs for the specified location.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getTheSubscriptionBillingSpecsForTheSpecifiedLocation(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.locations().listBillingSpecsWithResponse("East US 2", com.azure.core.util.Context.NONE);
    }
}

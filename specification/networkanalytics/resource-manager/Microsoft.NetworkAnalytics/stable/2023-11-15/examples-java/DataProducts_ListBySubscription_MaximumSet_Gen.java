
/**
 * Samples for DataProducts List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataProducts_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataProducts_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataProductsListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataProducts().list(com.azure.core.util.Context.NONE);
    }
}

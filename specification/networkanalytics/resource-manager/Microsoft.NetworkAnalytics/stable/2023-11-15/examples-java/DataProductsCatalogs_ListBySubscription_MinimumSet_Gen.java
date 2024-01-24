
/**
 * Samples for DataProductsCatalogs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataProductsCatalogs_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: DataProductsCatalogs_ListBySubscription_MinimumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataProductsCatalogsListBySubscriptionMinimumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataProductsCatalogs().list(com.azure.core.util.Context.NONE);
    }
}

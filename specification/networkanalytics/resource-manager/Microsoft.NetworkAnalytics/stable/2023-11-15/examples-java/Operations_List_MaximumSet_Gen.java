
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void
        operationsListMaximumSetGen(com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}

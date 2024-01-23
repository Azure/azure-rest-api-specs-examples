
/**
 * Samples for Functions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Function_Delete.json
     */
    /**
     * Sample code: Delete a function.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void deleteAFunction(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().deleteWithResponse("sjrg1637", "sj8653", "function8197", com.azure.core.util.Context.NONE);
    }
}

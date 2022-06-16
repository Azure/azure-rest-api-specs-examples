import com.azure.core.util.Context;

/** Samples for Functions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Delete.json
     */
    /**
     * Sample code: Delete a function.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void deleteAFunction(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().deleteWithResponse("sjrg1637", "sj8653", "function8197", Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for Functions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Get_JavaScript.json
     */
    /**
     * Sample code: Get a JavaScript function.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAJavaScriptFunction(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().getWithResponse("sjrg1637", "sj8653", "function8197", Context.NONE);
    }
}

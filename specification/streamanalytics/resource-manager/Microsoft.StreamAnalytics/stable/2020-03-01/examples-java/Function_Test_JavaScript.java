import com.azure.core.util.Context;

/** Samples for Functions Test. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Test_JavaScript.json
     */
    /**
     * Sample code: Test the connection for a JavaScript function.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void testTheConnectionForAJavaScriptFunction(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().test("sjrg1637", "sj8653", "function8197", null, Context.NONE);
    }
}

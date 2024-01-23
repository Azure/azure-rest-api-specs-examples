
/**
 * Samples for Functions Test.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Function_Test_JavaScript.json
     */
    /**
     * Sample code: Test the connection for a JavaScript function.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void testTheConnectionForAJavaScriptFunction(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().test("sjrg1637", "sj8653", "function8197", null, com.azure.core.util.Context.NONE);
    }
}

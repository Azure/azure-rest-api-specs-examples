
/**
 * Samples for Functions Test.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Function_Test_AzureML.json
     */
    /**
     * Sample code: Test the connection for an Azure ML function.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void testTheConnectionForAnAzureMLFunction(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().test("sjrg", "sjName", "function588", null, com.azure.core.util.Context.NONE);
    }
}

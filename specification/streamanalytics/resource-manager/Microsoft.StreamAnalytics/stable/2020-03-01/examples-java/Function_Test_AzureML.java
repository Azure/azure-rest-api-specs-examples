import com.azure.core.util.Context;

/** Samples for Functions Test. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Test_AzureML.json
     */
    /**
     * Sample code: Test the connection for an Azure ML function.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void testTheConnectionForAnAzureMLFunction(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().test("sjrg7", "sj9093", "function588", null, Context.NONE);
    }
}

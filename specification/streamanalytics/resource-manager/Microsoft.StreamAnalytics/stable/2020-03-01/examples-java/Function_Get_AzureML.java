import com.azure.core.util.Context;

/** Samples for Functions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Get_AzureML.json
     */
    /**
     * Sample code: Get an Azure ML function.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAnAzureMLFunction(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().getWithResponse("sjrg7", "sj9093", "function588", Context.NONE);
    }
}

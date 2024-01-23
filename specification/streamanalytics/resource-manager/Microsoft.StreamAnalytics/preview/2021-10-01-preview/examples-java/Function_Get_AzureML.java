
/**
 * Samples for Functions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Function_Get_AzureML.json
     */
    /**
     * Sample code: Get an Azure ML function.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAnAzureMLFunction(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().getWithResponse("sjrg7", "sj9093", "function588", com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Function_Get_JavaScript.json
     */
    /**
     * Sample code: Get a JavaScript function.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        getAJavaScriptFunction(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().getWithResponse("sjrg1637", "sj8653", "function8197", com.azure.core.util.Context.NONE);
    }
}

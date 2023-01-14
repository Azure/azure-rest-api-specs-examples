/** Samples for Predictions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/PredictionsGet.json
     */
    /**
     * Sample code: Predictions_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void predictionsGet(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.predictions().getWithResponse("TestHubRG", "sdkTestHub", "sdktest", com.azure.core.util.Context.NONE);
    }
}

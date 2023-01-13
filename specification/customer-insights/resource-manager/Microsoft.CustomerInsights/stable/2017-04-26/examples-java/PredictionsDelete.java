/** Samples for Predictions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/PredictionsDelete.json
     */
    /**
     * Sample code: Predictions_Delete.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void predictionsDelete(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.predictions().delete("TestHubRG", "sdkTestHub", "sdktest", com.azure.core.util.Context.NONE);
    }
}

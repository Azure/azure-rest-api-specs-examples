/** Samples for Predictions GetTrainingResults. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/PredictionsGetTrainingResults.json
     */
    /**
     * Sample code: Predictions_GetTrainingResults.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void predictionsGetTrainingResults(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .predictions()
            .getTrainingResultsWithResponse("TestHubRG", "sdkTestHub", "sdktest", com.azure.core.util.Context.NONE);
    }
}

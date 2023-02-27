/** Samples for Environments ListBySubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EnvironmentsListBySubscription.json
     */
    /**
     * Sample code: EnvironmentsBySubscription.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void environmentsBySubscription(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.environments().listBySubscriptionWithResponse(com.azure.core.util.Context.NONE);
    }
}

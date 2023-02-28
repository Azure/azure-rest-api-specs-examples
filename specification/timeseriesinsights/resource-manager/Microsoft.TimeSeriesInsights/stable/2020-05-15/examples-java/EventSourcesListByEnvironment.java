/** Samples for EventSources ListByEnvironment. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EventSourcesListByEnvironment.json
     */
    /**
     * Sample code: ListEventSourcesByEnvironment.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void listEventSourcesByEnvironment(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.eventSources().listByEnvironmentWithResponse("rg1", "env1", com.azure.core.util.Context.NONE);
    }
}

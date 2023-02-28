/** Samples for EventSources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EventSourcesGetEventHub.json
     */
    /**
     * Sample code: GetEventHubEventSource.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void getEventHubEventSource(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.eventSources().getWithResponse("rg1", "env1", "es1", com.azure.core.util.Context.NONE);
    }
}

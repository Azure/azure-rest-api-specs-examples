/** Samples for EventSources Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EventSourcesDelete.json
     */
    /**
     * Sample code: DeleteEventSource.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void deleteEventSource(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.eventSources().deleteWithResponse("rg1", "env1", "es1", com.azure.core.util.Context.NONE);
    }
}

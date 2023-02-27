/** Samples for Environments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EnvironmentsDelete.json
     */
    /**
     * Sample code: EnvironmentsDelete.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void environmentsDelete(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.environments().deleteByResourceGroupWithResponse("rg1", "env1", com.azure.core.util.Context.NONE);
    }
}

/** Samples for ReferenceDataSets ListByEnvironment. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsListByEnvironment.json
     */
    /**
     * Sample code: ReferenceDataSetsListByEnvironment.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void referenceDataSetsListByEnvironment(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.referenceDataSets().listByEnvironmentWithResponse("rg1", "env1", com.azure.core.util.Context.NONE);
    }
}

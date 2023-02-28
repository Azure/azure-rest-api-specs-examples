/** Samples for ReferenceDataSets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsDelete.json
     */
    /**
     * Sample code: ReferenceDataSetsDelete.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void referenceDataSetsDelete(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.referenceDataSets().deleteWithResponse("rg1", "env1", "rds1", com.azure.core.util.Context.NONE);
    }
}

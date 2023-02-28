/** Samples for ReferenceDataSets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsGet.json
     */
    /**
     * Sample code: ReferenceDataSetsGet.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void referenceDataSetsGet(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.referenceDataSets().getWithResponse("rg1", "env1", "rds1", com.azure.core.util.Context.NONE);
    }
}

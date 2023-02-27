/** Samples for AccessPolicies ListByEnvironment. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/AccessPoliciesListByEnvironment.json
     */
    /**
     * Sample code: AccessPoliciesByEnvironment.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void accessPoliciesByEnvironment(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.accessPolicies().listByEnvironmentWithResponse("rg1", "env1", com.azure.core.util.Context.NONE);
    }
}

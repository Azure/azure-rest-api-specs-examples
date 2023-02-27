/** Samples for AccessPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/AccessPoliciesDelete.json
     */
    /**
     * Sample code: AccessPoliciesDelete.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void accessPoliciesDelete(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.accessPolicies().deleteWithResponse("rg1", "env1", "ap1", com.azure.core.util.Context.NONE);
    }
}

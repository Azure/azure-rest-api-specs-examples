
/**
 * Samples for AccessPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/
     * AccessPoliciesGet.json
     */
    /**
     * Sample code: AccessPoliciesGet.
     * 
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void
        accessPoliciesGet(com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager.accessPolicies().getWithResponse("rg1", "env1", "ap1", com.azure.core.util.Context.NONE);
    }
}

/** Samples for Profiles GetEnrichingKpis. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ProfileGetEnrichingKpis.json
     */
    /**
     * Sample code: Profiles_GetEnrichingKpis.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void profilesGetEnrichingKpis(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .profiles()
            .getEnrichingKpisWithResponse(
                "TestHubRG", "sdkTestHub", "TestProfileType396", com.azure.core.util.Context.NONE);
    }
}

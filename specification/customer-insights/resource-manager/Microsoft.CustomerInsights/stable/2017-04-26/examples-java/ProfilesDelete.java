/** Samples for Profiles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ProfilesDelete.json
     */
    /**
     * Sample code: Profiles_Delete.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void profilesDelete(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .profiles()
            .delete("TestHubRG", "sdkTestHub", "TestProfileType396", null, com.azure.core.util.Context.NONE);
    }
}

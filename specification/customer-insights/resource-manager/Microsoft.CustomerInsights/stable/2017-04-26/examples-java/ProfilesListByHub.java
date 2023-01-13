/** Samples for Profiles ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ProfilesListByHub.json
     */
    /**
     * Sample code: Profiles_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void profilesListByHub(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.profiles().listByHub("TestHubRG", "sdkTestHub", null, com.azure.core.util.Context.NONE);
    }
}

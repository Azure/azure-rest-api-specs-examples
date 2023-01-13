/** Samples for AuthorizationPolicies ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesListByHub.json
     */
    /**
     * Sample code: AuthorizationPolicies_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void authorizationPoliciesListByHub(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.authorizationPolicies().listByHub("TestHubRG", "azSdkTestHub", com.azure.core.util.Context.NONE);
    }
}

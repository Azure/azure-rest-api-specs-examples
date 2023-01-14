/** Samples for AuthorizationPolicies RegeneratePrimaryKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesRegeneratePrimaryKey.json
     */
    /**
     * Sample code: AuthorizationPolicies_RegeneratePrimaryKey.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void authorizationPoliciesRegeneratePrimaryKey(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .authorizationPolicies()
            .regeneratePrimaryKeyWithResponse(
                "TestHubRG", "azSdkTestHub", "testPolicy4222", com.azure.core.util.Context.NONE);
    }
}

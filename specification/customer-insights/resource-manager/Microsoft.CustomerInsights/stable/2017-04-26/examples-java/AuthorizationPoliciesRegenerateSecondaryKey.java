/** Samples for AuthorizationPolicies RegenerateSecondaryKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesRegenerateSecondaryKey.json
     */
    /**
     * Sample code: AuthorizationPolicies_RegenerateSecondaryKey.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void authorizationPoliciesRegenerateSecondaryKey(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .authorizationPolicies()
            .regenerateSecondaryKeyWithResponse(
                "TestHubRG", "azSdkTestHub", "testPolicy4222", com.azure.core.util.Context.NONE);
    }
}

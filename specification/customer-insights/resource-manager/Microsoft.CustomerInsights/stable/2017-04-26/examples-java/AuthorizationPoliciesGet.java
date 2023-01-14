/** Samples for AuthorizationPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesGet.json
     */
    /**
     * Sample code: AuthorizationPolicies_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void authorizationPoliciesGet(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .authorizationPolicies()
            .getWithResponse("TestHubRG", "azSdkTestHub", "testPolicy4222", com.azure.core.util.Context.NONE);
    }
}

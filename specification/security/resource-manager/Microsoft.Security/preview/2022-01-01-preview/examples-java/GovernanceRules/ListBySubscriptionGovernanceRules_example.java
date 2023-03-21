/** Samples for GovernanceRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/ListBySubscriptionGovernanceRules_example.json
     */
    /**
     * Sample code: List governance rules by subscription scope.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listGovernanceRulesBySubscriptionScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .governanceRules()
            .list("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", com.azure.core.util.Context.NONE);
    }
}

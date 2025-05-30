
/**
 * Samples for GovernanceRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/
     * DeleteGovernanceRule_example.json
     */
    /**
     * Sample code: Delete a Governance rule over subscription scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        deleteAGovernanceRuleOverSubscriptionScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.governanceRules().delete("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", com.azure.core.util.Context.NONE);
    }
}

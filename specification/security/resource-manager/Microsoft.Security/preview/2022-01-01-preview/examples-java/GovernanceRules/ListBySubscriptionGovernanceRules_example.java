import com.azure.core.util.Context;

/** Samples for GovernanceRule List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/ListBySubscriptionGovernanceRules_example.json
     */
    /**
     * Sample code: List security governanceRules by subscription level scope.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecurityGovernanceRulesBySubscriptionLevelScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.governanceRules().list(Context.NONE);
    }
}

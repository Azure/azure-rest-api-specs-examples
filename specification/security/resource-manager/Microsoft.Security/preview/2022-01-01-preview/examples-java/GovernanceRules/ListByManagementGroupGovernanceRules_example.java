/** Samples for GovernanceRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/ListByManagementGroupGovernanceRules_example.json
     */
    /**
     * Sample code: List governance rules by management group scope.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listGovernanceRulesByManagementGroupScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .governanceRules()
            .list("providers/Microsoft.Management/managementGroups/contoso", com.azure.core.util.Context.NONE);
    }
}

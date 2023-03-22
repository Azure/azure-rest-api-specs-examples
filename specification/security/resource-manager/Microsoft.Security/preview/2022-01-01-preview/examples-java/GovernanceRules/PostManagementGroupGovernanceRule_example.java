/** Samples for GovernanceRules Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PostManagementGroupGovernanceRule_example.json
     */
    /**
     * Sample code: Execute governance rule over management group scope.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void executeGovernanceRuleOverManagementGroupScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .governanceRules()
            .execute(
                "providers/Microsoft.Management/managementGroups/contoso",
                "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
                null,
                com.azure.core.util.Context.NONE);
    }
}

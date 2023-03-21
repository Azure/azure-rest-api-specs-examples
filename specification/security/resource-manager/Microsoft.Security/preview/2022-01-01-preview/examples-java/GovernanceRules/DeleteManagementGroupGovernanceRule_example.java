/** Samples for GovernanceRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/DeleteManagementGroupGovernanceRule_example.json
     */
    /**
     * Sample code: Delete a Governance rule over management group scope.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteAGovernanceRuleOverManagementGroupScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .governanceRules()
            .delete(
                "providers/Microsoft.Management/managementGroups/contoso",
                "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
                com.azure.core.util.Context.NONE);
    }
}

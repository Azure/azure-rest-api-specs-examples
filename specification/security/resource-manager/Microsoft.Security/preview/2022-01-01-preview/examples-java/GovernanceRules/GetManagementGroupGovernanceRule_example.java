/** Samples for GovernanceRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/GetManagementGroupGovernanceRule_example.json
     */
    /**
     * Sample code: Get a governance rule over management group scope.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getAGovernanceRuleOverManagementGroupScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .governanceRules()
            .getWithResponse(
                "providers/Microsoft.Management/managementGroups/contoso",
                "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
                com.azure.core.util.Context.NONE);
    }
}

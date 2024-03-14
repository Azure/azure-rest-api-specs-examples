
/**
 * Samples for GovernanceRules OperationResults.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/
     * GetManagementGroupGovernanceRuleExecuteStatus_example.json
     */
    /**
     * Sample code: Get governance rules long run operation result over management group.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getGovernanceRulesLongRunOperationResultOverManagementGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.governanceRules().operationResultsWithResponse(
            "providers/Microsoft.Management/managementGroups/contoso", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
            "58b33f4f-c8c7-4b01-99cc-d437db4d40dd", com.azure.core.util.Context.NONE);
    }
}

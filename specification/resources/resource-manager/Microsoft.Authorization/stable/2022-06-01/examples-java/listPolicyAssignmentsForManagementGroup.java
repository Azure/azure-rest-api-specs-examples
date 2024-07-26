
/**
 * Samples for PolicyAssignments ListForManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/
     * listPolicyAssignmentsForManagementGroup.json
     */
    /**
     * Sample code: List policy assignments that apply to a management group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listPolicyAssignmentsThatApplyToAManagementGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments()
            .listForManagementGroup("TestManagementGroup", "atScope()", null, com.azure.core.util.Context.NONE);
    }
}

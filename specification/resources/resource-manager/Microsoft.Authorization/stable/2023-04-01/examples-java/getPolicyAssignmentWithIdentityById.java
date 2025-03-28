
/**
 * Samples for PolicyAssignments GetById.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * getPolicyAssignmentWithIdentityById.json
     */
    /**
     * Sample code: Retrieve a policy assignment with a managed identity by ID.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        retrieveAPolicyAssignmentWithAManagedIdentityByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments().getByIdWithResponse(
            "providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage",
            null, com.azure.core.util.Context.NONE);
    }
}

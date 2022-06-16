import com.azure.core.util.Context;

/** Samples for PolicyAssignments DeleteById. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/deletePolicyAssignmentById.json
     */
    /**
     * Sample code: Delete a policy assignment by ID.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPolicyAssignmentByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyAssignments()
            .deleteByIdWithResponse(
                "providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage",
                Context.NONE);
    }
}

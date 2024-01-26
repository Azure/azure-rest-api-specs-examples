
/** Samples for PolicyAssignments ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/
     * listPolicyAssignmentsForResourceGroup.json
     */
    /**
     * Sample code: List policy assignments that apply to a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listPolicyAssignmentsThatApplyToAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments()
            .listByResourceGroup("TestResourceGroup", "atScope()", null, com.azure.core.util.Context.NONE);
    }
}

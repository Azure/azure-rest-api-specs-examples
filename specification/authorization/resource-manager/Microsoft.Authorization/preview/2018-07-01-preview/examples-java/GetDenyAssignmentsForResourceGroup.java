import com.azure.core.util.Context;

/** Samples for DenyAssignments ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-07-01-preview/examples/GetDenyAssignmentsForResourceGroup.json
     */
    /**
     * Sample code: List deny assignments for resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDenyAssignmentsForResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getDenyAssignments()
            .listByResourceGroup("rgname", null, Context.NONE);
    }
}

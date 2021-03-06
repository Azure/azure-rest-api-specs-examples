import com.azure.core.util.Context;

/** Samples for Permissions ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-01-01-preview/examples/GetPermissions.json
     */
    /**
     * Sample code: List permissions for resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPermissionsForResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getPermissions()
            .listByResourceGroup("rgname", Context.NONE);
    }
}

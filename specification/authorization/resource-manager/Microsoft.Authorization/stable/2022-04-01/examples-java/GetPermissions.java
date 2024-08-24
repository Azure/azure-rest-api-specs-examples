
/**
 * Samples for Permissions ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetPermissions.
     * json
     */
    /**
     * Sample code: List permissions for resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPermissionsForResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getPermissions()
            .listByResourceGroup("rgname", com.azure.core.util.Context.NONE);
    }
}

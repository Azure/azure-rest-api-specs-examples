import com.azure.core.util.Context;

/** Samples for Permissions ListForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-01-01-preview/examples/GetResourcePermissions.json
     */
    /**
     * Sample code: List permissions for resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPermissionsForResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getPermissions()
            .listForResource(
                "rgname", "rpnamespace", "parentResourcePath", "resourceType", "resourceName", Context.NONE);
    }
}

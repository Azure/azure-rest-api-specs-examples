
/**
 * Samples for WorkspaceGlobalSchema Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceSchema.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceSchema.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceSchema(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGlobalSchemas().deleteWithResponse("rg1", "apimService1", "wks1", "schema1", "*",
            com.azure.core.util.Context.NONE);
    }
}

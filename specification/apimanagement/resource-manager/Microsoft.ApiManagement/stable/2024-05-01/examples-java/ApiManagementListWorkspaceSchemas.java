
/**
 * Samples for WorkspaceGlobalSchema ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceSchemas.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceSchemas.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceSchemas(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGlobalSchemas().listByService("rg1", "apimService1", "wks1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

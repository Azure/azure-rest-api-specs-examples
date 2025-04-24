
/**
 * Samples for WorkspaceBackend ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceBackends.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceBackends.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceBackends(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceBackends().listByWorkspace("rg1", "apimService1", "wks1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

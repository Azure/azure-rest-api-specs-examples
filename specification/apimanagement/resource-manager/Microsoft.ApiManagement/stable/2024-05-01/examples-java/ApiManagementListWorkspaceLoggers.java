
/**
 * Samples for WorkspaceLogger ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceLoggers.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceLoggers.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceLoggers(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceLoggers().listByWorkspace("rg1", "apimService1", "wks1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

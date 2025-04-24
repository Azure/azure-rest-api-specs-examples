
/**
 * Samples for WorkspaceLogger Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceLogger.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceLogger.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceLogger(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceLoggers().deleteWithResponse("rg1", "apimService1", "wks1", "loggerId", "*",
            com.azure.core.util.Context.NONE);
    }
}

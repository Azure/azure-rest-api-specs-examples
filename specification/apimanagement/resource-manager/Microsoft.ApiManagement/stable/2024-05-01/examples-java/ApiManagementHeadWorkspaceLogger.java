
/**
 * Samples for WorkspaceLogger GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceLogger.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceLogger.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceLogger(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceLoggers().getEntityTagWithResponse("rg1", "apimService1", "wks1", "templateLogger",
            com.azure.core.util.Context.NONE);
    }
}

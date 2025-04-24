
/**
 * Samples for WorkspaceLogger Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceLogger.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceLogger.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceLogger(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceLoggers().getWithResponse("rg1", "apimService1", "wks1", "templateLogger",
            com.azure.core.util.Context.NONE);
    }
}

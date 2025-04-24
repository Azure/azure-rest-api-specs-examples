
/**
 * Samples for WorkspaceBackend Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceBackend.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceBackend.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceBackend(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceBackends().getWithResponse("rg1", "apimService1", "wks1", "sfbackend",
            com.azure.core.util.Context.NONE);
    }
}

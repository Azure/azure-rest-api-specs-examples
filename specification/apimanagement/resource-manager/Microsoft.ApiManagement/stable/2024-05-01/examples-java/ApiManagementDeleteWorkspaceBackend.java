
/**
 * Samples for WorkspaceBackend Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceBackend.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceBackend.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceBackend(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceBackends().deleteWithResponse("rg1", "apimService1", "wks1", "sfbackend", "*",
            com.azure.core.util.Context.NONE);
    }
}

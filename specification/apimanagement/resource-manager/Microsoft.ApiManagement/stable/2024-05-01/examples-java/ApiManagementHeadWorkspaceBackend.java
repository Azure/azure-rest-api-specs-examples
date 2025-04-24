
/**
 * Samples for WorkspaceBackend GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceBackend.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceBackend.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceBackend(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceBackends().getEntityTagWithResponse("rg1", "apimService1", "wks1", "sfbackend",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkspaceTagApiLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceTagApiLink.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceTagApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceTagApiLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagApiLinks().deleteWithResponse("rg1", "apimService1", "wks1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}

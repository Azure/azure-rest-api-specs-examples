
/**
 * Samples for WorkspaceTagProductLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceTagProductLink.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceTagProductLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspaceTagProductLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagProductLinks().deleteWithResponse("rg1", "apimService1", "wks1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}

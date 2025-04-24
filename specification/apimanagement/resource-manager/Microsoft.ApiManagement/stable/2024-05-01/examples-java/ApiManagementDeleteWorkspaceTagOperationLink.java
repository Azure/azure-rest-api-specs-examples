
/**
 * Samples for WorkspaceTagOperationLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceTagOperationLink.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceTagOperationLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspaceTagOperationLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagOperationLinks().deleteWithResponse("rg1", "apimService1", "wks1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}

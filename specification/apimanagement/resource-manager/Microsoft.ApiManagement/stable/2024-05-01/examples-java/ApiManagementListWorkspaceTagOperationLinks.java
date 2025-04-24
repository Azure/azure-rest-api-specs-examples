
/**
 * Samples for WorkspaceTagOperationLink ListByProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceTagOperationLinks.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceTagOperationLinks.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListWorkspaceTagOperationLinks(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagOperationLinks().listByProduct("rg1", "apimService1", "wks1", "tag1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

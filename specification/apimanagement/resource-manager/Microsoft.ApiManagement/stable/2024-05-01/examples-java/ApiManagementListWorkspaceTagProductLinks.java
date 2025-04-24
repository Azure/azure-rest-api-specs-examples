
/**
 * Samples for WorkspaceTagProductLink ListByProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceTagProductLinks.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceTagProductLinks.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListWorkspaceTagProductLinks(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagProductLinks().listByProduct("rg1", "apimService1", "wks1", "tag1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

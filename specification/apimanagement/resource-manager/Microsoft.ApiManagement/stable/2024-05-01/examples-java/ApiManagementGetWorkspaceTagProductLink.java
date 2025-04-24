
/**
 * Samples for WorkspaceTagProductLink Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceTagProductLink.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceTagProductLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceTagProductLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagProductLinks().getWithResponse("rg1", "apimService1", "wks1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}

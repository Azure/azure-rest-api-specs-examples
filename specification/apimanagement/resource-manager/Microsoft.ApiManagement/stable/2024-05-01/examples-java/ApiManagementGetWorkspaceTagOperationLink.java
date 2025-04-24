
/**
 * Samples for WorkspaceTagOperationLink Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceTagOperationLink.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceTagOperationLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetWorkspaceTagOperationLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagOperationLinks().getWithResponse("rg1", "apimService1", "wks1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}

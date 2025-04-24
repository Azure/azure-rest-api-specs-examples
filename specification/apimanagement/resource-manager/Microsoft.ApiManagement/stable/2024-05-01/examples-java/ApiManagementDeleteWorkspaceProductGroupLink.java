
/**
 * Samples for WorkspaceProductGroupLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceProductGroupLink.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceProductGroupLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspaceProductGroupLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductGroupLinks().deleteWithResponse("rg1", "apimService1", "wks1", "testproduct", "link1",
            com.azure.core.util.Context.NONE);
    }
}

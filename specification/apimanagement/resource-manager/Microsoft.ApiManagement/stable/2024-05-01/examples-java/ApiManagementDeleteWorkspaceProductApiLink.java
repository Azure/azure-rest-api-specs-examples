
/**
 * Samples for WorkspaceProductApiLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceProductApiLink.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceProductApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspaceProductApiLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductApiLinks().deleteWithResponse("rg1", "apimService1", "wks1", "testproduct", "link1",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkspaceProductGroupLink Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceProductGroupLink.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceProductGroupLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetWorkspaceProductGroupLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductGroupLinks().getWithResponse("rg1", "apimService1", "wks1", "testproduct", "link1",
            com.azure.core.util.Context.NONE);
    }
}

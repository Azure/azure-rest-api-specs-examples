
/**
 * Samples for WorkspaceProductApiLink Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceProductApiLink.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceProductApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceProductApiLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductApiLinks().getWithResponse("rg1", "apimService1", "wks1", "testproduct", "link1",
            com.azure.core.util.Context.NONE);
    }
}

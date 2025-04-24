
/**
 * Samples for Workspace Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspace.json
     */
    /**
     * Sample code: ApiManagementGetWorkspace.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetWorkspace(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaces().getWithResponse("rg1", "apimService1", "wks1", com.azure.core.util.Context.NONE);
    }
}

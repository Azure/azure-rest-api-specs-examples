
/**
 * Samples for Workspace Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspace.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspace.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspace(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaces().deleteWithResponse("rg1", "apimService1", "wks1", "*", com.azure.core.util.Context.NONE);
    }
}

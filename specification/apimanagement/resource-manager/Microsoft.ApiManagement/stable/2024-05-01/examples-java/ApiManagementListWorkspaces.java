
/**
 * Samples for Workspace ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaces.json
     */
    /**
     * Sample code: ApiManagementListWorkspaces.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaces(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaces().listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}

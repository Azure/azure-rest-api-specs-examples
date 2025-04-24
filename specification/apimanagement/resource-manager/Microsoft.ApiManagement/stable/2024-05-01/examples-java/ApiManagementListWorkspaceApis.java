
/**
 * Samples for WorkspaceApi ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceApis.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceApis.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceApis(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApis().listByService("rg1", "apimService1", "wks1", null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkspaceGroup ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceGroups.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceGroups.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceGroups(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGroups().listByService("rg1", "apimService1", "wks1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

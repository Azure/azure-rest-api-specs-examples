
/**
 * Samples for WorkspaceApiVersionSet ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceApiVersionSets.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceApiVersionSets.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceApiVersionSets(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiVersionSets().listByService("rg1", "apimService1", "wks1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

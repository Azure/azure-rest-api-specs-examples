
/**
 * Samples for WorkspaceNamedValue RefreshSecret.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementRefreshWorkspaceNamedValue.json
     */
    /**
     * Sample code: ApiManagementRefreshWorkspaceNamedValue.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementRefreshWorkspaceNamedValue(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().refreshSecret("rg1", "apimService1", "wks1", "testprop2",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkspaceNamedValue Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceNamedValue.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceNamedValue.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceNamedValue(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().deleteWithResponse("rg1", "apimService1", "wks1", "testprop2", "*",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkspaceApiVersionSet Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceApiVersionSet.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceApiVersionSet.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspaceApiVersionSet(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiVersionSets().deleteWithResponse("rg1", "apimService1", "wks1", "a1", "*",
            com.azure.core.util.Context.NONE);
    }
}

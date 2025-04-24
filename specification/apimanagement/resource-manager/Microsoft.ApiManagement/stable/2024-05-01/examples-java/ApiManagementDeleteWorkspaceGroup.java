
/**
 * Samples for WorkspaceGroup Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceGroup.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceGroup.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceGroup(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGroups().deleteWithResponse("rg1", "apimService1", "wks1", "aadGroup", "*",
            com.azure.core.util.Context.NONE);
    }
}

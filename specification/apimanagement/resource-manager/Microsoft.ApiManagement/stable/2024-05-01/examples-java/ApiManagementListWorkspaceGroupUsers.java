
/**
 * Samples for WorkspaceGroupUser List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceGroupUsers.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceGroupUsers.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceGroupUsers(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGroupUsers().list("rg1", "apimService1", "wks1", "57d2ef278aa04f0888cba3f3", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

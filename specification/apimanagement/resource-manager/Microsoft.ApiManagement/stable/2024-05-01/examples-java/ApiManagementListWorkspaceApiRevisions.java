
/**
 * Samples for WorkspaceApiRevision ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceApiRevisions.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceApiRevisions.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceApiRevisions(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiRevisions().listByService("rg1", "apimService1", "wks1", "57d2ef278aa04f0888cba3f3", null,
            null, null, com.azure.core.util.Context.NONE);
    }
}

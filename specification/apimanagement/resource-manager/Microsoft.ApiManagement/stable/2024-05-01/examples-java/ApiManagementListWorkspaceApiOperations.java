
/**
 * Samples for WorkspaceApiOperation ListByApi.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceApiOperations.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceApiOperations.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceApiOperations(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiOperations().listByApi("rg1", "apimService1", "wks1", "57d2ef278aa04f0888cba3f3", null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}

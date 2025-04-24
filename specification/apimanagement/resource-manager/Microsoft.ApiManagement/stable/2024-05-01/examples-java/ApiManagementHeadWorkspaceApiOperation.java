
/**
 * Samples for WorkspaceApiOperation GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceApiOperation.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceApiOperation.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceApiOperation(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiOperations().getEntityTagWithResponse("rg1", "apimService1", "wks1",
            "57d2ef278aa04f0888cba3f3", "57d2ef278aa04f0ad01d6cdc", com.azure.core.util.Context.NONE);
    }
}

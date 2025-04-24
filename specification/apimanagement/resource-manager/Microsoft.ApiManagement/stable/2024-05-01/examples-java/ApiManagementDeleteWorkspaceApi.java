
/**
 * Samples for WorkspaceApi Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceApi.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceApi.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceApi(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApis().deleteWithResponse("rg1", "apimService1", "wks1", "echo-api", "*", null,
            com.azure.core.util.Context.NONE);
    }
}

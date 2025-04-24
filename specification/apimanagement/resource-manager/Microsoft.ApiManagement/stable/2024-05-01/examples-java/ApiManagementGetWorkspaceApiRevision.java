
/**
 * Samples for WorkspaceApi Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceApiRevision.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceApiRevision.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceApiRevision(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApis().getWithResponse("rg1", "apimService1", "wks1", "echo-api;rev=3",
            com.azure.core.util.Context.NONE);
    }
}

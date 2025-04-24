
/**
 * Samples for WorkspaceApiRelease Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceApiRelease.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceApiRelease.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceApiRelease(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiReleases().deleteWithResponse("rg1", "apimService1", "wks1", "5a5fcc09124a7fa9b89f2f1d",
            "testrev", "*", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkspaceApiRelease ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceApiReleases.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceApiReleases.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceApiReleases(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiReleases().listByService("rg1", "apimService1", "wks1", "a1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkspaceTag ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceTags.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceTags.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceTags(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTags().listByService("rg1", "apimService1", "wks1", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

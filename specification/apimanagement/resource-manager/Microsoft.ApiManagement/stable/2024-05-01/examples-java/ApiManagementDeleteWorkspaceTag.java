
/**
 * Samples for WorkspaceTag Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceTag.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceTag.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTags().deleteWithResponse("rg1", "apimService1", "wks1", "tagId1", "*",
            com.azure.core.util.Context.NONE);
    }
}

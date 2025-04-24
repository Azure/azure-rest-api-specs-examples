
/**
 * Samples for WorkspaceTag GetEntityState.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceTag.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceTag.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTags().getEntityStateWithResponse("rg1", "apimService1", "wks1", "59306a29e4bbd510dc24e5f9",
            com.azure.core.util.Context.NONE);
    }
}

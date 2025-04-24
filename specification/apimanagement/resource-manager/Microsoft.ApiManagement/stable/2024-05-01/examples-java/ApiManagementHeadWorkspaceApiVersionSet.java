
/**
 * Samples for WorkspaceApiVersionSet GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceApiVersionSet.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceApiVersionSet.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceApiVersionSet(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiVersionSets().getEntityTagWithResponse("rg1", "apimService1", "wks1", "vs1",
            com.azure.core.util.Context.NONE);
    }
}

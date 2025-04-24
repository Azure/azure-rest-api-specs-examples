
/**
 * Samples for Workspace GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspace.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspace.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspace(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaces().getEntityTagWithResponse("rg1", "apimService1", "wks", com.azure.core.util.Context.NONE);
    }
}

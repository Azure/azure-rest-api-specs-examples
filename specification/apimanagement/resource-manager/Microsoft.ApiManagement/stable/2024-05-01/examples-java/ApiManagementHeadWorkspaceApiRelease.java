
/**
 * Samples for WorkspaceApiRelease GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceApiRelease.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceApiRelease.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceApiRelease(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiReleases().getEntityTagWithResponse("rg1", "apimService1", "wks1", "a1",
            "5a7cb545298324c53224a799", com.azure.core.util.Context.NONE);
    }
}

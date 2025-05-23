
/**
 * Samples for WorkspaceSubscription GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceSubscription.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceSubscription.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceSubscription(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceSubscriptions().getEntityTagWithResponse("rg1", "apimService1", "wks1",
            "5931a769d8d14f0ad8ce13b8", com.azure.core.util.Context.NONE);
    }
}

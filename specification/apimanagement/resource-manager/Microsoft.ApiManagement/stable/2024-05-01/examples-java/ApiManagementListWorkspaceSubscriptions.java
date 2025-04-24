
/**
 * Samples for WorkspaceSubscription List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceSubscriptions.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceSubscriptions.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceSubscriptions(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceSubscriptions().list("rg1", "apimService1", "wks1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}

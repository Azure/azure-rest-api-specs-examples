
/**
 * Samples for WorkspaceSubscription RegeneratePrimaryKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementWorkspaceSubscriptionRegeneratePrimaryKey.json
     */
    /**
     * Sample code: ApiManagementWorkspaceSubscriptionRegeneratePrimaryKey.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementWorkspaceSubscriptionRegeneratePrimaryKey(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceSubscriptions().regeneratePrimaryKeyWithResponse("rg1", "apimService1", "wks1", "testsub",
            com.azure.core.util.Context.NONE);
    }
}

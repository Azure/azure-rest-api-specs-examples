
/**
 * Samples for WorkspaceSubscription RegenerateSecondaryKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementWorkspaceSubscriptionRegenerateSecondaryKey.json
     */
    /**
     * Sample code: ApiManagementWorkspaceSubscriptionRegenerateSecondaryKey.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementWorkspaceSubscriptionRegenerateSecondaryKey(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceSubscriptions().regenerateSecondaryKeyWithResponse("rg1", "apimService1", "wks1", "testsub",
            com.azure.core.util.Context.NONE);
    }
}

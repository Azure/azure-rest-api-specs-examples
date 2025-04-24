
/**
 * Samples for WorkspaceSubscription ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementWorkspaceSubscriptionListSecrets.json
     */
    /**
     * Sample code: ApiManagementWorkspaceSubscriptionListSecrets.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementWorkspaceSubscriptionListSecrets(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceSubscriptions().listSecretsWithResponse("rg1", "apimService1", "wks1",
            "5931a769d8d14f0ad8ce13b8", com.azure.core.util.Context.NONE);
    }
}

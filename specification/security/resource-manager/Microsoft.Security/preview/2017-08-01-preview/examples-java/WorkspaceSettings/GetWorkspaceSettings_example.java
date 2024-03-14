
/**
 * Samples for WorkspaceSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/
     * GetWorkspaceSettings_example.json
     */
    /**
     * Sample code: Get workspace settings on subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getWorkspaceSettingsOnSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.workspaceSettings().list(com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for WorkspaceSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/
     * GetWorkspaceSetting_example.json
     */
    /**
     * Sample code: Get a workspace setting on subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getAWorkspaceSettingOnSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.workspaceSettings().getWithResponse("default", com.azure.core.util.Context.NONE);
    }
}

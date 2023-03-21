/** Samples for WorkspaceSettings Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/DeleteWorkspaceSetting_example.json
     */
    /**
     * Sample code: Delete a workspace setting data for resource group.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteAWorkspaceSettingDataForResourceGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.workspaceSettings().deleteWithResponse("default", com.azure.core.util.Context.NONE);
    }
}

/** Samples for WorkspaceSettings Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/CreateWorkspaceSetting_example.json
     */
    /**
     * Sample code: Create a workspace setting data for subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void createAWorkspaceSettingDataForSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .workspaceSettings()
            .define("default")
            .withWorkspaceId(
                "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace")
            .withScope("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23")
            .create();
    }
}

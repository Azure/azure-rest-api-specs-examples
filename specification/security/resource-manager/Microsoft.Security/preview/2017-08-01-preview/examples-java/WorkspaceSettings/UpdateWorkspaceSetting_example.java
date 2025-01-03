
import com.azure.resourcemanager.security.models.WorkspaceSetting;

/**
 * Samples for WorkspaceSettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/
     * UpdateWorkspaceSetting_example.json
     */
    /**
     * Sample code: Update a workspace setting data for subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        updateAWorkspaceSettingDataForSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        WorkspaceSetting resource
            = manager.workspaceSettings().getWithResponse("default", com.azure.core.util.Context.NONE).getValue();
        resource.update().withWorkspaceId(
            "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace")
            .apply();
    }
}

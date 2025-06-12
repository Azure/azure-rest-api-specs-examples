
/**
 * Samples for OnlineExperimentationWorkspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_ListByResourceGroup.json
     */
    /**
     * Sample code: List OnlineExperimentationWorkspaces in a resource group.
     * 
     * @param manager Entry point to OnlineExperimentationManager.
     */
    public static void listOnlineExperimentationWorkspacesInAResourceGroup(
        com.azure.resourcemanager.onlineexperimentation.OnlineExperimentationManager manager) {
        manager.onlineExperimentationWorkspaces().listByResourceGroup("res9871", com.azure.core.util.Context.NONE);
    }
}

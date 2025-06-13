
/**
 * Samples for OnlineExperimentationWorkspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_ListBySubscription.json
     */
    /**
     * Sample code: List Online Experimentation Workspaces in a subscription.
     * 
     * @param manager Entry point to OnlineExperimentationManager.
     */
    public static void listOnlineExperimentationWorkspacesInASubscription(
        com.azure.resourcemanager.onlineexperimentation.OnlineExperimentationManager manager) {
        manager.onlineExperimentationWorkspaces().list(com.azure.core.util.Context.NONE);
    }
}

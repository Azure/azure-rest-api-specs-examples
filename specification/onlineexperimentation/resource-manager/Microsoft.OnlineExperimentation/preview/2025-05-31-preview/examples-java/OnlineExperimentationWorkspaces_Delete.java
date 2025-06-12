
/**
 * Samples for OnlineExperimentationWorkspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_Delete.json
     */
    /**
     * Sample code: Delete an Online Experimentation Workspace.
     * 
     * @param manager Entry point to OnlineExperimentationManager.
     */
    public static void deleteAnOnlineExperimentationWorkspace(
        com.azure.resourcemanager.onlineexperimentation.OnlineExperimentationManager manager) {
        manager.onlineExperimentationWorkspaces().delete("res9871", "expworkspace3", com.azure.core.util.Context.NONE);
    }
}

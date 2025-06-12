
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_OperationsList.json
     */
    /**
     * Sample code: List Online Experimentation Workspaces operations.
     * 
     * @param manager Entry point to OnlineExperimentationManager.
     */
    public static void listOnlineExperimentationWorkspacesOperations(
        com.azure.resourcemanager.onlineexperimentation.OnlineExperimentationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}

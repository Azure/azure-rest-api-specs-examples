
/**
 * Samples for WorkspaceConnections ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/WorkspaceConnection/listSecrets.json
     */
    /**
     * Sample code: GetWorkspaceConnection.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceConnection(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaceConnections().listSecretsWithResponse("test-rg", "workspace-1", "connection-1",
            com.azure.core.util.Context.NONE);
    }
}

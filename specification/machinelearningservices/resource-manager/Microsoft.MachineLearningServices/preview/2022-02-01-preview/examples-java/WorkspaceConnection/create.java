/** Samples for WorkspaceConnections Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/create.json
     */
    /**
     * Sample code: CreateWorkspaceConnection.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createWorkspaceConnection(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .workspaceConnections()
            .define("connection-1")
            .withExistingWorkspace("resourceGroup-1", "workspace-1")
            .withCategory("ACR")
            .withTarget("www.facebook.com")
            .withAuthType("PAT")
            .withValue("secrets")
            .create();
    }
}

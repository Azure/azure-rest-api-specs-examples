
/**
 * Samples for Workspaces ListNotebookKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Notebook/listKeys.json
     */
    /**
     * Sample code: List Workspace Keys.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listWorkspaceKeys(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().listNotebookKeysWithResponse("testrg123", "workspaces123",
            com.azure.core.util.Context.NONE);
    }
}

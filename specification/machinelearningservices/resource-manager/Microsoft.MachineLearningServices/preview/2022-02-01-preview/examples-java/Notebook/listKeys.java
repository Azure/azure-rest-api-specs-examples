import com.azure.core.util.Context;

/** Samples for Workspaces ListNotebookKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Notebook/listKeys.json
     */
    /**
     * Sample code: List Workspace Keys.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listWorkspaceKeys(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaces().listNotebookKeysWithResponse("testrg123", "workspaces123", Context.NONE);
    }
}

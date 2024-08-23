
/**
 * Samples for Workspaces PrepareNotebook.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Notebook/prepare.json
     */
    /**
     * Sample code: Prepare Notebook.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void prepareNotebook(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.workspaces().prepareNotebook("testrg123", "workspaces123", com.azure.core.util.Context.NONE);
    }
}

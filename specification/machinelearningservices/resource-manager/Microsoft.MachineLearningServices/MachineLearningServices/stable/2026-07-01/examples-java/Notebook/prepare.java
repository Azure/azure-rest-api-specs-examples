
/**
 * Samples for Workspaces PrepareNotebook.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Notebook/prepare.json
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

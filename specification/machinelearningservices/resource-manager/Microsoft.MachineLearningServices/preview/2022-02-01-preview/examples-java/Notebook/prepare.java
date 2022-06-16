import com.azure.core.util.Context;

/** Samples for Workspaces PrepareNotebook. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Notebook/prepare.json
     */
    /**
     * Sample code: Prepare Notebook.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void prepareNotebook(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.workspaces().prepareNotebook("testrg123", "workspaces123", Context.NONE);
    }
}
